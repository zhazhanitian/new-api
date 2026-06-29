package joyagent

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
	"github.com/QuantumNous/new-api/model"
	"github.com/QuantumNous/new-api/relay/channel"
	taskcommon "github.com/QuantumNous/new-api/relay/channel/task/taskcommon"
	relaycommon "github.com/QuantumNous/new-api/relay/common"
	"github.com/QuantumNous/new-api/service"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

// ── Request / Response structures ────────────────────────────────────────────

// mediaItem 对应阿里云/JoyAgent media 数组的单个元素
// type 可选值：first_frame / reference_image / video
type mediaItem struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}

type submitInput struct {
	Prompt string      `json:"prompt,omitempty"`
	Media  []mediaItem `json:"media,omitempty"`
}

type submitParameters struct {
	Resolution   string `json:"resolution,omitempty"`
	Ratio        string `json:"ratio,omitempty"`
	Duration     int    `json:"duration,omitempty"`
	Seed         int64  `json:"seed,omitempty"`
	AudioSetting string `json:"audio_setting,omitempty"`
}

type submitRequest struct {
	Model      string           `json:"model"`
	Input      submitInput      `json:"input"`
	Parameters submitParameters `json:"parameters,omitempty"`
}

type submitResponse struct {
	RequestID string       `json:"request_id"`
	Output    submitOutput `json:"output"`
	Code      int          `json:"code,omitempty"`
	Msg       string       `json:"msg,omitempty"`
}

type submitOutput struct {
	TaskID     string `json:"task_id"`
	TaskStatus string `json:"task_status"`
}

// queryRequest body 为平铺结构（非嵌套在 input 下），已由 JoyAgent 接口文档确认
type queryRequest struct {
	TaskID string `json:"task_id"`
}

type queryResponse struct {
	RequestID string      `json:"request_id"`
	Output    queryOutput `json:"output"`
}

type queryOutput struct {
	TaskID     string `json:"task_id"`
	TaskStatus string `json:"task_status"`
	VideoURL   string `json:"video_url,omitempty"`
	Code       string `json:"code,omitempty"`
	Message    string `json:"message,omitempty"`
}

// taskMetadata 从请求的 metadata 字段读取扩展参数
type taskMetadata struct {
	Seed         int64  `json:"seed"`
	Ratio        string `json:"ratio"`
	AudioSetting string `json:"audio_setting"`
}

// ── TaskAdaptor ───────────────────────────────────────────────────────────────

type TaskAdaptor struct {
	taskcommon.BaseBilling
	apiKey  string
	baseURL string
}

func (a *TaskAdaptor) Init(info *relaycommon.RelayInfo) {
	a.baseURL = info.ChannelBaseUrl
	a.apiKey = strings.TrimSpace(info.ApiKey)
}

func (a *TaskAdaptor) ValidateRequestAndSetAction(c *gin.Context, info *relaycommon.RelayInfo) *dto.TaskError {
	model := info.OriginModelName
	action := constant.TaskActionTextGenerate
	if strings.HasSuffix(model, "i2v") {
		action = constant.TaskActionGenerate
	} else if strings.HasSuffix(model, "r2v") {
		action = constant.TaskActionReferenceGenerate
	} else if strings.Contains(model, "video-edit") {
		action = constant.TaskActionGenerate
	}
	return relaycommon.ValidateBasicTaskRequest(c, info, action)
}

func (a *TaskAdaptor) BuildRequestURL(info *relaycommon.RelayInfo) (string, error) {
	// JoyAgent URL 路径中模型名的点号需替换为横杠
	// 例：happyhorse-1.0-t2v → happyhorse-1-0-t2v
	toolName := strings.ReplaceAll(info.OriginModelName, ".", "-")
	return fmt.Sprintf("%s/api/saas/plugin-u/v1/exec/%s", a.baseURL, toolName), nil
}

func (a *TaskAdaptor) BuildRequestHeader(_ *gin.Context, req *http.Request, _ *relaycommon.RelayInfo) error {
	req.Header.Set("Authorization", buildAuthHeader(a.apiKey))
	req.Header.Set("Content-Type", "application/json")
	return nil
}

func (a *TaskAdaptor) BuildRequestBody(c *gin.Context, info *relaycommon.RelayInfo) (io.Reader, error) {
	taskReq, err := relaycommon.GetTaskRequest(c)
	if err != nil {
		return nil, errors.Wrap(err, "get_task_request_failed")
	}

	var meta taskMetadata
	_ = taskcommon.UnmarshalMetadata(taskReq.Metadata, &meta)

	modelName := info.OriginModelName
	inp := submitInput{Prompt: taskReq.Prompt}

	switch info.Action {
	case constant.TaskActionTextGenerate:
		// t2v：仅 prompt，无 media

	case constant.TaskActionGenerate:
		if strings.HasSuffix(modelName, "i2v") {
			// i2v：首帧图，优先取 InputReference，其次 Images[0]
			imgURL := taskReq.InputReference
			if imgURL == "" && len(taskReq.Images) > 0 {
				imgURL = taskReq.Images[0]
			}
			if imgURL != "" {
				inp.Media = []mediaItem{{Type: "first_frame", URL: imgURL}}
			}
		} else {
			// video-edit：Images[0] 为待编辑视频，Images[1:] 为参考图（可选）
			for i, url := range taskReq.Images {
				t := "reference_image"
				if i == 0 {
					t = "video"
				}
				inp.Media = append(inp.Media, mediaItem{Type: t, URL: url})
			}
		}

	case constant.TaskActionReferenceGenerate:
		// r2v：Images 数组中每个 URL 均为参考图
		for _, url := range taskReq.Images {
			inp.Media = append(inp.Media, mediaItem{Type: "reference_image", URL: url})
		}
	}

	params := submitParameters{
		Resolution: sizeToResolution(taskReq.Size),
		Seed:       meta.Seed,
	}

	// ratio 仅 t2v/r2v 支持；i2v 无需（宽高比跟随首帧），video-edit 无需
	if meta.Ratio != "" && !strings.HasSuffix(modelName, "i2v") && !strings.Contains(modelName, "video-edit") {
		params.Ratio = meta.Ratio
	}

	// duration 仅 t2v/i2v/r2v 支持；video-edit 输出时长跟随输入视频
	if !strings.Contains(modelName, "video-edit") {
		if taskReq.Duration > 0 {
			params.Duration = taskReq.Duration
		}
	}

	// audio_setting 仅 video-edit 支持
	if meta.AudioSetting != "" && strings.Contains(modelName, "video-edit") {
		params.AudioSetting = meta.AudioSetting
	}

	body := submitRequest{
		Model:      modelName,
		Input:      inp,
		Parameters: params,
	}

	bs, err := common.Marshal(body)
	if err != nil {
		return nil, errors.Wrap(err, "marshal_request_failed")
	}
	return bytes.NewReader(bs), nil
}

func (a *TaskAdaptor) DoRequest(c *gin.Context, info *relaycommon.RelayInfo, requestBody io.Reader) (*http.Response, error) {
	return channel.DoTaskApiRequest(a, c, info, requestBody)
}

func (a *TaskAdaptor) DoResponse(c *gin.Context, resp *http.Response, info *relaycommon.RelayInfo) (taskID string, taskData []byte, taskErr *dto.TaskError) {
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", nil, service.TaskErrorWrapper(err, "read_response_failed", http.StatusInternalServerError)
	}
	defer resp.Body.Close()

	var submitResp submitResponse
	if err := common.Unmarshal(bodyBytes, &submitResp); err != nil {
		return "", nil, service.TaskErrorWrapper(
			errors.Wrapf(err, "body: %s", bodyBytes), "unmarshal_failed", http.StatusInternalServerError)
	}
	if submitResp.Code != 0 {
		return "", nil, service.TaskErrorWrapper(
			fmt.Errorf("joyagent error %d: %s", submitResp.Code, submitResp.Msg),
			"joyagent_error", http.StatusInternalServerError)
	}
	if submitResp.Output.TaskID == "" {
		return "", nil, service.TaskErrorWrapper(
			fmt.Errorf("joyagent: empty task_id in response, body: %s", bodyBytes),
			"empty_task_id", http.StatusInternalServerError)
	}

	ov := dto.NewOpenAIVideo()
	ov.ID = info.PublicTaskID
	ov.TaskID = info.PublicTaskID
	ov.CreatedAt = time.Now().Unix()
	ov.Model = info.OriginModelName
	c.JSON(http.StatusOK, ov)

	return submitResp.Output.TaskID, bodyBytes, nil
}

// FetchTask 调用 JoyAgent tasks-query 工具轮询任务状态
func (a *TaskAdaptor) FetchTask(baseURL, key string, body map[string]any, proxy string) (*http.Response, error) {
	taskID, _ := body["task_id"].(string)

	queryURL := fmt.Sprintf("%s/api/saas/plugin-u/v1/exec/%s", baseURL, queryToolName)
	payload, err := common.Marshal(queryRequest{TaskID: taskID})
	if err != nil {
		return nil, errors.Wrap(err, "marshal_query_request_failed")
	}

	req, err := http.NewRequest(http.MethodPost, queryURL, bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", buildAuthHeader(key))
	req.Header.Set("Content-Type", "application/json")

	client, err := service.GetHttpClientWithProxy(proxy)
	if err != nil {
		return nil, fmt.Errorf("get_http_client_failed: %w", err)
	}
	return client.Do(req)
}

func (a *TaskAdaptor) ParseTaskResult(respBody []byte) (*relaycommon.TaskInfo, error) {
	var qr queryResponse
	if err := common.Unmarshal(respBody, &qr); err != nil {
		return nil, errors.Wrap(err, "unmarshal_query_response_failed")
	}

	result := &relaycommon.TaskInfo{}

	if qr.Output.Code != "" {
		result.Code = 1
		result.Reason = fmt.Sprintf("%s: %s", qr.Output.Code, qr.Output.Message)
		result.Status = model.TaskStatusFailure
		result.Progress = "100%"
		return result, nil
	}

	switch qr.Output.TaskStatus {
	case "PENDING":
		result.Status = model.TaskStatusQueued
		result.Progress = "10%"
	case "RUNNING":
		result.Status = model.TaskStatusInProgress
		result.Progress = "50%"
	case "SUCCEEDED":
		result.Status = model.TaskStatusSuccess
		result.Progress = "100%"
		result.Url = qr.Output.VideoURL
	case "FAILED":
		result.Status = model.TaskStatusFailure
		result.Progress = "100%"
		result.Reason = qr.Output.Message
	case "CANCELED":
		result.Status = model.TaskStatusFailure
		result.Progress = "100%"
		result.Reason = "task canceled"
	default:
		result.Status = model.TaskStatusQueued
		result.Progress = "0%"
	}
	return result, nil
}

func (a *TaskAdaptor) ConvertToOpenAIVideo(originTask *model.Task) ([]byte, error) {
	var qr queryResponse
	if err := common.Unmarshal(originTask.Data, &qr); err != nil {
		return nil, errors.Wrap(err, "unmarshal_task_data_failed")
	}

	ov := dto.NewOpenAIVideo()
	ov.ID = originTask.TaskID
	ov.Status = originTask.Status.ToVideoStatus()
	ov.SetProgressStr(originTask.Progress)
	ov.CreatedAt = originTask.CreatedAt
	ov.CompletedAt = originTask.UpdatedAt
	ov.Model = originTask.Properties.OriginModelName

	if qr.Output.VideoURL != "" {
		ov.SetMetadata("url", qr.Output.VideoURL)
	}
	if qr.Output.Code != "" {
		ov.Error = &dto.OpenAIVideoError{
			Message: qr.Output.Message,
			Code:    qr.Output.Code,
		}
	}

	return common.Marshal(ov)
}

func (a *TaskAdaptor) GetModelList() []string { return ModelList }
func (a *TaskAdaptor) GetChannelName() string { return ChannelName }
