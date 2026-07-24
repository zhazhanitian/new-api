package dto

import (
	"encoding/json"
)

type TaskError struct {
	Code       string `json:"code"`
	Message    string `json:"message"`
	Data       any    `json:"data"`
	StatusCode int    `json:"-"`
	LocalError bool   `json:"-"`
	Error      error  `json:"-"`
}

type TaskData interface {
	SunoDataResponse | []SunoDataResponse | string | any
}

const TaskSuccessCode = "success"

type TaskResponse[T TaskData] struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

func (t *TaskResponse[T]) IsSuccess() bool {
	return t.Code == TaskSuccessCode
}

type TaskDto struct {
	ID         int64           `json:"id"`
	CreatedAt  int64           `json:"created_at"`
	UpdatedAt  int64           `json:"updated_at"`
	TaskID     string          `json:"task_id"`
	Platform   string          `json:"platform"`
	UserId     int             `json:"user_id"`
	Group      string          `json:"group"`
	ChannelId  int             `json:"channel_id"`
	Amount     string          `json:"amount"`   // 人民币金额，6位小数，如 "0.073000"
	Action     string          `json:"action"`
	Status     string          `json:"status"`
	FailReason string          `json:"fail_reason"`
	ResultURL  string          `json:"result_url,omitempty"` // 任务结果 URL（视频地址等）
	SubmitTime int64           `json:"submit_time"`
	StartTime  int64           `json:"start_time"`
	FinishTime int64           `json:"finish_time"`
	Progress   string          `json:"progress"`
	Properties any             `json:"properties"`
	Username   string          `json:"username,omitempty"`
	Data       json.RawMessage `json:"data"`
}

type FetchReq struct {
	IDs []string `json:"ids"`
}

// ImageTaskDto 是 /v2/image-tasks/:task_id 查询接口的响应体（干净版，面向外部调用方）
type ImageTaskDto struct {
	TaskID     string         `json:"task_id"`
	Status     string         `json:"status"`
	Progress   string         `json:"progress"`
	FailReason string         `json:"fail_reason,omitempty"`
	Model      string         `json:"model,omitempty"`
	Images     []ImageItem    `json:"images"`
	Amount     string         `json:"amount"`    // 人民币金额，6位小数字符串，如 "0.073000"
	Usage      ImageTaskUsage `json:"usage"`     // 标准化 token 用量，无数据时全为 0
	CreatedAt  int64          `json:"created_at"`
	SubmitTime int64          `json:"submit_time"`
	FinishTime int64          `json:"finish_time"`
}

// ImageItem 图片结果项，url 和 b64_json 互斥，有哪个填哪个
type ImageItem struct {
	URL     string `json:"url,omitempty"`
	B64Json string `json:"b64_json,omitempty"`
}

// ImageTaskUsage 标准化的 token 用量结构，兼容各渠道差异
// 按次计费渠道（如腾讯 VOD）三个字段均为 0
type ImageTaskUsage struct {
	InputTokens  int `json:"input_tokens"`
	OutputTokens int `json:"output_tokens"`
	TotalTokens  int `json:"total_tokens"`
}
