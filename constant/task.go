package constant

type TaskPlatform string

const (
	TaskPlatformSuno       TaskPlatform = "suno"
	TaskPlatformMidjourney              = "mj"
)

const (
	SunoActionMusic  = "MUSIC"
	SunoActionLyrics = "LYRICS"

	TaskActionGenerate          = "generate"
	TaskActionTextGenerate      = "textGenerate"
	TaskActionFirstTailGenerate = "firstTailGenerate"
	TaskActionReferenceGenerate = "referenceGenerate"
	TaskActionRemix             = "remixGenerate"
	TaskActionImageGenerate     = "imageGenerate"
	TaskAction3DGenerate        = "3dGenerate"
)

var SunoModel2Action = map[string]string{
	"suno_music":  SunoActionMusic,
	"suno_lyrics": SunoActionLyrics,
}

// APIMartSunoModelPath maps APIMart Suno model IDs to their upstream API path suffixes.
// The full upstream URL is: POST {baseURL}/v1/music/generations/{path}
// An empty path means POST {baseURL}/v1/music/generations (used for suno-music).
var APIMartSunoModelPath = map[string]string{
	"suno-music":          "",
	"suno-lyrics":         "lyrics",
	"suno-aligned-lyrics": "alignedLyrics",
	"suno-bpm":            "bpm",
	"suno-concat":         "concat",
	"suno-generate-video": "generateMp4",
	"suno-persona":        "persona",
	"suno-upload":         "uploadTask",
	"suno-upsample-tags":  "upsampleTags",
	"suno-vox":            "vox",
	"suno-wav":            "wav",
	"suno-crop":           "crop",
	"suno-fade-in":        "fadeIn",
	"suno-fade-out":       "fadeOut",
	"suno-remove-section": "removeSection",
	"suno-sounds":         "sounds",
	"suno-create-voice":   "createVoice",
	"suno-adjust-speed":   "adjustSpeed",
	"suno-add-instrumental": "addInstrumental",
	"suno-add-stem":       "addStem",
	"suno-add-vocals":     "addVocals",
	"suno-cover":          "coverSong",
	"suno-extend":         "extend",
	"suno-mashup":         "mashup",
	"suno-midi":           "midi",
	"suno-remaster":       "remaster",
	"suno-replace-section": "replaceMusic",
	"suno-sample":         "sample",
	"suno-inspo":          "inspo",
	"suno-stems":          "stems",
	"suno-stems-all":      "stemsAll",
}
