package apimart_suno

// allVersions lists every Suno model version supported by APIMart.
var allVersions = []string{"v3.5", "v4", "v4.5", "v4.5+", "v4.5-all", "v5", "v5.5"}

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
	// (Only suno-inspo.)
	UsesAudioURLs bool

	// UsesAudioURL marks tools that accept a single audio_url instead of task_id.
	// (Only suno-create-voice.)
	UsesAudioURL bool

	// NoTaskID marks tools that do not reference a source task at all.
	// (suno-music, suno-lyrics, suno-upload, suno-sounds, suno-upsample-tags,
	//  suno-inspo, suno-create-voice)
	NoTaskID bool
}

// toolDefs maps each model ID to its ToolDef.
var toolDefs = map[string]ToolDef{
	"suno-music": {
		Path:              "",
		SupportedVersions: allVersions,
		VersionRequired:   true,
		NoTaskID:          true,
	},
	"suno-lyrics": {
		Path:     "lyrics",
		NoTaskID: true,
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
	"suno-vox": {
		Path:           "vox",
		RequiredFields: []string{"task_id"},
	},
	"suno-wav": {
		Path:           "wav",
		RequiredFields: []string{"task_id"},
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
		SupportedVersions: []string{"v5", "v5.5"},
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
		SupportedVersions: []string{"v5", "v5.5"},
		RequiredFields:    []string{"task_id"},
	},
	"suno-add-stem": {
		Path:              "addStem",
		SupportedVersions: []string{"v5.5"},
		RequiredFields:    []string{"task_id"},
	},
	"suno-add-vocals": {
		Path:              "addVocals",
		SupportedVersions: []string{"v5", "v5.5"},
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
		Path:              "remaster",
		SupportedVersions: []string{"v4.5+", "v5", "v5.5"},
		RequiredFields:    []string{"task_id"},
	},
	"suno-replace-section": {
		Path:              "replaceMusic",
		SupportedVersions: []string{"v4", "v4.5+", "v5", "v5.5"},
		RequiredFields:    []string{"task_id", "start_s", "end_s"},
	},
	"suno-sample": {
		Path:              "sample",
		SupportedVersions: allVersions,
		RequiredFields:    []string{"task_id", "start_s", "end_s"},
	},
	"suno-inspo": {
		Path:              "inspo",
		SupportedVersions: []string{"v4", "v4.5", "v4.5+", "v4.5-all", "v5", "v5.5"},
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
}

// GetToolDef returns the ToolDef for the given model ID, or (zero, false) if not found.
func GetToolDef(modelID string) (ToolDef, bool) {
	def, ok := toolDefs[modelID]
	return def, ok
}
