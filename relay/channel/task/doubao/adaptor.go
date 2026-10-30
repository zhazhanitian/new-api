package doubao

import (
	"bytes"
	"context"
	"encoding/json"
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
	"github.com/QuantumNous/new-api/relay/channel/volcengine"
	relaycommon "github.com/QuantumNous/new-api/relay/common"
	relayconstant "github.com/QuantumNous/new-api/relay/constant"
	"github.com/QuantumNous/new-api/service"
	portraitSvc "github.com/QuantumNous/new-api/service/portrait"
	"github.com/QuantumNous/new-api/setting/portrait_setting"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/samber/lo"
)

// ============================
// Request / Response structures
// ============================

type ContentItem struct {
	Type      string     `json:"type,omitempty"`
	Text      string     `json:"text,omitempty"`
	ImageURL  *MediaURL  `json:"image_url,omitempty"`
	VideoURL  *MediaURL  `json:"video_url,omitempty"`
	AudioURL  *MediaURL  `json:"audio_url,omitempty"`
	DraftTask *DraftTask `json:"draft_task,omitempty"` // Seedance 2.5：样片任务 → 正式视频
	Role      string     `json:"role,omitempty"`
}

type MediaURL struct {
	URL string `json:"url,omitempty"`
}

// DraftTask 样片任务引用（仅 Seedance 2.5）。
type DraftTask struct {
	ID string `json:"id,omitempty"`
}

type requestPayload struct {
	Model                 string         `json:"model"`
	Content               []ContentItem  `json:"content,omitempty"`
	CallbackURL           string         `json:"callback_url,omitempty"`
	ReturnLastFrame       *dto.BoolValue `json:"return_last_frame,omitempty"`
	ServiceTier           string         `json:"service_tier,omitempty"`
	ExecutionExpiresAfter *dto.IntValue  `json:"execution_expires_after,omitempty"`
	GenerateAudio         *dto.BoolValue `json:"generate_audio,omitempty"`
	Draft                 *dto.BoolValue `json:"draft,omitempty"`                    // Seedance 2.5：样片模式
	OmniReferenceTaskType string         `json:"omni_reference_task_type,omitempty"` // Seedance 2.5：auto/reference/edit/extend
	OutputFormat          string         `json:"output_format,omitempty"`            // Seedance 2.5：mp4/mov
	Tools                 []struct {
		Type string `json:"type,omitempty"`
	} `json:"tools,omitempty"`
	Resolution  string         `json:"resolution,omitempty"`
	Ratio       string         `json:"ratio,omitempty"`
	Duration    *dto.IntValue  `json:"duration,omitempty"`
	Frames      *dto.IntValue  `json:"frames,omitempty"`
	Seed        *dto.IntValue  `json:"seed,omitempty"`
	CameraFixed *dto.BoolValue `json:"camera_fixed,omitempty"`
	Watermark   *dto.BoolValue `json:"watermark,omitempty"`
}

type responsePayload struct {
	ID string `json:"id"` // task_id
}

type responseTask struct {
	ID      string `json:"id"`
	Model   string `json:"model"`
	Status  string `json:"status"`
	Content struct {
		VideoURL string `json:"video_url"`
	} `json:"content"`
	Seed            int    `json:"seed"`
	Resolution      string `json:"resolution"`
	Duration        int    `json:"duration"`
	Ratio           string `json:"ratio"`
	FramesPerSecond int    `json:"framespersecond"`
	ServiceTier     string `json:"service_tier"`
	Tools           []struct {
		Type string `json:"type"`
	} `json:"tools"`
	Usage struct {
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
		ToolUsage        struct {
			WebSearch int `json:"web_search"`
		} `json:"tool_usage"`
	} `json:"usage"`
	Error struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
}

// ============================
// Adaptor implementation
// ============================

type TaskAdaptor struct {
	taskcommon.BaseBilling
	ChannelType int
	apiKey      string
	baseURL     string
}

func (a *TaskAdaptor) Init(info *relaycommon.RelayInfo) {
	a.ChannelType = info.ChannelType
	a.baseURL = info.ChannelBaseUrl
	a.apiKey = info.ApiKey
}

// ValidateRequestAndSetAction parses body, validates fields and sets default action.
func (a *TaskAdaptor) ValidateRequestAndSetAction(c *gin.Context, info *relaycommon.RelayInfo) (taskErr *dto.TaskError) {
	if c.GetInt("relay_mode") == relayconstant.RelayMode3DTaskSubmit {
		// 3D 任务：请求已由 relay_task.go 的前置处理解析到 context，直接校验
		return validate3DRequest(c, info)
	}
	// Seedance：prompt 必填（提示词优选外层 prompt）
	if taskErr = relaycommon.ValidateBasicTaskRequest(c, info, constant.TaskActionGenerate); taskErr != nil {
		return taskErr
	}
	// 人像素材：请求里的 asset:// 必须在当前选中渠道就绪。放在预扣费之前，失败按 400 返回。
	if req, err := relaycommon.GetTaskRequest(c); err == nil {
		for _, item := range taskcommon.NormalizeVideoContent(req) {
			if _, err := portraitSvc.RewriteAssetURL(info.UserId, portrait_setting.AliasForChannelType(info.ChannelType), item.URL); err != nil {
				return service.TaskErrorWrapperLocal(err, "portrait_asset_unavailable", http.StatusBadRequest)
			}
		}
	}
	return nil
}

// rewritePortraitAssets 把 content 里的 asset://逻辑ID 改写成当前渠道账号的上游 ID。
// 素材必须属于该用户，且当前渠道的绑定为 Active，否则报错，不透传给上游。
func rewritePortraitAssets(userId, channelType int, content []ContentItem) error {
	alias := portrait_setting.AliasForChannelType(channelType)
	for i := range content {
		for _, media := range []*MediaURL{content[i].ImageURL, content[i].VideoURL, content[i].AudioURL} {
			if media == nil {
				continue
			}
			rewritten, err := portraitSvc.RewriteAssetURL(userId, alias, media.URL)
			if err != nil {
				return err
			}
			media.URL = rewritten
		}
	}
	return nil
}

// BuildRequestURL constructs the upstream URL.
func (a *TaskAdaptor) BuildRequestURL(_ *relaycommon.RelayInfo) (string, error) {
	return fmt.Sprintf("%s/api/v3/contents/generations/tasks", a.baseURL), nil
}

// BuildRequestHeader sets required headers.
func (a *TaskAdaptor) BuildRequestHeader(_ *gin.Context, req *http.Request, _ *relaycommon.RelayInfo) error {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+a.apiKey)
	return nil
}

// EstimateBilling 检测视频输入折扣、分辨率调价、有声/无声折扣，返回 OtherRatios。
// 火山 3D 模型按 ModelRatio 走 token 计费，无需 OtherRatios，直接返回 nil。
func (a *TaskAdaptor) EstimateBilling(c *gin.Context, info *relaycommon.RelayInfo) map[string]float64 {
	if c.GetInt("relay_mode") == relayconstant.RelayMode3DTaskSubmit {
		return nil
	}
	req, err := relaycommon.GetTaskRequest(c)
	if err != nil {
		return nil
	}
	relaycommon.NormalizeVideoResolutionRatio(&req)

	otherRatios := make(map[string]float64)

	billingModel := billingModelName(info)
	if normalizedContentHasVideo(taskcommon.NormalizeVideoContent(req)) {
		if ratio, ok := GetVideoInputRatio(billingModel); ok {
			otherRatios["video_input"] = ratio
		}
	}

	// 分辨率只读收拢后的外层 resolution。
	resolution := doubaoResolutionFromSize(req.Resolution)
	if resolution == "" {
		resolution = strings.TrimSpace(req.Resolution)
	}
	if resolution != "" {
		if resRatio, ok := GetResolutionRatio(billingModel, resolution); ok {
			otherRatios["resolution"] = resRatio
		}
	}

	// 若模型支持有声/无声区分，且明确关闭音频，则应用无声折扣。
	// 不传 generate_audio 时按火山默认有声，不打折。
	if ratio, ok := GetSilentVideoRatio(billingModel); ok {
		if !hasAudioInMetadata(req.Metadata) {
			otherRatios["silent_video"] = ratio
		}
	}

	if len(otherRatios) == 0 {
		return nil
	}
	return otherRatios
}

func normalizedContentHasVideo(items []taskcommon.VideoContentItem) bool {
	for _, item := range items {
		if item.Type == taskcommon.VideoContentVideo && strings.TrimSpace(item.URL) != "" {
			return true
		}
	}
	return false
}

// hasAudioInMetadata 判断这次请求是否按有声视频计费。
// 火山默认 generate_audio=true。字段缺失时按有声；只有明确为 false 时才按无声。
func hasAudioInMetadata(metadata map[string]interface{}) bool {
	if metadata == nil {
		return true
	}
	val, ok := metadata["generate_audio"]
	if !ok {
		return true
	}
	switch v := val.(type) {
	case bool:
		return v
	case map[string]interface{}:
		// dto.BoolValue 序列化后为 {"value": true/false}
		if bv, ok := v["value"].(bool); ok {
			return bv
		}
	}
	return false
}

// durationSpecified 表示调用方显式给了时长。0 是未设置，-1 是火山智能时长。
func durationSpecified(sec int) bool {
	return sec > 0 || sec == -1
}

// BuildRequestBody converts request into Doubao specific format.
func (a *TaskAdaptor) BuildRequestBody(c *gin.Context, info *relaycommon.RelayInfo) (io.Reader, error) {
	if c.GetInt("relay_mode") == relayconstant.RelayMode3DTaskSubmit {
		body, taskErr := build3DRequestBody(c, info)
		if taskErr != nil {
			return nil, taskErr.Error
		}
		return body, nil
	}
	req, err := relaycommon.GetTaskRequest(c)
	if err != nil {
		return nil, err
	}

	body, err := a.convertToRequestPayload(&req)
	if err != nil {
		return nil, errors.Wrap(err, "convert request payload failed")
	}
	if err = rewritePortraitAssets(info.UserId, info.ChannelType, body.Content); err != nil {
		return nil, err
	}
	if info.IsModelMapped {
		body.Model = info.UpstreamModelName
	} else {
		info.UpstreamModelName = body.Model
	}
	data, err := common.Marshal(body)
	if err != nil {
		return nil, err
	}
	url, _ := a.BuildRequestURL(info)
	logger.LogInfo(c, fmt.Sprintf("[doubao-video] submit request: url=%s body=%s", url, string(data)))
	return bytes.NewReader(data), nil
}

// DoRequest delegates to common helper.
func (a *TaskAdaptor) DoRequest(c *gin.Context, info *relaycommon.RelayInfo, requestBody io.Reader) (*http.Response, error) {
	return channel.DoTaskApiRequest(a, c, info, requestBody)
}

// DoResponse handles upstream response, returns taskID etc.
func (a *TaskAdaptor) DoResponse(c *gin.Context, resp *http.Response, info *relaycommon.RelayInfo) (taskID string, taskData []byte, taskErr *dto.TaskError) {
	if c.GetInt("relay_mode") == relayconstant.RelayMode3DTaskSubmit {
		return do3DSubmitResponse(c, resp, info)
	}

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		taskErr = service.TaskErrorWrapper(err, "read_response_body_failed", http.StatusInternalServerError)
		return
	}
	_ = resp.Body.Close()
	logger.LogInfo(c, fmt.Sprintf("[doubao-video] submit response: status=%d body=%s", resp.StatusCode, string(responseBody)))

	// Parse Doubao response
	var dResp responsePayload
	if err := common.Unmarshal(responseBody, &dResp); err != nil {
		taskErr = service.TaskErrorWrapper(errors.Wrapf(err, "body: %s", responseBody), "unmarshal_response_body_failed", http.StatusInternalServerError)
		return
	}

	if dResp.ID == "" {
		taskErr = service.TaskErrorWrapper(fmt.Errorf("task_id is empty"), "invalid_response", http.StatusInternalServerError)
		return
	}

	ov := dto.NewOpenAIVideo()
	ov.ID = info.PublicTaskID
	ov.TaskID = info.PublicTaskID
	ov.CreatedAt = time.Now().Unix()
	ov.Model = info.OriginModelName

	c.JSON(http.StatusOK, ov)
	return dResp.ID, responseBody, nil
}

// FetchTask fetch task status
func (a *TaskAdaptor) FetchTask(baseUrl, key string, body map[string]any, proxy string) (*http.Response, error) {
	taskID, ok := body["task_id"].(string)
	if !ok {
		return nil, fmt.Errorf("invalid task_id")
	}

	uri := fmt.Sprintf("%s/api/v3/contents/generations/tasks/%s", baseUrl, taskID)

	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+key)

	client, err := service.GetHttpClientWithProxy(proxy)
	if err != nil {
		return nil, fmt.Errorf("new proxy http client failed: %w", err)
	}
	return client.Do(req)
}

func (a *TaskAdaptor) GetModelList() []string {
	return ModelList
}

func (a *TaskAdaptor) GetChannelName() string {
	return ChannelName
}

func (a *TaskAdaptor) convertToRequestPayload(req *relaycommon.TaskSubmitReq) (*requestPayload, error) {
	r := requestPayload{
		Model: req.Model,
	}

	metadata := req.Metadata
	if err := taskcommon.UnmarshalMetadata(metadata, &r); err != nil {
		return nil, errors.Wrap(err, "unmarshal metadata failed")
	}
	relaycommon.NormalizeVideoResolutionRatio(req)
	r.Content = doubaoContentFromNormalized(taskcommon.NormalizeVideoContent(*req))

	// 视频只读外层 resolution、ratio。size 与 metadata 里的同名字段已收拢。
	if res := doubaoResolutionFromSize(req.Resolution); res != "" {
		r.Resolution = res
	} else if strings.TrimSpace(req.Resolution) != "" {
		r.Resolution = strings.TrimSpace(req.Resolution)
	}
	if strings.TrimSpace(req.Ratio) != "" {
		r.Ratio = strings.TrimSpace(req.Ratio)
	}

	// duration 优先级：seconds（OpenAI 风格）> duration（标准字段）> metadata（已在上方 UnmarshalMetadata 处理）
	// 显式指定 duration 时清除 frames，避免上游 API 因 frames 优先级更高而忽略 duration。
	// -1 是火山的智能时长，视频编辑必须传 -1，不能当成未设置丢掉。
	if sec, err := strconv.Atoi(req.Seconds); err == nil && durationSpecified(sec) {
		r.Duration = lo.ToPtr(dto.IntValue(sec))
		r.Frames = nil
	} else if durationSpecified(req.Duration) {
		r.Duration = lo.ToPtr(dto.IntValue(req.Duration))
		r.Frames = nil
	}

	// seed：外层字段覆盖 metadata（与 size / duration 一致）
	if req.Seed != nil {
		r.Seed = lo.ToPtr(dto.IntValue(*req.Seed))
	}

	// 文本处理（prompt 必填）：
	// 与 content 中已有 text 做比对（忽略首尾空白）；已存在则不重复追加，否则追加一条。
	// 推荐只传外层 prompt，保持一条提示词；content 中的其他 text 会被保留一并提交。
	prompt := strings.TrimSpace(req.Prompt)
	if prompt != "" {
		already := false
		for _, item := range r.Content {
			if item.Type == "text" && strings.TrimSpace(item.Text) == prompt {
				already = true
				break
			}
		}
		if !already {
			r.Content = append(r.Content, ContentItem{
				Type: "text",
				Text: prompt,
			})
		}
	}

	return &r, nil
}

func doubaoContentFromNormalized(items []taskcommon.VideoContentItem) []ContentItem {
	content := make([]ContentItem, 0, len(items))
	for _, item := range items {
		role := item.Role
		switch item.Type {
		case taskcommon.VideoContentImage:
			if strings.TrimSpace(item.URL) == "" {
				continue
			}
			if role == "" {
				role = taskcommon.VideoRoleReferenceImage
			}
			content = append(content, ContentItem{
				Type:     item.Type,
				Role:     role,
				ImageURL: &MediaURL{URL: item.URL},
			})
		case taskcommon.VideoContentVideo:
			if strings.TrimSpace(item.URL) == "" {
				continue
			}
			if role == "" {
				role = taskcommon.VideoRoleReferenceVideo
			}
			content = append(content, ContentItem{
				Type:     item.Type,
				Role:     role,
				VideoURL: &MediaURL{URL: item.URL},
			})
		case taskcommon.VideoContentAudio:
			if strings.TrimSpace(item.URL) == "" {
				continue
			}
			if role == "" {
				role = taskcommon.VideoRoleReferenceAudio
			}
			content = append(content, ContentItem{
				Type:     item.Type,
				Role:     role,
				AudioURL: &MediaURL{URL: item.URL},
			})
		case taskcommon.VideoContentText:
			if strings.TrimSpace(item.Text) == "" {
				continue
			}
			content = append(content, ContentItem{Type: item.Type, Text: item.Text})
		case taskcommon.VideoContentDraft:
			if item.DraftTaskID == "" {
				continue
			}
			content = append(content, ContentItem{
				Type:      item.Type,
				DraftTask: &DraftTask{ID: item.DraftTaskID},
			})
		}
	}
	return content
}

func (a *TaskAdaptor) ParseTaskResult(respBody []byte) (*relaycommon.TaskInfo, error) {
	common.SysLog(fmt.Sprintf("[doubao] poll response: body=%s", string(respBody)))

	// 通过响应体中的 model 字段判断是否为 3D 任务
	var modelHint struct {
		Model string `json:"model"`
	}
	_ = common.Unmarshal(respBody, &modelHint)
	if is3DModel(modelHint.Model) {
		return parse3DTaskResult(respBody)
	}

	resTask := responseTask{}
	if err := common.Unmarshal(respBody, &resTask); err != nil {
		return nil, errors.Wrap(err, "unmarshal task result failed")
	}

	taskResult := relaycommon.TaskInfo{
		Code: 0,
	}

	// Map Doubao status to internal status
	switch resTask.Status {
	case "pending", "queued":
		taskResult.Status = model.TaskStatusQueued
		taskResult.Progress = "10%"
	case "processing", "running":
		taskResult.Status = model.TaskStatusInProgress
		taskResult.Progress = "50%"
	case "succeeded":
		taskResult.Status = model.TaskStatusSuccess
		taskResult.Progress = "100%"
		taskResult.Url = resTask.Content.VideoURL
		// 解析 usage 信息用于按倍率计费
		taskResult.CompletionTokens = resTask.Usage.CompletionTokens
		taskResult.TotalTokens = resTask.Usage.TotalTokens
	case "failed":
		taskResult.Status = model.TaskStatusFailure
		taskResult.Progress = "100%"
		taskResult.Reason = resTask.Error.Message
	default:
		// Unknown status, treat as processing
		taskResult.Status = model.TaskStatusInProgress
		taskResult.Progress = "30%"
	}

	return &taskResult, nil
}

func (a *TaskAdaptor) ConvertToOpenAIVideo(originTask *model.Task) ([]byte, error) {
	var dResp responseTask
	if err := common.Unmarshal(originTask.Data, &dResp); err != nil {
		return nil, errors.Wrap(err, "unmarshal doubao task data failed")
	}

	openAIVideo := dto.NewOpenAIVideo()
	openAIVideo.ID = originTask.TaskID
	openAIVideo.TaskID = originTask.TaskID
	openAIVideo.Status = originTask.Status.ToVideoStatus()
	openAIVideo.SetProgressStr(originTask.Progress)
	openAIVideo.SetMetadata("url", dResp.Content.VideoURL)
	openAIVideo.CreatedAt = originTask.CreatedAt
	openAIVideo.CompletedAt = originTask.UpdatedAt
	openAIVideo.Model = originTask.Properties.OriginModelName

	if dResp.Status == "failed" {
		openAIVideo.Error = &dto.OpenAIVideoError{
			Message: dResp.Error.Message,
			Code:    dResp.Error.Code,
		}
	}

	return common.Marshal(openAIVideo)
}

// doubaoResolutionFromSize 从 size 字符串推导豆包 resolution 档位。
// 支持两种格式：
//   - "WxH"（如 "1920x1080"）：按短边决定档位
//   - "Np"（如 "1080p"、"720p"）：直接规范化为豆包枚举值
//
// 返回空字符串表示无法识别，保持原有值不变。
func doubaoResolutionFromSize(size string) string {
	size = strings.ToLower(strings.TrimSpace(size))

	// "Np" 格式：直接映射
	switch size {
	case "4k", "2160p":
		return "4k"
	case "1080p":
		return "1080p"
	case "720p":
		return "720p"
	case "480p":
		return "480p"
	}

	// "WxH" 格式
	parts := strings.FieldsFunc(size, func(r rune) bool { return r == 'x' || r == '*' })
	if len(parts) != 2 {
		return ""
	}
	w, errW := strconv.Atoi(strings.TrimSpace(parts[0]))
	h, errH := strconv.Atoi(strings.TrimSpace(parts[1]))
	if errW != nil || errH != nil || w <= 0 || h <= 0 {
		return ""
	}
	short := w
	if h < w {
		short = h
	}
	switch {
	case short >= 2160:
		return "4k"
	case short >= 1080:
		return "1080p"
	case short >= 720:
		return "720p"
	default:
		return "480p"
	}
}

// doubaoRatioFromSize 从 "WxH" 格式推导宽高比，匹配豆包支持的枚举值。
// 仅 WxH 格式可推导，Np 格式返回空字符串（无法确定宽高比）。
// 豆包支持枚举：16:9、9:16、1:1、4:3、3:4、21:9。
func doubaoRatioFromSize(size string) string {
	size = strings.ToLower(strings.TrimSpace(size))
	parts := strings.FieldsFunc(size, func(r rune) bool { return r == 'x' || r == '*' })
	if len(parts) != 2 {
		return ""
	}
	w, errW := strconv.Atoi(strings.TrimSpace(parts[0]))
	h, errH := strconv.Atoi(strings.TrimSpace(parts[1]))
	if errW != nil || errH != nil || w <= 0 || h <= 0 {
		return ""
	}

	// 计算浮点比值，就近匹配豆包支持的宽高比枚举
	ratio := float64(w) / float64(h)
	type candidate struct {
		ratio float64
		label string
	}
	candidates := []candidate{
		{16.0 / 9.0, "16:9"},
		{9.0 / 16.0, "9:16"},
		{1.0, "1:1"},
		{4.0 / 3.0, "4:3"},
		{3.0 / 4.0, "3:4"},
		{21.0 / 9.0, "21:9"},
	}
	best := ""
	bestDiff := 1e9
	for _, c := range candidates {
		diff := ratio - c.ratio
		if diff < 0 {
			diff = -diff
		}
		if diff < bestDiff {
			bestDiff = diff
			best = c.label
		}
	}
	return best
}

// ============================================================
// BackgroundTaskAdaptor — synchronous image generation support
// Volcengine seedream models use a synchronous API; we wrap
// them as async tasks using a background goroutine.
// ============================================================

// IsBackgroundSubmit returns true for Volcengine image generation models (seedream).
// These models use a synchronous API unlike the seedance video models (async).
func (a *TaskAdaptor) IsBackgroundSubmit(modelName string) bool {
	return strings.Contains(strings.ToLower(modelName), "seedream")
}

// ExecuteBackgroundTask calls the Volcengine image generation API synchronously.
// 请求结构通过 volcengine.ConvertToImageGenRequest 构建，与同步接口（/v1/images/generations）
// 走完全相同的请求格式，返回 dto.ImageResponse（也与同步接口一致）。
// MUST NOT write to gin.Context — called from a background goroutine.
func (a *TaskAdaptor) ExecuteBackgroundTask(info *relaycommon.RelayInfo) (resultURL string, responseBody []byte, err error) {
	req := info.BackgroundTaskReq
	if req == nil {
		return "", nil, fmt.Errorf("background task request not set in relay info")
	}

	// 将 TaskSubmitReq 还原为 dto.ImageRequest，再走与同步接口相同的转换逻辑
	imgReq := taskReqToImageReq(req, info)
	// 调用与 /v1/images/generations 同步接口完全相同的转换函数，确保请求格式一致
	upstream := volcengine.ConvertToImageGenRequest(imgReq)

	data, err := common.Marshal(upstream)
	if err != nil {
		return "", nil, fmt.Errorf("marshal image request failed: %w", err)
	}

	// ── 打印上游请求入参（排查用） ────────────────────────────────────────────
	logger.LogInfo(context.Background(), fmt.Sprintf("[BackgroundTask] upstream request: model=%s url=%s/api/v3/images/generations body=%s",
		upstream.Model, a.baseURL, string(data)))

	apiURL := fmt.Sprintf("%s/api/v3/images/generations", a.baseURL)
	httpReq, err := http.NewRequest(http.MethodPost, apiURL, bytes.NewReader(data))
	if err != nil {
		return "", nil, fmt.Errorf("create http request failed: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Accept", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+a.apiKey)

	var client *http.Client
	if proxy := info.ChannelSetting.Proxy; proxy != "" {
		client, err = service.GetHttpClientWithProxy(proxy)
		if err != nil {
			return "", nil, fmt.Errorf("create proxy http client failed: %w", err)
		}
	} else {
		client = service.GetHttpClient()
	}

	resp, err := client.Do(httpReq)
	if err != nil {
		return "", nil, fmt.Errorf("image generation request failed: %w", err)
	}
	defer resp.Body.Close()

	responseBody, err = io.ReadAll(resp.Body)
	if err != nil {
		return "", nil, fmt.Errorf("read response body failed: %w", err)
	}

	// ── 打印上游原始响应（排查用） ────────────────────────────────────────────
	logger.LogInfo(context.Background(), fmt.Sprintf("[BackgroundTask] upstream response: status=%d body=%s",
		resp.StatusCode, string(responseBody)))

	if resp.StatusCode != http.StatusOK {
		return "", responseBody, fmt.Errorf("image generation API returned status %d: %s", resp.StatusCode, string(responseBody))
	}

	// 解析为与同步接口相同的 dto.ImageResponse 结构
	var imgResp dto.ImageResponse
	if err = common.Unmarshal(responseBody, &imgResp); err != nil {
		return "", responseBody, fmt.Errorf("unmarshal image response failed: %w", err)
	}

	// ── 打印解析结果（排查用） ────────────────────────────────────────────────
	logger.LogInfo(context.Background(), fmt.Sprintf("[BackgroundTask] parsed images count=%d", len(imgResp.Data)))

	if len(imgResp.Data) == 0 {
		return "", responseBody, fmt.Errorf("image generation returned empty data")
	}

	// 取第一张图作为 resultURL（多张图通过 task.Data 存储完整 ImageResponse 返回）
	first := imgResp.Data[0]
	if first.Url != "" {
		resultURL = first.Url
	} else if first.B64Json != "" {
		resultURL = "data:image/png;base64," + first.B64Json
	}

	return resultURL, responseBody, nil
}

// taskReqToImageReq 将 TaskSubmitReq 还原为 dto.ImageRequest，
// 这是 ValidateImageTaskRequest 的逆操作，确保后台异步路径与同步路径使用相同的请求结构。
func taskReqToImageReq(req *relaycommon.TaskSubmitReq, info *relaycommon.RelayInfo) dto.ImageRequest {
	modelName := info.UpstreamModelName
	if modelName == "" {
		modelName = req.Model
	}

	imgReq := dto.ImageRequest{
		Model:          modelName,
		Prompt:         req.Prompt,
		Size:           req.Size,
		ResponseFormat: req.ResponseFormat,
	}

	// N → *uint
	if req.N > 0 {
		n := uint(req.N)
		imgReq.N = &n
	}

	// output_format: string → json.RawMessage
	if req.OutputFormat != "" {
		if b, merr := json.Marshal(req.OutputFormat); merr == nil {
			imgReq.OutputFormat = b
		}
	}

	// image（图生图）：与同步接口对齐 —— 单张传 string，多张传 []string
	// volcengineImageRequest.Image 是 json.RawMessage，两种格式均可接收
	if len(req.Images) == 1 {
		if b, merr := json.Marshal(req.Images[0]); merr == nil {
			imgReq.Image = b
		}
	} else if len(req.Images) > 1 {
		if b, merr := json.Marshal(req.Images); merr == nil {
			imgReq.Image = b
		}
	}

	// Metadata 中的渠道专属参数（guidance_scale、watermark 等）→ ExtraFields
	if len(req.Metadata) > 0 {
		if b, merr := common.Marshal(req.Metadata); merr == nil {
			imgReq.ExtraFields = b
		}
	}

	return imgReq
}
