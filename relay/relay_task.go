package relay

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/QuantumNous/new-api/common"
	"github.com/QuantumNous/new-api/constant"
	"github.com/QuantumNous/new-api/dto"
	"github.com/QuantumNous/new-api/logger"
	"github.com/QuantumNous/new-api/model"
	"github.com/QuantumNous/new-api/relay/channel"
	"github.com/QuantumNous/new-api/relay/channel/task/taskcommon"
	relaycommon "github.com/QuantumNous/new-api/relay/common"
	relayconstant "github.com/QuantumNous/new-api/relay/constant"
	"github.com/QuantumNous/new-api/relay/helper"
	"github.com/QuantumNous/new-api/service"
	"github.com/QuantumNous/new-api/setting/operation_setting"
	"github.com/gin-gonic/gin"
)

type TaskSubmitResult struct {
	UpstreamTaskID string
	TaskData       []byte
	Platform       constant.TaskPlatform
	Quota          int

	// BackgroundAdaptor is non-nil when the upstream image API is synchronous
	// and the controller should insert a "queued" task, respond immediately,
	// then spawn a goroutine via RunBackgroundTask.
	// RelayTaskSubmit returns early (after billing pre-consume) without calling
	// BuildRequestBody / DoRequest / DoResponse.
	BackgroundAdaptor channel.BackgroundTaskAdaptor
}

// ResolveOriginTask 处理基于已有任务的提交（remix / continuation）：
// 查找原始任务、从中提取模型名称、将渠道锁定到原始任务的渠道
// （通过 info.LockedChannel，重试时复用同一渠道并轮换 key），
// 以及提取 OtherRatios（时长、分辨率）。
// 该函数在控制器的重试循环之前调用一次，其结果通过 info 字段和上下文持久化。
func ResolveOriginTask(c *gin.Context, info *relaycommon.RelayInfo) *dto.TaskError {
	// 检测 remix action
	path := c.Request.URL.Path
	if strings.Contains(path, "/v1/videos/") && strings.HasSuffix(path, "/remix") {
		info.Action = constant.TaskActionRemix
	}

	// 提取 remix 任务的 video_id
	if info.Action == constant.TaskActionRemix {
		videoID := c.Param("video_id")
		if strings.TrimSpace(videoID) == "" {
			return service.TaskErrorWrapperLocal(fmt.Errorf("video_id is required"), "invalid_request", http.StatusBadRequest)
		}
		info.OriginTaskID = videoID
	}

	if info.OriginTaskID == "" {
		return nil
	}

	// 查找原始任务
	originTask, exist, err := model.GetByTaskId(info.UserId, info.OriginTaskID)
	if err != nil {
		return service.TaskErrorWrapper(err, "get_origin_task_failed", http.StatusInternalServerError)
	}
	if !exist {
		return service.TaskErrorWrapperLocal(errors.New("task_origin_not_exist"), "task_not_exist", http.StatusBadRequest)
	}

	// 从原始任务推导模型名称
	if info.OriginModelName == "" {
		if originTask.Properties.OriginModelName != "" {
			info.OriginModelName = originTask.Properties.OriginModelName
		} else if originTask.Properties.UpstreamModelName != "" {
			info.OriginModelName = originTask.Properties.UpstreamModelName
		} else {
			var taskData map[string]interface{}
			_ = common.Unmarshal(originTask.Data, &taskData)
			if m, ok := taskData["model"].(string); ok && m != "" {
				info.OriginModelName = m
			}
		}
	}

	// 锁定到原始任务的渠道（重试时复用同一渠道，轮换 key）
	ch, err := model.GetChannelById(originTask.ChannelId, true)
	if err != nil {
		return service.TaskErrorWrapperLocal(err, "channel_not_found", http.StatusBadRequest)
	}
	if ch.Status != common.ChannelStatusEnabled {
		return service.TaskErrorWrapperLocal(errors.New("the channel of the origin task is disabled"), "task_channel_disable", http.StatusBadRequest)
	}
	info.LockedChannel = ch

	if originTask.ChannelId != info.ChannelId {
		key, _, newAPIError := ch.GetNextEnabledKey()
		if newAPIError != nil {
			return service.TaskErrorWrapper(newAPIError, "channel_no_available_key", newAPIError.StatusCode)
		}
		common.SetContextKey(c, constant.ContextKeyChannelKey, key)
		common.SetContextKey(c, constant.ContextKeyChannelType, ch.Type)
		common.SetContextKey(c, constant.ContextKeyChannelBaseUrl, ch.GetBaseURL())
		common.SetContextKey(c, constant.ContextKeyChannelId, originTask.ChannelId)

		info.ChannelBaseUrl = ch.GetBaseURL()
		info.ChannelId = originTask.ChannelId
		info.ChannelType = ch.Type
		info.ApiKey = key
	}

	// 提取 remix 参数（时长、分辨率 → OtherRatios）
	if info.Action == constant.TaskActionRemix {
		if originTask.PrivateData.BillingContext != nil {
			// 新的 remix 逻辑：直接从原始任务的 BillingContext 中提取 OtherRatios（如果存在）
			for s, f := range originTask.PrivateData.BillingContext.OtherRatios {
				info.PriceData.AddOtherRatio(s, f)
			}
		} else {
			// 旧的 remix 逻辑：直接从 task data 解析 seconds 和 size（如果存在）
			var taskData map[string]interface{}
			_ = common.Unmarshal(originTask.Data, &taskData)
			secondsStr, _ := taskData["seconds"].(string)
			seconds, _ := strconv.Atoi(secondsStr)
			if seconds <= 0 {
				seconds = 4
			}
			sizeStr, _ := taskData["size"].(string)
			if info.PriceData.OtherRatios == nil {
				info.PriceData.OtherRatios = map[string]float64{}
			}
			info.PriceData.OtherRatios["seconds"] = float64(seconds)
			info.PriceData.OtherRatios["size"] = 1
			if sizeStr == "1792x1024" || sizeStr == "1024x1792" {
				info.PriceData.OtherRatios["size"] = 1.666667
			}
		}
	}

	return nil
}

// RelayTaskSubmit 完成 task 提交的全部流程（每次尝试调用一次）：
// 刷新渠道元数据 → 确定 platform/adaptor → 验证请求 →
// 估算计费(EstimateBilling) → 计算价格 → 预扣费（仅首次）→
// 构建/发送/解析上游请求 → 提交后计费调整(AdjustBillingOnSubmit)。
// 控制器负责 defer Refund 和成功后 Settle。
func RelayTaskSubmit(c *gin.Context, info *relaycommon.RelayInfo) (*TaskSubmitResult, *dto.TaskError) {
	info.InitChannelMeta(c)

	// 1. 确定 platform → 创建适配器 → 验证请求
	platform := constant.TaskPlatform(c.GetString("platform"))
	if platform == "" {
		platform = GetTaskPlatform(c)
	}
	adaptor := GetTaskAdaptor(platform)
	if adaptor == nil {
		return nil, service.TaskErrorWrapperLocal(fmt.Errorf("invalid api platform: %s", platform), "invalid_api_platform", http.StatusBadRequest)
	}
	adaptor.Init(info)
	// 前置处理：v2 图片任务使用 ImageRequest 格式，在 adaptor 之前统一转换为 TaskSubmitReq
	if c.GetInt("relay_mode") == relayconstant.RelayModeImageTaskSubmit {
		if taskErr := relaycommon.ValidateImageTaskRequest(c, info); taskErr != nil {
			return nil, taskErr
		}
	}
	if taskErr := adaptor.ValidateRequestAndSetAction(c, info); taskErr != nil {
		return nil, taskErr
	}

	// 2. 确定模型名称
	modelName := info.OriginModelName
	if modelName == "" {
		modelName = service.CoverTaskActionToModelName(platform, info.Action)
	}

	// 2.5 应用渠道的模型映射（与同步任务对齐）
	info.OriginModelName = modelName
	info.UpstreamModelName = modelName
	if err := helper.ModelMappedHelper(c, info, nil); err != nil {
		return nil, service.TaskErrorWrapperLocal(err, "model_mapping_failed", http.StatusBadRequest)
	}

	// 3. 预生成公开 task ID（仅首次）
	if info.PublicTaskID == "" {
		info.PublicTaskID = model.GenerateTaskID()
	}

	// 4. 价格计算：基础模型价格
	info.OriginModelName = modelName
	priceData, err := helper.ModelPriceHelperPerCall(c, info)
	if err != nil {
		return nil, service.TaskErrorWrapper(err, "model_price_error", http.StatusBadRequest)
	}
	info.PriceData = priceData

	// 5. 计费估算：让适配器根据用户请求提供 OtherRatios（时长、分辨率等）
	//    必须在 ModelPriceHelperPerCall 之后调用（它会重建 PriceData）。
	//    ResolveOriginTask 可能已在 remix 路径中预设了 OtherRatios，此处合并。
	if estimatedRatios := adaptor.EstimateBilling(c, info); len(estimatedRatios) > 0 {
		for k, v := range estimatedRatios {
			info.PriceData.AddOtherRatio(k, v)
		}
	}

	// 6. 将 OtherRatios 应用到基础额度
	if !common.StringsContains(constant.TaskPricePatches, modelName) {
		for _, ra := range info.PriceData.OtherRatios {
			if ra != 1.0 {
				info.PriceData.Quota = int(float64(info.PriceData.Quota) * ra)
			}
		}
	}

	// 7. 预扣费（仅首次 — 重试时 info.Billing 已存在，跳过）
	if info.Billing == nil && !info.PriceData.FreeModel {
		info.ForcePreConsume = true
		if apiErr := service.PreConsumeBilling(c, info.PriceData.Quota, info); apiErr != nil {
			return nil, service.TaskErrorFromAPIError(apiErr)
		}
	}

	// 7.5 Background-mode detection: sync-API-as-async (e.g. Volcengine image gen).
	// Return early; the controller will insert a "queued" task, send the HTTP
	// response immediately, and then spawn RunBackgroundTask in a goroutine.
	if bgAdaptor, ok := adaptor.(channel.BackgroundTaskAdaptor); ok && bgAdaptor.IsBackgroundSubmit(info.UpstreamModelName) {
		// Snapshot the request now — gin.Context won't be available in the goroutine.
		if req, reqErr := relaycommon.GetTaskRequest(c); reqErr == nil {
			info.BackgroundTaskReq = &req
		}
		return &TaskSubmitResult{
			Platform:          platform,
			Quota:             info.PriceData.Quota,
			BackgroundAdaptor: bgAdaptor,
		}, nil
	}

	// 8. 构建请求体
	requestBody, err := adaptor.BuildRequestBody(c, info)
	if err != nil {
		return nil, service.TaskErrorWrapper(err, "build_request_failed", http.StatusInternalServerError)
	}

	// 9. 发送请求
	resp, err := adaptor.DoRequest(c, info, requestBody)
	if err != nil {
		return nil, service.TaskErrorWrapper(err, "do_request_failed", http.StatusInternalServerError)
	}
	if resp != nil && resp.StatusCode != http.StatusOK {
		responseBody, _ := io.ReadAll(resp.Body)
		return nil, service.TaskErrorWrapper(fmt.Errorf("%s", string(responseBody)), "fail_to_fetch_task", resp.StatusCode)
	}

	// 10. 返回 OtherRatios 给下游（header 必须在 DoResponse 写 body 之前设置）
	otherRatios := info.PriceData.OtherRatios
	if otherRatios == nil {
		otherRatios = map[string]float64{}
	}
	ratiosJSON, _ := common.Marshal(otherRatios)
	c.Header("X-New-Api-Other-Ratios", string(ratiosJSON))

	// 11. 解析响应
	upstreamTaskID, taskData, taskErr := adaptor.DoResponse(c, resp, info)
	if taskErr != nil {
		return nil, taskErr
	}

	// 11. 提交后计费调整：让适配器根据上游实际返回调整 OtherRatios
	finalQuota := info.PriceData.Quota
	if adjustedRatios := adaptor.AdjustBillingOnSubmit(info, taskData); len(adjustedRatios) > 0 {
		// 基于调整后的 ratios 重新计算 quota
		finalQuota = recalcQuotaFromRatios(info, adjustedRatios)
		info.PriceData.OtherRatios = adjustedRatios
		info.PriceData.Quota = finalQuota
	}

	return &TaskSubmitResult{
		UpstreamTaskID: upstreamTaskID,
		TaskData:       taskData,
		Platform:       platform,
		Quota:          finalQuota,
	}, nil
}

// recalcQuotaFromRatios 根据 adjustedRatios 重新计算 quota。
// 公式: baseQuota × ∏(ratio) — 其中 baseQuota 是不含 OtherRatios 的基础额度。
func recalcQuotaFromRatios(info *relaycommon.RelayInfo, ratios map[string]float64) int {
	// 从 PriceData 获取不含 OtherRatios 的基础价格
	baseQuota := info.PriceData.Quota
	// 先除掉原有的 OtherRatios 恢复基础额度
	for _, ra := range info.PriceData.OtherRatios {
		if ra != 1.0 && ra > 0 {
			baseQuota = int(float64(baseQuota) / ra)
		}
	}
	// 应用新的 ratios
	result := float64(baseQuota)
	for _, ra := range ratios {
		if ra != 1.0 {
			result *= ra
		}
	}
	return int(result)
}

// RunBackgroundTask executes a synchronous upstream image API call inside a
// background goroutine. It updates the task status in the DB and handles
// billing refund on failure. It MUST NOT use the gin.Context for HTTP writes.
func RunBackgroundTask(task *model.Task, adaptor channel.BackgroundTaskAdaptor, info *relaycommon.RelayInfo) {
	ctx := context.Background()

	resultURL, rawBody, err := adaptor.ExecuteBackgroundTask(info)

	task.FinishTime = time.Now().Unix()

	if err != nil {
		logger.LogError(ctx, fmt.Sprintf("RunBackgroundTask failed (task=%s): %v", task.TaskID, err))
		task.Status = model.TaskStatusFailure
		task.FailReason = err.Error()
		task.Progress = "0%"
		if updateErr := task.Update(); updateErr != nil {
			logger.LogError(ctx, fmt.Sprintf("RunBackgroundTask: update task failed (task=%s): %v", task.TaskID, updateErr))
		}
		service.RefundTaskQuota(ctx, task, "background image task failed: "+err.Error())
		return
	}

	task.Status = model.TaskStatusSuccess
	task.Progress = taskcommon.ProgressComplete
	if resultURL != "" {
		task.PrivateData.ResultURL = resultURL
	}
	if len(rawBody) > 0 {
		task.Data = rawBody
	}
	if updateErr := task.Update(); updateErr != nil {
		logger.LogError(ctx, fmt.Sprintf("RunBackgroundTask: update task failed (task=%s): %v", task.TaskID, updateErr))
	}

	// ── 差额结算（与轮询路径 settleTaskBillingOnComplete 对齐）──────────────────
	//
	// 规则：
	//   1. token 计费模型（hasRatioSetting）→ 按实际 TotalTokens 重算
	//   2. 固定价/按次计费模型（PerCallBilling）→ 按实际生成图片数占预估图片数的比例退款
	//      背景：sequential_image_generation=auto 模式，上游可能生成少于 n 张图，
	//             预扣是按 n 张估算的，需要按实际生成数补差。
	//   3. 无调整 → 保持预扣额度不变
	//
	// 重算后必须再次 task.Update()，确保 task.Quota 写回 DB，
	// 否则查询接口返回的 amount 仍是预估值。
	if len(rawBody) > 0 {
		var imgResp dto.ImageResponse
		if jsonErr := json.Unmarshal(rawBody, &imgResp); jsonErr == nil {
			settled := false

			// 优先：token 计费模型，按实际 TotalTokens 重算（内部自动跳过非 ratio 模型）
			if imgResp.Usage != nil && imgResp.Usage.TotalTokens > 0 {
				beforeQuota := task.Quota
				service.RecalculateTaskQuotaByTokens(ctx, task, imgResp.Usage.TotalTokens)
				if task.Quota != beforeQuota {
					settled = true
				}
			}

			// 按实际生成图片数多退少补（兜底，覆盖所有计费模式）
			// 条件：token 重算未生效，且实际图片数与请求数不符
			// 适用范围：fixed-price、tiered_expr（param("n") 线性乘法）等所有计费模型。
			// 原理：预扣 quota 与 n 成正比，实际 quota = preQuota × actualN / nRequested。
			if !settled {
				nRequested := 0
				if info.BackgroundTaskReq != nil {
					nRequested = info.BackgroundTaskReq.N
				}
				actualImages := len(imgResp.Data)
				if nRequested > 0 && actualImages > 0 && actualImages != nRequested {
					// 多退少补：实际费用 = 预扣总额 × 实际张数 / 请求张数
					actualQuota := task.Quota * actualImages / nRequested
					reason := fmt.Sprintf("image count adjust: requested=%d actual=%d", nRequested, actualImages)
					service.RecalculateTaskQuota(ctx, task, actualQuota, reason)
					settled = true
				}
			}

			// 差额结算更新了 task.Quota（内存），写回 DB 保证 amount 准确
			if settled {
				if updateErr := task.Update(); updateErr != nil {
					logger.LogError(ctx, fmt.Sprintf("RunBackgroundTask: update task quota after settle failed (task=%s): %v", task.TaskID, updateErr))
				}
			}
		}
	}
}

var fetchRespBuilders = map[int]func(c *gin.Context) (respBody []byte, taskResp *dto.TaskError){
	relayconstant.RelayModeSunoFetchByID:      sunoFetchByIDRespBodyBuilder,
	relayconstant.RelayModeSunoFetch:          sunoFetchRespBodyBuilder,
	relayconstant.RelayModeVideoFetchByID:     videoFetchByIDRespBodyBuilder,
	relayconstant.RelayModeImageTaskFetchByID: imageTaskFetchByIDRespBodyBuilder,
}

func RelayTaskFetch(c *gin.Context, relayMode int) (taskResp *dto.TaskError) {
	respBuilder, ok := fetchRespBuilders[relayMode]
	if !ok {
		taskResp = service.TaskErrorWrapperLocal(errors.New("invalid_relay_mode"), "invalid_relay_mode", http.StatusBadRequest)
	}

	respBody, taskErr := respBuilder(c)
	if taskErr != nil {
		return taskErr
	}
	if len(respBody) == 0 {
		respBody = []byte("{\"code\":\"success\",\"data\":null}")
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	_, err := io.Copy(c.Writer, bytes.NewBuffer(respBody))
	if err != nil {
		taskResp = service.TaskErrorWrapper(err, "copy_response_body_failed", http.StatusInternalServerError)
		return
	}
	return
}

func sunoFetchRespBodyBuilder(c *gin.Context) (respBody []byte, taskResp *dto.TaskError) {
	userId := c.GetInt("id")
	var condition = struct {
		IDs    []any  `json:"ids"`
		Action string `json:"action"`
	}{}
	err := c.BindJSON(&condition)
	if err != nil {
		taskResp = service.TaskErrorWrapper(err, "invalid_request", http.StatusBadRequest)
		return
	}
	var tasks []any
	if len(condition.IDs) > 0 {
		taskModels, err := model.GetByTaskIds(userId, condition.IDs)
		if err != nil {
			taskResp = service.TaskErrorWrapper(err, "get_tasks_failed", http.StatusInternalServerError)
			return
		}
		for _, task := range taskModels {
			tasks = append(tasks, TaskModel2Dto(task))
		}
	} else {
		tasks = make([]any, 0)
	}
	respBody, err = common.Marshal(dto.TaskResponse[[]any]{
		Code: "success",
		Data: tasks,
	})
	return
}

func sunoFetchByIDRespBodyBuilder(c *gin.Context) (respBody []byte, taskResp *dto.TaskError) {
	taskId := c.Param("id")
	userId := c.GetInt("id")

	originTask, exist, err := model.GetByTaskId(userId, taskId)
	if err != nil {
		taskResp = service.TaskErrorWrapper(err, "get_task_failed", http.StatusInternalServerError)
		return
	}
	if !exist {
		taskResp = service.TaskErrorWrapperLocal(errors.New("task_not_exist"), "task_not_exist", http.StatusBadRequest)
		return
	}

	respBody, err = common.Marshal(dto.TaskResponse[any]{
		Code: "success",
		Data: TaskModel2Dto(originTask),
	})
	return
}

func videoFetchByIDRespBodyBuilder(c *gin.Context) (respBody []byte, taskResp *dto.TaskError) {
	taskId := c.Param("task_id")
	if taskId == "" {
		taskId = c.GetString("task_id")
	}
	userId := c.GetInt("id")

	originTask, exist, err := model.GetByTaskId(userId, taskId)
	if err != nil {
		taskResp = service.TaskErrorWrapper(err, "get_task_failed", http.StatusInternalServerError)
		return
	}
	if !exist {
		taskResp = service.TaskErrorWrapperLocal(errors.New("task_not_exist"), "task_not_exist", http.StatusBadRequest)
		return
	}

	isOpenAIVideoAPI := strings.HasPrefix(c.Request.RequestURI, "/v1/videos/")

	// Gemini/Vertex 支持实时查询：用户 fetch 时直接从上游拉取最新状态
	if realtimeResp := tryRealtimeFetch(originTask, isOpenAIVideoAPI); len(realtimeResp) > 0 {
		respBody = realtimeResp
		return
	}

	// OpenAI Video API 格式: 走各 adaptor 的 ConvertToOpenAIVideo
	if isOpenAIVideoAPI {
		adaptor := GetTaskAdaptor(originTask.Platform)
		if adaptor == nil {
			taskResp = service.TaskErrorWrapperLocal(fmt.Errorf("invalid channel id: %d", originTask.ChannelId), "invalid_channel_id", http.StatusBadRequest)
			return
		}
		if converter, ok := adaptor.(channel.OpenAIVideoConverter); ok {
			openAIVideoData, err := converter.ConvertToOpenAIVideo(originTask)
			if err != nil {
				taskResp = service.TaskErrorWrapper(err, "convert_to_openai_video_failed", http.StatusInternalServerError)
				return
			}
			respBody = openAIVideoData
			return
		}
		taskResp = service.TaskErrorWrapperLocal(fmt.Errorf("not_implemented:%s", originTask.Platform), "not_implemented", http.StatusNotImplemented)
		return
	}

	// 通用 TaskDto 格式
	respBody, err = common.Marshal(dto.TaskResponse[any]{
		Code: "success",
		Data: TaskModel2Dto(originTask),
	})
	if err != nil {
		taskResp = service.TaskErrorWrapper(err, "marshal_response_failed", http.StatusInternalServerError)
	}
	return
}

// tryRealtimeFetch 尝试从上游实时拉取 Gemini/Vertex 任务状态。
// 仅当渠道类型为 Gemini 或 Vertex 时触发；其他渠道或出错时返回 nil。
// 当非 OpenAI Video API 时，还会构建自定义格式的响应体。
func tryRealtimeFetch(task *model.Task, isOpenAIVideoAPI bool) []byte {
	channelModel, err := model.GetChannelById(task.ChannelId, true)
	if err != nil {
		return nil
	}
	if channelModel.Type != constant.ChannelTypeVertexAi && channelModel.Type != constant.ChannelTypeGemini {
		return nil
	}

	baseURL := constant.ChannelBaseURLs[channelModel.Type]
	if channelModel.GetBaseURL() != "" {
		baseURL = channelModel.GetBaseURL()
	}
	proxy := channelModel.GetSetting().Proxy
	adaptor := GetTaskAdaptor(constant.TaskPlatform(strconv.Itoa(channelModel.Type)))
	if adaptor == nil {
		return nil
	}

	resp, err := adaptor.FetchTask(baseURL, channelModel.Key, map[string]any{
		"task_id": task.GetUpstreamTaskID(),
		"action":  task.Action,
	}, proxy)
	if err != nil || resp == nil {
		return nil
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil
	}

	ti, err := adaptor.ParseTaskResult(body)
	if err != nil || ti == nil {
		return nil
	}

	snap := task.Snapshot()

	// 将上游最新状态更新到 task
	if ti.Status != "" {
		task.Status = model.TaskStatus(ti.Status)
	}
	if ti.Progress != "" {
		task.Progress = ti.Progress
	}
	if ti.Url != "" {
		// data: URI（base64）与普通 URL 均写入 ResultURL，供查询接口读取
		task.PrivateData.ResultURL = ti.Url
	} else if task.Status == model.TaskStatusSuccess {
		// No URL from adaptor — construct proxy URL using public task ID
		task.PrivateData.ResultURL = taskcommon.BuildProxyURL(task.TaskID)
	}

	if !snap.Equal(task.Snapshot()) {
		_, _ = task.UpdateWithStatus(snap.Status)
	}

	// OpenAI Video API 由调用者的 ConvertToOpenAIVideo 分支处理
	if isOpenAIVideoAPI {
		return nil
	}

	// 非 OpenAI Video API: 构建自定义格式响应
	format := detectVideoFormat(body)
	out := map[string]any{
		"error":    nil,
		"format":   format,
		"metadata": nil,
		"status":   mapTaskStatusToSimple(task.Status),
		"task_id":  task.TaskID,
		"url":      task.GetResultURL(),
	}
	respBody, _ := common.Marshal(dto.TaskResponse[any]{
		Code: "success",
		Data: out,
	})
	return respBody
}

// detectVideoFormat 从 Gemini/Vertex 原始响应中探测视频格式
func detectVideoFormat(rawBody []byte) string {
	var raw map[string]any
	if err := common.Unmarshal(rawBody, &raw); err != nil {
		return "mp4"
	}
	respObj, ok := raw["response"].(map[string]any)
	if !ok {
		return "mp4"
	}
	vids, ok := respObj["videos"].([]any)
	if !ok || len(vids) == 0 {
		return "mp4"
	}
	v0, ok := vids[0].(map[string]any)
	if !ok {
		return "mp4"
	}
	mt, ok := v0["mimeType"].(string)
	if !ok || mt == "" || strings.Contains(mt, "mp4") {
		return "mp4"
	}
	return mt
}

// mapTaskStatusToSimple 将内部 TaskStatus 映射为简化状态字符串
func mapTaskStatusToSimple(status model.TaskStatus) string {
	switch status {
	case model.TaskStatusSuccess:
		return "succeeded"
	case model.TaskStatusFailure:
		return "failed"
	case model.TaskStatusQueued, model.TaskStatusSubmitted:
		return "queued"
	default:
		return "processing"
	}
}

func TaskModel2Dto(task *model.Task) *dto.TaskDto {
	usd := float64(task.Quota) / common.QuotaPerUnit
	cny := usd * operation_setting.USDExchangeRate
	amount := fmt.Sprintf("%.6f", cny)

	return &dto.TaskDto{
		ID:         task.ID,
		CreatedAt:  task.CreatedAt,
		UpdatedAt:  task.UpdatedAt,
		TaskID:     task.TaskID,
		Platform:   string(task.Platform),
		UserId:     task.UserId,
		Group:      task.Group,
		ChannelId:  task.ChannelId,
		Amount:     amount,
		Action:     task.Action,
		Status:     string(task.Status),
		FailReason: task.FailReason,
		ResultURL:  task.GetResultURL(),
		SubmitTime: task.SubmitTime,
		StartTime:  task.StartTime,
		FinishTime: task.FinishTime,
		Progress:   task.Progress,
		Properties: task.Properties,
		Username:   task.Username,
		Data:       task.Data,
	}
}

// imageTaskFetchByIDRespBodyBuilder 是 /v2/image-tasks/:task_id 的查询构建器
func imageTaskFetchByIDRespBodyBuilder(c *gin.Context) (respBody []byte, taskResp *dto.TaskError) {
	taskId := c.Param("task_id")
	userId := c.GetInt("id")

	task, exist, err := model.GetByTaskId(userId, taskId)
	if err != nil {
		taskResp = service.TaskErrorWrapper(err, "get_task_failed", http.StatusInternalServerError)
		return
	}
	if !exist {
		taskResp = service.TaskErrorWrapperLocal(errors.New("task_not_exist"), "task_not_exist", http.StatusNotFound)
		return
	}

	respBody, err = common.Marshal(dto.TaskResponse[any]{
		Code: "success",
		Data: TaskModel2ImageTaskDto(task),
	})
	if err != nil {
		taskResp = service.TaskErrorWrapper(err, "marshal_response_failed", http.StatusInternalServerError)
	}
	return
}

// TaskModel2ImageTaskDto 将 model.Task 转换为面向外部的 ImageTaskDto
// 去掉内部敏感字段，金额换算为人民币，usage 从 task.Data 原始响应体中提取并标准化
func TaskModel2ImageTaskDto(task *model.Task) *dto.ImageTaskDto {
	// quota → CNY amount
	usd := float64(task.Quota) / common.QuotaPerUnit
	cny := usd * operation_setting.USDExchangeRate
	amount := fmt.Sprintf("%.6f", cny)

	result := &dto.ImageTaskDto{
		TaskID:     task.TaskID,
		Status:     string(task.Status),
		Progress:   task.Progress,
		FailReason: task.FailReason,
		Model:      task.Properties.OriginModelName,
		Amount:     amount,
		Usage:      extractImageTaskUsage(task.Data),
		CreatedAt:  task.CreatedAt,
		SubmitTime: task.SubmitTime,
		FinishTime: task.FinishTime,
		Images:     []dto.ImageItem{},
	}

	// images：优先从 task.Data 中解析多张图片（后台异步任务存储完整 ImageResponse）；
	// 退化为单张 PrivateData.ResultURL。
	// 注意：不能用 task.GetResultURL()，该方法在 ResultURL 为空时会 fallback 到
	// FailReason，会把错误信息错误地写入 images[].url。
	if items := extractImagesFromTaskData(task.Data); len(items) > 0 {
		result.Images = items
	} else if resultURL := task.PrivateData.ResultURL; resultURL != "" {
		if strings.HasPrefix(resultURL, "data:") {
			// 剥离 "data:image/xxx;base64," 前缀，只保留纯 base64 数据
			if idx := strings.Index(resultURL, ";base64,"); idx != -1 {
				result.Images = append(result.Images, dto.ImageItem{B64Json: resultURL[idx+8:]})
			}
		} else {
			result.Images = append(result.Images, dto.ImageItem{URL: resultURL})
		}
	}
	return result
}

// extractImagesFromTaskData 尝试将 task.Data 解析为 dto.ImageResponse 并提取图片列表。
// 用于后台异步图片任务（存储完整的上游 ImageResponse JSON）。
// 如果 task.Data 不含有效的 data 数组，返回 nil（调用方退化为 ResultURL 模式）。
func extractImagesFromTaskData(taskData json.RawMessage) []dto.ImageItem {
	if len(taskData) == 0 {
		return nil
	}
	var imgResp dto.ImageResponse
	if err := json.Unmarshal(taskData, &imgResp); err != nil || len(imgResp.Data) == 0 {
		return nil
	}
	items := make([]dto.ImageItem, 0, len(imgResp.Data))
	for _, d := range imgResp.Data {
		if d.Url != "" {
			items = append(items, dto.ImageItem{URL: d.Url})
		} else if d.B64Json != "" {
			items = append(items, dto.ImageItem{B64Json: d.B64Json})
		}
	}
	return items
}

// extractImageTaskUsage 从上游原始响应体中提取并标准化 token 用量
//
// 渠道差异处理：
//   - 火山(Doubao)：task.Data["usage"]["completion_tokens"] → output_tokens
//   - 京东：task.Data["usage_metadata"]["input_tokens"/"output_tokens"]
//   - 腾讯 VOD：无 usage 字段 → 全返回 0（按次计费）
func extractImageTaskUsage(taskData json.RawMessage) dto.ImageTaskUsage {
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(taskData, &raw); err != nil {
		return dto.ImageTaskUsage{}
	}

	// 优先取 "usage"，其次取 "usage_metadata"（京东）
	usageRaw, ok := raw["usage"]
	if !ok {
		usageRaw, ok = raw["usage_metadata"]
	}
	if !ok {
		return dto.ImageTaskUsage{}
	}

	var fields map[string]int
	_ = json.Unmarshal(usageRaw, &fields)

	outputTokens := fields["output_tokens"]
	if outputTokens == 0 {
		outputTokens = fields["completion_tokens"] // 火山兼容
	}
	return dto.ImageTaskUsage{
		InputTokens:  fields["input_tokens"],
		OutputTokens: outputTokens,
		TotalTokens:  fields["total_tokens"],
	}
}
