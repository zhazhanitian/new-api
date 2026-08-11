package tencentvod

// ──────────────────────────────────────────────────────────────────────────────
// 腾讯混元 3D 任务适配器
// 参考：docs/3D文档/3D模型渠道接入规范.md §5、§6.3
// ──────────────────────────────────────────────────────────────────────────────

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/QuantumNous/new-api/common"
	"github.com/QuantumNous/new-api/constant"
	"github.com/QuantumNous/new-api/dto"
	"github.com/QuantumNous/new-api/logger"
	"github.com/QuantumNous/new-api/model"
	relaycommon "github.com/QuantumNous/new-api/relay/common"
	"github.com/QuantumNous/new-api/service"
	"github.com/QuantumNous/new-api/setting/operation_setting"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

// ──────────────────────────────────────────────────────────────────────────────
// 腾讯 AI3D 请求结构体
// ──────────────────────────────────────────────────────────────────────────────

// ai3DViewImage 对应 Tencent ViewImage（多视角图）
type ai3DViewImage struct {
	ViewType       string `json:"ViewType,omitempty"`
	ViewImageUrl   string `json:"ViewImageUrl,omitempty"`
	ViewImageBase64 string `json:"ViewImageBase64,omitempty"`
}

// ai3DFile3D 对应 Tencent File3D（带预览图的 3D 文件）
type ai3DFile3D struct {
	Type            string `json:"Type"`
	Url             string `json:"Url"`
	PreviewImageUrl string `json:"PreviewImageUrl,omitempty"`
}

// ai3DInputFile3D 对应 Tencent InputFile3D（绑骨蒙皮/文生动作输入文件）
type ai3DInputFile3D struct {
	Url  string `json:"Url"`
	Type string `json:"Type"`
}

// ai3DImageRef 对应 Tencent Image 结构（Base64 + Url）
type ai3DImageRef struct {
	Base64 string `json:"Base64,omitempty"`
	Url    string `json:"Url,omitempty"`
}

// ── 极速版提交请求 ────────────────────────────────────────────────────────────
type ai3DRapidSubmitReq struct {
	Prompt         string  `json:"Prompt,omitempty"`
	ImageBase64    string  `json:"ImageBase64,omitempty"`
	ImageUrl       string  `json:"ImageUrl,omitempty"`
	ResultFormat   string  `json:"ResultFormat,omitempty"`
	EnablePBR      *bool   `json:"EnablePBR,omitempty"`
	EnableGeometry *bool   `json:"EnableGeometry,omitempty"`
}

// ── 专业版提交请求 ────────────────────────────────────────────────────────────
type ai3DProSubmitReq struct {
	Model           string          `json:"Model,omitempty"`
	Prompt          string          `json:"Prompt,omitempty"`
	ImageBase64     string          `json:"ImageBase64,omitempty"`
	ImageUrl        string          `json:"ImageUrl,omitempty"`
	MultiViewImages []ai3DViewImage `json:"MultiViewImages,omitempty"`
	EnablePBR       *bool           `json:"EnablePBR,omitempty"`
	FaceCount       *int            `json:"FaceCount,omitempty"`
	GenerateType    string          `json:"GenerateType,omitempty"`
	PolygonType     string          `json:"PolygonType,omitempty"`
	ResultFormat    string          `json:"ResultFormat,omitempty"`
}

// ── 智能拓扑提交请求 ─────────────────────────────────────────────────────────
type ai3DReduceFaceSubmitReq struct {
	File3D      ai3DFile3D `json:"File3D"`
	PolygonType string     `json:"PolygonType,omitempty"`
	FaceLevel   string     `json:"FaceLevel,omitempty"`
}

// ── 纹理生成提交请求 ─────────────────────────────────────────────────────────
type ai3DTextureSubmitReq struct {
	File3D          ai3DFile3D      `json:"File3D"`
	Model           string          `json:"Model,omitempty"`
	MultiViewImages []ai3DViewImage `json:"MultiViewImages,omitempty"`
	Prompt          string          `json:"Prompt,omitempty"`
	Image           *ai3DImageRef   `json:"Image,omitempty"`
	EnablePBR       *bool           `json:"EnablePBR,omitempty"`
	EnableKeepUV    *bool           `json:"EnableKeepUV,omitempty"`
	TextureSize     *int            `json:"TextureSize,omitempty"`
}

// ── 3D 人物生成提交请求 ───────────────────────────────────────────────────────
type ai3DProfileSubmitReq struct {
	Profile  *ai3DImageRef `json:"Profile,omitempty"`
	Template string        `json:"Template,omitempty"`
}

// ── 绑骨蒙皮提交请求 ─────────────────────────────────────────────────────────
type ai3DAutoRiggingSubmitReq struct {
	File3D     ai3DInputFile3D `json:"File3D"`
	MotionType *int            `json:"MotionType,omitempty"`
}

// ── 文生动作提交请求 ─────────────────────────────────────────────────────────
type ai3DMotionSubmitReq struct {
	Prompt           string           `json:"Prompt"`
	Model            string           `json:"Model,omitempty"`
	RetargetFile     *ai3DInputFile3D `json:"RetargetFile,omitempty"`
	Duration         *int             `json:"Duration,omitempty"`
	EnableMesh       *bool            `json:"EnableMesh,omitempty"`
	EnableRewrite    *bool            `json:"EnableRewrite,omitempty"`
	EnableDurationEst *bool           `json:"EnableDurationEst,omitempty"`
}

// ── 提交响应（所有 AI3D 模型统一格式）──────────────────────────────────────
type ai3DSubmitResponse struct {
	Response struct {
		JobId     string `json:"JobId"`
		RequestId string `json:"RequestId"`
		Error     *struct {
			Code    string `json:"Code"`
			Message string `json:"Message"`
		} `json:"Error,omitempty"`
	} `json:"Response"`
}

// ── 查询请求 ─────────────────────────────────────────────────────────────────
type ai3DQueryReq struct {
	JobId string `json:"JobId"`
}

// ── 查询响应（通用基础字段）──────────────────────────────────────────────────
type ai3DQueryResponse struct {
	Response struct {
		Status       string       `json:"Status"`
		ErrorCode    string       `json:"ErrorCode"`
		ErrorMessage string       `json:"ErrorMessage"`
		// ResultFile3Ds 所有模型均有此字段
		ResultFile3Ds []ai3DFile3D `json:"ResultFile3Ds"`
		// 专业版额外字段
		ResultCreditConsumed float64 `json:"ResultCreditConsumed,omitempty"`
		ResultCreditDetails  string  `json:"ResultCreditDetails,omitempty"`
		RequestId            string  `json:"RequestId"`
	} `json:"Response"`
}

// ──────────────────────────────────────────────────────────────────────────────
// validate3DRequest
// ──────────────────────────────────────────────────────────────────────────────

// validate3DRequest 3D 请求已由 relay_task.go 前置解析到 context，复用公共校验逻辑。
func validate3DRequest(c *gin.Context, info *relaycommon.RelayInfo) *dto.TaskError {
	return relaycommon.Validate3DTaskRequest(c, info)
}

// ──────────────────────────────────────────────────────────────────────────────
// buildAI3DRequestBody：统一字段 → 腾讯 AI3D 请求 JSON
// ──────────────────────────────────────────────────────────────────────────────

func buildAI3DRequestBody(c *gin.Context, info *relaycommon.RelayInfo) (io.Reader, *dto.TaskError) {
	req, err := relaycommon.Get3DTaskRequest(c)
	if err != nil {
		return nil, service.TaskErrorWrapperLocal(err, "get_3d_request_failed", http.StatusBadRequest)
	}

	modelName := info.UpstreamModelName
	if modelName == "" {
		modelName = req.Model
	}

	meta, ok := getAI3DModelMeta(modelName)
	if !ok {
		return nil, service.TaskErrorWrapperLocal(
			fmt.Errorf("unsupported AI3D model: %s", modelName),
			"unsupported_model", http.StatusBadRequest,
		)
	}

	// 解析统一元数据
	reqMeta := parseAI3DMetadata(req.Metadata)

	var payload interface{}
	switch modelName {
	case "hunyuan-3d-rapid":
		payload = buildRapidPayload(req, reqMeta)
	case "hunyuan-3d-pro-3.0", "hunyuan-3d-pro-3.1":
		payload = buildProPayload(meta, req, reqMeta, info)
	case "hunyuan-3d-reduce-face":
		payload = buildReduceFacePayload(req, reqMeta)
	case "hunyuan-3d-texture-3.0", "hunyuan-3d-texture-3.1":
		payload = buildTexturePayload(meta, req, reqMeta)
	case "hunyuan-3d-profile":
		payload = buildProfilePayload(req, reqMeta)
	case "hunyuan-3d-auto-rigging":
		payload = buildAutoRiggingPayload(req, reqMeta)
	case "hunyuan-3d-motion":
		payload = buildMotionPayload(meta, req, reqMeta)
	default:
		return nil, service.TaskErrorWrapperLocal(
			fmt.Errorf("unsupported AI3D model: %s", modelName),
			"unsupported_model", http.StatusBadRequest,
		)
	}

	// 计算预估积分并存入 RelayInfo
	estimatedCredits := estimateAI3DCredits(modelName, req, reqMeta)
	info.EstimatedCredits = float64(estimatedCredits)
	info.QuotaPerUnit = common.QuotaPerUnit
	if meta.IsProVersion {
		info.AllowCompleteAdjustment = true
	}

	data, marshalErr := common.Marshal(payload)
	if marshalErr != nil {
		return nil, service.TaskErrorWrapperLocal(marshalErr, "marshal_ai3d_request_failed", http.StatusInternalServerError)
	}
	logger.LogInfo(c, fmt.Sprintf("[tencent-ai3d] submit request: model=%s action=%s body=%s", modelName, meta.SubmitAction, string(data)))
	return bytes.NewReader(data), nil
}

// ── 各模型请求体构造函数 ──────────────────────────────────────────────────────

func buildRapidPayload(req dto.OpenAI3DRequest, meta *dto.OpenAI3DMetadata) *ai3DRapidSubmitReq {
	p := &ai3DRapidSubmitReq{
		Prompt:    meta.Prompt,
		EnablePBR: meta.EnablePBR,
	}
	// 主图（第一张）
	if len(meta.Images) > 0 {
		img := meta.Images[0]
		if strings.HasPrefix(img.URL, "data:") {
			p.ImageBase64 = extractBase64(img.URL)
		} else {
			p.ImageUrl = img.URL
		}
	}
	// 文件格式
	if meta.FileFormat != "" {
		p.ResultFormat = strings.ToUpper(meta.FileFormat)
	}
	// geometry_only
	if meta.GeometryOnly != nil && *meta.GeometryOnly {
		t := true
		p.EnableGeometry = &t
	}
	_ = req
	return p
}

func buildProPayload(entry ai3DModelMetaEntry, req dto.OpenAI3DRequest, meta *dto.OpenAI3DMetadata, info *relaycommon.RelayInfo) *ai3DProSubmitReq {
	p := &ai3DProSubmitReq{
		Prompt:    meta.Prompt,
		EnablePBR: meta.EnablePBR,
	}
	// Model version (3.0 / 3.1)
	if entry.ModelVersion != "" {
		p.Model = entry.ModelVersion
	}
	// 主图（第一张，无 view 或 view == front）
	mainImgs, multiViewImgs := splitImagesForPro(meta.Images)
	if len(mainImgs) > 0 {
		if strings.HasPrefix(mainImgs[0].URL, "data:") {
			p.ImageBase64 = extractBase64(mainImgs[0].URL)
		} else {
			p.ImageUrl = mainImgs[0].URL
		}
	}
	// 多视角图
	for _, img := range multiViewImgs {
		vi := ai3DViewImage{ViewType: img.View}
		if strings.HasPrefix(img.URL, "data:") {
			vi.ViewImageBase64 = extractBase64(img.URL)
		} else {
			vi.ViewImageUrl = img.URL
		}
		p.MultiViewImages = append(p.MultiViewImages, vi)
	}
	// generate_mode → GenerateType; geometry_only 优先
	if meta.GeometryOnly != nil && *meta.GeometryOnly {
		p.GenerateType = "Geometry"
	} else if meta.GenerateMode != "" {
		p.GenerateType = generateModeToTencent(meta.GenerateMode)
	}
	// polygon_type（仅 LowPoly 模式有效，透传不过滤）
	if meta.PolygonType != "" {
		p.PolygonType = tencentPolygonType(meta.PolygonType)
	}
	// file_format → ResultFormat（仅用户显式传才发送）
	if meta.FileFormat != "" {
		p.ResultFormat = strings.ToUpper(meta.FileFormat)
	}
	// face_count（仅 Normal/Sketch 下有效，透传不过滤）
	if meta.FaceCount != nil {
		p.FaceCount = meta.FaceCount
	}
	_ = req
	_ = info
	return p
}

func buildReduceFacePayload(req dto.OpenAI3DRequest, meta *dto.OpenAI3DMetadata) *ai3DReduceFaceSubmitReq {
	p := &ai3DReduceFaceSubmitReq{
		PolygonType: tencentPolygonType(meta.PolygonType),
		FaceLevel:   meta.FaceLevel,
	}
	if meta.InputFile != nil {
		p.File3D = ai3DFile3D{
			Type: strings.ToUpper(meta.InputFile.Type),
			Url:  meta.InputFile.URL,
		}
	}
	_ = req
	return p
}

func buildTexturePayload(entry ai3DModelMetaEntry, req dto.OpenAI3DRequest, meta *dto.OpenAI3DMetadata) *ai3DTextureSubmitReq {
	p := &ai3DTextureSubmitReq{
		Prompt:       meta.Prompt,
		EnablePBR:    meta.EnablePBR,
		EnableKeepUV: meta.EnableKeepUV,
		TextureSize:  meta.TextureSize,
	}
	if entry.ModelVersion != "" {
		p.Model = entry.ModelVersion
	}
	if meta.InputFile != nil {
		p.File3D = ai3DFile3D{
			Type: strings.ToUpper(meta.InputFile.Type),
			Url:  meta.InputFile.URL,
		}
	}
	// 参考图（images[0] 作为纹理参考图，无 view 的第一张）
	if len(meta.Images) > 0 {
		img := meta.Images[0]
		ref := &ai3DImageRef{}
		if strings.HasPrefix(img.URL, "data:") {
			ref.Base64 = extractBase64(img.URL)
		} else {
			ref.Url = img.URL
		}
		p.Image = ref
	}
	// 多视角图（3.1 版本支持）
	for _, img := range meta.Images {
		if img.View == "" {
			continue
		}
		vi := ai3DViewImage{ViewType: img.View}
		if strings.HasPrefix(img.URL, "data:") {
			vi.ViewImageBase64 = extractBase64(img.URL)
		} else {
			vi.ViewImageUrl = img.URL
		}
		p.MultiViewImages = append(p.MultiViewImages, vi)
	}
	_ = req
	return p
}

func buildProfilePayload(req dto.OpenAI3DRequest, meta *dto.OpenAI3DMetadata) *ai3DProfileSubmitReq {
	p := &ai3DProfileSubmitReq{Template: meta.Template}
	if len(meta.Images) > 0 {
		img := meta.Images[0]
		ref := &ai3DImageRef{}
		if strings.HasPrefix(img.URL, "data:") {
			ref.Base64 = extractBase64(img.URL)
		} else {
			ref.Url = img.URL
		}
		p.Profile = ref
	}
	_ = req
	return p
}

func buildAutoRiggingPayload(req dto.OpenAI3DRequest, meta *dto.OpenAI3DMetadata) *ai3DAutoRiggingSubmitReq {
	p := &ai3DAutoRiggingSubmitReq{MotionType: meta.MotionType}
	if meta.InputFile != nil {
		p.File3D = ai3DInputFile3D{
			Url:  meta.InputFile.URL,
			Type: strings.ToUpper(meta.InputFile.Type),
		}
	}
	_ = req
	return p
}

func buildMotionPayload(entry ai3DModelMetaEntry, req dto.OpenAI3DRequest, meta *dto.OpenAI3DMetadata) *ai3DMotionSubmitReq {
	p := &ai3DMotionSubmitReq{
		Prompt:            meta.Prompt,
		Duration:          meta.Duration,
		EnableMesh:        meta.EnableMesh,
		EnableRewrite:     meta.EnableRewrite,
		EnableDurationEst: meta.EnableDurationEst,
	}
	if entry.ModelVersion != "" {
		p.Model = entry.ModelVersion
	}
	if meta.InputFile != nil {
		p.RetargetFile = &ai3DInputFile3D{
			Url:  meta.InputFile.URL,
			Type: strings.ToUpper(meta.InputFile.Type),
		}
	}
	_ = req
	return p
}

// ──────────────────────────────────────────────────────────────────────────────
// doAI3DSubmitResponse：处理 AI3D 提交响应
// ──────────────────────────────────────────────────────────────────────────────

func doAI3DSubmitResponse(c *gin.Context, resp *http.Response, info *relaycommon.RelayInfo) (taskID string, taskData []byte, taskErr *dto.TaskError) {
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		taskErr = service.TaskErrorWrapper(err, "read_response_body_failed", http.StatusInternalServerError)
		return
	}
	_ = resp.Body.Close()
	logger.LogInfo(c, fmt.Sprintf("[tencent-ai3d] submit response: status=%d body=%s", resp.StatusCode, string(responseBody)))

	var submitResp ai3DSubmitResponse
	if err := common.Unmarshal(responseBody, &submitResp); err != nil {
		taskErr = service.TaskErrorWrapper(errors.Wrapf(err, "body: %s", responseBody), "unmarshal_response_failed", http.StatusInternalServerError)
		return
	}
	if submitResp.Response.Error != nil {
		taskErr = service.TaskErrorWrapper(
			fmt.Errorf("%s: %s", submitResp.Response.Error.Code, submitResp.Response.Error.Message),
			submitResp.Response.Error.Code, http.StatusInternalServerError,
		)
		return
	}
	if submitResp.Response.JobId == "" {
		taskErr = service.TaskErrorWrapper(
			fmt.Errorf("JobId is empty, body: %s", responseBody),
			"invalid_response", http.StatusInternalServerError,
		)
		return
	}

	respOut := map[string]interface{}{
		"task_id": info.PublicTaskID,
		"status":  "QUEUED",
		"object":  "3d.generation",
		"model":   info.OriginModelName,
	}
	respBytes, _ := common.Marshal(respOut)
	c.Data(http.StatusOK, "application/json", respBytes)
	return submitResp.Response.JobId, responseBody, nil
}

// ──────────────────────────────────────────────────────────────────────────────
// fetchAI3DTask：轮询 AI3D 任务状态
// ──────────────────────────────────────────────────────────────────────────────

// fetchAI3DTask 调用对应模型的查询接口，轮询任务状态。
// body 中必须有 "task_id"（JobId）和 "model"（原始模型名）。
func fetchAI3DTask(key string, body map[string]any, proxy string) (*http.Response, error) {
	jobID, _ := body["task_id"].(string)
	modelName, _ := body["model"].(string)

	meta, ok := getAI3DModelMeta(modelName)
	if !ok {
		return nil, fmt.Errorf("unknown AI3D model: %s", modelName)
	}

	parts := strings.Split(key, "|")
	var secretId, secretKey string
	if len(parts) == 3 {
		secretId = strings.TrimSpace(parts[1])
		secretKey = strings.TrimSpace(parts[2])
	} else if len(parts) == 2 {
		secretId = strings.TrimSpace(parts[0])
		secretKey = strings.TrimSpace(parts[1])
	}

	queryReq := ai3DQueryReq{JobId: jobID}
	payload, err := common.Marshal(queryReq)
	if err != nil {
		return nil, errors.Wrap(err, "marshal AI3D query payload failed")
	}

	ts := time.Now().Unix()
	auth := buildAuthorization(secretId, secretKey, meta.QueryAction, string(payload), ts, tencentAI3DHost, tencentAI3DService)

	endpoint := "https://" + tencentAI3DHost
	req, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Host", tencentAI3DHost)
	req.Header.Set("X-TC-Action", meta.QueryAction)
	req.Header.Set("X-TC-Version", tencentAI3DVersion)
	req.Header.Set("X-TC-Timestamp", fmt.Sprintf("%d", ts))
	req.Header.Set("X-TC-Region", tencentAI3DRegion)
	req.Header.Set("Authorization", auth)

	common.SysLog(fmt.Sprintf("[tencent-ai3d] query request: action=%s jobId=%s", meta.QueryAction, jobID))

	client, err := service.GetHttpClientWithProxy(proxy)
	if err != nil {
		return nil, fmt.Errorf("get http client failed: %w", err)
	}
	return client.Do(req)
}

// ──────────────────────────────────────────────────────────────────────────────
// parseAI3DTaskResult：AI3D 查询响应 → TaskInfo
// ──────────────────────────────────────────────────────────────────────────────

func parseAI3DTaskResult(respBody []byte) (*relaycommon.TaskInfo, error) {
	common.SysLog(fmt.Sprintf("[tencent-ai3d] query response: %s", string(respBody)))

	var queryResp ai3DQueryResponse
	if err := common.Unmarshal(respBody, &queryResp); err != nil {
		return nil, errors.Wrap(err, "unmarshal AI3D query response failed")
	}

	result := &relaycommon.TaskInfo{}

	switch queryResp.Response.Status {
	case "WAIT":
		result.Status = model.TaskStatusQueued
		result.Progress = "10%"
	case "RUN":
		result.Status = model.TaskStatusInProgress
		result.Progress = "50%"
	case "DONE":
		if queryResp.Response.ErrorCode != "" {
			result.Status = model.TaskStatusFailure
			result.Progress = "100%"
			result.Reason = queryResp.Response.ErrorMessage
			if result.Reason == "" {
				result.Reason = queryResp.Response.ErrorCode
			}
		} else {
			result.Status = model.TaskStatusSuccess
			result.Progress = "100%"
			if len(queryResp.Response.ResultFile3Ds) > 0 {
				result.Url = queryResp.Response.ResultFile3Ds[0].Url
			}
		}
	case "FAIL":
		result.Status = model.TaskStatusFailure
		result.Progress = "100%"
		result.Reason = queryResp.Response.ErrorMessage
		if result.Reason == "" {
			result.Reason = queryResp.Response.ErrorCode
		}
	default:
		result.Status = model.TaskStatusQueued
		result.Progress = "10%"
	}
	return result, nil
}

// ──────────────────────────────────────────────────────────────────────────────
// estimateAI3DCredits：提交前预估积分
// 参考：docs/3D文档/3D模型渠道接入规范.md §6.3
// ──────────────────────────────────────────────────────────────────────────────

func estimateAI3DCredits(modelName string, req dto.OpenAI3DRequest, meta *dto.OpenAI3DMetadata) int {
	switch modelName {
	case "hunyuan-3d-rapid":
		credits := creditsRapidBase
		if meta.EnablePBR != nil && *meta.EnablePBR {
			credits += creditsRapidPBR
		}
		return credits

	case "hunyuan-3d-pro-3.0", "hunyuan-3d-pro-3.1":
		credits := creditsProNormal
		// 基础积分由 generate_mode/geometry_only 决定
		if meta.GeometryOnly != nil && *meta.GeometryOnly {
			credits = creditsProGeometry
		} else {
			switch meta.GenerateMode {
			case "low_poly":
				credits = creditsProLowPoly
			case "sketch":
				credits = creditsProSketch
			default:
				credits = creditsProNormal
			}
		}
		// 多视角图
		if hasMultiViewImages(meta.Images) {
			credits += creditsProMultiView
		}
		// PBR
		if meta.EnablePBR != nil && *meta.EnablePBR {
			credits += creditsProPBR
		}
		// face_count（保守预扣，专业版最终按实际积分校准）
		if meta.FaceCount != nil {
			credits += creditsProFaceCount
		}
		// file_format（用户显式传才增积分）
		if meta.FileFormat != "" {
			credits += creditsProFileFormat
		}
		return credits

	case "hunyuan-3d-reduce-face":
		return creditsReduceFace
	case "hunyuan-3d-texture-3.0", "hunyuan-3d-texture-3.1":
		return creditsTexture
	case "hunyuan-3d-profile":
		return creditsProfile
	case "hunyuan-3d-auto-rigging":
		return creditsAutoRigging
	case "hunyuan-3d-motion":
		return creditsMotion
	}
	return 1
}

// hasMultiViewImages 判断 images 数组中是否有带 view 字段的多视角图。
func hasMultiViewImages(images []dto.Image3D) bool {
	for _, img := range images {
		if img.View != "" {
			return true
		}
	}
	return false
}

// ──────────────────────────────────────────────────────────────────────────────
// AdjustBillingOnComplete：AI3D 专业版完成时积分校准
// ──────────────────────────────────────────────────────────────────────────────

// adjustAI3DBillingOnComplete 专业版使用查询返回的实际积分重算额度，实现多退少补。
// 非专业版（AllowCompleteAdjustment=false）返回 0，保持预扣额度不变。
func adjustAI3DBillingOnComplete(task *model.Task) int {
	bc := task.PrivateData.BillingContext
	if bc == nil || !bc.AllowCompleteAdjustment {
		return 0
	}

	var queryResp ai3DQueryResponse
	if err := common.Unmarshal(task.Data, &queryResp); err != nil {
		common.SysLog(fmt.Sprintf("[tencent-ai3d] AdjustBillingOnComplete: unmarshal task.Data failed for task %s: %v", task.TaskID, err))
		return 0
	}

	actualCredits := queryResp.Response.ResultCreditConsumed
	if actualCredits <= 0 {
		common.SysLog(fmt.Sprintf("[tencent-ai3d] AdjustBillingOnComplete: no ResultCreditConsumed for task %s, keeping pre-charged quota", task.TaskID))
		return 0
	}

	// 使用提交时快照的 ModelPrice/GroupRatio/QuotaPerUnit，避免受全局配置变更影响
	quotaPerUnit := bc.QuotaPerUnit
	if quotaPerUnit <= 0 {
		quotaPerUnit = common.QuotaPerUnit
	}
	actualQuota := int(actualCredits * bc.ModelPrice * quotaPerUnit * bc.GroupRatio)

	common.SysLog(fmt.Sprintf("[tencent-ai3d] AdjustBillingOnComplete: task=%s estimatedCredits=%.1f actualCredits=%.1f preQuota=%d finalQuota=%d",
		task.TaskID, bc.EstimatedCredits, actualCredits, task.Quota, actualQuota))

	return actualQuota
}

// ──────────────────────────────────────────────────────────────────────────────
// ConvertToOpenAI3DTask：model.Task → 统一 OpenAI3DTask 查询响应 JSON
// ──────────────────────────────────────────────────────────────────────────────

// ConvertToOpenAI3DTask 实现 channel.Task3DConverter 接口。
func (a *TaskAdaptor) ConvertToOpenAI3DTask(originTask *model.Task) ([]byte, error) {
	var queryResp ai3DQueryResponse
	if err := common.Unmarshal(originTask.Data, &queryResp); err != nil {
		// 可能是提交响应（submit 阶段存的是 ai3DSubmitResponse）
		var submitResp ai3DSubmitResponse
		if err2 := common.Unmarshal(originTask.Data, &submitResp); err2 != nil {
			return nil, errors.Wrap(err, "unmarshal tencent ai3d task data failed")
		}
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
		SubmitTime: originTask.SubmitTime,
		Files:      []dto.Task3DFile{},
	}
	if originTask.FinishTime > 0 {
		result.FinishTime = originTask.FinishTime
	}

	// 从查询响应提取文件列表
	for _, f := range queryResp.Response.ResultFile3Ds {
		if f.Url == "" {
			continue
		}
		file := dto.Task3DFile{
			URL:        f.Url,
			Format:     strings.ToLower(f.Type),
			PreviewURL: f.PreviewImageUrl,
		}
		result.Files = append(result.Files, file)
	}

	// usage.credits：专业版优先使用查询响应中的实际积分；其余版本使用提交时预估积分
	var credits float64
	if queryResp.Response.ResultCreditConsumed > 0 {
		credits = queryResp.Response.ResultCreditConsumed
	} else if bc := originTask.PrivateData.BillingContext; bc != nil {
		credits = bc.EstimatedCredits
	}
	result.Usage = &dto.Task3DUsage{Credits: credits}

	return common.Marshal(result)
}

// ──────────────────────────────────────────────────────────────────────────────
// 工具函数
// ──────────────────────────────────────────────────────────────────────────────

// parseAI3DMetadata 从 map[string]interface{} 解析统一 3D 元数据。
func parseAI3DMetadata(meta map[string]interface{}) *dto.OpenAI3DMetadata {
	if meta == nil {
		return &dto.OpenAI3DMetadata{}
	}
	b, _ := common.Marshal(meta)
	var m dto.OpenAI3DMetadata
	_ = common.Unmarshal(b, &m)
	return &m
}

// splitImagesForPro 将 images 数组分为主图（无 view 或 view=="front"）和多视角图。
// 规格：接入规范 §5.2 专业版主图/多视角规则。
func splitImagesForPro(images []dto.Image3D) (main []dto.Image3D, multiView []dto.Image3D) {
	for _, img := range images {
		if img.View == "" || strings.EqualFold(img.View, "front") {
			main = append(main, img)
		} else {
			multiView = append(multiView, img)
		}
	}
	return
}

// generateModeToTencent 将统一 generate_mode 映射为腾讯 GenerateType 值。
func generateModeToTencent(mode string) string {
	switch strings.ToLower(mode) {
	case "low_poly":
		return "LowPoly"
	case "sketch":
		return "Sketch"
	default:
		return "Normal"
	}
}

// tencentPolygonType 将统一 polygon_type 映射为腾讯 PolygonType 值。
func tencentPolygonType(pt string) string {
	switch strings.ToLower(pt) {
	case "quadrilateral":
		return "quadrilateral"
	default:
		return pt
	}
}

// extractBase64 从 data URL（"data:image/xxx;base64,..."）中提取 base64 部分。
func extractBase64(dataURL string) string {
	if idx := strings.Index(dataURL, ","); idx >= 0 {
		return dataURL[idx+1:]
	}
	return dataURL
}

// buildAI3DRequestHeader 构造 AI3D 请求的 TC3 签名头。
// 由 tencentvod/task_adaptor.go 的 BuildRequestHeader 在 3D 模式下调用。
func buildAI3DRequestHeader(secretId, secretKey, modelName string, bodyBytes []byte, req *http.Request) error {
	meta, ok := getAI3DModelMeta(modelName)
	if !ok {
		return fmt.Errorf("unknown AI3D model for header: %s", modelName)
	}
	action := meta.SubmitAction
	ts := time.Now().Unix()
	auth := buildAuthorization(secretId, secretKey, action, string(bodyBytes), ts, tencentAI3DHost, tencentAI3DService)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Host", tencentAI3DHost)
	req.Header.Set("X-TC-Action", action)
	req.Header.Set("X-TC-Version", tencentAI3DVersion)
	req.Header.Set("X-TC-Timestamp", fmt.Sprintf("%d", ts))
	req.Header.Set("X-TC-Region", tencentAI3DRegion)
	req.Header.Set("Authorization", auth)
	return nil
}

// getAI3DQueryActionFromTaskData 从 task.Data 中提取 model 字段，再查询对应的 QueryAction。
// 用于需要从已存储数据中恢复查询信息的场景（不常用，备用）。
func getAI3DQueryActionFromTaskData(_ []byte) string { return "" }

// validateAI3DEstimatedCredits 日志校验：打印预估积分摘要。
func validateAI3DEstimatedCredits(model string, credits int) {
	common.SysLog(fmt.Sprintf("[tencent-ai3d] estimated credits: model=%s credits=%d", model, credits))
}

// saveActualCreditsToContext is a placeholder; actual credits are stored via BillingContext snapshot.
func saveActualCreditsToContext(_ string, _ float64) {}

// isAI3DAction returns true when the relay action indicates a 3D submit.
func isAI3DAction(action string) bool {
	return action == constant.TaskAction3DGenerate
}
