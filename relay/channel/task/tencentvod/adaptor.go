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
	"github.com/QuantumNous/new-api/relay/channel"
	relaycommon "github.com/QuantumNous/new-api/relay/common"
	"github.com/QuantumNous/new-api/service"
	"github.com/QuantumNous/new-api/types"
	"github.com/gin-gonic/gin"
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

	tencentModelName := modelToTencentName(info.OriginModelName)
	body := createTaskRequest{
		SubAppId:     a.subAppId,
		ModelName:    tencentModelName,
		ModelVersion: qualityToModelVersion(tencentModelName, quality),
		Prompt:       request.Prompt,
	}
	if tencentModelName == "OG" {
		if extInfo := buildOGExtInfo(size); extInfo != "" {
			body.ExtInfo = extInfo
		}
	} else {
		body.OutputConfig = &outputConfig{AspectRatio: sizeToAspectRatio(size)}
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
	auth := buildAuthorization(a.secretId, a.secretKey, "CreateAigcImageTask", string(bodyBytes), ts)

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
	auth := buildAuthorization(a.secretId, a.secretKey, "DescribeTaskDetail", string(payload), ts)

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
