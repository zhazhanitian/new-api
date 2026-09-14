package apimart_suno

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"slices"
	"strings"
	"time"

	"github.com/QuantumNous/new-api/common"
	"github.com/QuantumNous/new-api/dto"
	"github.com/QuantumNous/new-api/logger"
	"github.com/QuantumNous/new-api/model"
	relaycommon "github.com/QuantumNous/new-api/relay/common"
	"github.com/QuantumNous/new-api/service"
	"github.com/gin-gonic/gin"
)

// apimartSunoHTTPClient 直连 APIMart，不读取环境代理变量。
// 渠道级代理（channel.Proxy）在 DoRequest/FetchTask 中单独处理。
// 不设 Proxy 字段是为了绕过 http_proxy/HTTP_PROXY 等环境变量——
// 开发机上的 VPN/代理客户端（如 Clash）对 HTTPS CONNECT 隧道处理不当，会导致 EOF。
var apimartSunoHTTPClient = &http.Client{
	Transport: &http.Transport{
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 10,
		IdleConnTimeout:     90 * time.Second,
	},
}

// TaskAdaptor implements channel.TaskAdaptor for the APIMart Suno API.
type TaskAdaptor struct {
	ChannelType int
}

// ── Lifecycle ──────────────────────────────────────────────────────────────────

func (a *TaskAdaptor) Init(info *relaycommon.RelayInfo) {
	a.ChannelType = info.ChannelType
}

// ── Billing (delegated to billing.go) ─────────────────────────────────────────

func (a *TaskAdaptor) EstimateBilling(c *gin.Context, info *relaycommon.RelayInfo) map[string]float64 {
	return EstimateBilling(c, info)
}

func (a *TaskAdaptor) AdjustBillingOnSubmit(info *relaycommon.RelayInfo, taskData []byte) map[string]float64 {
	return AdjustBillingOnSubmit(info, taskData)
}

func (a *TaskAdaptor) AdjustBillingOnComplete(task *model.Task, taskResult *relaycommon.TaskInfo) int {
	return AdjustBillingOnComplete(task, taskResult)
}

// ── Validation ────────────────────────────────────────────────────────────────

// ValidateRequestAndSetAction parses the request body, looks up the tool definition
// for the model ID, validates required fields and version, then stores the parsed
// request in the gin context for BuildRequestBody.
func (a *TaskAdaptor) ValidateRequestAndSetAction(c *gin.Context, info *relaycommon.RelayInfo) *dto.TaskError {
	var req dto.APIMartSunoRequest
	if err := common.UnmarshalBodyReusable(c, &req); err != nil {
		return service.TaskErrorWrapperLocal(err, "invalid_request", http.StatusBadRequest)
	}

	modelID := info.OriginModelName
	def, ok := GetToolDef(modelID)
	if !ok {
		return service.TaskErrorWrapperLocal(
			fmt.Errorf("unknown suno model: %s", modelID),
			"unknown_model", http.StatusBadRequest,
		)
	}

	// version 与 custom_model_id 互斥（所有工具）
	if req.Version != "" && req.CustomModelID != "" {
		return service.TaskErrorWrapperLocal(
			fmt.Errorf("version and custom_model_id are mutually exclusive; provide one or the other"),
			"invalid_request", http.StatusBadRequest,
		)
	}

	// suno-music 专项：version 或 custom_model_id 至少提供一个
	if modelID == "suno-music" && req.Version == "" && req.CustomModelID == "" {
		return service.TaskErrorWrapperLocal(
			fmt.Errorf("suno-music requires either version (v6/v6-wild/v6-mini) or custom_model_id"),
			"version_required", http.StatusBadRequest,
		)
	}

	// Version dimension validation
	if req.Version != "" && def.SupportedVersions == nil {
		// Tool has no version dimension; ignore the field silently
		req.Version = ""
	}
	if req.Version != "" && def.SupportedVersions != nil {
		if !slices.Contains(def.SupportedVersions, req.Version) {
			return service.TaskErrorWrapperLocal(
				fmt.Errorf("version %q is not supported for model %s; supported: %v",
					req.Version, modelID, def.SupportedVersions),
				"unsupported_version", http.StatusBadRequest,
			)
		}
	}

	// Required field validation
	if err := validateRequiredFields(&req, def); err != nil {
		return service.TaskErrorWrapperLocal(err, "invalid_request", http.StatusBadRequest)
	}

	// mashup: exactly 2 task_ids
	if def.UsesTaskIDs {
		if len(req.TaskIDs) != 2 {
			return service.TaskErrorWrapperLocal(
				fmt.Errorf("mashup requires exactly 2 task_ids, got %d", len(req.TaskIDs)),
				"invalid_request", http.StatusBadRequest,
			)
		}
	}

	// inspo: 1-4 audio_urls
	if modelID == "suno-inspo" {
		if len(req.AudioURLs) < 1 || len(req.AudioURLs) > 4 {
			return service.TaskErrorWrapperLocal(
				fmt.Errorf("inspo requires 1-4 audio_urls, got %d", len(req.AudioURLs)),
				"invalid_request", http.StatusBadRequest,
			)
		}
	}

	// create-model: 6-24 audio_urls
	if modelID == "suno-create-model" {
		if len(req.AudioURLs) < 6 || len(req.AudioURLs) > 24 {
			return service.TaskErrorWrapperLocal(
				fmt.Errorf("create-model requires 6-24 audio_urls, got %d", len(req.AudioURLs)),
				"invalid_request", http.StatusBadRequest,
			)
		}
	}

	// download: format / formats 至少一个非空
	if modelID == "suno-download" {
		if len(req.Formats) == 0 && strings.TrimSpace(req.Format) == "" {
			return service.TaskErrorWrapperLocal(
				fmt.Errorf("suno-download requires at least one format (use 'formats' or 'format')"),
				"invalid_request", http.StatusBadRequest,
			)
		}
	}

	// cover: when custom=false (or inferred false), gpt_description is required
	if modelID == "suno-cover" {
		customVal := req.Custom != nil && *req.Custom
		hasGPTDesc := strings.TrimSpace(req.GptDescription) != ""
		hasPromptOrTags := strings.TrimSpace(req.Prompt) != "" || strings.TrimSpace(req.Tags) != ""
		if !customVal && !hasGPTDesc && !hasPromptOrTags {
			return service.TaskErrorWrapperLocal(
				fmt.Errorf("cover song requires gpt_description (or prompt/tags for custom mode)"),
				"invalid_request", http.StatusBadRequest,
			)
		}
	}

	// sounds: bpm range 1-300
	if modelID == "suno-sounds" && req.BPM != nil {
		if *req.BPM < 1 || *req.BPM > 300 {
			return service.TaskErrorWrapperLocal(
				fmt.Errorf("bpm must be between 1 and 300, got %d", *req.BPM),
				"invalid_request", http.StatusBadRequest,
			)
		}
	}

	// adjust_speed: 0.25-4
	if modelID == "suno-adjust-speed" && req.Speed != nil {
		if *req.Speed < 0.25 || *req.Speed > 4.0 {
			return service.TaskErrorWrapperLocal(
				fmt.Errorf("speed must be between 0.25 and 4, got %g", *req.Speed),
				"invalid_request", http.StatusBadRequest,
			)
		}
	}

	// style_weight / weirdness_constraint / weirdness / audio_weight: 0.00-1.00
	if err := validateRatio("style_weight", req.StyleWeight); err != nil {
		return service.TaskErrorWrapperLocal(err, "invalid_request", http.StatusBadRequest)
	}
	if err := validateRatio("weirdness_constraint", req.WeirdnessConstraint); err != nil {
		return service.TaskErrorWrapperLocal(err, "invalid_request", http.StatusBadRequest)
	}
	if err := validateRatio("weirdness", req.Weirdness); err != nil {
		return service.TaskErrorWrapperLocal(err, "invalid_request", http.StatusBadRequest)
	}
	if err := validateRatio("audio_weight", req.AudioWeight); err != nil {
		return service.TaskErrorWrapperLocal(err, "invalid_request", http.StatusBadRequest)
	}

	info.Action = modelID
	c.Set("task_request", &req)
	return nil
}

// validateRequiredFields checks that all fields listed in def.RequiredFields are present.
func validateRequiredFields(req *dto.APIMartSunoRequest, def ToolDef) error {
	for _, field := range def.RequiredFields {
		switch field {
		case "task_id":
			if strings.TrimSpace(req.TaskID) == "" {
				return fmt.Errorf("task_id is required")
			}
		case "task_ids":
			if len(req.TaskIDs) == 0 {
				return fmt.Errorf("task_ids is required")
			}
		case "audio_urls":
			if len(req.AudioURLs) == 0 {
				return fmt.Errorf("audio_urls is required")
			}
		case "audio_url":
			if strings.TrimSpace(req.AudioURL) == "" {
				return fmt.Errorf("audio_url is required")
			}
		case "audio_file_path":
			if strings.TrimSpace(req.AudioFilePath) == "" {
				return fmt.Errorf("audio_file_path is required")
			}
		case "tags":
			if strings.TrimSpace(req.Tags) == "" {
				return fmt.Errorf("tags is required")
			}
		case "name":
			if strings.TrimSpace(req.Name) == "" {
				return fmt.Errorf("name is required")
			}
		case "prompt":
			if strings.TrimSpace(req.Prompt) == "" {
				return fmt.Errorf("prompt is required")
			}
		case "version":
			if strings.TrimSpace(req.Version) == "" {
				return fmt.Errorf("version is required")
			}
		case "start_s":
			if req.StartS == nil {
				return fmt.Errorf("start_s is required")
			}
		case "end_s":
			if req.EndS == nil {
				return fmt.Errorf("end_s is required")
			}
		case "continue_at":
			if req.ContinueAt == nil {
				return fmt.Errorf("continue_at is required")
			}
		case "duration_s":
			if req.DurationS == nil {
				return fmt.Errorf("duration_s is required")
			}
		case "speed":
			if req.Speed == nil {
				return fmt.Errorf("speed is required")
			}
		}
	}
	return nil
}

func validateRatio(name string, val *float64) error {
	if val == nil {
		return nil
	}
	if *val < 0.0 || *val > 1.0 {
		return fmt.Errorf("%s must be between 0.00 and 1.00, got %g", name, *val)
	}
	return nil
}

// ── Request Building ──────────────────────────────────────────────────────────

// BuildRequestURL constructs the upstream APIMart endpoint URL.
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

// BuildRequestHeader sets the APIMart authentication header.
func (a *TaskAdaptor) BuildRequestHeader(_ *gin.Context, req *http.Request, info *relaycommon.RelayInfo) error {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+info.ApiKey)
	return nil
}

// resolveUpstreamTaskID 将 MAPI 内部 task_id 转换为上游 APIMart task_id。
// 若 DB 中查不到，则原样返回（兜底，让上游报错）。
func resolveUpstreamTaskID(userID int, internalID string) string {
	if internalID == "" {
		return internalID
	}
	task, exist, err := model.GetByTaskId(userID, internalID)
	if err != nil || !exist {
		return internalID
	}
	return task.GetUpstreamTaskID()
}

// translateTaskIDs 返回一个副本，其中 task_id / task_ids 字段已替换为上游 ID。
func translateTaskIDs(req *dto.APIMartSunoRequest, def ToolDef, userID int) *dto.APIMartSunoRequest {
	if def.UsesTaskIDs {
		ids := make([]string, len(req.TaskIDs))
		for i, id := range req.TaskIDs {
			ids[i] = resolveUpstreamTaskID(userID, id)
		}
		reqCopy := *req
		reqCopy.TaskIDs = ids
		return &reqCopy
	}
	if !def.NoTaskID && !def.UsesAudioURLs && !def.UsesAudioURL && req.TaskID != "" {
		upstream := resolveUpstreamTaskID(userID, req.TaskID)
		if upstream != req.TaskID {
			reqCopy := *req
			reqCopy.TaskID = upstream
			return &reqCopy
		}
	}
	return req
}

// BuildRequestBody converts the parsed APIMartSunoRequest into the upstream JSON body.
func (a *TaskAdaptor) BuildRequestBody(c *gin.Context, info *relaycommon.RelayInfo) (io.Reader, error) {
	v, ok := c.Get("task_request")
	if !ok {
		return nil, fmt.Errorf("task_request not found in context")
	}
	req, ok := v.(*dto.APIMartSunoRequest)
	if !ok {
		return nil, fmt.Errorf("task_request has unexpected type")
	}

	modelID := info.OriginModelName
	def, _ := GetToolDef(modelID)

	// 将用户传入的内部 task_id(s) 翻译为 APIMart 上游 task_id(s)
	req = translateTaskIDs(req, def, info.UserId)

	body := buildUpstreamBody(req, modelID, def)
	data, err := common.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("marshal request body: %w", err)
	}

	upstreamURL, _ := a.BuildRequestURL(info)
	logger.LogInfo(c, fmt.Sprintf("[apimart-suno] submit request: model=%s url=%s body=%s", modelID, upstreamURL, string(data)))

	return bytes.NewReader(data), nil
}

// buildUpstreamBody constructs the upstream request payload for the given tool.
func buildUpstreamBody(req *dto.APIMartSunoRequest, modelID string, def ToolDef) map[string]any {
	body := make(map[string]any)

	// APIMart 所有工具均需要 model 字段，固定值为 "suno"
	body["model"] = "suno"

	// Version (only for versioned tools)
	if def.SupportedVersions != nil && req.Version != "" {
		body["version"] = req.Version
	}

	// custom_model_id（与 version 互斥，version 已在上面按 SupportedVersions!=nil 判断）
	if req.CustomModelID != "" {
		body["custom_model_id"] = req.CustomModelID
	}

	// Tool-specific source reference
	if def.UsesTaskIDs {
		body["task_ids"] = req.TaskIDs
		if len(req.AudioIndexes) > 0 {
			body["audio_indexes"] = req.AudioIndexes
		}
	} else if def.UsesAudioURLs {
		body["audio_urls"] = req.AudioURLs
	} else if def.UsesAudioURL {
		body["audio_url"] = req.AudioURL
	} else if !def.NoTaskID {
		body["task_id"] = req.TaskID
		if req.AudioIndex != nil {
			body["audio_index"] = *req.AudioIndex
		}
	}

	// V6 通用字段（所有支持的工具均透传）
	if req.MaxMode != nil {
		body["max_mode"] = *req.MaxMode
	}
	if req.Variety != "" {
		body["variety"] = req.Variety
	}
	if req.AudioFormat != "" {
		body["audio_format"] = req.AudioFormat
	}

	// Tool-specific body builders
	switch modelID {
	case "suno-music":
		buildMusicBody(req, body)
	case "suno-lyrics":
		buildLyricsBody(req, body)
	case "suno-upload":
		body["audioFilePath"] = req.AudioFilePath
	case "suno-upsample-tags":
		body["tags"] = req.Tags
	case "suno-sounds":
		body["prompt"] = req.Prompt
		if req.Type != "" {
			body["type"] = req.Type
		}
		if req.BPM != nil {
			body["bpm"] = *req.BPM
		}
		if req.Key != "" {
			body["key"] = req.Key
		}
	case "suno-adjust-speed":
		if req.Speed != nil {
			body["speed"] = *req.Speed
		}
		if req.KeepPitch != nil {
			body["keep_pitch"] = *req.KeepPitch
		}
		if req.Title != "" {
			body["title"] = req.Title
		}
	case "suno-crop":
		if req.StartS != nil {
			body["start_s"] = *req.StartS
		}
		if req.EndS != nil {
			body["end_s"] = *req.EndS
		}
	case "suno-fade-in", "suno-fade-out":
		if req.DurationS != nil {
			body["duration_s"] = *req.DurationS
		}
		if req.Title != "" {
			body["title"] = req.Title
		}
	case "suno-remove-section":
		if req.StartS != nil {
			body["start_s"] = *req.StartS
		}
		if req.EndS != nil {
			body["end_s"] = *req.EndS
		}
	case "suno-replace-section":
		if req.StartS != nil {
			body["start_s"] = *req.StartS
		}
		if req.EndS != nil {
			body["end_s"] = *req.EndS
		}
		if req.InfillLyrics != "" {
			body["infill_lyrics"] = req.InfillLyrics
		}
		addContextFields(req, body)
		addOperationWeirdness(req, body)
	case "suno-remaster":
		if req.VariationCategory != "" {
			body["variation_category"] = req.VariationCategory
		}
	case "suno-extend":
		if req.ContinueAt != nil {
			body["continue_at"] = *req.ContinueAt
		}
		addCustomGenerationFields(req, body, false)
		addOperationWeirdness(req, body)
	case "suno-sample":
		if req.StartS != nil {
			body["start_s"] = *req.StartS
		}
		if req.EndS != nil {
			body["end_s"] = *req.EndS
		}
		if req.Instrumental != nil {
			body["instrumental"] = *req.Instrumental
		}
		addCustomGenerationFields(req, body, false)
		addOperationWeirdness(req, body)
	case "suno-cover", "suno-mashup":
		addCustomGenerationFields(req, body, false)
		addOperationWeirdness(req, body)
	case "suno-add-vocals", "suno-add-instrumental":
		// 支持：custom, prompt, gpt_description, title, tags, negative_tags,
		//       style_weight, weirdness, audio_weight, vocal_gender
		// 不支持：auto_lyrics, persona_id, instrumental, duration_s
		buildAddVocalsInstrumentalBody(req, body)
		addOperationWeirdness(req, body)
	case "suno-add-stem":
		// 与 add-vocals 相比，还不支持 vocal_gender
		buildAddStemBody(req, body)
		addOperationWeirdness(req, body)
	case "suno-inspo":
		// 不支持 custom / instrumental / gpt_description / persona_id / duration_s
		buildInspoBody(req, body)
		addOperationWeirdness(req, body)
	case "suno-persona":
		body["name"] = req.Name
		if req.Description != "" {
			body["describe"] = req.Description // 上游字段名为 describe
		}
		if req.Styles != "" {
			body["styles"] = req.Styles
		}
		// V6: vox_audio_id 已废弃，不再透传
		if req.VocalStartS != nil {
			body["vocal_start_s"] = *req.VocalStartS
		}
		if req.VocalEndS != nil {
			body["vocal_end_s"] = *req.VocalEndS
		}
	case "suno-stems":
		if req.StemType != "" {
			body["stem_type"] = req.StemType
		}
		// suno-stems-all 不发 stem_type
	case "suno-download":
		// 合并 format（单数）和 formats（数组）
		formats := req.Formats
		if req.Format != "" {
			// 避免重复
			found := false
			for _, f := range formats {
				if strings.EqualFold(f, req.Format) {
					found = true
					break
				}
			}
			if !found {
				formats = append(formats, req.Format)
			}
		}
		if len(formats) > 0 {
			body["formats"] = formats
		}
	case "suno-upload-cover":
		// 上传公网 URL 翻唱
		addCustomGenerationFields(req, body, false)
		addOperationWeirdness(req, body)
		if req.DurationS != nil {
			body["duration_s"] = *req.DurationS
		}
	case "suno-upload-extend":
		// 上传公网 URL 延伸（固定自定义模式，不发 custom/instrumental/gpt_description）
		if req.ContinueAt != nil {
			body["continue_at"] = *req.ContinueAt
		}
		if req.Prompt != "" {
			body["prompt"] = req.Prompt
		}
		if req.Tags != "" {
			body["tags"] = req.Tags
		}
		if req.Title != "" {
			body["title"] = req.Title
		}
		if req.NegativeTags != "" {
			body["negative_tags"] = req.NegativeTags
		}
		if req.DurationS != nil {
			body["duration_s"] = *req.DurationS
		}
		if req.AutoLyrics != nil {
			body["auto_lyrics"] = *req.AutoLyrics
		}
		if req.VocalGender != "" {
			body["vocal_gender"] = req.VocalGender
		}
		if req.PersonaID != "" {
			body["persona_id"] = req.PersonaID
		}
		addRatioFields(req, body)
		addOperationWeirdness(req, body)
	case "suno-create-model":
		// audio_urls 已在通用层处理（UsesAudioURLs=true）
		body["name"] = req.Name
	}

	return body
}

// buildMusicBody handles suno-music which uses "style" instead of "tags".
// suno-music uses weirdness_constraint (not weirdness) per V6 spec.
func buildMusicBody(req *dto.APIMartSunoRequest, body map[string]any) {
	if req.Custom != nil {
		body["custom"] = *req.Custom
	}
	if req.Instrumental != nil {
		body["instrumental"] = *req.Instrumental
	}
	if req.Prompt != "" {
		body["prompt"] = req.Prompt
	}
	if req.Title != "" {
		body["title"] = req.Title
	}
	// suno-music uses "style" (not "tags") for style tags
	if req.Style != "" {
		body["style"] = req.Style
	}
	if req.NegativeTags != "" {
		body["negative_tags"] = req.NegativeTags
	}
	if req.AutoLyrics != nil {
		body["auto_lyrics"] = *req.AutoLyrics
	}
	if req.PersonaID != "" {
		body["persona_id"] = req.PersonaID
	}
	if req.VocalGender != "" {
		body["vocal_gender"] = req.VocalGender
	}
	// Duration (integer, custom=true only)
	if req.Duration != nil {
		body["duration"] = *req.Duration
	}
	// suno-music: use weirdness_constraint (V6 spec)
	addRatioFields(req, body)
	// weirdness_constraint is already handled in addRatioFields for suno-music
}

func buildLyricsBody(req *dto.APIMartSunoRequest, body map[string]any) {
	if req.Prompt != "" {
		body["prompt"] = req.Prompt
	}
	if req.LyricsModel != "" {
		body["lyrics_model"] = req.LyricsModel
	}
}

// addCustomGenerationFields adds custom/prompt/gpt_description and style fields.
// These are shared by cover/extend/sample/mashup/inspo/add_* tools.
func addCustomGenerationFields(req *dto.APIMartSunoRequest, body map[string]any, usesStyle bool) {
	if req.Custom != nil {
		body["custom"] = *req.Custom
	}
	if req.Prompt != "" {
		body["prompt"] = req.Prompt
	}
	if req.GptDescription != "" {
		body["gpt_description"] = req.GptDescription
	}
	if req.Title != "" {
		body["title"] = req.Title
	}
	if usesStyle {
		if req.Style != "" {
			body["style"] = req.Style
		}
	} else {
		if req.Tags != "" {
			body["tags"] = req.Tags
		}
	}
	if req.NegativeTags != "" {
		body["negative_tags"] = req.NegativeTags
	}
	if req.AutoLyrics != nil {
		body["auto_lyrics"] = *req.AutoLyrics
	}
	if req.PersonaID != "" {
		body["persona_id"] = req.PersonaID
	}
	if req.VocalGender != "" {
		body["vocal_gender"] = req.VocalGender
	}
	if req.Instrumental != nil {
		body["instrumental"] = *req.Instrumental
	}
	if req.DurationS != nil {
		body["duration_s"] = *req.DurationS
	}
	addRatioFields(req, body)
}

// addContextFields adds prompt/title/tags used as song context (replace_section).
func addContextFields(req *dto.APIMartSunoRequest, body map[string]any) {
	if req.Prompt != "" {
		body["prompt"] = req.Prompt
	}
	if req.Title != "" {
		body["title"] = req.Title
	}
	if req.Tags != "" {
		body["tags"] = req.Tags
	}
	if req.NegativeTags != "" {
		body["negative_tags"] = req.NegativeTags
	}
}

// addRatioFields sends style_weight, audio_weight.
// weirdness handling is split: suno-music uses weirdness_constraint, operation tools use weirdness.
func addRatioFields(req *dto.APIMartSunoRequest, body map[string]any) {
	if req.StyleWeight != nil {
		body["style_weight"] = *req.StyleWeight
	}
	if req.AudioWeight != nil {
		body["audio_weight"] = *req.AudioWeight
	}
	// weirdness_constraint for suno-music (main generation)
	if req.WeirdnessConstraint != nil {
		body["weirdness_constraint"] = *req.WeirdnessConstraint
	}
}

// buildAddVocalsInstrumentalBody builds request for suno-add-vocals and suno-add-instrumental.
// Supported: custom, prompt, gpt_description, title, tags, negative_tags,
//
//	style_weight, weirdness, audio_weight, vocal_gender
//
// NOT supported: auto_lyrics, persona_id, instrumental, duration_s
func buildAddVocalsInstrumentalBody(req *dto.APIMartSunoRequest, body map[string]any) {
	if req.Custom != nil {
		body["custom"] = *req.Custom
	}
	if req.Prompt != "" {
		body["prompt"] = req.Prompt
	}
	if req.GptDescription != "" {
		body["gpt_description"] = req.GptDescription
	}
	if req.Title != "" {
		body["title"] = req.Title
	}
	if req.Tags != "" {
		body["tags"] = req.Tags
	}
	if req.NegativeTags != "" {
		body["negative_tags"] = req.NegativeTags
	}
	if req.VocalGender != "" {
		body["vocal_gender"] = req.VocalGender
	}
	addRatioFields(req, body)
}

// buildAddStemBody builds request for suno-add-stem.
// Same as add-vocals/instrumental but without vocal_gender.
func buildAddStemBody(req *dto.APIMartSunoRequest, body map[string]any) {
	if req.Custom != nil {
		body["custom"] = *req.Custom
	}
	if req.Prompt != "" {
		body["prompt"] = req.Prompt
	}
	if req.GptDescription != "" {
		body["gpt_description"] = req.GptDescription
	}
	if req.Title != "" {
		body["title"] = req.Title
	}
	if req.Tags != "" {
		body["tags"] = req.Tags
	}
	if req.NegativeTags != "" {
		body["negative_tags"] = req.NegativeTags
	}
	addRatioFields(req, body)
}

// buildInspoBody builds request for suno-inspo.
// NOT supported: custom, instrumental, gpt_description, persona_id, duration_s
// audio_urls is handled by the common layer (UsesAudioURLs=true).
func buildInspoBody(req *dto.APIMartSunoRequest, body map[string]any) {
	if req.Prompt != "" {
		body["prompt"] = req.Prompt
	}
	if req.Title != "" {
		body["title"] = req.Title
	}
	if req.Tags != "" {
		body["tags"] = req.Tags
	}
	if req.NegativeTags != "" {
		body["negative_tags"] = req.NegativeTags
	}
	if req.VocalGender != "" {
		body["vocal_gender"] = req.VocalGender
	}
	if req.AutoLyrics != nil {
		body["auto_lyrics"] = *req.AutoLyrics
	}
	addRatioFields(req, body)
}

// addOperationWeirdness sends "weirdness" for operation/editing tools (V6 preferred field).
// weirdness_constraint is accepted as compat alias; if client sends Weirdness, prefer it.
func addOperationWeirdness(req *dto.APIMartSunoRequest, body map[string]any) {
	if req.Weirdness != nil {
		body["weirdness"] = *req.Weirdness
	} else if req.WeirdnessConstraint != nil {
		// compat: client sent old name, forward as weirdness
		body["weirdness"] = *req.WeirdnessConstraint
		// remove weirdness_constraint if it was set by addRatioFields
		delete(body, "weirdness_constraint")
	}
}

// ── HTTP ──────────────────────────────────────────────────────────────────────

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
	client := apimartSunoHTTPClient
	if info.ChannelSetting.Proxy != "" {
		proxyClient, proxyErr := service.NewProxyHttpClient(info.ChannelSetting.Proxy)
		if proxyErr != nil {
			return nil, fmt.Errorf("new proxy client: %w", proxyErr)
		}
		client = proxyClient
	}
	return client.Do(req)
}

// DoResponse parses the APIMart submit response and returns the upstream task ID.
func (a *TaskAdaptor) DoResponse(c *gin.Context, resp *http.Response, info *relaycommon.RelayInfo) (taskID string, taskData []byte, taskErr *dto.TaskError) {
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		taskErr = service.TaskErrorWrapper(err, "read_response_body_failed", http.StatusInternalServerError)
		return
	}

	logger.LogInfo(c, fmt.Sprintf("[apimart-suno] submit response: model=%s status=%d body=%s", info.OriginModelName, resp.StatusCode, string(responseBody)))

	if resp.StatusCode != http.StatusOK {
		taskErr = service.TaskErrorWrapper(
			fmt.Errorf("upstream returned status %d: %s", resp.StatusCode, string(responseBody)),
			"upstream_error", resp.StatusCode,
		)
		return
	}

	var submitResp dto.APIMartSunoSubmitResponse
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

	// Return the public task ID to the client (never expose the upstream ID)
	clientResp := dto.APIMartSunoSubmitResponse{
		Code:    200,
		Message: "success",
		Data: []dto.APIMartSunoSubmitTaskEntry{
			{Status: "submitted", TaskID: info.PublicTaskID},
		},
	}
	c.JSON(http.StatusOK, clientResp)

	return upstreamTaskID, responseBody, nil
}

// ── Polling ───────────────────────────────────────────────────────────────────

// FetchTask issues GET /v1/music/tasks/:task_id to APIMart.
func (a *TaskAdaptor) FetchTask(baseUrl, key string, body map[string]any, proxy string) (*http.Response, error) {
	upstreamTaskID, ok := body["task_id"].(string)
	if !ok || upstreamTaskID == "" {
		return nil, fmt.Errorf("FetchTask: missing task_id in body")
	}

	url := strings.TrimRight(baseUrl, "/") + "/v1/music/tasks/" + upstreamTaskID
	logger.LogInfo(context.Background(), fmt.Sprintf("[apimart-suno] fetch task request: url=%s", url))

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("FetchTask: new request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+key)

	client := apimartSunoHTTPClient
	if proxy != "" {
		var proxyErr error
		client, proxyErr = service.NewProxyHttpClient(proxy)
		if proxyErr != nil {
			return nil, fmt.Errorf("FetchTask: get http client: %w", proxyErr)
		}
	}
	return client.Do(req)
}

// ParseTaskResult parses the APIMart single-task query response.
func (a *TaskAdaptor) ParseTaskResult(respBody []byte) (*relaycommon.TaskInfo, error) {
	return ParseTaskResult(respBody)
}

// ── Registry ──────────────────────────────────────────────────────────────────

func (a *TaskAdaptor) GetModelList() []string {
	return ModelList
}

func (a *TaskAdaptor) GetChannelName() string {
	return ChannelName
}
