package service

import (
	"bytes"
	"encoding/json"
	"io"
	"mime"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/QuantumNous/new-api/common"
	"github.com/QuantumNous/new-api/constant"
	"github.com/QuantumNous/new-api/dto"
	"github.com/QuantumNous/new-api/logger"
	"github.com/QuantumNous/new-api/model"
	relaycommon "github.com/QuantumNous/new-api/relay/common"
	"github.com/gin-gonic/gin"
)

const maxRequestLogBytes = 128 * 1024

var secretJSONKey = regexp.MustCompile(`(?i)(api[_-]?key|authorization|secret|access[_-]?token|access[_-]?key)$`)

// SnapshotUserRequest 读取调用方原始请求，供任务日志保存。
// multipart 只保留文本字段和文件名、大小，不把文件内容写进日志。
func SnapshotUserRequest(c *gin.Context) []byte {
	if c == nil || c.Request == nil {
		return nil
	}
	contentType := c.GetHeader("Content-Type")
	mediaType, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		mediaType = contentType
	}
	switch {
	case strings.HasPrefix(mediaType, "multipart/"):
		return snapshotMultipart(c)
	case mediaType == "application/x-www-form-urlencoded":
		body, bodyErr := readBodyBytes(c)
		if bodyErr != nil || len(body) == 0 {
			return nil
		}
		return snapshotForm(body)
	default:
		body, bodyErr := readBodyBytes(c)
		if bodyErr != nil || len(body) == 0 {
			return nil
		}
		return body
	}
}

func readBodyBytes(c *gin.Context) ([]byte, error) {
	storage, err := common.GetBodyStorage(c)
	if err != nil {
		return nil, err
	}
	body, err := storage.Bytes()
	if err != nil {
		return nil, err
	}
	return append([]byte(nil), body...), nil
}

func snapshotMultipart(c *gin.Context) []byte {
	form, err := c.MultipartForm()
	if err != nil || form == nil {
		return nil
	}
	out := make(map[string]any, len(form.Value)+1)
	for key, values := range form.Value {
		if len(values) == 1 {
			out[key] = values[0]
			continue
		}
		out[key] = values
	}
	if len(form.File) > 0 {
		files := make(map[string]any, len(form.File))
		for key, headers := range form.File {
			metas := make([]map[string]any, 0, len(headers))
			for _, header := range headers {
				if header == nil {
					continue
				}
				metas = append(metas, map[string]any{
					"filename": header.Filename,
					"size":     header.Size,
				})
			}
			files[key] = metas
		}
		out["_files"] = files
	}
	payload, err := json.Marshal(out)
	if err != nil {
		return nil
	}
	return payload
}

func snapshotForm(body []byte) []byte {
	values, err := url.ParseQuery(string(body))
	if err != nil {
		return body
	}
	out := make(map[string]any, len(values))
	for key, items := range values {
		if len(items) == 1 {
			out[key] = items[0]
			continue
		}
		out[key] = items
	}
	payload, err := json.Marshal(out)
	if err != nil {
		return body
	}
	return payload
}

// CaptureReader 读完请求体，并还给调用方一个可再次读取的副本。
// 返回的字节是发给上游的原文，日志里再单独做脱敏。
func CaptureReader(r io.Reader) ([]byte, io.Reader, error) {
	if r == nil {
		return nil, nil, nil
	}
	body, err := io.ReadAll(r)
	if err != nil {
		return nil, nil, err
	}
	return body, bytes.NewReader(body), nil
}

// AttachTaskRequestLogs 把用户请求、上游请求和上游响应写到任务上。空内容不写。
func AttachTaskRequestLogs(task *model.Task, userBody, upstreamBody, upstreamResp []byte) {
	if task == nil {
		return
	}
	if sanitized := SanitizeRequestLog(userBody); len(sanitized) > 0 {
		task.UserRequest = sanitized
	}
	if sanitized := SanitizeRequestLog(upstreamBody); len(sanitized) > 0 {
		task.UpstreamRequest = sanitized
	}
	if sanitized := SanitizeRequestLog(upstreamResp); len(sanitized) > 0 {
		task.UpstreamResponse = sanitized
	}
}

// RecordFailedTaskSubmit 在任务提交失败时补一条失败记录。
// 渠道还没选定的错误不记，避免把选路失败刷进任务日志。
func RecordFailedTaskSubmit(c *gin.Context, info *relaycommon.RelayInfo, taskErr *dto.TaskError) {
	if c == nil || info == nil || info.ChannelMeta == nil || taskErr == nil || info.UserId == 0 {
		return
	}
	platform := constant.TaskPlatform(strconv.Itoa(info.ChannelType))
	task := model.InitTask(platform, info)
	task.Action = info.Action
	task.Status = model.TaskStatusFailure
	task.Progress = "0%"
	task.FinishTime = time.Now().Unix()
	task.Quota = 0
	task.FailReason = taskErr.Message
	if task.FailReason == "" && taskErr.Error != nil {
		task.FailReason = taskErr.Error.Error()
	}
	AttachTaskRequestLogs(task, info.UserRequestBody, info.UpstreamRequestBody, info.UpstreamResponseBody)
	if err := task.Insert(); err != nil {
		logger.LogWarn(c, "RecordFailedTaskSubmit: "+err.Error())
	}
}

// SanitizeRequestLog 去掉密钥和图片二进制，保留提示词、参数和图片地址。
// 非 JSON 会包成 {"raw": "..."}。内容为空时返回 nil。
func SanitizeRequestLog(raw []byte) json.RawMessage {
	raw = bytes.TrimSpace(raw)
	if len(raw) == 0 {
		return nil
	}
	var value any
	if err := json.Unmarshal(raw, &value); err != nil {
		return wrapRawLog(raw)
	}
	encoded, err := json.Marshal(redactLogValue("", value))
	if err != nil {
		return wrapRawLog(raw)
	}
	if len(encoded) > maxRequestLogBytes {
		return wrapRawLog(encoded)
	}
	return encoded
}

func redactLogValue(key string, value any) any {
	switch typed := value.(type) {
	case map[string]any:
		out := make(map[string]any, len(typed))
		for itemKey, itemValue := range typed {
			if secretJSONKey.MatchString(itemKey) {
				out[itemKey] = "[redacted]"
				continue
			}
			out[itemKey] = redactLogValue(itemKey, itemValue)
		}
		return out
	case []any:
		out := make([]any, len(typed))
		for i, item := range typed {
			out[i] = redactLogValue(key, item)
		}
		return out
	case string:
		if isBinaryLogString(typed) {
			return "[omitted binary, " + strconv.Itoa(len(typed)) + " chars]"
		}
		return typed
	default:
		return value
	}
}

func isBinaryLogString(value string) bool {
	if strings.HasPrefix(value, "data:") && len(value) > 128 {
		return true
	}
	if len(value) < 1024 {
		return false
	}
	limit := len(value)
	if limit > 256 {
		limit = 256
	}
	matched := 0
	for i := 0; i < limit; i++ {
		ch := value[i]
		switch {
		case ch >= 'A' && ch <= 'Z', ch >= 'a' && ch <= 'z', ch >= '0' && ch <= '9', ch == '+', ch == '/', ch == '=', ch == '-', ch == '_':
			matched++
		}
	}
	return matched*100/limit >= 98
}

func wrapRawLog(raw []byte) json.RawMessage {
	preview := raw
	truncated := false
	if len(preview) > 4000 {
		preview = preview[:4000]
		truncated = true
	}
	if !utf8.Valid(preview) {
		preview = bytes.ToValidUTF8(preview, []byte("?"))
	}
	payload, err := json.Marshal(map[string]any{
		"raw":       string(preview),
		"truncated": truncated,
		"bytes":     len(raw),
	})
	if err != nil {
		return nil
	}
	return payload
}
