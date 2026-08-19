package apimart_suno

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/QuantumNous/new-api/common"
	"github.com/QuantumNous/new-api/dto"
	"github.com/QuantumNous/new-api/logger"
	"github.com/QuantumNous/new-api/model"
	relaycommon "github.com/QuantumNous/new-api/relay/common"
)

// ParseTaskResult implements channel.TaskAdaptor and parses the APIMart single-task
// query response from GET /v1/music/tasks/:task_id.
func ParseTaskResult(respBody []byte) (*relaycommon.TaskInfo, error) {
	logger.LogInfo(context.Background(), fmt.Sprintf("[apimart-suno] fetch task response: body=%s", string(respBody)))

	var resp dto.APIMartSunoFetchResponse
	if err := common.Unmarshal(respBody, &resp); err != nil {
		return nil, fmt.Errorf("unmarshal APIMart fetch response: %w", err)
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

	info := &relaycommon.TaskInfo{
		TaskID: resp.Data.ID,
	}

	// Map APIMart status to platform status
	switch resp.Data.Status {
	case "submitted":
		info.Status = string(model.TaskStatusSubmitted)
		info.Progress = "10%"
	case "pending":
		info.Status = string(model.TaskStatusQueued)
		info.Progress = "50%"
	case "completed":
		info.Status = string(model.TaskStatusSuccess)
		info.Progress = "100%"
		// Extract primary audio URL for quick access
		if len(resp.Data.Result.Music) > 0 {
			info.Url = resp.Data.Result.Music[0].AudioURL
		} else if resp.Data.Result.UpsampledTags != "" {
			// upsampleTags result stored as URL-like string for compatibility
			info.Url = resp.Data.Result.UpsampledTags
		}
	case "failed":
		info.Status = string(model.TaskStatusFailure)
		info.Progress = "100%"
		if resp.Data.Error != nil {
			info.Reason = resp.Data.Error.Message
		}
	default:
		// Unknown status: treat as in-progress to keep polling
		info.Status = string(model.TaskStatusInProgress)
		if resp.Data.Progress > 0 {
			info.Progress = fmt.Sprintf("%d%%", resp.Data.Progress)
		}
	}

	return info, nil
}

// buildTaskData serialises the raw fetch response for storage in task.Data.
// This preserves the full upstream response for auditing and future result retrieval.
func buildTaskData(respBody []byte) json.RawMessage {
	return json.RawMessage(respBody)
}
