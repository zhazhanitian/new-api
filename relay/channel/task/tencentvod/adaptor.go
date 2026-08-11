package tencentvod

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/QuantumNous/new-api/common"
	"github.com/QuantumNous/new-api/dto"
	"github.com/QuantumNous/new-api/pkg/billingexpr"
	"github.com/QuantumNous/new-api/relay/channel"
	relaycommon "github.com/QuantumNous/new-api/relay/common"
	"github.com/QuantumNous/new-api/service"
	"github.com/QuantumNous/new-api/types"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/sjson"
)

// Adaptor implements channel.Adaptor for synchronous image generation via
// /v1/images/generations. It submits a CreateAigcImageTask, then polls
// DescribeTaskDetail internally for up to 90 seconds before returning.
type Adaptor struct {
	subAppId  int64
	secretId  string
	secretKey string
	baseURL   string
}

func (a *Adaptor) Init(info *relaycommon.RelayInfo) {
	a.baseURL = info.ChannelBaseUrl
	parts := strings.Split(info.ApiKey, "|")
	if len(parts) == 3 {
		var n int64
		fmt.Sscan(strings.TrimSpace(parts[0]), &n)
		a.subAppId = n
		a.secretId = strings.TrimSpace(parts[1])
		a.secretKey = strings.TrimSpace(parts[2])
	}
}

func (a *Adaptor) GetRequestURL(info *relaycommon.RelayInfo) (string, error) {
	return a.baseURL, nil
}

func (a *Adaptor) SetupRequestHeader(c *gin.Context, req *http.Header, info *relaycommon.RelayInfo) error {
	channel.SetupApiRequestHeader(info, c, req)
	req.Set("Content-Type", "application/json")
	return nil
}

func (a *Adaptor) ConvertImageRequest(c *gin.Context, info *relaycommon.RelayInfo, request dto.ImageRequest) (any, error) {
	if a.subAppId == 0 {
		return nil, errors.New("TencentVOD: SubAppId 未配置，请将渠道 API Key 格式设置为 subAppId|secretId|secretKey")
	}
	if strings.TrimSpace(request.Prompt) == "" {
		return nil, errors.New("tencentvod: prompt is required")
	}

	quality := request.Quality
	size := request.Size

	lookupModel := resolveModelName(info)
	tencentModelName := modelToTencentName(lookupModel)

	// 各模型的 ModelVersion 来源不同，不能统一通过 quality 推断：
	// - OG: quality 映射到精度档位（image2_low/medium/high），由 qualityToModelVersion 处理
	// - GG/Kling/Vidu/Qwen/Hunyuan: 版本固定由原始模型名决定，quality 另独立映射到 Resolution
	var modelVersion string
	if tencentModelName == "GG" {
		modelVersion = getGGVersion(lookupModel)
	} else if tencentModelName == "Kling" {
		modelVersion = getKlingVersion(lookupModel)
	} else if tencentModelName == "Vidu" {
		modelVersion = getViduVersion(lookupModel)
	} else if tencentModelName == "Qwen" {
		modelVersion = getQwenVersion(lookupModel)
	} else if tencentModelName == "Hunyuan" {
		modelVersion = getHunyuanVersion(lookupModel)
	} else {
		modelVersion = qualityToModelVersion(tencentModelName, quality)
	}

	body := createTaskRequest{
		SubAppId:     a.subAppId,
		ModelName:    tencentModelName,
		ModelVersion: modelVersion,
		Prompt:       request.Prompt,
	}

	// ── 尺寸/分辨率处理（各模型路径不同）────────────────────────────────────
	// OG:      ExtInfo.AdditionalParameters.size 传自定义像素尺寸（[655,360–8,294,400] 像素）
	// GG:      OutputConfig.Resolution（大写 1K/2K/4K）+ AspectRatio
	// Kling:   OutputConfig.Resolution（小写 1k/2k/4k）+ AspectRatio
	//          scene 版本：跳过 OutputConfig，仅设 SceneType 和扩图 ExtInfo
	// Vidu:    OutputConfig.Resolution（1080p/2K/4K）+ AspectRatio
	// Qwen:    ExtInfo.AdditionalParameters.size 传自定义像素尺寸（[512×512–2048×2048] 像素）
	// Hunyuan: ExtInfo.AdditionalParameters.size 传自定义像素尺寸（各维度 [512,2048]，乘积≤1M）
	// 其他:    OutputConfig.AspectRatio（fallback）
	if tencentModelName == "OG" {
		if extInfo := buildOGExtInfo(size); extInfo != "" {
			body.ExtInfo = extInfo
		}
	} else if tencentModelName == "GG" {
		aspectRatio, arErr := getGGAspectRatio(modelVersion, size)
		if arErr != nil {
			// 非致命错误：记录日志，使用回退值继续
			common.SysLog(fmt.Sprintf("TencentVOD GG ConvertImageRequest: aspect ratio fallback: %v", arErr))
		}
		resolution, resErr := getGGResolution(modelVersion, quality)
		if resErr != nil {
			// 非致命错误：记录日志，使用回退值继续
			common.SysLog(fmt.Sprintf("TencentVOD GG ConvertImageRequest: resolution fallback: %v", resErr))
		}
		body.OutputConfig = &outputConfig{AspectRatio: aspectRatio, Resolution: resolution}
	} else if tencentModelName == "Kling" {
		aspectRatio, arErr := getKlingAspectRatio(modelVersion, size)
		if arErr != nil {
			common.SysLog(fmt.Sprintf("TencentVOD Kling ConvertImageRequest: aspect ratio fallback: %v", arErr))
		}
		resolution, resErr := getKlingResolution(modelVersion, quality)
		if resErr != nil {
			common.SysLog(fmt.Sprintf("TencentVOD Kling ConvertImageRequest: resolution fallback: %v", resErr))
		}
		// scene 版本：getKlingAspectRatio/getKlingResolution 均返回 ""，跳过 OutputConfig
		if aspectRatio != "" || resolution != "" {
			body.OutputConfig = &outputConfig{AspectRatio: aspectRatio, Resolution: resolution}
		}
	} else if tencentModelName == "Vidu" {
		aspectRatio, arErr := getViduAspectRatio(size)
		if arErr != nil {
			common.SysLog(fmt.Sprintf("TencentVOD Vidu ConvertImageRequest: aspect ratio fallback: %v", arErr))
		}
		resolution, resErr := getViduResolution(quality)
		if resErr != nil {
			common.SysLog(fmt.Sprintf("TencentVOD Vidu ConvertImageRequest: resolution fallback: %v", resErr))
		}
		body.OutputConfig = &outputConfig{AspectRatio: aspectRatio, Resolution: resolution}
	} else if tencentModelName == "Qwen" {
		// Qwen 不支持 OutputConfig，使用 ExtInfo 传自定义像素尺寸
		// 特殊说明：Qwen 0925 合法总像素范围 [512×512=261,632, 2048×2048=4,194,304]
		if extInfo := buildOGExtInfo(size); extInfo != "" {
			body.ExtInfo = extInfo
		}
	} else if tencentModelName == "Hunyuan" {
		// Hunyuan 不支持 OutputConfig，使用 ExtInfo 传自定义像素尺寸
		// 特殊说明：Hunyuan 3.0 宽高均在 [512, 2048] 像素范围内，宽高乘积 ≤ 1024×1024
		if extInfo := buildOGExtInfo(size); extInfo != "" {
			body.ExtInfo = extInfo
		}
	} else {
		body.OutputConfig = &outputConfig{AspectRatio: sizeToAspectRatio(size)}
	}

	// OutputImageCount：将请求的 n 映射到 OutputConfig.OutputImageCount。
	// 仅 OG（上限 8）和 Kling（上限 9）支持；其他模型忽略 n 字段。
	// n <= 1 时不设此字段，使用 API 默认值（1 张）。
	if request.N != nil && int(*request.N) > 1 {
		maxCount := 0
		switch tencentModelName {
		case "OG":
			maxCount = 8
		case "Kling":
			maxCount = 9
		}
		if maxCount > 0 {
			count := int(*request.N)
			if count > maxCount {
				common.SysLog(fmt.Sprintf(
					"TencentVOD %s: OutputImageCount %d 超出上限 %d，已截断至 %d 张",
					tencentModelName, count, maxCount, maxCount,
				))
				count = maxCount
			}
			if body.OutputConfig == nil {
				body.OutputConfig = &outputConfig{}
			}
			body.OutputConfig.OutputImageCount = count
		}
	}

	// OutputFormat：将请求的 output_format 映射到 OutputConfig.OutputFormat。
	// 可选值：jpeg / png；不传时跟随模型默认值。
	if len(request.OutputFormat) > 0 {
		var ofStr string
		if err := common.Unmarshal(request.OutputFormat, &ofStr); err == nil && ofStr != "" {
			if body.OutputConfig == nil {
				body.OutputConfig = &outputConfig{}
			}
			body.OutputConfig.OutputFormat = ofStr
		}
	}

	// extended params from extra_fields (json.RawMessage)
	if len(request.ExtraFields) > 0 {
		var extraMap map[string]interface{}
		if err := common.Unmarshal(request.ExtraFields, &extraMap); err == nil {
			if v, ok := extraMap["negative_prompt"].(string); ok {
				body.NegativePrompt = v
			}
			if v, ok := extraMap["seed"].(float64); ok {
				body.Seed = int64(v)
			}
			// EnhancePrompt 自动优化提示词，仅 GG 系列支持，取值 "Enabled" / "Disabled"
			if v, ok := extraMap["enhance_prompt"].(string); ok && tencentModelName == "GG" {
				body.EnhancePrompt = v
			}
			// SceneType：Kling scene 和 Hunyuan 3.0 支持，通过 extra_fields.scene_type 传入
			// Kling 取值 "image_expand"；Hunyuan 3.0 取值 "3d_panorama"
			if v, ok := extraMap["scene_type"].(string); ok &&
				(tencentModelName == "Kling" || tencentModelName == "Hunyuan") {
				body.SceneType = v
			}
			// Kling 扩图参数：仅 SceneType="image_expand" 时有效，通过 extra_fields.expansion 传入
			// 格式：{"up": 0.1, "down": 0.2, "left": 0.3, "right": 0.4}，各值范围 [0, 2]
			if tencentModelName == "Kling" {
				if expansion, ok := extraMap["expansion"].(map[string]interface{}); ok {
					if extInfo := buildKlingExpansionExtInfoFromMap(expansion); extInfo != "" {
						body.ExtInfo = extInfo
					}
				}
			}
		}
	}

	// 参考图：优先从 images（数组）读取，再兼容 image（单图字符串）
	var refURLs []string
	if len(request.Images) > 0 {
		var urls []string
		if err := common.Unmarshal(request.Images, &urls); err == nil {
			refURLs = urls
		}
	}
	if len(refURLs) == 0 && len(request.Image) > 0 {
		var singleURL string
		if err := common.Unmarshal(request.Image, &singleURL); err == nil && singleURL != "" {
			refURLs = []string{singleURL}
		}
	}

	// 各模型均有参考图数量上限，超出时截断并记录日志，避免上游报错。
	// GG 2.5: 3张；GG 3.0/3.1: 14张；Kling 2.1: 4张；Kling 3.0: 1张；
	// Kling 3.0-Omni/O1: 10张；Kling scene: 1张；Vidu q2: 7张；Qwen 0925: 1张；Hunyuan 3.0: 3张
	maxRef := getMaxRefImages(tencentModelName, modelVersion)
	if maxRef > 0 && len(refURLs) > maxRef {
		common.SysLog(fmt.Sprintf(
			"TencentVOD %s %s: 参考图数量 %d 超出上限 %d，已截断至 %d 张",
			tencentModelName, modelVersion, len(refURLs), maxRef, maxRef,
		))
		refURLs = refURLs[:maxRef]
	}

	for _, url := range refURLs {
		if strings.TrimSpace(url) != "" {
			body.FileInfos = append(body.FileInfos, fileInfo{Type: "Url", Url: url})
		}
	}

	return body, nil
}

func (a *Adaptor) DoRequest(c *gin.Context, info *relaycommon.RelayInfo, requestBody io.Reader) (any, error) {
	bodyBytes, err := io.ReadAll(requestBody)
	if err != nil {
		return nil, err
	}

	ts := time.Now().Unix()
	auth := buildAuthorization(a.secretId, a.secretKey, "CreateAigcImageTask", string(bodyBytes), ts, tencentVODHost, tencentVODService)

	req, err := http.NewRequest(http.MethodPost, a.baseURL, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Host", tencentVODHost)
	req.Header.Set("X-TC-Action", "CreateAigcImageTask")
	req.Header.Set("X-TC-Version", "2018-07-17")
	req.Header.Set("X-TC-Timestamp", fmt.Sprintf("%d", ts))
	req.Header.Set("Authorization", auth)

	client := service.GetHttpClient()
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var createResp createTaskResponse
	if err := common.Unmarshal(respBodyBytes, &createResp); err != nil {
		return nil, fmt.Errorf("tencentvod: unmarshal create response: %w", err)
	}
	if createResp.Response.Error != nil {
		return nil, fmt.Errorf("tencentvod: %s: %s",
			createResp.Response.Error.Code, createResp.Response.Error.Message)
	}

	taskID := createResp.Response.TaskId

	// ── internal polling (up to 90s) ─────────────────────────────────
	const (
		pollInterval = 3 * time.Second
		maxWait      = 90 * time.Second
	)
	deadline := time.Now().Add(maxWait)

	for time.Now().Before(deadline) {
		time.Sleep(pollInterval)

		descResp, pollErr := a.describeTask(taskID)
		if pollErr != nil {
			return nil, pollErr
		}
		if descResp.Response.Error != nil {
			return nil, fmt.Errorf("tencentvod: %s: %s",
				descResp.Response.Error.Code, descResp.Response.Error.Message)
		}

	// Use AigcImageTask.Status when available; fall back to top-level Status
	taskStatus := descResp.Response.Status
	if descResp.Response.AigcImageTask != nil && descResp.Response.AigcImageTask.Status != "" {
		taskStatus = descResp.Response.AigcImageTask.Status
	}
	switch taskStatus {
	case "SUCCESS", "FINISH":
		// Store result in context; DoResponse will read it.
		c.Set("tencentvod_sync_result", descResp)
		return nil, nil
	case "FAIL":
		var msg string
		if descResp.Response.AigcImageTask != nil {
			msg = descResp.Response.AigcImageTask.ErrMsg
			if msg == "" {
				msg = descResp.Response.AigcImageTask.ErrCodeExt
			}
		}
		return nil, fmt.Errorf("tencentvod: image generation failed: %s", msg)
	}
	}
	return nil, errors.New("tencentvod: image generation timed out after 90s")
}

func (a *Adaptor) DoResponse(c *gin.Context, resp *http.Response, info *relaycommon.RelayInfo) (any, *types.NewAPIError) {
	// The actual result was already resolved in DoRequest (sync polling).
	// resp here is a synthetic wrapper; the real data is in info.Request context.
	// Because DoRequest returns the describeTaskResponse directly (not via http.Response),
	// we need to pull it out of the gin context.
	val, exists := c.Get("tencentvod_sync_result")
	if !exists {
		return nil, types.NewError(errors.New("tencentvod: missing sync result"), types.ErrorCodeBadResponse)
	}

	descResp, ok := val.(*describeTaskResponse)
	if !ok {
		return nil, types.NewError(errors.New("tencentvod: invalid sync result type"), types.ErrorCodeBadResponse)
	}

	imageResponse := dto.ImageResponse{
		Created: common.GetTimestamp(),
		Data:    []dto.ImageData{},
	}

	// 从原始请求中回填 quality / size / output_format / background
	if imageReq, ok := info.Request.(*dto.ImageRequest); ok {
		imageResponse.Quality = imageReq.Quality
		imageResponse.Size = imageReq.Size
		if len(imageReq.OutputFormat) > 0 {
			var ofStr string
			if err := common.Unmarshal(imageReq.OutputFormat, &ofStr); err == nil && ofStr != "" {
				imageResponse.OutputFormat = ofStr
			}
		}
		if len(imageReq.Background) > 0 {
			var bgStr string
			if err := common.Unmarshal(imageReq.Background, &bgStr); err == nil && bgStr != "" {
				imageResponse.Background = bgStr
			}
		}
	}

	if descResp.Response.AigcImageTask != nil && descResp.Response.AigcImageTask.Output != nil {
		for _, fi := range descResp.Response.AigcImageTask.Output.FileInfos {
			imageResponse.Data = append(imageResponse.Data, dto.ImageData{Url: fi.FileUrl})
		}
	}
	if len(imageResponse.Data) == 0 {
		return nil, types.NewError(errors.New("tencentvod: no images in response"), types.ErrorCodeBadResponseBody)
	}

	// 修正计费：多退少补——实际生成数与请求 n 不符时按实际数量计费
	// 同时覆盖两条计费路径（price-based 和 tiered_expr）
	actualN := len(imageResponse.Data)
	requestedN := 1
	if imageReq, ok := info.Request.(*dto.ImageRequest); ok && imageReq.N != nil {
		requestedN = int(*imageReq.N)
	}
	if actualN != requestedN {
		// 路径 1：price-based 计费（UsePrice=true）→ 修正 OtherRatios["n"]
		// 仅限 UsePrice 模型，ratio 模型不用 OtherRatios["n"] 计费，注入会错误地乘以 n 倍
		if info.PriceData.UsePrice {
			info.PriceData.AddOtherRatio("n", float64(actualN))
		}
		// 路径 2：tiered_expr 计费 → 修正 BillingRequestInput.Body 里的 n
		// 仅在 TieredBillingSnapshot 存在时才有效（其他模式 TryTieredSettle 会直接跳过）
		if info.TieredBillingSnapshot != nil &&
			info.BillingRequestInput != nil && len(info.BillingRequestInput.Body) > 0 {
			if newBody, sjsonErr := sjson.SetBytes(info.BillingRequestInput.Body, "n", actualN); sjsonErr == nil {
				updated := billingexpr.RequestInput{
					Headers: info.BillingRequestInput.Headers,
					Body:    newBody,
				}
				info.BillingRequestInput = &updated
			}
		}
	}

	// 腾讯 VOD 接口不返回 token 用量，按生图惯例用 prompt 字符数做占位估算
	promptTokens := 0
	if imageReq, ok := info.Request.(*dto.ImageRequest); ok {
		promptTokens = len([]rune(imageReq.Prompt))
	}
	imageResponse.Usage = &dto.ImageUsage{
		InputTokens: promptTokens,
		InputTokensDetails: &dto.ImageTokensDetails{
			TextTokens: promptTokens,
		},
		OutputTokens: 1,
		OutputTokensDetails: &dto.ImageTokensDetails{
			ImageTokens: 1,
		},
		TotalTokens: promptTokens + 1,
	}

	respBytes, err := common.Marshal(imageResponse)
	if err != nil {
		return nil, types.NewError(err, types.ErrorCodeBadResponseBody)
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	_, _ = c.Writer.Write(respBytes)

	return &dto.Usage{PromptTokens: promptTokens, TotalTokens: promptTokens + 1}, nil
}

// describeTask calls DescribeTaskDetail and returns the parsed response.
func (a *Adaptor) describeTask(taskID string) (*describeTaskResponse, error) {
	reqBody := describeTaskRequest{TaskId: taskID}
	if a.subAppId != 0 {
		reqBody.SubAppId = a.subAppId
	}
	payload, err := common.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	ts := time.Now().Unix()
	auth := buildAuthorization(a.secretId, a.secretKey, "DescribeTaskDetail", string(payload), ts, tencentVODHost, tencentVODService)

	req, err := http.NewRequest(http.MethodPost, a.baseURL, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Host", tencentVODHost)
	req.Header.Set("X-TC-Action", "DescribeTaskDetail")
	req.Header.Set("X-TC-Version", "2018-07-17")
	req.Header.Set("X-TC-Timestamp", fmt.Sprintf("%d", ts))
	req.Header.Set("Authorization", auth)

	client := service.GetHttpClient()
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var descResp describeTaskResponse
	if err := common.Unmarshal(respBodyBytes, &descResp); err != nil {
		return nil, fmt.Errorf("tencentvod: unmarshal describe response: %w", err)
	}
	return &descResp, nil
}

// ── Stub methods for unused channel.Adaptor methods ─────────────────────────

func (a *Adaptor) ConvertOpenAIRequest(*gin.Context, *relaycommon.RelayInfo, *dto.GeneralOpenAIRequest) (any, error) {
	return nil, errors.New("tencentvod: ConvertOpenAIRequest not supported")
}
func (a *Adaptor) ConvertRerankRequest(*gin.Context, int, dto.RerankRequest) (any, error) {
	return nil, errors.New("tencentvod: ConvertRerankRequest not supported")
}
func (a *Adaptor) ConvertEmbeddingRequest(*gin.Context, *relaycommon.RelayInfo, dto.EmbeddingRequest) (any, error) {
	return nil, errors.New("tencentvod: ConvertEmbeddingRequest not supported")
}
func (a *Adaptor) ConvertAudioRequest(*gin.Context, *relaycommon.RelayInfo, dto.AudioRequest) (io.Reader, error) {
	return nil, errors.New("tencentvod: ConvertAudioRequest not supported")
}
func (a *Adaptor) ConvertOpenAIResponsesRequest(*gin.Context, *relaycommon.RelayInfo, dto.OpenAIResponsesRequest) (any, error) {
	return nil, errors.New("tencentvod: ConvertOpenAIResponsesRequest not supported")
}
func (a *Adaptor) ConvertClaudeRequest(*gin.Context, *relaycommon.RelayInfo, *dto.ClaudeRequest) (any, error) {
	return nil, errors.New("tencentvod: ConvertClaudeRequest not supported")
}
func (a *Adaptor) ConvertGeminiRequest(*gin.Context, *relaycommon.RelayInfo, *dto.GeminiChatRequest) (any, error) {
	return nil, errors.New("tencentvod: ConvertGeminiRequest not supported")
}

func (a *Adaptor) GetModelList() []string { return ModelList }
func (a *Adaptor) GetChannelName() string  { return ChannelName }
