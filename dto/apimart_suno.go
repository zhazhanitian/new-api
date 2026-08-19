package dto

// APIMartSunoRequest is the unified inbound request body for all APIMart Suno tools.
// Clients set `model` to the tool model ID (e.g. "suno-music") and populate only the
// fields relevant to that tool. Validation and field selection are performed in the
// adaptor layer.
type APIMartSunoRequest struct {
	Model   string `json:"model,omitempty"`
	Version string `json:"version,omitempty"`

	// Custom generation fields (music / cover / extend / sample / mashup / inspo / sounds / add_*)
	Custom              *bool    `json:"custom,omitempty"`
	Instrumental        *bool    `json:"instrumental,omitempty"`
	Prompt              string   `json:"prompt,omitempty"`
	Title               string   `json:"title,omitempty"`
	Tags                string   `json:"tags,omitempty"`
	Style               string   `json:"style,omitempty"` // suno-music only (upstream uses "style" not "tags")
	NegativeTags        string   `json:"negative_tags,omitempty"`
	GptDescription      string   `json:"gpt_description,omitempty"`
	AutoLyrics          *bool    `json:"auto_lyrics,omitempty"`
	PersonaID           string   `json:"persona_id,omitempty"`
	VocalGender         string   `json:"vocal_gender,omitempty"`
	StyleWeight         *float64 `json:"style_weight,omitempty"`
	WeirdnessConstraint *float64 `json:"weirdness_constraint,omitempty"`
	AudioWeight         *float64 `json:"audio_weight,omitempty"`

	// Source task reference (most editing tools)
	TaskID      string `json:"task_id,omitempty"`
	AudioIndex  *int   `json:"audio_index,omitempty"` // 1-based, default 1
	TaskIDs     []string `json:"task_ids,omitempty"` // mashup only (exactly 2)
	AudioIndexes []int  `json:"audio_indexes,omitempty"` // mashup parallel to task_ids

	// Audio URL inputs (inspo / create_voice)
	AudioURLs []string `json:"audio_urls,omitempty"` // inspo: 1-4 public URLs
	AudioURL  string   `json:"audio_url,omitempty"`  // create_voice: single MP3/WAV URL

	// Time-range fields
	StartS     *float64 `json:"start_s,omitempty"` // crop / remove_section / replace_section / sample / vox
	EndS       *float64 `json:"end_s,omitempty"`
	ContinueAt *int     `json:"continue_at,omitempty"` // extend: seconds

	// Upload
	AudioFilePath string `json:"audio_file_path,omitempty"` // upload

	// Sounds
	Type string   `json:"type,omitempty"` // sounds: "one-shot" | "loop"
	BPM  *int     `json:"bpm,omitempty"`  // sounds: 1-300
	Key  string   `json:"key,omitempty"`  // sounds: musical key

	// Speed adjustment
	Speed     *float64 `json:"speed,omitempty"`      // adjust_speed: 0.25-4
	KeepPitch *bool    `json:"keep_pitch,omitempty"` // adjust_speed: default true

	// Remaster
	VariationCategory string `json:"variation_category,omitempty"` // subtle | normal | high

	// Stems
	StemType string `json:"stem_type,omitempty"` // stems: e.g. "lead_vocal"

	// Persona
	Name         string `json:"name,omitempty"`
	Description  string `json:"description,omitempty"`
	Styles       string `json:"styles,omitempty"`
	VoxAudioID   string `json:"vox_audio_id,omitempty"`
	VocalStartS  *float64 `json:"vocal_start_s,omitempty"`
	VocalEndS    *float64 `json:"vocal_end_s,omitempty"`

	// Vox extraction
	// VocalStartS / VocalEndS already defined above

	// Lyrics
	LyricsModel string `json:"lyrics_model,omitempty"` // "classic" | "remi"

	// Fade
	DurationS *float64 `json:"duration_s,omitempty"` // fade_in / fade_out

	// Infill (replace_section)
	InfillLyrics string `json:"infill_lyrics,omitempty"`
}

// APIMartSunoSubmitResponse is the response from a successful task submission.
// APIMart returns data as an array even for a single task.
type APIMartSunoSubmitResponse struct {
	Code    int                          `json:"code"`
	Message string                       `json:"message,omitempty"`
	Data    []APIMartSunoSubmitTaskEntry `json:"data"`
}

type APIMartSunoSubmitTaskEntry struct {
	Status string `json:"status"`
	TaskID string `json:"task_id"`
}

func (r *APIMartSunoSubmitResponse) IsSuccess() bool {
	return r.Code == 200
}

// APIMartSunoFetchResponse is the response from GET /v1/music/tasks/:task_id.
type APIMartSunoFetchResponse struct {
	Code    int                     `json:"code"`
	Message string                  `json:"message,omitempty"`
	Data    APIMartSunoTaskData     `json:"data"`
}

type APIMartSunoTaskData struct {
	ID       string  `json:"id"`
	Status   string  `json:"status"`   // submitted | pending | completed | failed
	Progress int     `json:"progress"` // 0-100
	Result   APIMartSunoResult `json:"result"`
	Error    *APIMartSunoError `json:"error,omitempty"`
}

type APIMartSunoResult struct {
	Music []APIMartSunoMusicItem `json:"music,omitempty"`
	// upsampleTags result
	UpsampledTags string `json:"upsampled_tags,omitempty"`
	// Other tool-specific result fields stored as raw data
}

type APIMartSunoMusicItem struct {
	AudioID       string  `json:"audio_id,omitempty"`
	Title         string  `json:"title,omitempty"`
	Duration      float64 `json:"duration,omitempty"`
	Lyrics        string  `json:"lyrics,omitempty"`
	Tags          string  `json:"tags,omitempty"`
	AudioURL      string  `json:"audio_url,omitempty"`
	ImageURL      string  `json:"image_url,omitempty"`
	ImageLargeURL string  `json:"image_large_url,omitempty"`
	VideoURL      string  `json:"video_url,omitempty"`
	Status        string  `json:"status,omitempty"`
}

type APIMartSunoError struct {
	Message string `json:"message"`
}
