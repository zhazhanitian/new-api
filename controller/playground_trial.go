package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/QuantumNous/new-api/common"
	"github.com/QuantumNous/new-api/middleware"
	"github.com/QuantumNous/new-api/model"
	relaycommon "github.com/QuantumNous/new-api/relay/common"
	"github.com/QuantumNous/new-api/trial"
	"github.com/QuantumNous/new-api/types"

	"github.com/gin-gonic/gin"
)

type trialCaptureWriter struct {
	gin.ResponseWriter
	body bytes.Buffer
}

func (w *trialCaptureWriter) Write(data []byte) (int, error) {
	w.body.Write(data)
	return w.ResponseWriter.Write(data)
}

func (w *trialCaptureWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func (w *trialCaptureWriter) Flush() {
	if flusher, ok := w.ResponseWriter.(http.Flusher); ok {
		flusher.Flush()
	}
}

func PlaygroundTrial(c *gin.Context) {
	if c.GetBool("use_access_token") {
		c.JSON(http.StatusForbidden, gin.H{"error": gin.H{"message": "暂不支持使用 access token"}})
		return
	}
	kind := trial.APIKind(c.GetString(trial.CtxAPIKind))
	proofID := c.GetInt(trial.CtxProofID)
	if proofID > 0 {
		c.Header("X-Model-Trial-Proof", strconv.Itoa(proofID))
	}
	userId := c.GetInt("id")
	userCache, err := model.GetUserCache(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"message": err.Error()}})
		return
	}
	userCache.WriteContext(c)

	format := types.RelayFormatOpenAI
	switch kind {
	case trial.APIKindResponses:
		format = types.RelayFormatOpenAIResponses
	case trial.APIKindImageGeneration:
		format = types.RelayFormatOpenAIImage
	}
	relayInfo, err := relaycommon.GenRelayInfo(c, format, nil, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": gin.H{"message": err.Error()}})
		return
	}
	tempToken := &model.Token{
		UserId: userId,
		Name:   fmt.Sprintf("playground-%s", relayInfo.UsingGroup),
		Group:  relayInfo.UsingGroup,
	}
	if err := middleware.SetupContextForToken(c, tempToken); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"message": err.Error()}})
		return
	}

	capture := &trialCaptureWriter{ResponseWriter: c.Writer}
	c.Writer = capture
	if trial.IsAsync(kind) {
		RelayTask(c)
	} else {
		Relay(c, format)
	}
	if status := capture.Status(); status == 0 || status >= http.StatusBadRequest {
		return
	}
	if proofID == 0 {
		return
	}
	if trial.IsAsync(kind) {
		if taskID := extractTrialTaskID(capture.body.Bytes()); taskID != "" {
			_ = model.AttachTrialProofTask(proofID, taskID)
		}
		return
	}
	_ = model.MarkTrialProofVerified(proofID)
}

func PlaygroundTrialTask(c *gin.Context) {
	taskID := c.Param("task_id")
	userId := c.GetInt("id")
	task, exists, err := model.GetByTaskId(userId, taskID)
	if err != nil {
		common.ApiError(c, err)
		return
	}
	if !exists || task == nil {
		common.ApiErrorMsg(c, "任务不存在")
		return
	}
	proofID, _ := strconv.Atoi(c.Query("proof_id"))
	if task.Status == model.TaskStatusSuccess && proofID > 0 {
		proof, proofErr := model.GetTrialProof(proofID, userId)
		if proofErr == nil && proof != nil && proof.TaskID == task.TaskID {
			_ = model.MarkTrialProofVerified(proof.Id)
		}
	}
	common.ApiSuccess(c, gin.H{
		"task_id":     task.TaskID,
		"status":      task.Status,
		"progress":    task.Progress,
		"fail_reason": task.FailReason,
		"data":        json.RawMessage(task.Data),
		"result_url":  task.PrivateData.ResultURL,
	})
}

func extractTrialTaskID(raw []byte) string {
	if len(bytes.TrimSpace(raw)) == 0 {
		return ""
	}
	var payload map[string]any
	if err := json.Unmarshal(raw, &payload); err != nil {
		return ""
	}
	if id := trialString(payload["task_id"]); id != "" {
		return id
	}
	switch data := payload["data"].(type) {
	case map[string]any:
		if id := trialString(data["task_id"]); id != "" {
			return id
		}
		if id := trialString(data["id"]); id != "" {
			return id
		}
	case []any:
		if len(data) == 0 {
			return ""
		}
		if item, ok := data[0].(map[string]any); ok {
			return trialString(item["task_id"])
		}
	}
	return ""
}

func trialString(value any) string {
	text, ok := value.(string)
	if !ok || text == "" || text == "success" {
		return ""
	}
	return text
}

var errTrialNotVerified = errors.New("请先测试通过后再保存示例")
