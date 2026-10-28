package tencentvod

import (
	"math"
	"strings"
)

// VideoModelList exposes all Tencent VOD AIGC video models.
// Naming convention: {family}-video-{version}, consistent with the image side ({family}-image-{version}).
// All 25 versions are included, strictly following the official ModelVersion values from:
// https://cloud.tencent.com/document/product/266/126239
var VideoModelList = []string{
	// Kling 系列生视频模型（腾讯 VOD ModelName="Kling"）
	"kling-video-1.6",
	"kling-video-2.0",
	"kling-video-2.1",
	"kling-video-2.5",
	"kling-video-2.6",
	"kling-video-o1",
	"kling-video-3.0",
	"kling-video-3.0-omni",
	// Vidu 系列生视频模型（腾讯 VOD ModelName="Vidu"）
	"vidu-video-q2",
	"vidu-video-q2-pro",
	"vidu-video-q2-turbo",
	"vidu-video-q3",
	"vidu-video-q3-pro",
	"vidu-video-q3-turbo",
	// Hailuo（海螺）系列生视频模型（腾讯 VOD ModelName="Hailuo"）
	"hailuo-video-02",
	"hailuo-video-2.3",
	"hailuo-video-2.3-fast",
	"hailuo-video-h3",
	// Google Veo 系列生视频模型（腾讯 VOD ModelName="GV"）
	"veo-video-3.1",
	"veo-video-3.1-fast",
	// OpenAI Sora 系列生视频模型（腾讯 VOD ModelName="OS"）
	"sora-video-2.0",
	// Hunyuan 系列生视频模型（腾讯 VOD ModelName="Hunyuan"）
	"hunyuan-video-1.5",
	// Mingmou 系列生视频模型（腾讯 VOD ModelName="Mingmou"）
	"mingmou-video-1.0",
	// PixVerse 系列生视频模型（腾讯 VOD ModelName="PixVerse"）
	"pixverse-video-v5.6",
	"pixverse-video-v6",
	"pixverse-video-c1",
}

// videoModelNameMap maps internal model name to Tencent VOD API ModelName field.
var videoModelNameMap = map[string]string{
	"kling-video-1.6":      "Kling",
	"kling-video-2.0":      "Kling",
	"kling-video-2.1":      "Kling",
	"kling-video-2.5":      "Kling",
	"kling-video-2.6":      "Kling",
	"kling-video-o1":       "Kling",
	"kling-video-3.0":      "Kling",
	"kling-video-3.0-omni": "Kling",

	"vidu-video-q2":       "Vidu",
	"vidu-video-q2-pro":   "Vidu",
	"vidu-video-q2-turbo": "Vidu",
	"vidu-video-q3":       "Vidu",
	"vidu-video-q3-pro":   "Vidu",
	"vidu-video-q3-turbo": "Vidu",

	"hailuo-video-02":       "Hailuo",
	"hailuo-video-2.3":      "Hailuo",
	"hailuo-video-2.3-fast": "Hailuo",
	"hailuo-video-h3":       "Hailuo",

	"veo-video-3.1":      "GV",
	"veo-video-3.1-fast": "GV",

	"sora-video-2.0": "OS",

	"hunyuan-video-1.5": "Hunyuan",

	"mingmou-video-1.0": "Mingmou",

	"pixverse-video-v5.6": "PixVerse",
	"pixverse-video-v6":   "PixVerse",
	"pixverse-video-c1":   "PixVerse",
}

// videoModelVersionMap maps internal model name to Tencent VOD API ModelVersion field.
// Strictly follows official documentation values (case-sensitive).
var videoModelVersionMap = map[string]string{
	"kling-video-1.6":      "1.6",
	"kling-video-2.0":      "2.0",
	"kling-video-2.1":      "2.1",
	"kling-video-2.5":      "2.5",
	"kling-video-2.6":      "2.6",
	"kling-video-o1":       "O1",
	"kling-video-3.0":      "3.0",
	"kling-video-3.0-omni": "3.0-Omni",

	"vidu-video-q2":       "q2",
	"vidu-video-q2-pro":   "q2-pro",
	"vidu-video-q2-turbo": "q2-turbo",
	"vidu-video-q3":       "q3",
	"vidu-video-q3-pro":   "q3-pro",
	"vidu-video-q3-turbo": "q3-turbo",

	"hailuo-video-02":       "02",
	"hailuo-video-2.3":      "2.3",
	"hailuo-video-2.3-fast": "2.3-fast",
	"hailuo-video-h3":       "H3",

	"veo-video-3.1":      "3.1",
	"veo-video-3.1-fast": "3.1-fast",

	"sora-video-2.0": "2.0",

	"hunyuan-video-1.5": "1.5",

	"mingmou-video-1.0": "1.0",

	"pixverse-video-v5.6": "v5.6",
	"pixverse-video-v6":   "v6",
	"pixverse-video-c1":   "c1",
}

// videoModelSet is built from VideoModelList for O(1) isVideoModel lookups.
var videoModelSet = func() map[string]struct{} {
	m := make(map[string]struct{}, len(VideoModelList))
	for _, name := range VideoModelList {
		m[name] = struct{}{}
	}
	return m
}()

// isVideoModel returns true if the given origin model name is a TencentVOD video model.
func isVideoModel(originModelName string) bool {
	_, ok := videoModelSet[originModelName]
	return ok
}

// getVideoModelName returns the Tencent VOD ModelName for the given internal model name.
func getVideoModelName(originModelName string) string {
	if v, ok := videoModelNameMap[originModelName]; ok {
		return v
	}
	return originModelName
}

// getVideoModelVersion returns the Tencent VOD ModelVersion for the given internal model name.
func getVideoModelVersion(originModelName string) string {
	if v, ok := videoModelVersionMap[originModelName]; ok {
		return v
	}
	return ""
}

// ──────────────────────────────────────────────────────────────────────────────
// Duration constraints per ModelName
// ──────────────────────────────────────────────────────────────────────────────

// getVideoDuration returns the clamped/snapped Duration value (in seconds) to send
// to the Tencent VOD API, or 0 if Duration should not be sent for this model.
//
// Rules by ModelName:
//
//	Kling:   clamp to [3, 15]; 0 means "use API default (5)"
//	Hailuo:  forward the requested duration; 0 means "use API default (6)"
//	Vidu:    clamp to [1, 10]
//	GV(Veo): always 8 (only legal value), ignore user input
//	OS(Sora): snap to nearest of {4, 8, 12}; 0 means "use API default (8)"
//	PixVerse: clamp to [1, 15]
//	Hunyuan / Mingmou: return 0 (do not send Duration)
func getVideoDuration(tencentModelName string, requestedSeconds float64) float64 {
	switch tencentModelName {
	case "Kling":
		if requestedSeconds <= 0 {
			return 0 // let API use default (5s)
		}
		return clampF(requestedSeconds, 3, 15)
	case "Hailuo":
		// Do not snap. H3 accepts 4–15; other versions accept 6 or 10.
		// Illegal values are rejected by Tencent and returned as-is.
		if requestedSeconds <= 0 {
			return 0
		}
		return requestedSeconds
	case "Vidu":
		if requestedSeconds <= 0 {
			return 0
		}
		return clampF(requestedSeconds, 1, 10)
	case "GV":
		return 8 // only legal value
	case "OS":
		if requestedSeconds <= 0 {
			return 0 // let API use default (8s)
		}
		return snapToNearest(requestedSeconds, []float64{4, 8, 12})
	case "PixVerse":
		if requestedSeconds <= 0 {
			return 0 // let API use default (5s)
		}
		return clampF(requestedSeconds, 1, 15)
	default:
		// Hunyuan, Mingmou: do not send Duration
		return 0
	}
}

// ──────────────────────────────────────────────────────────────────────────────
// Resolution constraints per ModelName
// ──────────────────────────────────────────────────────────────────────────────

// validVideoResolutions lists the legal Resolution values per ModelName.
// Note: PixVerse uses all-lowercase (540p/720p/1080p/2k/4k), all others use uppercase P/K.
var validVideoResolutions = map[string][]string{
	"Kling":    {"720P", "1080P"},
	"Hailuo":   {"768P", "1080P"},
	"Vidu":     {"720P", "1080P"},
	"GV":       {"720P", "1080P"},
	"OS":       {"720P"},
	"PixVerse": {"540p", "720p", "1080p", "2k", "4k"},
	// Hunyuan, Mingmou: no Resolution field
}

// defaultVideoResolution returns the default Resolution value for a model.
var defaultVideoResolution = map[string]string{
	"Kling":    "720P",
	"Hailuo":   "768P",
	"Vidu":     "720P",
	"GV":       "720P",
	"OS":       "720P",
	"PixVerse": "720p",
}

// getVideoResolution converts a size string to the model-specific Resolution value.
//
// Priority: size field > metadata.resolution (caller is responsible for fallback).
// Returns "" if the model does not support Resolution (Hunyuan, Mingmou).
//
// Size format handling:
//   - "WxH" (e.g. "1920x1080"): classify by short side (≥1080→1080P/1080p, ≥720→720P/720p)
//   - "Np"/"NK" (e.g. "720p", "1080P", "2k"): normalize then match model's valid list
//   - "" : return model default
func getVideoResolution(tencentModelName, size string) string {
	validList, ok := validVideoResolutions[tencentModelName]
	if !ok {
		return "" // model does not use Resolution
	}
	def := defaultVideoResolution[tencentModelName]

	if size == "" {
		return def
	}

	// "WxH" pixel dimension
	if strings.Contains(size, "x") {
		parts := strings.SplitN(size, "x", 2)
		if len(parts) == 2 {
			w := parseIntSafe(parts[0])
			h := parseIntSafe(parts[1])
			short := h
			if w < h {
				short = w
			}
			if short >= 1080 {
				size = "1080P"
			} else if short >= 720 {
				size = "720P"
			} else {
				return def
			}
		}
	}

	// Normalize: PixVerse wants lowercase, everything else uppercase
	isPixVerse := tencentModelName == "PixVerse"
	normalized := normalizeResolutionToken(size, isPixVerse)

	// Exact match
	for _, v := range validList {
		if v == normalized {
			return v
		}
	}

	// Hailuo H3 also accepts 2K/4K. Do not snap an explicit value onto 768P/1080P.
	if tencentModelName == "Hailuo" && normalized != "" {
		return normalized
	}

	// Nearest match (by resolution order index)
	return nearestResolution(normalized, validList, def, isPixVerse)
}

// normalizeResolutionToken converts a raw size token to the canonical form for comparison.
// PixVerse: "720P"→"720p", "2K"→"2k"
// Others:   "720p"→"720P", "2k"→"2K"
func normalizeResolutionToken(token string, pixverse bool) string {
	if pixverse {
		return strings.ToLower(strings.TrimSpace(token))
	}
	// Uppercase P/K, e.g. "720p"→"720P", "1080p"→"1080P", "2k"→"2K"
	t := strings.TrimSpace(token)
	t = strings.ReplaceAll(t, "p", "P")
	t = strings.ReplaceAll(t, "k", "K")
	return t
}

// nearestResolution picks the closest value from validList by pixel count.
var resolutionPixels = map[string]int{
	"540p":  540,
	"720p":  720,
	"720P":  720,
	"768P":  768,
	"1080p": 1080,
	"1080P": 1080,
	"2k":    2048,
	"2K":    2048,
	"4k":    4096,
	"4K":    4096,
}

func nearestResolution(target string, validList []string, def string, pixverse bool) string {
	targetPx, ok := resolutionPixels[target]
	if !ok {
		return def
	}
	best := def
	bestDiff := math.MaxInt32
	for _, v := range validList {
		px, exists := resolutionPixels[v]
		if !exists {
			continue
		}
		diff := abs(px - targetPx)
		if diff < bestDiff {
			bestDiff = diff
			best = v
		}
	}
	return best
}

// getBillingDuration returns the effective duration (seconds) that the pricing
// expression used for pre-charge. It converts getVideoDuration's "0 = use API
// default" sentinel into the concrete default value that the expression applies.
//
// Parameters:
//   - sentDur: the Duration value actually sent to the Tencent API (0 = not sent / API default).
//   - requestedDur: the raw duration from the user's request body (0 = not specified).
//
// Model rules:
//
//	Kling:    sentDur=0 → 5 (expression default)
//	Hailuo:   sentDur=0 → 6 (API default). An explicit duration is billed as sent.
//	Vidu:     sentDur=0 → 5 (expression default)
//	GV:       always 8 (only legal value)
//	OS(Sora): sentDur=0 → 8 (expression default)
//	PixVerse: sentDur=0 → 5 (expression default)
//	Hunyuan / Mingmou: adapter never sends Duration; expression uses requestedDur directly
//	  (let d = d0 <= 0 ? 5 : d0), so reconciliation is based on request intent.
//
// A return value of 0 means duration reconciliation should be skipped for
// this task (no reliable billing duration is known at submit time).
func getBillingDuration(tencentModelName string, sentDur, requestedDur float64) float64 {
	if sentDur > 0 {
		return sentDur
	}
	// sentDur == 0: either API default was used, or the model never sends Duration.
	switch tencentModelName {
	case "Kling":
		return 5 // expression: d0 <= 0 ? 5
	case "Vidu":
		return 5 // expression: d0 <= 0 ? 5
	case "OS":
		return 8 // expression: d0 <= 0 ? 8
	case "GV":
		return 8 // always sends 8; sentDur>0 path normally handles it
	case "PixVerse":
		return 5 // expression: d0 <= 0 ? 5
	case "Hailuo":
		return 6 // API default when Duration is omitted
	case "Hunyuan", "Mingmou":
		// Adapter does not send Duration to Tencent; expression bills the request intent:
		//   let d = d0 <= 0 ? 5 : d0
		// Track this so AdjustBillingOnComplete can reconcile against actual output duration.
		if requestedDur <= 0 {
			return 5
		}
		return requestedDur
	default:
		return 0
	}
}

// frameInterpolateSupported returns true if the model supports FrameInterpolate.
// Currently Vidu-exclusive.
func frameInterpolateSupported(tencentModelName string) bool {
	return tencentModelName == "Vidu"
}

// sceneTypeSupported returns true if the model supports the SceneType field.
func sceneTypeSupported(tencentModelName string) bool {
	switch tencentModelName {
	case "Kling", "Vidu":
		return true
	}
	return false
}

// ──────────────────────────────────────────────────────────────────────────────
// Utility helpers
// ──────────────────────────────────────────────────────────────────────────────

func clampF(v, min, max float64) float64 {
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}

func snapToNearest(v float64, candidates []float64) float64 {
	if len(candidates) == 0 {
		return v
	}
	best := candidates[0]
	bestDiff := math.Abs(v - best)
	for _, c := range candidates[1:] {
		if d := math.Abs(v - c); d < bestDiff {
			bestDiff = d
			best = c
		}
	}
	return best
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func parseIntSafe(s string) int {
	s = strings.TrimSpace(s)
	n := 0
	for _, c := range s {
		if c >= '0' && c <= '9' {
			n = n*10 + int(c-'0')
		}
	}
	return n
}
