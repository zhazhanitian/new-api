package apimart_image

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"strings"

	"github.com/QuantumNous/new-api/common"
	"github.com/QuantumNous/new-api/constant"
	"github.com/QuantumNous/new-api/dto"
	"github.com/QuantumNous/new-api/model"
	"github.com/QuantumNous/new-api/pkg/billingexpr"
	"github.com/QuantumNous/new-api/relay/channel"
	taskcommon "github.com/QuantumNous/new-api/relay/channel/task/taskcommon"
	relaycommon "github.com/QuantumNous/new-api/relay/common"
	"github.com/QuantumNous/new-api/service"

	"github.com/gin-gonic/gin"
)

const ChannelName = "APIMartImage"

var ModelList = []string{
	"gpt-image-2.5-flare",
	"gpt-image-2.5-sunburst",
	"gpt-image-2.6-flare",
	"gpt-image-2.6-sunburst",
	"flux-3-image",
	"flux-3-image-klein",
	"mai-image-2.6",
	"mai-image-2.6-flash",
	"qwen-image-3.0",
	"qwen-image-3.0-pro",
	"muse-image-ext",
	"hy-image-v3.5-preview",
}

// TaskAdaptor 把四字段图片请求转成 APIMart 的异步生图任务。
type TaskAdaptor struct {
	taskcommon.BaseBilling
	ChannelType int
}

func (a *TaskAdaptor) Init(info *relaycommon.RelayInfo) {
	a.ChannelType = info.ChannelType
}

func (a *TaskAdaptor) ValidateRequestAndSetAction(c *gin.Context, info *relaycommon.RelayInfo) *dto.TaskError {
	if _, exists := c.Get("task_request"); !exists {
		if taskErr := relaycommon.ValidateImageTaskRequest(c, info); taskErr != nil {
			return taskErr
		}
	}
	info.Action = constant.TaskActionImageGenerate
	req, err := relaycommon.GetTaskRequest(c)
	if err != nil {
		return service.TaskErrorWrapperLocal(err, "invalid_request", http.StatusBadRequest)
	}
	// 只有按 token 结算的模型才在完成时改预扣。按张计费的模型保持原来的按次或表达式。
	if isTokenBilledModel(req.Model) {
		info.AllowCompleteAdjustment = true
	}
	if err = validateExtras(req); err != nil {
		return service.TaskErrorWrapperLocal(err, "invalid_request", http.StatusBadRequest)
	}
	return nil
}

func (a *TaskAdaptor) BuildRequestURL(info *relaycommon.RelayInfo) (string, error) {
	base := "https://apib.ai"
	if info != nil && info.ChannelBaseUrl != "" {
		base = strings.TrimRight(info.ChannelBaseUrl, "/")
	}
	return base + "/v1/images/generations", nil
}

func (a *TaskAdaptor) BuildRequestHeader(_ *gin.Context, req *http.Request, info *relaycommon.RelayInfo) error {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+info.ApiKey)
	return nil
}

func (a *TaskAdaptor) BuildRequestBody(c *gin.Context, info *relaycommon.RelayInfo) (io.Reader, error) {
	req, err := relaycommon.GetTaskRequest(c)
	if err != nil {
		return nil, err
	}
	modelName := info.UpstreamModelName
	if modelName == "" {
		modelName = req.Model
	}
	body := buildUpstreamBody(modelName, req)
	data, err := common.Marshal(body)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(data), nil
}

func (a *TaskAdaptor) DoRequest(c *gin.Context, info *relaycommon.RelayInfo, requestBody io.Reader) (*http.Response, error) {
	return channel.DoTaskApiRequest(a, c, info, requestBody)
}

func (a *TaskAdaptor) DoResponse(_ *gin.Context, resp *http.Response, _ *relaycommon.RelayInfo) (string, []byte, *dto.TaskError) {
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", nil, service.TaskErrorWrapper(err, "read_response_body_failed", http.StatusInternalServerError)
	}
	var submit struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    []struct {
			TaskID string `json:"task_id"`
		} `json:"data"`
	}
	if err = common.Unmarshal(responseBody, &submit); err != nil {
		return "", nil, service.TaskErrorWrapper(err, "unmarshal_response_failed", http.StatusInternalServerError)
	}
	if submit.Code != 0 && submit.Code != http.StatusOK {
		msg := submit.Message
		if msg == "" {
			msg = fmt.Sprintf("upstream returned code %d", submit.Code)
		}
		return "", nil, service.TaskErrorWrapperLocal(fmt.Errorf("%s", msg), "task_submit_failed", http.StatusBadRequest)
	}
	if len(submit.Data) == 0 || submit.Data[0].TaskID == "" {
		return "", nil, service.TaskErrorWrapperLocal(fmt.Errorf("upstream returned empty task_id"), "task_submit_failed", http.StatusBadRequest)
	}
	return submit.Data[0].TaskID, responseBody, nil
}

func (a *TaskAdaptor) FetchTask(baseURL, key string, body map[string]any, proxy string) (*http.Response, error) {
	taskID, _ := body["task_id"].(string)
	if taskID == "" {
		return nil, fmt.Errorf("missing task_id")
	}
	url := strings.TrimRight(baseURL, "/") + "/v1/tasks/" + taskID
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+key)
	client, err := service.GetHttpClientWithProxy(proxy)
	if err != nil {
		return nil, err
	}
	return client.Do(req)
}

func (a *TaskAdaptor) ParseTaskResult(respBody []byte) (*relaycommon.TaskInfo, error) {
	return parseTaskResult(respBody)
}

func (a *TaskAdaptor) GetModelList() []string { return ModelList }

func (a *TaskAdaptor) GetChannelName() string { return ChannelName }

// AdjustBillingOnComplete 只对登记了 token 单价的模型按 usage 结算。
// 同一个渠道里按张计费的模型返回 0，继续用预扣表达式，并按实际张数多退少补。
func (a *TaskAdaptor) AdjustBillingOnComplete(task *model.Task, result *relaycommon.TaskInfo) int {
	if task == nil || result == nil {
		return 0
	}
	if isTokenBilledModel(billingModelName(task)) && result.TokenUsage != nil {
		return quotaFromPlatformExpr(task, *result.TokenUsage)
	}
	return imageCountQuota(task, result)
}

type upstreamBody struct {
	Model             string   `json:"model"`
	Prompt            string   `json:"prompt"`
	N                 int      `json:"n,omitempty"`
	Size              string   `json:"size,omitempty"`
	Resolution        string   `json:"resolution,omitempty"`
	Quality           string   `json:"quality,omitempty"`
	OutputFormat      string   `json:"output_format,omitempty"`
	OutputCompression int      `json:"output_compression,omitempty"`
	ImageURLs         []string `json:"image_urls,omitempty"`
	Background        string   `json:"background,omitempty"`
	Moderation        string   `json:"moderation,omitempty"`
	MaskURL           string   `json:"mask_url,omitempty"`
	MarkupURL         string   `json:"markup_url,omitempty"`
	Complexity        string   `json:"complexity,omitempty"`
	AspectRatio       string   `json:"aspect_ratio,omitempty"`
	SafetyTolerance   *int     `json:"safety_tolerance,omitempty"`
	Grounding         *bool    `json:"grounding,omitempty"`
	Seed              *int64   `json:"seed,omitempty"`
	Steps             *int     `json:"steps,omitempty"`
	Width             *int     `json:"width,omitempty"`
	Height            *int     `json:"height,omitempty"`
	WebGrounding      *bool    `json:"web_grounding,omitempty"`
	AutoAspectRatio   *bool    `json:"auto_aspect_ratio,omitempty"`
	NegativePrompt    string   `json:"negative_prompt,omitempty"`
	PromptExtend      *bool    `json:"prompt_extend,omitempty"`
	PromptExtendMode  string   `json:"prompt_extend_mode,omitempty"`
	NsfwCheck         *bool    `json:"nsfw_check,omitempty"`
}

func buildUpstreamBody(modelName string, req relaycommon.TaskSubmitReq) upstreamBody {
	switch {
	case isFluxModel(modelName):
		return buildFluxUpstreamBody(modelName, req)
	case isMaiModel(modelName):
		return buildMaiUpstreamBody(modelName, req)
	case isQwenImage3(modelName):
		return buildQwenImage3UpstreamBody(modelName, req)
	case isMuseModel(modelName):
		return buildMuseUpstreamBody(modelName, req)
	case isHyImageModel(modelName):
		return buildHyImageUpstreamBody(modelName, req)
	}
	body := upstreamBody{
		Model:        modelName,
		Prompt:       req.Prompt,
		N:            req.N,
		OutputFormat: req.OutputFormat,
		Quality:      req.Quality,
	}
	if body.Quality == "" {
		body.Quality = relaycommon.ImageUpstreamDefaultQuality(modelName)
	}
	if req.Size != "" {
		body.Size = req.Size
	} else if req.Ratio != "" {
		body.Size = req.Ratio
	}
	if req.Size == "" && (req.Resolution != "" || isTokenBilledModel(modelName)) {
		resolution := req.Resolution
		if resolution == "" {
			resolution = "1k"
		}
		if isTokenBilledModel(modelName) {
			resolution = strings.ToLower(resolution)
		}
		body.Resolution = resolution
	}
	if req.OutputCompression > 0 {
		body.OutputCompression = req.OutputCompression
	}
	images := append([]string(nil), req.Images...)
	if limit := relaycommon.ImageRefLimit(modelName); limit > 0 && len(images) > limit {
		common.SysLog(fmt.Sprintf("APIMart %s: 参考图数量 %d 超出上限 %d，已截断", modelName, len(images), limit))
		images = images[:limit]
	}
	body.ImageURLs = images
	body.Background = metaString(req.Metadata, "background")
	body.Moderation = metaString(req.Metadata, "moderation")
	body.MaskURL = metaString(req.Metadata, "mask_url")
	body.MarkupURL = metaString(req.Metadata, "markup_url")
	body.Complexity = metaString(req.Metadata, "complexity")
	return body
}

func isFluxModel(modelName string) bool {
	switch strings.TrimSpace(modelName) {
	case "flux-3-image", "flux-3-image-klein":
		return true
	default:
		return false
	}
}

func isFluxKlein(modelName string) bool {
	return strings.TrimSpace(modelName) == "flux-3-image-klein"
}

// buildFluxUpstreamBody 只发 FLUX 认识的字段。像素 size、quality 和 GPT 扩展字段会让上游返回 400。
func buildFluxUpstreamBody(modelName string, req relaycommon.TaskSubmitReq) upstreamBody {
	body := upstreamBody{
		Model:  modelName,
		Prompt: req.Prompt,
		N:      req.N,
	}
	if req.Ratio != "" {
		body.AspectRatio = req.Ratio
	}
	body.Resolution = fluxResolution(req.Resolution)
	body.ImageURLs = limitImages(modelName, req.Images)
	if isFluxKlein(modelName) {
		if format := fluxOutputFormat(req.OutputFormat); format != "" {
			body.OutputFormat = format
		}
		if seed, ok, err := metaInt(req.Metadata, "seed"); err == nil && ok {
			body.Seed = &seed
		}
		if steps, ok, err := metaInt(req.Metadata, "steps"); err == nil && ok {
			n := int(steps)
			body.Steps = &n
		}
		return body
	}
	if tolerance, ok, err := metaInt(req.Metadata, "safety_tolerance"); err == nil && ok {
		n := int(tolerance)
		body.SafetyTolerance = &n
	}
	if grounding, ok, err := metaBool(req.Metadata, "grounding"); err == nil && ok {
		body.Grounding = &grounding
	}
	return body
}

func fluxResolution(value string) string {
	value = strings.TrimSpace(value)
	if value == "" || strings.EqualFold(value, "1k") {
		return "1k"
	}
	if strings.EqualFold(value, "768") || strings.EqualFold(value, "768sq") {
		return "768sq"
	}
	return strings.ToLower(value)
}

func isMaiModel(modelName string) bool {
	switch strings.TrimSpace(modelName) {
	case "mai-image-2.6", "mai-image-2.6-flash":
		return true
	default:
		return false
	}
}

func isQwenImage3(modelName string) bool {
	switch strings.TrimSpace(modelName) {
	case "qwen-image-3.0", "qwen-image-3.0-pro":
		return true
	default:
		return false
	}
}

func isMuseModel(modelName string) bool {
	return strings.TrimSpace(modelName) == "muse-image-ext"
}

func isHyImageModel(modelName string) bool {
	return strings.TrimSpace(modelName) == "hy-image-v3.5-preview"
}

func limitImages(modelName string, images []string) []string {
	images = append([]string(nil), images...)
	if limit := relaycommon.ImageRefLimit(modelName); limit > 0 && len(images) > limit {
		common.SysLog(fmt.Sprintf("APIMart %s: 参考图数量 %d 超出上限 %d，已截断", modelName, len(images), limit))
		return images[:limit]
	}
	return images
}

func buildMaiUpstreamBody(modelName string, req relaycommon.TaskSubmitReq) upstreamBody {
	body := upstreamBody{Model: modelName, Prompt: req.Prompt, N: req.N}
	width, height, paired := pixelPair(req.Metadata)
	if paired {
		body.Width = &width
		body.Height = &height
	} else if req.Size != "" {
		body.Size = req.Size
	} else if req.Ratio != "" {
		body.Size = req.Ratio
	}
	if !paired && req.Size == "" {
		resolution := req.Resolution
		if resolution == "" {
			resolution = "1K"
		}
		body.Resolution = resolution
	}
	body.ImageURLs = limitImages(modelName, req.Images)
	if grounding, ok, err := metaBool(req.Metadata, "web_grounding"); err == nil && ok {
		body.WebGrounding = &grounding
	}
	if auto, ok, err := metaBool(req.Metadata, "auto_aspect_ratio"); err == nil && ok {
		body.AutoAspectRatio = &auto
	}
	return body
}

func buildQwenImage3UpstreamBody(modelName string, req relaycommon.TaskSubmitReq) upstreamBody {
	body := upstreamBody{Model: modelName, Prompt: req.Prompt, N: req.N}
	if req.Size != "" {
		body.Size = req.Size
	} else if req.Ratio != "" {
		body.Size = req.Ratio
	}
	if req.Size == "" {
		resolution := req.Resolution
		if resolution == "" {
			resolution = "1K"
		}
		body.Resolution = resolution
	}
	body.ImageURLs = limitImages(modelName, req.Images)
	body.NegativePrompt = metaString(req.Metadata, "negative_prompt")
	if extend, ok, err := metaBool(req.Metadata, "prompt_extend"); err == nil && ok {
		body.PromptExtend = &extend
	}
	if mode := metaString(req.Metadata, "prompt_extend_mode"); mode != "" {
		body.PromptExtendMode = mode
	}
	if check, ok, err := metaBool(req.Metadata, "nsfw_check"); err == nil && ok {
		body.NsfwCheck = &check
	}
	return body
}

func buildMuseUpstreamBody(modelName string, req relaycommon.TaskSubmitReq) upstreamBody {
	body := upstreamBody{Model: modelName, Prompt: req.Prompt, N: req.N}
	if req.Ratio != "" {
		body.AspectRatio = req.Ratio
	}
	body.ImageURLs = limitImages(modelName, req.Images)
	return body
}

func buildHyImageUpstreamBody(modelName string, req relaycommon.TaskSubmitReq) upstreamBody {
	body := upstreamBody{Model: modelName, Prompt: req.Prompt, N: req.N}
	width, height, paired := pixelPair(req.Metadata)
	if paired {
		body.Width = &width
		body.Height = &height
	} else if req.Size != "" {
		body.Size = req.Size
	} else if req.Ratio != "" {
		body.Size = req.Ratio
	}
	if !paired && req.Size == "" {
		resolution := req.Resolution
		if resolution == "" {
			resolution = "1K"
		}
		body.Resolution = resolution
	}
	body.ImageURLs = limitImages(modelName, req.Images)
	body.NegativePrompt = metaString(req.Metadata, "negative_prompt")
	if seed, ok, err := metaInt(req.Metadata, "seed"); err == nil && ok {
		body.Seed = &seed
	}
	return body
}

func pixelPair(meta map[string]interface{}) (int, int, bool) {
	width, widthOK, widthErr := metaInt(meta, "width")
	height, heightOK, heightErr := metaInt(meta, "height")
	if widthErr != nil || heightErr != nil || widthOK != heightOK || !widthOK {
		return 0, 0, false
	}
	return int(width), int(height), true
}

func fluxOutputFormat(value string) string {
	format := strings.ToLower(strings.TrimSpace(value))
	if format == "jpg" {
		return "jpeg"
	}
	return format
}

// tokenBilledModels 只标记哪些模型要在完成时用实际 token 重跑平台表达式。
// 单价不写在这里。APIMart 的 Credits 换成平台 $/1M tokens 后，由后台表达式决定售价。
var tokenBilledModels = map[string]struct{}{
	"gpt-image-2.5-flare":    {},
	"gpt-image-2.5-sunburst": {},
	"gpt-image-2.6-flare":    {},
	"gpt-image-2.6-sunburst": {},
}

func isTokenBilledModel(modelName string) bool {
	if _, ok := tokenBilledModels[modelName]; ok {
		return true
	}
	// 对外模型名是 OG-image-*，渠道映射后再变成上游的 gpt-image-*。
	if strings.HasPrefix(modelName, "OG-image-") {
		_, ok := tokenBilledModels["gpt-image-"+strings.TrimPrefix(modelName, "OG-image-")]
		return ok
	}
	return false
}

func billingModelName(task *model.Task) string {
	if task == nil {
		return ""
	}
	if bc := task.PrivateData.BillingContext; bc != nil && bc.OriginModelName != "" {
		return bc.OriginModelName
	}
	return task.Properties.OriginModelName
}

// quotaFromPlatformExpr 用提交时冻结的阶梯表达式结算。
// 表达式系数是平台自己的 $/1M tokens，不是 APIMart Credits。
// p 文本输入，cr 缓存文本，img 图片输入，cc 缓存图片输入，img_o 图片输出。
func quotaFromPlatformExpr(task *model.Task, usage relaycommon.TaskTokenUsage) int {
	bc := task.PrivateData.BillingContext
	if bc == nil || strings.TrimSpace(bc.BillingExpr) == "" {
		return 0
	}
	groupRatio := bc.GroupRatio
	if groupRatio <= 0 {
		groupRatio = 1
	}
	quotaPerUnit := bc.QuotaPerUnit
	if quotaPerUnit <= 0 {
		quotaPerUnit = common.QuotaPerUnit
	}
	expr := strings.TrimSpace(bc.BillingExpr)
	snap := &billingexpr.BillingSnapshot{
		BillingMode:  "tiered_expr",
		ExprString:   expr,
		ExprHash:     billingexpr.ExprHashString(expr),
		GroupRatio:   groupRatio,
		QuotaPerUnit: quotaPerUnit,
		ExprVersion:  billingexpr.ExprVersion(expr),
	}
	result, err := billingexpr.ComputeTieredQuota(snap, billingexpr.TokenParams{
		P:    float64(usage.TextInputTokens),
		CR:   float64(usage.CachedTextInputTokens),
		Img:  float64(usage.ImageInputTokens),
		CC:   float64(usage.CachedImageInputTokens),
		ImgO: float64(usage.ImageOutputTokens),
	})
	if err != nil {
		return 0
	}
	return result.ActualQuotaAfterGroup
}

func imageCountQuota(task *model.Task, result *relaycommon.TaskInfo) int {
	if task.Quota <= 0 || task.PrivateData.BillingContext == nil {
		return 0
	}
	expected := task.PrivateData.BillingContext.RequestedN
	actual := result.ActualImages
	if actual == 0 {
		actual = len(result.Urls)
	}
	if expected <= 1 || actual <= 0 || actual == expected {
		return 0
	}
	return task.Quota * actual / expected
}

// splitTokenUsage 把上游 usage 拆成结算桶。
// cached_tokens 视为 text_tokens 的子集，cached_image_tokens 视为 image_tokens 的子集。
// 输出侧没有单独的文本单价，text_tokens 并入图片输出。
func splitTokenUsage(text, cachedText, image, cachedImage, outputImage, outputText, outputTokens int) *relaycommon.TaskTokenUsage {
	if cachedImage < 0 {
		cachedImage = 0
	}
	if cachedImage > image {
		cachedImage = image
	}
	if cachedText < 0 {
		cachedText = 0
	}
	if cachedText > text {
		cachedText = text
	}
	if outputImage == 0 && outputText == 0 {
		outputImage = outputTokens
	} else {
		outputImage += outputText
	}
	usage := &relaycommon.TaskTokenUsage{
		TextInputTokens:        text - cachedText,
		CachedTextInputTokens:  cachedText,
		ImageInputTokens:       image - cachedImage,
		CachedImageInputTokens: cachedImage,
		ImageOutputTokens:      outputImage,
	}
	if usage.TextInputTokens == 0 && usage.CachedTextInputTokens == 0 &&
		usage.ImageInputTokens == 0 && usage.CachedImageInputTokens == 0 &&
		usage.ImageOutputTokens == 0 {
		return nil
	}
	return usage
}

func metaString(meta map[string]interface{}, key string) string {
	if meta == nil {
		return ""
	}
	value, _ := meta[key].(string)
	return value
}

func metaPresent(meta map[string]interface{}, key string) (interface{}, bool) {
	if meta == nil {
		return nil, false
	}
	value, ok := meta[key]
	if !ok || value == nil {
		return nil, false
	}
	return value, true
}

func metaInt(meta map[string]interface{}, key string) (int64, bool, error) {
	raw, ok := metaPresent(meta, key)
	if !ok {
		return 0, false, nil
	}
	switch n := raw.(type) {
	case float64:
		if n != math.Trunc(n) || n > 4294967295 || n < -4294967295 {
			return 0, true, fmt.Errorf("%s must be an integer", key)
		}
		return int64(n), true, nil
	case int:
		return int64(n), true, nil
	case int64:
		return n, true, nil
	case json.Number:
		parsed, err := n.Int64()
		if err != nil {
			return 0, true, fmt.Errorf("%s must be an integer", key)
		}
		return parsed, true, nil
	default:
		return 0, true, fmt.Errorf("%s must be an integer", key)
	}
}

func metaBool(meta map[string]interface{}, key string) (bool, bool, error) {
	raw, ok := metaPresent(meta, key)
	if !ok {
		return false, false, nil
	}
	value, ok := raw.(bool)
	if !ok {
		return false, true, fmt.Errorf("%s must be a boolean", key)
	}
	return value, true, nil
}

func validateExtras(req relaycommon.TaskSubmitReq) error {
	switch {
	case isFluxModel(req.Model):
		return validateFluxExtras(req)
	case isMaiModel(req.Model):
		return validateMaiExtras(req)
	case isQwenImage3(req.Model):
		return validateQwenImage3Extras(req)
	case isMuseModel(req.Model):
		return validateMuseExtras(req)
	case isHyImageModel(req.Model):
		return validateHyImageExtras(req)
	}
	is26 := strings.Contains(req.Model, "image-2.6")
	format := strings.ToLower(req.OutputFormat)
	if bg := metaString(req.Metadata, "background"); strings.EqualFold(bg, "transparent") && format != "png" && format != "webp" {
		return fmt.Errorf("transparent background requires output_format png or webp")
	}
	if mod := metaString(req.Metadata, "moderation"); mod != "" && mod != "low" && mod != "auto" {
		return fmt.Errorf("moderation must be low or auto")
	}
	mask := metaString(req.Metadata, "mask_url")
	markup := metaString(req.Metadata, "markup_url")
	if mask != "" && len(req.Images) == 0 {
		return fmt.Errorf("mask_url requires a reference image")
	}
	if markup != "" && !is26 {
		return fmt.Errorf("markup_url is only supported by gpt-image-2.6")
	}
	if markup != "" && mask != "" {
		return fmt.Errorf("markup_url and mask_url cannot be used together")
	}
	if complexity := metaString(req.Metadata, "complexity"); complexity != "" {
		if !is26 {
			return fmt.Errorf("complexity is only supported by gpt-image-2.6")
		}
		switch complexity {
		case "low", "medium", "high":
		default:
			return fmt.Errorf("complexity must be low, medium, or high")
		}
	}
	if req.OutputCompression > 0 && format != "jpeg" && format != "webp" {
		return fmt.Errorf("output_compression requires output_format jpeg or webp")
	}
	return nil
}

func validateFluxExtras(req relaycommon.TaskSubmitReq) error {
	klein := isFluxKlein(req.Model)
	if klein {
		if _, ok := metaPresent(req.Metadata, "grounding"); ok {
			return fmt.Errorf("grounding is only supported by flux-3-image")
		}
		if _, ok := metaPresent(req.Metadata, "safety_tolerance"); ok {
			return fmt.Errorf("safety_tolerance is only supported by flux-3-image")
		}
		switch fluxOutputFormat(req.OutputFormat) {
		case "", "jpeg", "png":
		default:
			return fmt.Errorf("output_format %q is not supported for %s", req.OutputFormat, req.Model)
		}
		if _, ok := metaPresent(req.Metadata, "seed"); ok {
			seed, _, err := metaInt(req.Metadata, "seed")
			if err != nil {
				return err
			}
			if seed < 0 || seed > 4294967295 {
				return fmt.Errorf("seed must be between 0 and 4294967295")
			}
		}
		if _, ok := metaPresent(req.Metadata, "steps"); ok {
			steps, _, err := metaInt(req.Metadata, "steps")
			if err != nil {
				return err
			}
			if steps < 1 || steps > 16 {
				return fmt.Errorf("steps must be between 1 and 16")
			}
		}
		return nil
	}
	if req.OutputFormat != "" {
		return fmt.Errorf("output_format is not supported by flux-3-image")
	}
	if _, ok := metaPresent(req.Metadata, "seed"); ok {
		return fmt.Errorf("seed is only supported by flux-3-image-klein")
	}
	if _, ok := metaPresent(req.Metadata, "steps"); ok {
		return fmt.Errorf("steps is only supported by flux-3-image-klein")
	}
	if _, ok := metaPresent(req.Metadata, "grounding"); ok {
		if _, _, err := metaBool(req.Metadata, "grounding"); err != nil {
			return err
		}
	}
	if _, ok := metaPresent(req.Metadata, "safety_tolerance"); ok {
		tolerance, _, err := metaInt(req.Metadata, "safety_tolerance")
		if err != nil {
			return err
		}
		if tolerance < 0 || tolerance > 4 {
			return fmt.Errorf("safety_tolerance must be between 0 and 4")
		}
	}
	return nil
}

func validateMaiExtras(req relaycommon.TaskSubmitReq) error {
	if _, ok := metaPresent(req.Metadata, "web_grounding"); ok {
		if _, _, err := metaBool(req.Metadata, "web_grounding"); err != nil {
			return err
		}
	}
	if _, ok := metaPresent(req.Metadata, "auto_aspect_ratio"); ok {
		if _, _, err := metaBool(req.Metadata, "auto_aspect_ratio"); err != nil {
			return err
		}
	}
	return validatePixelPair(req.Metadata, 768, 0, 2359296)
}

func validateQwenImage3Extras(req relaycommon.TaskSubmitReq) error {
	extend, extendOK, err := metaBool(req.Metadata, "prompt_extend")
	if err != nil {
		return err
	}
	if _, ok := metaPresent(req.Metadata, "nsfw_check"); ok {
		if _, _, err = metaBool(req.Metadata, "nsfw_check"); err != nil {
			return err
		}
	}
	mode := metaString(req.Metadata, "prompt_extend_mode")
	if mode != "" && mode != "direct" && mode != "agent" {
		return fmt.Errorf("prompt_extend_mode must be direct or agent")
	}
	if mode != "" && (!extendOK || !extend) {
		return fmt.Errorf("prompt_extend_mode requires prompt_extend true")
	}
	if mode == "agent" && len(req.Images) > 0 {
		return fmt.Errorf("prompt_extend_mode agent is only supported for text-to-image")
	}
	return nil
}

func validateMuseExtras(req relaycommon.TaskSubmitReq) error {
	if req.OutputFormat != "" {
		return fmt.Errorf("output_format is not supported by muse-image-ext")
	}
	if _, ok := metaPresent(req.Metadata, "seed"); ok {
		return fmt.Errorf("seed is not supported by muse-image-ext")
	}
	if metaString(req.Metadata, "negative_prompt") != "" {
		return fmt.Errorf("negative_prompt is not supported by muse-image-ext")
	}
	return nil
}

func validateHyImageExtras(req relaycommon.TaskSubmitReq) error {
	if _, ok := metaPresent(req.Metadata, "seed"); ok {
		seed, _, err := metaInt(req.Metadata, "seed")
		if err != nil {
			return err
		}
		if seed < 0 {
			return fmt.Errorf("seed must be a non-negative integer")
		}
	}
	return validatePixelPair(req.Metadata, 256, 8192, 8294400)
}

func validatePixelPair(meta map[string]interface{}, minEdge, maxEdge, maxPixels int) error {
	width, widthOK, widthErr := metaInt(meta, "width")
	if widthErr != nil {
		return widthErr
	}
	height, heightOK, heightErr := metaInt(meta, "height")
	if heightErr != nil {
		return heightErr
	}
	if widthOK != heightOK {
		return fmt.Errorf("width and height must be provided together")
	}
	if !widthOK {
		return nil
	}
	if width < int64(minEdge) || height < int64(minEdge) {
		return fmt.Errorf("width and height must be at least %d", minEdge)
	}
	if maxEdge > 0 && (width > int64(maxEdge) || height > int64(maxEdge)) {
		return fmt.Errorf("width and height must be at most %d", maxEdge)
	}
	if maxPixels > 0 && (height > int64(maxPixels) || width > int64(maxPixels)/height) {
		return fmt.Errorf("width times height must be at most %d", maxPixels)
	}
	return nil
}

func parseTaskResult(respBody []byte) (*relaycommon.TaskInfo, error) {
	var resp struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    struct {
			ID       string `json:"id"`
			Status   string `json:"status"`
			Progress int    `json:"progress"`
			Result   struct {
				Images []struct {
					URL json.RawMessage `json:"url"`
				} `json:"images"`
			} `json:"result"`
			Error *struct {
				Message string `json:"message"`
			} `json:"error"`
			Usage struct {
				TotalTokens      int `json:"total_tokens"`
				InputTokens      int `json:"input_tokens"`
				OutputTokens     int `json:"output_tokens"`
				CompletionTokens int `json:"completion_tokens"`
				InputDetails     struct {
					CachedTokens      int `json:"cached_tokens"`
					ImageTokens       int `json:"image_tokens"`
					TextTokens        int `json:"text_tokens"`
					CachedImageTokens int `json:"cached_image_tokens"`
				} `json:"input_tokens_details"`
				OutputDetails struct {
					ImageTokens int `json:"image_tokens"`
					TextTokens  int `json:"text_tokens"`
				} `json:"output_tokens_details"`
			} `json:"usage"`
		} `json:"data"`
	}
	if err := common.Unmarshal(respBody, &resp); err != nil {
		return nil, fmt.Errorf("unmarshal APIMart image task: %w", err)
	}
	if resp.Code != 0 && resp.Code != http.StatusOK {
		msg := resp.Message
		if msg == "" {
			msg = fmt.Sprintf("upstream returned code %d", resp.Code)
		}
		return &relaycommon.TaskInfo{Status: string(model.TaskStatusFailure), Reason: msg}, nil
	}
	info := &relaycommon.TaskInfo{TaskID: resp.Data.ID}
	switch strings.ToLower(resp.Data.Status) {
	case "submitted", "pending":
		info.Status = string(model.TaskStatusSubmitted)
		info.Progress = "10%"
	case "processing":
		info.Status = string(model.TaskStatusInProgress)
		info.Progress = progressText(resp.Data.Progress, "50%")
	case "completed":
		info.Status = string(model.TaskStatusSuccess)
		info.Progress = "100%"
		for _, image := range resp.Data.Result.Images {
			info.Urls = append(info.Urls, imageURLs(image.URL)...)
		}
		if len(info.Urls) > 0 {
			info.Url = info.Urls[0]
		}
		info.ActualImages = len(info.Urls)
		info.TotalTokens = resp.Data.Usage.TotalTokens
		info.CompletionTokens = resp.Data.Usage.CompletionTokens
		if info.CompletionTokens == 0 {
			info.CompletionTokens = resp.Data.Usage.OutputTokens
		}
		info.TokenUsage = splitTokenUsage(
			resp.Data.Usage.InputDetails.TextTokens,
			resp.Data.Usage.InputDetails.CachedTokens,
			resp.Data.Usage.InputDetails.ImageTokens,
			resp.Data.Usage.InputDetails.CachedImageTokens,
			resp.Data.Usage.OutputDetails.ImageTokens,
			resp.Data.Usage.OutputDetails.TextTokens,
			resp.Data.Usage.OutputTokens,
		)
	case "failed":
		info.Status = string(model.TaskStatusFailure)
		info.Progress = "100%"
		if resp.Data.Error != nil {
			info.Reason = resp.Data.Error.Message
		}
		if info.Reason == "" {
			info.Reason = resp.Message
		}
	default:
		info.Status = string(model.TaskStatusInProgress)
		info.Progress = progressText(resp.Data.Progress, "10%")
	}
	return info, nil
}

func progressText(progress int, fallback string) string {
	if progress <= 0 {
		return fallback
	}
	return fmt.Sprintf("%d%%", progress)
}

func imageURLs(raw json.RawMessage) []string {
	if len(raw) == 0 {
		return nil
	}
	var many []string
	if err := json.Unmarshal(raw, &many); err == nil {
		var urls []string
		for _, item := range many {
			if strings.TrimSpace(item) != "" {
				urls = append(urls, item)
			}
		}
		return urls
	}
	var one string
	if err := json.Unmarshal(raw, &one); err == nil && strings.TrimSpace(one) != "" {
		return []string{one}
	}
	return nil
}
