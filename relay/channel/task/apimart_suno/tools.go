package apimart_suno

// allVersions lists every Suno V6 model version supported by APIMart.
var allVersions = []string{"v6", "v6-wild", "v6-mini"}

// ToolDef describes the upstream routing and validation rules for one APIMart Suno tool.
type ToolDef struct {
	// Path is the APIMart API path suffix after /v1/music/generations.
	// An empty string means the request goes directly to /v1/music/generations.
	Path string

	// SupportedVersions is the list of accepted "version" values.
	// nil means the tool has no version dimension (do not send "version" to upstream).
	SupportedVersions []string

	// VersionRequired means the "version" field MUST be provided by the client.
	// (Currently only suno-music enforces this.)
	VersionRequired bool

	// RequiredFields lists the field names the client must provide (non-empty / non-nil).
	RequiredFields []string

	// UsesTaskIDs marks tools that accept task_ids[] instead of a single task_id.
	// (Only suno-mashup.)
	UsesTaskIDs bool

	// UsesAudioURLs marks tools that accept audio_urls[] instead of task_id.
	// (suno-inspo, suno-create-model)
	UsesAudioURLs bool

	// UsesAudioURL marks tools that accept a single audio_url instead of task_id.
	// (suno-create-voice, suno-upload-cover, suno-upload-extend)
	UsesAudioURL bool

	// NoTaskID marks tools that do not reference a source task at all.
	NoTaskID bool
}

// toolDefs maps each model ID to its ToolDef.
var toolDefs = map[string]ToolDef{
	"suno-music": {
		Path:              "",
		SupportedVersions: allVersions,
		// VersionRequired 不再强制：允许仅传 custom_model_id（V6 文档：两者至少提供一个，互斥）
		// 具体校验在 ValidateRequestAndSetAction 中处理
		VersionRequired: false,
		NoTaskID:        true,
	},
	"suno-lyrics": {
		Path:           "lyrics",
		RequiredFields: []string{"prompt"},
		NoTaskID:       true,
	},
	"suno-aligned-lyrics": {
		Path:           "alignedLyrics",
		RequiredFields: []string{"task_id"},
	},
	"suno-bpm": {
		Path:           "bpm",
		RequiredFields: []string{"task_id"},
	},
	"suno-concat": {
		Path:           "concat",
		RequiredFields: []string{"task_id"},
	},
	"suno-generate-video": {
		Path:           "generateMp4",
		RequiredFields: []string{"task_id"},
	},
	"suno-persona": {
		Path:           "persona",
		RequiredFields: []string{"task_id", "name"},
	},
	"suno-upload": {
		Path:           "uploadTask",
		RequiredFields: []string{"audio_file_path"},
		NoTaskID:       true,
	},
	"suno-upsample-tags": {
		Path:           "upsampleTags",
		RequiredFields: []string{"tags"},
		NoTaskID:       true,
	},
	"suno-crop": {
		Path:           "crop",
		RequiredFields: []string{"task_id", "start_s", "end_s"},
	},
	"suno-fade-in": {
		Path:           "fadeIn",
		RequiredFields: []string{"task_id", "duration_s"},
	},
	"suno-fade-out": {
		Path:           "fadeOut",
		RequiredFields: []string{"task_id", "duration_s"},
	},
	"suno-remove-section": {
		Path:           "removeSection",
		RequiredFields: []string{"task_id", "start_s", "end_s"},
	},
	"suno-sounds": {
		Path:              "sounds",
		SupportedVersions: allVersions,
		RequiredFields:    []string{"prompt"},
		NoTaskID:          true,
	},
	"suno-create-voice": {
		Path:           "createVoice",
		RequiredFields: []string{"audio_url"},
		UsesAudioURL:   true,
		NoTaskID:       true,
	},
	"suno-adjust-speed": {
		Path:           "adjustSpeed",
		RequiredFields: []string{"task_id", "speed"},
	},
	"suno-add-instrumental": {
		Path:              "addInstrumental",
		SupportedVersions: allVersions,
		RequiredFields:    []string{"task_id"},
	},
	"suno-add-stem": {
		Path:              "addStem",
		SupportedVersions: allVersions,
		RequiredFields:    []string{"task_id"},
	},
	"suno-add-vocals": {
		Path:              "addVocals",
		SupportedVersions: allVersions,
		RequiredFields:    []string{"task_id"},
	},
	"suno-cover": {
		Path:              "coverSong",
		SupportedVersions: allVersions,
		RequiredFields:    []string{"task_id"},
	},
	"suno-extend": {
		Path:              "extend",
		SupportedVersions: allVersions,
		RequiredFields:    []string{"task_id", "continue_at"},
	},
	"suno-mashup": {
		Path:              "mashup",
		SupportedVersions: allVersions,
		RequiredFields:    []string{"task_ids"},
		UsesTaskIDs:       true,
	},
	"suno-midi": {
		Path:           "midi",
		RequiredFields: []string{"task_id"},
	},
	"suno-remaster": {
		// V6: no version dimension
		Path:           "remaster",
		RequiredFields: []string{"task_id"},
	},
	"suno-replace-section": {
		Path:              "replaceMusic",
		SupportedVersions: allVersions,
		RequiredFields:    []string{"task_id", "start_s", "end_s"},
	},
	"suno-sample": {
		Path:              "sample",
		SupportedVersions: allVersions,
		RequiredFields:    []string{"task_id", "start_s", "end_s"},
	},
	"suno-inspo": {
		Path:              "inspo",
		SupportedVersions: allVersions,
		RequiredFields:    []string{"audio_urls"},
		UsesAudioURLs:     true,
		NoTaskID:          true,
	},
	"suno-stems": {
		Path:           "stems",
		RequiredFields: []string{"task_id"},
	},
	"suno-stems-all": {
		Path:           "stemsAll",
		RequiredFields: []string{"task_id"},
	},
	// ── V6 新增工具 ─────────────────────────────────────────────────────────────
	"suno-download": {
		// 下载音频文件（替代已废弃的 wav 接口）
		Path:           "download",
		RequiredFields: []string{"task_id"},
	},
	"suno-upload-cover": {
		// 上传公网音频 URL 并翻唱（无需先 uploadTask）
		Path:              "uploadCover",
		SupportedVersions: allVersions,
		RequiredFields:    []string{"audio_url"},
		UsesAudioURL:      true,
		NoTaskID:          true,
	},
	"suno-upload-extend": {
		// 上传公网音频 URL 并延伸（无需先 uploadTask）
		Path:              "uploadExtend",
		SupportedVersions: allVersions,
		RequiredFields:    []string{"audio_url", "continue_at"},
		UsesAudioURL:      true,
		NoTaskID:          true,
	},
	"suno-create-model": {
		// 创建自定义模型（6-24 条参考音频）
		Path:          "createModel",
		RequiredFields: []string{"name"},
		UsesAudioURLs: true,
		NoTaskID:      true,
	},
}

// GetToolDef returns the ToolDef for the given model ID, or (zero, false) if not found.
func GetToolDef(modelID string) (ToolDef, bool) {
	def, ok := toolDefs[modelID]
	return def, ok
}
