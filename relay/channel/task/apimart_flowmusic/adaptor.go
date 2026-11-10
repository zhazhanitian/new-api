package apimart_flowmusic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/QuantumNous/new-api/common"
	"github.com/QuantumNous/new-api/dto"
	"github.com/QuantumNous/new-api/logger"
	"github.com/QuantumNous/new-api/model"
	"github.com/QuantumNous/new-api/relay/channel"
	relaycommon "github.com/QuantumNous/new-api/relay/common"
	"github.com/QuantumNous/new-api/service"
	"github.com/gin-gonic/gin"
)

// upstreamModel is the model name APIMart expects on every Flow Music tool.
const upstreamModel = "flowmusic"

const contextKeyRequest = "task_request"

var _ channel.TaskAdaptor = (*TaskAdaptor)(nil)

// flowMusicHTTPClient talks to APIMart directly and ignores environment proxies.
// A channel-level proxy is applied separately in DoRequest and FetchTask.
var flowMusicHTTPClient = &http.Client{
	Transport: &http.Transport{
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 10,
		IdleConnTimeout:     90 * time.Second,
	},
}

// TaskAdaptor implements channel.TaskAdaptor for the APIMart Flow Music API.
type TaskAdaptor struct {
	ChannelType int
}

type submitResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
	Data    []struct {
		Status string `json:"status"`
		TaskID string `json:"task_id"`
	} `json:"data"`
}

func (r *submitResponse) IsSuccess() bool {
	return r.Code == 200
}

func (a *TaskAdaptor) Init(info *relaycommon.RelayInfo) {
	a.ChannelType = info.ChannelType
}

func (a *TaskAdaptor) EstimateBilling(c *gin.Context, info *relaycommon.RelayInfo) map[string]float64 {
	return EstimateBilling(c, info)
}

func (a *TaskAdaptor) AdjustBillingOnSubmit(info *relaycommon.RelayInfo, taskData []byte) map[string]float64 {
	return AdjustBillingOnSubmit(info, taskData)
}

func (a *TaskAdaptor) AdjustBillingOnComplete(task *model.Task, taskResult *relaycommon.TaskInfo) int {
	return AdjustBillingOnComplete(task, taskResult)
}

// ValidateRequestAndSetAction accepts any JSON object for a known model ID.
// Parameter ranges and version values are left for the upstream to reject.
func (a *TaskAdaptor) ValidateRequestAndSetAction(c *gin.Context, info *relaycommon.RelayInfo) *dto.TaskError {
	var raw map[string]json.RawMessage
	if err := common.UnmarshalBodyReusable(c, &raw); err != nil {
		return service.TaskErrorWrapperLocal(err, "invalid_request", http.StatusBadRequest)
	}
	if raw == nil {
		return service.TaskErrorWrapperLocal(fmt.Errorf("request body must be a JSON object"), "invalid_request", http.StatusBadRequest)
	}

	modelID := info.OriginModelName
	if _, ok := GetToolDef(modelID); !ok {
		return service.TaskErrorWrapperLocal(
			fmt.Errorf("unknown flow music model: %s", modelID),
			"unknown_model", http.StatusBadRequest,
		)
	}

	info.Action = modelID
	c.Set(contextKeyRequest, raw)
	return nil
}

func (a *TaskAdaptor) BuildRequestURL(info *relaycommon.RelayInfo) (string, error) {
	def, ok := GetToolDef(info.OriginModelName)
	if !ok {
		return "", fmt.Errorf("unknown model: %s", info.OriginModelName)
	}
	base := strings.TrimRight(info.ChannelBaseUrl, "/")
	if def.Path == "" {
		return base + "/v1/music/generations", nil
	}
	return base + "/v1/music/generations/" + def.Path, nil
}

func (a *TaskAdaptor) BuildRequestHeader(_ *gin.Context, req *http.Request, info *relaycommon.RelayInfo) error {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+info.ApiKey)
	return nil
}

// BuildRequestBody forwards the client JSON object.
// model is rewritten to "flowmusic", and task_id is removed because it only locks the channel.
func (a *TaskAdaptor) BuildRequestBody(c *gin.Context, info *relaycommon.RelayInfo) (io.Reader, error) {
	v, ok := c.Get(contextKeyRequest)
	if !ok {
		return nil, fmt.Errorf("task_request not found in context")
	}
	raw, ok := v.(map[string]json.RawMessage)
	if !ok {
		return nil, fmt.Errorf("task_request has unexpected type")
	}

	body := make(map[string]json.RawMessage, len(raw))
	for k, val := range raw {
		body[k] = val
	}
	delete(body, "task_id")
	modelJSON, err := json.Marshal(upstreamModel)
	if err != nil {
		return nil, fmt.Errorf("marshal upstream model: %w", err)
	}
	body["model"] = modelJSON
	if def, ok := GetToolDef(info.OriginModelName); ok && def.HasVersion && versionMissing(body["version"]) {
		versionJSON, marshalErr := json.Marshal(defaultVersion)
		if marshalErr != nil {
			return nil, fmt.Errorf("marshal default version: %w", marshalErr)
		}
		body["version"] = versionJSON
	}

	data, err := common.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("marshal request body: %w", err)
	}

	upstreamURL, _ := a.BuildRequestURL(info)
	logger.LogInfo(c, fmt.Sprintf("[apimart-flowmusic] submit request: model=%s url=%s body=%s", info.OriginModelName, upstreamURL, string(data)))
	return bytes.NewReader(data), nil
}

func versionMissing(raw json.RawMessage) bool {
	trimmed := bytes.TrimSpace(raw)
	return len(trimmed) == 0 || string(trimmed) == `""` || string(trimmed) == "null"
}

func (a *TaskAdaptor) DoRequest(c *gin.Context, info *relaycommon.RelayInfo, requestBody io.Reader) (*http.Response, error) {
	fullURL, err := a.BuildRequestURL(info)
	if err != nil {
		return nil, fmt.Errorf("build request url: %w", err)
	}
	req, err := http.NewRequest(http.MethodPost, fullURL, requestBody)
	if err != nil {
		return nil, fmt.Errorf("new request: %w", err)
	}
	if err = a.BuildRequestHeader(c, req, info); err != nil {
		return nil, fmt.Errorf("build request header: %w", err)
	}
	client := flowMusicHTTPClient
	if info.ChannelSetting.Proxy != "" {
		proxyClient, proxyErr := service.NewProxyHttpClient(info.ChannelSetting.Proxy)
		if proxyErr != nil {
			return nil, fmt.Errorf("new proxy client: %w", proxyErr)
		}
		client = proxyClient
	}
	return client.Do(req)
}

func (a *TaskAdaptor) DoResponse(c *gin.Context, resp *http.Response, info *relaycommon.RelayInfo) (taskID string, taskData []byte, taskErr *dto.TaskError) {
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		taskErr = service.TaskErrorWrapper(err, "read_response_body_failed", http.StatusInternalServerError)
		return
	}

	logger.LogInfo(c, fmt.Sprintf("[apimart-flowmusic] submit response: model=%s status=%d body=%s", info.OriginModelName, resp.StatusCode, string(responseBody)))

	if resp.StatusCode != http.StatusOK {
		taskErr = service.TaskErrorWrapper(
			fmt.Errorf("upstream returned status %d: %s", resp.StatusCode, string(responseBody)),
			"upstream_error", resp.StatusCode,
		)
		return
	}

	var submitResp submitResponse
	if err = common.Unmarshal(responseBody, &submitResp); err != nil {
		taskErr = service.TaskErrorWrapper(err, "unmarshal_response_failed", http.StatusInternalServerError)
		return
	}
	if !submitResp.IsSuccess() {
		msg := submitResp.Message
		if msg == "" {
			msg = fmt.Sprintf("upstream returned code %d", submitResp.Code)
		}
		taskErr = service.TaskErrorWrapperLocal(fmt.Errorf("%s", msg), "task_submit_failed", http.StatusBadRequest)
		return
	}
	if len(submitResp.Data) == 0 || submitResp.Data[0].TaskID == "" {
		taskErr = service.TaskErrorWrapperLocal(fmt.Errorf("upstream returned empty task_id"), "task_submit_failed", http.StatusInternalServerError)
		return
	}

	upstreamTaskID := submitResp.Data[0].TaskID
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data": []gin.H{
			{"status": "submitted", "task_id": info.PublicTaskID},
		},
	})
	return upstreamTaskID, responseBody, nil
}

func (a *TaskAdaptor) FetchTask(baseUrl, key string, body map[string]any, proxy string) (*http.Response, error) {
	upstreamTaskID, ok := body["task_id"].(string)
	if !ok || upstreamTaskID == "" {
		return nil, fmt.Errorf("FetchTask: missing task_id in body")
	}

	url := strings.TrimRight(baseUrl, "/") + "/v1/music/tasks/" + upstreamTaskID
	logger.LogInfo(context.Background(), fmt.Sprintf("[apimart-flowmusic] fetch task request: url=%s", url))

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("FetchTask: new request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+key)

	client := flowMusicHTTPClient
	if proxy != "" {
		var proxyErr error
		client, proxyErr = service.NewProxyHttpClient(proxy)
		if proxyErr != nil {
			return nil, fmt.Errorf("FetchTask: get http client: %w", proxyErr)
		}
	}
	return client.Do(req)
}

func (a *TaskAdaptor) ParseTaskResult(respBody []byte) (*relaycommon.TaskInfo, error) {
	return ParseTaskResult(respBody)
}

func (a *TaskAdaptor) GetModelList() []string {
	return ModelList
}

func (a *TaskAdaptor) GetChannelName() string {
	return ChannelName
}
