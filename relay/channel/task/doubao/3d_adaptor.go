package doubao

// ──────────────────────────────────────────────────────────────────────────────
// 火山方舟 3D 任务适配器
// 参考：docs/3D文档/3D模型渠道接入规范.md §4
// ──────────────────────────────────────────────────────────────────────────────

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/QuantumNous/new-api/common"
	"github.com/QuantumNous/new-api/dto"
	"github.com/QuantumNous/new-api/logger"
	"github.com/QuantumNous/new-api/model"
	relaycommon "github.com/QuantumNous/new-api/relay/common"
	"github.com/QuantumNous/new-api/service"
	"github.com/QuantumNous/new-api/setting/operation_setting"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/samber/lo"
)

// ──────────────────────────────────────────────────────────────────────────────
// 上游请求/响应结构（仅 3D 专用）
// ──────────────────────────────────────────────────────────────────────────────

// request3DPayload 火山 3D 提交请求体，与视频共用 Content + Model 结构。
type request3DPayload struct {
	Model   string        `json:"model"`
	Content []ContentItem `json:"content,omitempty"`
	Seed    *int64        `json:"seed,omitempty"`
}

// response3DTask 火山 3D 查询响应体。
type response3DTask struct {
	ID          string `json:"id"`
	Model       string `json:"model"`
	Status      string `json:"status"`
	FileFormat  string `json:"fileformat,omitempty"`  // seed3d 专用顶层字段
	SubdivLevel string `json:"subdivisionlevel,omitempty"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
	Content     struct {
		FileURL string `json:"file_url"`
	} `json:"content"`
	Error struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
	Usage struct {
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}

// ──────────────────────────────────────────────────────────────────────────────
// 解析元数据
// ──────────────────────────────────────────────────────────────────────────────

// parseMetadata 从 map 中解析 3D 元数据（将 interface{} 值 JSON-round-trip 到强类型结构）。
func parseMetadata(meta map[string]interface{}) *dto.OpenAI3DMetadata {
	if meta == nil {
		return &dto.OpenAI3DMetadata{}
	}
	b, _ := json.Marshal(meta)
	var m dto.OpenAI3DMetadata
	_ = json.Unmarshal(b, &m)
	return &m
}

// ──────────────────────────────────────────────────────────────────────────────
// validate3DRequest：3D 请求校验
// ──────────────────────────────────────────────────────────────────────────────

// validate3DRequest 3D 请求已由 relay_task.go 前置解析到 context（Validate3DTaskRequest），
// 直接复用公共校验逻辑（幂等，使用 UnmarshalBodyReusable）。
func validate3DRequest(c *gin.Context, info *relaycommon.RelayInfo) *dto.TaskError {
	return relaycommon.Validate3DTaskRequest(c, info)
}

// ──────────────────────────────────────────────────────────────────────────────
// build3DRequestBody：统一字段 → 火山 3D 请求 JSON
// ──────────────────────────────────────────────────────────────────────────────

// build3DRequestBody 将统一 OpenAI3DRequest 转换为火山 3D 提交请求 JSON。
// 根据模型名称选择不同的字段映射分支（seed3d / hyper3d / hitem3d）。
func build3DRequestBody(c *gin.Context, info *relaycommon.RelayInfo) (io.Reader, *dto.TaskError) {
	req, err := relaycommon.Get3DTaskRequest(c)
	if err != nil {
		return nil, service.TaskErrorWrapperLocal(err, "get_3d_request_failed", http.StatusBadRequest)
	}
	meta := parseMetadata(req.Metadata)

	modelName := info.UpstreamModelName
	if modelName == "" {
		modelName = req.Model
	}

	var payload *request3DPayload
	var buildErr error
	switch modelName {
	case ModelSeed3D:
		payload, buildErr = buildSeed3DPayload(modelName, meta)
	case ModelHyper3D:
		payload, buildErr = buildHyper3DPayload(modelName, meta)
	case ModelHitem3D:
		payload, buildErr = buildHitem3DPayload(modelName, meta)
	default:
		return nil, service.TaskErrorWrapperLocal(
			fmt.Errorf("unsupported 3D model: %s", modelName),
			"unsupported_model", http.StatusBadRequest,
		)
	}
	if buildErr != nil {
		return nil, service.TaskErrorWrapperLocal(buildErr, "build_3d_request_failed", http.StatusInternalServerError)
	}

	// 记录提交时的文件格式到 RelayInfo，controller 会存入 task.PrivateData.ReqFileFormat
	// seed3d 查询响应有 fileformat 字段，无需快照；影眸/数美无该字段，需要快照
	if meta.FileFormat != "" && modelName != ModelSeed3D {
		info.ReqFileFormat = meta.FileFormat
	}

	data, marshalErr := common.Marshal(payload)
	if marshalErr != nil {
		return nil, service.TaskErrorWrapperLocal(marshalErr, "marshal_3d_request_failed", http.StatusInternalServerError)
	}
	logger.LogInfo(c, fmt.Sprintf("[doubao-3d] submit request: model=%s body=%s", modelName, string(data)))
	return bytes.NewReader(data), nil
}

// ── seed3d ────────────────────────────────────────────────────────────────────

func buildSeed3DPayload(modelName string, meta *dto.OpenAI3DMetadata) (*request3DPayload, error) {
	p := &request3DPayload{Model: modelName}
	for _, img := range meta.Images {
		p.Content = append(p.Content, ContentItem{
			Type:     "image_url",
			ImageURL: &MediaURL{URL: img.URL},
		})
	}
	var cmds []string
	if meta.FileFormat != "" {
		cmds = append(cmds, fmt.Sprintf("--fileformat %s", meta.FileFormat))
	}
	if meta.QualityLevel != "" {
		cmds = append(cmds, fmt.Sprintf("--subdivisionlevel %s", meta.QualityLevel))
	}
	if len(cmds) > 0 {
		p.Content = append(p.Content, ContentItem{
			Type: "text",
			Text: strings.Join(cmds, " "),
		})
	}
	return p, nil
}

// ── hyper3d（影眸）────────────────────────────────────────────────────────────

func buildHyper3DPayload(modelName string, meta *dto.OpenAI3DMetadata) (*request3DPayload, error) {
	p := &request3DPayload{Model: modelName}
	if meta.Seed != nil {
		p.Seed = meta.Seed
	}
	for _, img := range meta.Images {
		p.Content = append(p.Content, ContentItem{
			Type:     "image_url",
			ImageURL: &MediaURL{URL: img.URL},
		})
	}
	cmds := buildHyper3DTextCommands(meta)
	var textParts []string
	if meta.Prompt != "" {
		textParts = append(textParts, meta.Prompt)
	}
	textParts = append(textParts, cmds...)
	if len(textParts) > 0 {
		p.Content = append(p.Content, ContentItem{
			Type: "text",
			Text: strings.Join(textParts, " "),
		})
	}
	return p, nil
}

func buildHyper3DTextCommands(meta *dto.OpenAI3DMetadata) []string {
	var cmds []string
	if meta.Material != "" {
		cmds = append(cmds, fmt.Sprintf("--material %s", capitalizeFirst(meta.Material)))
	}
	if meta.MeshMode != "" {
		cmds = append(cmds, fmt.Sprintf("--mesh_mode %s", capitalizeFirst(meta.MeshMode)))
	}
	// face_count 优先于 quality_level
	if meta.FaceCount != nil {
		cmds = append(cmds, fmt.Sprintf("--quality_override %d", *meta.FaceCount))
	} else if meta.QualityLevel != "" {
		cmds = append(cmds, fmt.Sprintf("--subdivisionlevel %s", meta.QualityLevel))
	}
	if meta.FileFormat != "" {
		cmds = append(cmds, fmt.Sprintf("--fileformat %s", meta.FileFormat))
	}
	if meta.HdTexture != nil {
		cmds = append(cmds, fmt.Sprintf("--hd_texture %v", *meta.HdTexture))
	}
	if meta.Addons != "" {
		addonVal := meta.Addons
		if strings.EqualFold(addonVal, "high_pack") {
			addonVal = "HighPack"
		}
		cmds = append(cmds, fmt.Sprintf("--addons %s", addonVal))
	}
	if meta.UseOriginalAlpha != nil {
		cmds = append(cmds, fmt.Sprintf("--use_original_alpha %v", *meta.UseOriginalAlpha))
	}
	if meta.TaPose != nil {
		cmds = append(cmds, fmt.Sprintf("--TAPose %v", *meta.TaPose))
	}
	if len(meta.BboxCondition) == 3 {
		cmds = append(cmds, fmt.Sprintf("--bbox_condition [%d,%d,%d]",
			meta.BboxCondition[0], meta.BboxCondition[1], meta.BboxCondition[2]))
	}
	return cmds
}

// ── hitem3d（数美）────────────────────────────────────────────────────────────

func buildHitem3DPayload(modelName string, meta *dto.OpenAI3DMetadata) (*request3DPayload, error) {
	p := &request3DPayload{Model: modelName}

	hasView := lo.SomeBy(meta.Images, func(img dto.Image3D) bool { return img.View != "" })
	if hasView {
		// 按 front/back/left/right 顺序构造位图，并重排图片
		viewOrder := []string{"front", "back", "left", "right"}
		viewMap := make(map[string]dto.Image3D)
		for _, img := range meta.Images {
			if _, ok := viewBitIndex[img.View]; ok {
				viewMap[img.View] = img
			}
		}
		bitStr := make([]byte, 4)
		var orderedImages []dto.Image3D
		for i, v := range viewOrder {
			if img, ok := viewMap[v]; ok {
				bitStr[i] = '1'
				orderedImages = append(orderedImages, img)
			} else {
				bitStr[i] = '0'
			}
		}
		for _, img := range orderedImages {
			p.Content = append(p.Content, ContentItem{
				Type:     "image_url",
				ImageURL: &MediaURL{URL: img.URL},
			})
		}
		cmds := []string{fmt.Sprintf("--multi_images_bit %s", string(bitStr))}
		cmds = append(cmds, buildHitem3DTextCommands(meta)...)
		p.Content = append(p.Content, ContentItem{Type: "text", Text: strings.Join(cmds, " ")})
	} else {
		// 无 view：按原数组顺序写入，不传 --multi_images_bit
		for _, img := range meta.Images {
			p.Content = append(p.Content, ContentItem{
				Type:     "image_url",
				ImageURL: &MediaURL{URL: img.URL},
			})
		}
		cmds := buildHitem3DTextCommands(meta)
		if len(cmds) > 0 {
			p.Content = append(p.Content, ContentItem{Type: "text", Text: strings.Join(cmds, " ")})
		}
	}
	return p, nil
}

func buildHitem3DTextCommands(meta *dto.OpenAI3DMetadata) []string {
	var cmds []string
	if meta.FileFormat != "" {
		if fmtInt, ok := fileFormatIntMap[strings.ToLower(meta.FileFormat)]; ok {
			cmds = append(cmds, fmt.Sprintf("--fileformat %d", fmtInt))
		} else {
			cmds = append(cmds, fmt.Sprintf("--fileformat %s", meta.FileFormat))
		}
	}
	if meta.FaceCount != nil {
		cmds = append(cmds, fmt.Sprintf("--face %d", *meta.FaceCount))
	}
	if meta.GeometryOnly != nil {
		if *meta.GeometryOnly {
			cmds = append(cmds, "--request_type 1")
		} else {
			cmds = append(cmds, "--request_type 3")
		}
	}
	if meta.Resolution != "" {
		cmds = append(cmds, fmt.Sprintf("--resolution %s", meta.Resolution))
	}
	return cmds
}

// ──────────────────────────────────────────────────────────────────────────────
// parse3DTaskResult：火山查询响应 → TaskInfo（供 ParseTaskResult 分支调用）
// ──────────────────────────────────────────────────────────────────────────────

func parse3DTaskResult(respBody []byte) (*relaycommon.TaskInfo, error) {
	var r response3DTask
	if err := common.Unmarshal(respBody, &r); err != nil {
		return nil, errors.Wrap(err, "unmarshal 3d task result failed")
	}

	info := &relaycommon.TaskInfo{}
	switch r.Status {
	case "queued", "pending":
		info.Status = model.TaskStatusQueued
		info.Progress = "10%"
	case "running", "processing":
		info.Status = model.TaskStatusInProgress
		info.Progress = "50%"
	case "succeeded":
		info.Status = model.TaskStatusSuccess
		info.Progress = "100%"
		info.Url = r.Content.FileURL
		info.CompletionTokens = r.Usage.CompletionTokens
		info.TotalTokens = r.Usage.TotalTokens
	case "failed":
		info.Status = model.TaskStatusFailure
		info.Progress = "100%"
		info.Reason = r.Error.Message
	case "cancelled":
		info.Status = model.TaskStatusUnknown
		info.Progress = "100%"
	default:
		info.Status = model.TaskStatusInProgress
		info.Progress = "30%"
	}
	return info, nil
}

// ──────────────────────────────────────────────────────────────────────────────
// do3DSubmitResponse：处理 3D 提交成功响应
// ──────────────────────────────────────────────────────────────────────────────

// do3DSubmitResponse 解析火山 3D 提交响应，返回上游 task_id。
// 同时向客户端写入统一 3D 任务提交响应（task_id, status=QUEUED）。
func do3DSubmitResponse(c *gin.Context, resp *http.Response, info *relaycommon.RelayInfo) (taskID string, taskData []byte, taskErr *dto.TaskError) {
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		taskErr = service.TaskErrorWrapper(err, "read_response_body_failed", http.StatusInternalServerError)
		return
	}
	_ = resp.Body.Close()
	logger.LogInfo(c, fmt.Sprintf("[doubao-3d] submit response: status=%d body=%s", resp.StatusCode, string(responseBody)))

	var result struct {
		ID string `json:"id"`
	}
	if err := common.Unmarshal(responseBody, &result); err != nil {
		taskErr = service.TaskErrorWrapper(errors.Wrapf(err, "body: %s", responseBody), "unmarshal_response_failed", http.StatusInternalServerError)
		return
	}
	if result.ID == "" {
		taskErr = service.TaskErrorWrapper(fmt.Errorf("task_id is empty, body: %s", responseBody), "invalid_response", http.StatusInternalServerError)
		return
	}

	submitResp := map[string]interface{}{
		"task_id": info.PublicTaskID,
		"status":  "QUEUED",
		"object":  "3d.generation",
		"model":   info.OriginModelName,
	}
	respBytes, _ := common.Marshal(submitResp)
	c.Data(http.StatusOK, "application/json", respBytes)
	return result.ID, responseBody, nil
}

// ──────────────────────────────────────────────────────────────────────────────
// ConvertToOpenAI3DTask：model.Task → 统一 OpenAI3DTask 查询响应 JSON
// ──────────────────────────────────────────────────────────────────────────────

// ConvertToOpenAI3DTask 将平台内部 Task 转换为统一的 OpenAI3DTask 查询响应。
func (a *TaskAdaptor) ConvertToOpenAI3DTask(originTask *model.Task) ([]byte, error) {
	var r response3DTask
	if err := common.Unmarshal(originTask.Data, &r); err != nil {
		return nil, errors.Wrap(err, "unmarshal doubao 3d task data failed")
	}

	usd := float64(originTask.Quota) / common.QuotaPerUnit
	cny := usd * operation_setting.USDExchangeRate
	amount := fmt.Sprintf("%.6f", cny)

	result := &dto.OpenAI3DTask{
		TaskID:     originTask.TaskID,
		Object:     "3d.generation",
		Model:      originTask.Properties.OriginModelName,
		Status:     string(originTask.Status),
		Progress:   originTask.Progress,
		FailReason: originTask.FailReason,
		Amount:     amount,
		CreatedAt:  r.CreatedAt,
		SubmitTime: originTask.SubmitTime,
		Files:      []dto.Task3DFile{},
	}
	if originTask.FinishTime > 0 {
		result.FinishTime = originTask.FinishTime
	}

	if r.Content.FileURL != "" {
		// seed3d 查询响应有 fileformat 字段；影眸/数美 用提交快照 ReqFileFormat
		fileFormat := r.FileFormat
		if fileFormat == "" {
			fileFormat = originTask.PrivateData.ReqFileFormat
		}
		if fileFormat == "" {
			fileFormat = "zip"
		}
		result.Files = append(result.Files, dto.Task3DFile{
			URL:    r.Content.FileURL,
			Format: strings.ToLower(fileFormat),
		})
	}

	result.Usage = &dto.Task3DUsage{
		OutputTokens: r.Usage.CompletionTokens,
		TotalTokens:  r.Usage.TotalTokens,
	}

	return common.Marshal(result)
}

// ──────────────────────────────────────────────────────────────────────────────
// 工具函数
// ──────────────────────────────────────────────────────────────────────────────

// capitalizeFirst 将字符串首字母大写（用于 material/mesh_mode 等枚举转换）。
func capitalizeFirst(s string) string {
	if s == "" {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}
