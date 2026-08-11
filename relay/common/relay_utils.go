package common

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/QuantumNous/new-api/common"
	"github.com/QuantumNous/new-api/constant"
	"github.com/QuantumNous/new-api/dto"

	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
)

type HasPrompt interface {
	GetPrompt() string
}

type HasImage interface {
	HasImage() bool
}

func GetFullRequestURL(baseURL string, requestURL string, channelType int) string {
	fullRequestURL := fmt.Sprintf("%s%s", baseURL, requestURL)

	if strings.HasPrefix(baseURL, "https://gateway.ai.cloudflare.com") {
		switch channelType {
		case constant.ChannelTypeOpenAI:
			fullRequestURL = fmt.Sprintf("%s%s", baseURL, strings.TrimPrefix(requestURL, "/v1"))
		case constant.ChannelTypeAzure:
			fullRequestURL = fmt.Sprintf("%s%s", baseURL, strings.TrimPrefix(requestURL, "/openai/deployments"))
		}
	}
	return fullRequestURL
}

func GetAPIVersion(c *gin.Context) string {
	query := c.Request.URL.Query()
	apiVersion := query.Get("api-version")
	if apiVersion == "" {
		apiVersion = c.GetString("api_version")
	}
	return apiVersion
}

func createTaskError(err error, code string, statusCode int, localError bool) *dto.TaskError {
	return &dto.TaskError{
		Code:       code,
		Message:    err.Error(),
		StatusCode: statusCode,
		LocalError: localError,
		Error:      err,
	}
}

func storeTaskRequest(c *gin.Context, info *RelayInfo, action string, requestObj TaskSubmitReq) {
	info.Action = action
	c.Set("task_request", requestObj)
}
func GetTaskRequest(c *gin.Context) (TaskSubmitReq, error) {
	v, exists := c.Get("task_request")
	if !exists {
		return TaskSubmitReq{}, fmt.Errorf("request not found in context")
	}
	req, ok := v.(TaskSubmitReq)
	if !ok {
		return TaskSubmitReq{}, fmt.Errorf("invalid task request type")
	}
	return req, nil
}

func validatePrompt(prompt string) *dto.TaskError {
	if strings.TrimSpace(prompt) == "" {
		return createTaskError(fmt.Errorf("prompt is required"), "invalid_request", http.StatusBadRequest, true)
	}
	return nil
}

func validateMultipartTaskRequest(c *gin.Context, info *RelayInfo, action string) (TaskSubmitReq, error) {
	var req TaskSubmitReq
	if _, err := c.MultipartForm(); err != nil {
		return req, err
	}

	formData := c.Request.PostForm
	req = TaskSubmitReq{
		Prompt:   formData.Get("prompt"),
		Model:    formData.Get("model"),
		Mode:     formData.Get("mode"),
		Image:    formData.Get("image"),
		Size:     formData.Get("size"),
		Metadata: make(map[string]interface{}),
	}

	if durationStr := formData.Get("seconds"); durationStr != "" {
		if duration, err := strconv.Atoi(durationStr); err == nil {
			req.Duration = duration
		}
	}

	if images := formData["images"]; len(images) > 0 {
		req.Images = images
	}

	for key, values := range formData {
		if len(values) > 0 && !isKnownTaskField(key) {
			if intVal, err := strconv.Atoi(values[0]); err == nil {
				req.Metadata[key] = intVal
			} else if floatVal, err := strconv.ParseFloat(values[0], 64); err == nil {
				req.Metadata[key] = floatVal
			} else {
				req.Metadata[key] = values[0]
			}
		}
	}
	return req, nil
}

func ValidateMultipartDirect(c *gin.Context, info *RelayInfo) *dto.TaskError {
	var prompt string
	var model string
	var seconds int
	var size string
	var hasInputReference bool

	var req TaskSubmitReq
	if err := common.UnmarshalBodyReusable(c, &req); err != nil {
		return createTaskError(err, "invalid_json", http.StatusBadRequest, true)
	}

	prompt = req.Prompt
	model = req.Model
	size = req.Size
	seconds, _ = strconv.Atoi(req.Seconds)
	if seconds == 0 {
		seconds = req.Duration
	}
	if req.InputReference != "" {
		req.Images = []string{req.InputReference}
	}

	if strings.TrimSpace(req.Model) == "" {
		return createTaskError(fmt.Errorf("model field is required"), "missing_model", http.StatusBadRequest, true)
	}

	if req.HasImage() {
		hasInputReference = true
	}

	if taskErr := validatePrompt(prompt); taskErr != nil {
		return taskErr
	}

	action := constant.TaskActionTextGenerate
	if hasInputReference {
		action = constant.TaskActionGenerate
	}
	if strings.HasPrefix(model, "sora-2") {

		if size == "" {
			size = "720x1280"
		}

		if seconds <= 0 {
			seconds = 4
		}

		if model == "sora-2" && !lo.Contains([]string{"720x1280", "1280x720"}, size) {
			return createTaskError(fmt.Errorf("sora-2 size is invalid"), "invalid_size", http.StatusBadRequest, true)
		}
		if model == "sora-2-pro" && !lo.Contains([]string{"720x1280", "1280x720", "1792x1024", "1024x1792"}, size) {
			return createTaskError(fmt.Errorf("sora-2 size is invalid"), "invalid_size", http.StatusBadRequest, true)
		}
		// OtherRatios 已移到 Sora adaptor 的 EstimateBilling 中设置
	}

	storeTaskRequest(c, info, action, req)

	return nil
}

func isKnownTaskField(field string) bool {
	knownFields := map[string]bool{
		"prompt":          true,
		"model":           true,
		"mode":            true,
		"image":           true,
		"images":          true,
		"size":            true,
		"duration":        true,
		"input_reference": true, // Sora 特有字段
	}
	return knownFields[field]
}

func ValidateBasicTaskRequest(c *gin.Context, info *RelayInfo, action string) *dto.TaskError {
	// 已由前置步骤（ValidateImageTaskRequest）解析并存入 context，跳过重复解析
	if _, exists := c.Get("task_request"); exists {
		if info.Action == "" {
			info.Action = action
		}
		return nil
	}
	var err error
	contentType := c.GetHeader("Content-Type")
	var req TaskSubmitReq
	if strings.HasPrefix(contentType, "multipart/form-data") {
		req, err = validateMultipartTaskRequest(c, info, action)
		if err != nil {
			return createTaskError(err, "invalid_multipart_form", http.StatusBadRequest, true)
		}
	}
	// 为了metadata字段的兼容性，统一UnmarshalBodyReusable
	if err := common.UnmarshalBodyReusable(c, &req); err != nil {
		return createTaskError(err, "invalid_request", http.StatusBadRequest, true)
	}

	if taskErr := validatePrompt(req.Prompt); taskErr != nil {
		return taskErr
	}

	if len(req.Images) == 0 && strings.TrimSpace(req.Image) != "" {
		// 兼容单图上传
		req.Images = []string{req.Image}
	}

	storeTaskRequest(c, info, action, req)
	return nil
}

// Validate3DTaskRequest 专用于 POST /v1/3d/generations。
// 解析统一 3D 请求体（{"model":"...","metadata":{...}}），校验 model 字段非空，
// 并将解析结果以 "3d_request" 键存入 gin 上下文，供适配器通过 Get3DTaskRequest 取出。
func Validate3DTaskRequest(c *gin.Context, info *RelayInfo) *dto.TaskError {
	var req dto.OpenAI3DRequest
	if err := common.UnmarshalBodyReusable(c, &req); err != nil {
		return createTaskError(err, "invalid_request", http.StatusBadRequest, true)
	}
	if strings.TrimSpace(req.Model) == "" {
		return createTaskError(fmt.Errorf("model is required"), "invalid_request", http.StatusBadRequest, true)
	}
	info.Action = constant.TaskAction3DGenerate
	c.Set("3d_request", req)
	return nil
}

// Get3DTaskRequest 从 gin 上下文中取出 Validate3DTaskRequest 存入的 3D 请求。
func Get3DTaskRequest(c *gin.Context) (dto.OpenAI3DRequest, error) {
	v, exists := c.Get("3d_request")
	if !exists {
		return dto.OpenAI3DRequest{}, fmt.Errorf("3d request not found in context")
	}
	req, ok := v.(dto.OpenAI3DRequest)
	if !ok {
		return dto.OpenAI3DRequest{}, fmt.Errorf("invalid 3d request type in context")
	}
	return req, nil
}

// ValidateImageTaskRequest 专用于 POST /v2/image-tasks（ImageRequest → TaskSubmitReq）。
// 解析 OpenAI 兼容的 ImageRequest，转换为内部 TaskSubmitReq 后存入 context；
// 下游 adaptor 通过 GetTaskRequest 取出，无需感知入口格式。
func ValidateImageTaskRequest(c *gin.Context, info *RelayInfo) *dto.TaskError {
	var imgReq dto.ImageRequest
	if err := common.UnmarshalBodyReusable(c, &imgReq); err != nil {
		return createTaskError(err, "invalid_request", http.StatusBadRequest, true)
	}
	if strings.TrimSpace(imgReq.Prompt) == "" {
		return createTaskError(fmt.Errorf("prompt is required"), "invalid_request", http.StatusBadRequest, true)
	}

	req := TaskSubmitReq{
		Model:          imgReq.Model,
		Prompt:         imgReq.Prompt,
		Size:           imgReq.Size,
		Quality:        imgReq.Quality,
		ResponseFormat: imgReq.ResponseFormat,
		Metadata:       make(map[string]interface{}),
	}
	if imgReq.N != nil {
		req.N = int(*imgReq.N)
	}
	// style: json.RawMessage → string
	if len(imgReq.Style) > 0 {
		var s string
		if json.Unmarshal(imgReq.Style, &s) == nil {
			req.Style = s
		}
	}
	// output_format: json.RawMessage → string
	if len(imgReq.OutputFormat) > 0 {
		var s string
		if json.Unmarshal(imgReq.OutputFormat, &s) == nil {
			req.OutputFormat = s
		}
	}
	// output_compression: json.RawMessage → int
	if len(imgReq.OutputCompression) > 0 {
		var v int
		if json.Unmarshal(imgReq.OutputCompression, &v) == nil {
			req.OutputCompression = v
		}
	}
	// images: json.RawMessage → []string（静默失败，不影响主流程）
	if len(imgReq.Images) > 0 {
		var urls []string
		if json.Unmarshal(imgReq.Images, &urls) == nil {
			req.Images = urls
		}
	}
	// image → 合并进 Images，兼容字符串和数组两种格式：
	//   字符串："image": "https://example.com/img.jpg"
	//   数组：  "image": ["https://example.com/img.jpg"] 或多张
	if len(imgReq.Image) > 0 && len(req.Images) == 0 {
		var s string
		if json.Unmarshal(imgReq.Image, &s) == nil && s != "" {
			req.Images = []string{s}
		} else {
			var urls []string
			if json.Unmarshal(imgReq.Image, &urls) == nil && len(urls) > 0 {
				req.Images = urls
			}
		}
	}
	// extra_fields → Metadata（渠道私有扩展参数通道，与同步接口 ConvertImageRequest 读取 ExtraFields 保持对称）
	if len(imgReq.ExtraFields) > 0 {
		var extraMap map[string]interface{}
		if json.Unmarshal(imgReq.ExtraFields, &extraMap) == nil {
			for k, v := range extraMap {
				req.Metadata[k] = v
			}
		}
	}

	storeTaskRequest(c, info, constant.TaskActionImageGenerate, req)
	return nil
}
