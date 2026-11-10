package apimart_flowmusic

// defaultVersion is Lyria 3 Pro. Clients may omit version; the gateway fills it in.
const defaultVersion = "lyria-3-pro"

// ToolDef maps one public model ID to its upstream path under /v1/music/generations.
// An empty Path is POST {base}/v1/music/generations.
// HasVersion marks tools whose version enum is lyria-3-pro (default) or lyria-3.5.
type ToolDef struct {
	Path       string
	HasVersion bool
}

var toolDefs = map[string]ToolDef{
	"flowmusic":          {Path: "", HasVersion: true},
	"flowmusic-lyrics":   {Path: "lyricsFlowMusic"},
	"flowmusic-extend":   {Path: "extendFlowMusic", HasVersion: true},
	"flowmusic-replace":  {Path: "replaceFlowMusic", HasVersion: true},
	"flowmusic-cover":    {Path: "coverFlowMusic", HasVersion: true},
	"flowmusic-stems":    {Path: "stemsFlowMusic"},
	"flowmusic-upload":   {Path: "uploadAudioFlowMusic"},
	"flowmusic-download": {Path: "downloadAudioFlowMusic"},
	"flowmusic-video":    {Path: "videoClipFlowMusic"},
}

func GetToolDef(modelID string) (ToolDef, bool) {
	def, ok := toolDefs[modelID]
	return def, ok
}
