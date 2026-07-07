package tencentvod

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

type outputConfig struct {
	AspectRatio string `json:"AspectRatio,omitempty"`
}

// fileInfo 对应腾讯 VOD API 的 AigcImageTaskInputFileInfo
// Type 取值：File（VOD文件ID）| Url（外部可访问URL）| Base64（Base64字符串）
type fileInfo struct {
	Type   string `json:"Type,omitempty"`
	FileId string `json:"FileId,omitempty"`
	Url    string `json:"Url,omitempty"`
	Base64 string `json:"Base64,omitempty"`
}

type createTaskRequest struct {
	SubAppId       int64         `json:"SubAppId,omitempty"`
	ModelName      string        `json:"ModelName"`
	ModelVersion   string        `json:"ModelVersion"`
	Prompt         string        `json:"Prompt"`
	FileInfos      []fileInfo    `json:"FileInfos,omitempty"`
	OutputConfig   *outputConfig `json:"OutputConfig,omitempty"`
	Seed           int64         `json:"Seed,omitempty"`
	NegativePrompt string        `json:"NegativePrompt,omitempty"`
	ExtInfo        string        `json:"ExtInfo,omitempty"`
}

type createTaskResponse struct {
	Response struct {
		TaskId    string `json:"TaskId"`
		RequestId string `json:"RequestId"`
		Error     *struct {
			Code    string `json:"Code"`
			Message string `json:"Message"`
		} `json:"Error,omitempty"`
	} `json:"Response"`
}

type describeTaskRequest struct {
	SubAppId int64  `json:"SubAppId,omitempty"`
	TaskId   string `json:"TaskId"`
}

type describeTaskResponse struct {
	Response struct {
		RequestId string `json:"RequestId"`
		Status    string `json:"Status"`
		Error     *struct {
			Code    string `json:"Code"`
			Message string `json:"Message"`
		} `json:"Error,omitempty"`
		AigcImageTask *struct {
			Status     string `json:"Status"`
			ErrCode    int    `json:"ErrCode"`
			ErrCodeExt string `json:"ErrCodeExt"`
			ErrMsg     string `json:"Message"`
			Output     *struct {
				FileInfos []struct {
					FileUrl string `json:"FileUrl"`
				} `json:"FileInfos"`
			} `json:"Output,omitempty"`
		} `json:"AigcImageTask,omitempty"`
	} `json:"Response"`
}

// metadata fields clients may include inside the "metadata" object
type taskMetadata struct {
	Seed           int64  `json:"seed"`
	NegativePrompt string `json:"negative_prompt"`
}

// ── TaskAdaptor ───────────────────────────────────────────────────────────────

type TaskAdaptor struct {
	taskcommon.BaseBilling
	subAppId  int64
	secretId  string
	secretKey string
	baseURL   string
}

func (a *TaskAdaptor) Init(info *relaycommon.RelayInfo) {
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

func (a *TaskAdaptor) ValidateRequestAndSetAction(c *gin.Context, info *relaycommon.RelayInfo) *dto.TaskError {
	return relaycommon.ValidateBasicTaskRequest(c, info, constant.TaskActionImageGenerate)
}

func (a *TaskAdaptor) BuildRequestURL(info *relaycommon.RelayInfo) (string, error) {
	return a.baseURL, nil
}

func (a *TaskAdaptor) BuildRequestHeader(c *gin.Context, req *http.Request, info *relaycommon.RelayInfo) error {
	bodyBytes, err := io.ReadAll(req.Body)
	if err != nil {
		return err
	}
	req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	ts := time.Now().Unix()
	auth := buildAuthorization(a.secretId, a.secretKey, "CreateAigcImageTask", string(bodyBytes), ts)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Host", tencentVODHost)
	req.Header.Set("X-TC-Action", "CreateAigcImageTask")
	req.Header.Set("X-TC-Version", "2018-07-17")
	req.Header.Set("X-TC-Timestamp", fmt.Sprintf("%d", ts))
	req.Header.Set("Authorization", auth)
	return nil
}

func (a *TaskAdaptor) BuildRequestBody(c *gin.Context, info *relaycommon.RelayInfo) (io.Reader, error) {
	if a.subAppId == 0 {
		return nil, fmt.Errorf("TencentVOD: SubAppId 未配置，请将渠道 API Key 格式设置为 subAppId|secretId|secretKey")
	}

	taskReq, err := relaycommon.GetTaskRequest(c)
	if err != nil {
		return nil, err
	}

	var meta taskMetadata
	_ = taskcommon.UnmarshalMetadata(taskReq.Metadata, &meta)

	// quality comes from top-level field; fall back to metadata["quality"] for compatibility
	quality := taskReq.Quality
	if quality == "" {
		quality, _ = taskReq.Metadata["quality"].(string)
	}
	tencentModelName := modelToTencentName(info.OriginModelName)
	body := createTaskRequest{
		SubAppId:       a.subAppId,
		ModelName:      tencentModelName,
		ModelVersion:   qualityToModelVersion(tencentModelName, quality),
		Prompt:         taskReq.Prompt,
		Seed:           meta.Seed,
		NegativePrompt: meta.NegativePrompt,
	}
	// OG model uses ExtInfo for pixel size; other models use OutputConfig.AspectRatio
	if tencentModelName == "OG" {
		if extInfo := buildOGExtInfo(taskReq.Size); extInfo != "" {
			body.ExtInfo = extInfo
		}
	} else {
		body.OutputConfig = &outputConfig{AspectRatio: sizeToAspectRatio(taskReq.Size)}
	}

	// 参考图：从 taskReq.Images 读取（单图 image 字段已在 ValidateBasicTaskRequest 里自动合并进 Images）
	for _, url := range taskReq.Images {
		if strings.TrimSpace(url) != "" {
			body.FileInfos = append(body.FileInfos, fileInfo{Type: "Url", Url: url})
		}
	}

	data, err := common.Marshal(body)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(data), nil
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

	var createResp createTaskResponse
	if err := common.Unmarshal(bodyBytes, &createResp); err != nil {
		return "", nil, service.TaskErrorWrapper(
			errors.Wrapf(err, "body: %s", bodyBytes), "unmarshal_failed", http.StatusInternalServerError)
	}
	if createResp.Response.Error != nil {
		return "", nil, service.TaskErrorWrapper(
			fmt.Errorf("%s: %s", createResp.Response.Error.Code, createResp.Response.Error.Message),
			createResp.Response.Error.Code, http.StatusInternalServerError)
	}

	ov := dto.NewOpenAIVideo()
	ov.ID = info.PublicTaskID
	ov.TaskID = info.PublicTaskID
	ov.CreatedAt = time.Now().Unix()
	ov.Model = info.OriginModelName
	c.JSON(http.StatusOK, ov)

	return createResp.Response.TaskId, bodyBytes, nil
}

// FetchTask calls DescribeTaskDetail to poll task status.
func (a *TaskAdaptor) FetchTask(baseUrl, key string, body map[string]any, proxy string) (*http.Response, error) {
	taskID, _ := body["task_id"].(string)

	// Re-parse credentials from the stored key
	parts := strings.Split(key, "|")
	var subAppId int64
	var secretId, secretKey string
	if len(parts) == 3 {
		fmt.Sscan(strings.TrimSpace(parts[0]), &subAppId)
		secretId = strings.TrimSpace(parts[1])
		secretKey = strings.TrimSpace(parts[2])
	}

	reqBody := describeTaskRequest{TaskId: taskID}
	if subAppId != 0 {
		reqBody.SubAppId = subAppId
	}
	payload, err := common.Marshal(reqBody)
	if err != nil {
		return nil, errors.Wrap(err, "marshal fetch payload failed")
	}

	ts := time.Now().Unix()
	auth := buildAuthorization(secretId, secretKey, "DescribeTaskDetail", string(payload), ts)

	req, err := http.NewRequest(http.MethodPost, baseUrl, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Host", tencentVODHost)
	req.Header.Set("X-TC-Action", "DescribeTaskDetail")
	req.Header.Set("X-TC-Version", "2018-07-17")
	req.Header.Set("X-TC-Timestamp", fmt.Sprintf("%d", ts))
	req.Header.Set("Authorization", auth)

	client, err := service.GetHttpClientWithProxy(proxy)
	if err != nil {
		return nil, fmt.Errorf("get http client failed: %w", err)
	}
	return client.Do(req)
}

func (a *TaskAdaptor) ParseTaskResult(respBody []byte) (*relaycommon.TaskInfo, error) {
	var descResp describeTaskResponse
	if err := common.Unmarshal(respBody, &descResp); err != nil {
		return nil, errors.Wrap(err, "unmarshal describe response failed")
	}

	result := &relaycommon.TaskInfo{}

	if descResp.Response.Error != nil {
		result.Code = 1
		result.Reason = fmt.Sprintf("%s: %s", descResp.Response.Error.Code, descResp.Response.Error.Message)
		result.Status = model.TaskStatusFailure
		result.Progress = "100%"
		return result, nil
	}

	// Use AigcImageTask.Status when available; fall back to top-level Status
	taskStatus := descResp.Response.Status
	if descResp.Response.AigcImageTask != nil && descResp.Response.AigcImageTask.Status != "" {
		taskStatus = descResp.Response.AigcImageTask.Status
	}

	switch taskStatus {
	case "PROCESSING", "WAITING":
		result.Status = model.TaskStatusInProgress
		result.Progress = "50%"
	case "SUCCESS", "FINISH":
		// Tencent returns Status=FINISH for both success and failure; check ErrCode to distinguish.
		common.SysLog(fmt.Sprintf("TencentVOD ParseTaskResult: status=%s errCode=%d errCodeExt=%s",
			taskStatus,
			func() int {
				if descResp.Response.AigcImageTask != nil {
					return descResp.Response.AigcImageTask.ErrCode
				}
				return -1
			}(),
			func() string {
				if descResp.Response.AigcImageTask != nil {
					return descResp.Response.AigcImageTask.ErrCodeExt
				}
				return "nil"
			}()))
		if descResp.Response.AigcImageTask != nil &&
			(descResp.Response.AigcImageTask.ErrCode != 0 || descResp.Response.AigcImageTask.ErrCodeExt != "") {
			result.Status = model.TaskStatusFailure
			result.Progress = "100%"
			result.Reason = descResp.Response.AigcImageTask.ErrMsg
			if result.Reason == "" {
				result.Reason = descResp.Response.AigcImageTask.ErrCodeExt
			}
		} else {
			result.Status = model.TaskStatusSuccess
			result.Progress = "100%"
			if descResp.Response.AigcImageTask != nil &&
				descResp.Response.AigcImageTask.Output != nil &&
				len(descResp.Response.AigcImageTask.Output.FileInfos) > 0 {
				result.Url = descResp.Response.AigcImageTask.Output.FileInfos[0].FileUrl
			}
		}
	case "FAIL":
		result.Status = model.TaskStatusFailure
		result.Progress = "100%"
		if descResp.Response.AigcImageTask != nil {
			result.Reason = descResp.Response.AigcImageTask.ErrMsg
			if result.Reason == "" {
				result.Reason = descResp.Response.AigcImageTask.ErrCodeExt
			}
		}
	default:
		result.Status = model.TaskStatusQueued
		result.Progress = "10%"
	}
	return result, nil
}

func (a *TaskAdaptor) ConvertToOpenAIVideo(originTask *model.Task) ([]byte, error) {
	var descResp describeTaskResponse
	if err := common.Unmarshal(originTask.Data, &descResp); err != nil {
		return nil, errors.Wrap(err, "unmarshal task data failed")
	}

	ov := dto.NewOpenAIVideo()
	ov.ID = originTask.TaskID
	ov.Status = originTask.Status.ToVideoStatus()
	ov.SetProgressStr(originTask.Progress)
	ov.CreatedAt = originTask.CreatedAt
	ov.CompletedAt = originTask.UpdatedAt
	ov.Model = originTask.Properties.OriginModelName

	if descResp.Response.AigcImageTask != nil &&
		descResp.Response.AigcImageTask.Output != nil &&
		len(descResp.Response.AigcImageTask.Output.FileInfos) > 0 {
		ov.SetMetadata("url", descResp.Response.AigcImageTask.Output.FileInfos[0].FileUrl)
	}
	if descResp.Response.Error != nil {
		ov.Error = &dto.OpenAIVideoError{
			Message: descResp.Response.Error.Message,
			Code:    descResp.Response.Error.Code,
		}
	}

	return common.Marshal(ov)
}

func (a *TaskAdaptor) GetModelList() []string { return ModelList }
func (a *TaskAdaptor) GetChannelName() string  { return ChannelName }
