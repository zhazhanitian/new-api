package apimart_flowmusic

import (
	"context"
	"fmt"

	"github.com/QuantumNous/new-api/common"
	"github.com/QuantumNous/new-api/logger"
	"github.com/QuantumNous/new-api/model"
	relaycommon "github.com/QuantumNous/new-api/relay/common"
)

type fetchResponse struct {
	Code    int        `json:"code"`
	Message string     `json:"message"`
	Data    *fetchData `json:"data"`
}

type fetchData struct {
	ID       string       `json:"id"`
	Status   string       `json:"status"`
	Progress int          `json:"progress"`
	Result   *fetchResult `json:"result"`
	Error    *fetchError  `json:"error"`
}

type fetchResult struct {
	Music []fetchMusicItem `json:"music"`
}

type fetchMusicItem struct {
	AudioURL string `json:"audio_url"`
	VideoURL string `json:"video_url"`
	FileURL  string `json:"file_url"`
}

type fetchError struct {
	Message string `json:"message"`
}

// ParseTaskResult maps GET /v1/music/tasks/:task_id into a platform task status.
// The raw body is stored separately by the poller, so lyrics and zip fields stay intact.
func ParseTaskResult(respBody []byte) (*relaycommon.TaskInfo, error) {
	logger.LogInfo(context.Background(), fmt.Sprintf("[apimart-flowmusic] fetch task response: body=%s", string(respBody)))

	var resp fetchResponse
	if err := common.Unmarshal(respBody, &resp); err != nil {
		return nil, fmt.Errorf("unmarshal APIMart flow music fetch response: %w", err)
	}

	if resp.Code != 200 {
		msg := resp.Message
		if msg == "" {
			msg = fmt.Sprintf("upstream returned code %d", resp.Code)
		}
		return &relaycommon.TaskInfo{
			Status: string(model.TaskStatusFailure),
			Reason: msg,
		}, nil
	}

	data := resp.Data
	if data == nil {
		return &relaycommon.TaskInfo{
			Status: string(model.TaskStatusFailure),
			Reason: "upstream returned empty task data",
		}, nil
	}

	info := &relaycommon.TaskInfo{
		TaskID: data.ID,
	}

	switch data.Status {
	case "submitted":
		info.Status = string(model.TaskStatusSubmitted)
		info.Progress = "10%"
	case "pending":
		info.Status = string(model.TaskStatusQueued)
		info.Progress = "50%"
	case "processing":
		info.Status = string(model.TaskStatusInProgress)
		if data.Progress > 0 {
			info.Progress = fmt.Sprintf("%d%%", data.Progress)
		}
	case "completed":
		info.Status = string(model.TaskStatusSuccess)
		info.Progress = "100%"
		info.Url = primaryResultURL(data.Result)
	case "failed":
		info.Status = string(model.TaskStatusFailure)
		info.Progress = "100%"
		if data.Error != nil {
			info.Reason = data.Error.Message
		}
	case "unknown":
		info.Status = string(model.TaskStatusInProgress)
		if data.Progress > 0 {
			info.Progress = fmt.Sprintf("%d%%", data.Progress)
		}
	default:
		info.Status = string(model.TaskStatusInProgress)
		if data.Progress > 0 {
			info.Progress = fmt.Sprintf("%d%%", data.Progress)
		}
	}

	return info, nil
}

func primaryResultURL(result *fetchResult) string {
	if result == nil || len(result.Music) == 0 {
		return ""
	}
	item := result.Music[0]
	if item.AudioURL != "" {
		return item.AudioURL
	}
	if item.VideoURL != "" {
		return item.VideoURL
	}
	return item.FileURL
}
