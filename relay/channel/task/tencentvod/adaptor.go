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
// DescribeTaskDetail internally until the image is ready or the wait limit is reached.
type Adaptor struct {
	subAppId  int64
	secretId  string
	secretKey string
	baseURL   string
	syncBody  []byte
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

	var extra map[string]interface{}
	if len(request.ExtraFields) > 0 {
		_ = common.Unmarshal(request.ExtraFields, &extra)
	}
	n := 0
	if request.N != nil {
		n = int(*request.N)
	}
	var outputFormat string
	if len(request.OutputFormat) > 0 {
		_ = common.Unmarshal(request.OutputFormat, &outputFormat)
	}
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
		} else {
			var urls []string
			if err := common.Unmarshal(request.Image, &urls); err == nil {
				refURLs = urls
			}
		}
	}

	body := buildImageCreateRequest(imageBuildInput{
		SubAppId:     a.subAppId,
		LookupModel:  resolveModelName(info),
		Prompt:       request.Prompt,
		Size:         request.Size,
		Resolution:   relaycommon.ExtraString(request.Extra, "resolution"),
		Ratio:        relaycommon.ExtraString(request.Extra, "ratio"),
		Quality:      request.Quality,
		N:            n,
		OutputFormat: outputFormat,
		Images:       refURLs,
		Extra:        extra,
	})
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

	// 同步接口要在同一次请求里带回图片。腾讯生图本身是异步任务，
	// gpt-image-2 经常超过 90 秒，这里把内部等待放到 10 分钟。
	const (
		pollInterval = 3 * time.Second
		maxWait      = 10 * time.Minute
	)
	deadline := time.Now().Add(maxWait)
	ctx := c.Request.Context()

	for time.Now().Before(deadline) {
		timer := time.NewTimer(pollInterval)
		select {
		case <-ctx.Done():
			timer.Stop()
			return nil, ctx.Err()
		case <-timer.C:
		}

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
	return nil, fmt.Errorf("tencentvod: image generation timed out after %s", maxWait)
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

	var imageURLs []string
	if descResp.Response.AigcImageTask != nil && descResp.Response.AigcImageTask.Output != nil {
		for _, fi := range descResp.Response.AigcImageTask.Output.FileInfos {
			if fi.FileUrl == "" {
				continue
			}
			imageResponse.Data = append(imageResponse.Data, dto.ImageData{Url: fi.FileUrl})
			imageURLs = append(imageURLs, fi.FileUrl)
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

	// 任务日志存 DescribeTaskDetail 原文。
	// 给调用方的响应先留着，等结算完成后补上 amount 再写出。
	service.KeepSyncImageRecord(c, a.syncBody, imageURLs)
	service.KeepSyncImageClientResponse(c, &imageResponse)

	return &dto.Usage{}, nil
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
	a.syncBody = append([]byte(nil), respBodyBytes...)
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
func (a *Adaptor) GetChannelName() string { return ChannelName }
