package zhongji

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/QuantumNous/new-api/relay/channel"
	"github.com/QuantumNous/new-api/relay/channel/task/doubao"
	relaycommon "github.com/QuantumNous/new-api/relay/common"
	"github.com/QuantumNous/new-api/service"

	"github.com/gin-gonic/gin"
)

// TaskAdaptor 复用火山豆包视频适配器。
// 中纪把同一套请求体挂在 /v1 下，状态词与火山不完全相同。
// 厂商文档写 queued / in_progress / completed / failed，示例响应却是 succeeded。
// 提供的密钥已脱敏，无法在线核对，因此两种写法都接受。
type TaskAdaptor struct {
	doubao.TaskAdaptor
	baseURL string
}

func (a *TaskAdaptor) Init(info *relaycommon.RelayInfo) {
	a.TaskAdaptor.Init(info)
	if info != nil {
		a.baseURL = strings.TrimRight(info.ChannelBaseUrl, "/")
	}
}

func (a *TaskAdaptor) BuildRequestURL(_ *relaycommon.RelayInfo) (string, error) {
	return a.baseURL + "/v1/contents/generations/tasks", nil
}

// DoRequest 必须由外层实现。嵌入的火山 DoRequest 会把接收者收成火山适配器，
// 随后 BuildRequestURL 仍会拼出 /api/v3。
func (a *TaskAdaptor) DoRequest(c *gin.Context, info *relaycommon.RelayInfo, requestBody io.Reader) (*http.Response, error) {
	return channel.DoTaskApiRequest(a, c, info, requestBody)
}

func (a *TaskAdaptor) FetchTask(baseUrl, key string, body map[string]any, proxy string) (*http.Response, error) {
	taskID, ok := body["task_id"].(string)
	if !ok {
		return nil, fmt.Errorf("invalid task_id")
	}

	uri := fmt.Sprintf("%s/v1/contents/generations/tasks/%s", strings.TrimRight(baseUrl, "/"), taskID)
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

func (a *TaskAdaptor) ParseTaskResult(respBody []byte) (*relaycommon.TaskInfo, error) {
	return a.TaskAdaptor.ParseTaskResult(normalizeStatus(respBody))
}

func (a *TaskAdaptor) GetChannelName() string {
	return "zhongji-video"
}

// normalizeStatus 把中纪文档里的状态词改写成火山适配器已经识别的词。
// 只替换 status 字段，其余 JSON 原样保留。
func normalizeStatus(body []byte) []byte {
	var payload map[string]json.RawMessage
	if err := json.Unmarshal(body, &payload); err != nil {
		return body
	}
	raw, ok := payload["status"]
	if !ok {
		return body
	}
	var status string
	if err := json.Unmarshal(raw, &status); err != nil {
		return body
	}
	switch strings.ToLower(strings.TrimSpace(status)) {
	case "in_progress":
		status = "processing"
	case "completed":
		status = "succeeded"
	default:
		return body
	}
	encoded, err := json.Marshal(status)
	if err != nil {
		return body
	}
	payload["status"] = encoded
	out, err := json.Marshal(payload)
	if err != nil {
		return body
	}
	return out
}
