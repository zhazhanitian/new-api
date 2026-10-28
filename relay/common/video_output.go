package common

import (
	"strconv"
	"strings"
)

// NormalizeVideoResolutionRatio 把旧视频写法收成外层 Resolution、Ratio。
//
// 视频模块之后只读这两个外层字段。Size 只留给图片。
// metadata.resolution、metadata.ratio 已废弃，仅在这里读取。
//
// 外层已有值时不被旧字段覆盖。
func NormalizeVideoResolutionRatio(req *TaskSubmitReq) {
	if req == nil {
		return
	}
	resolution := strings.TrimSpace(req.Resolution)
	ratio := strings.TrimSpace(req.Ratio)
	metaResolution, metaRatio := legacyVideoResolutionRatio(req.Metadata)
	size := strings.TrimSpace(req.Size)

	if resolution == "" {
		switch {
		case isWxHSize(size):
			resolution = resolutionTierFromWxH(size)
		case isResolutionToken(size):
			resolution = size
		case metaResolution != "":
			resolution = metaResolution
		}
	}
	if ratio == "" {
		switch {
		case metaRatio != "":
			ratio = metaRatio
		case isAspectRatioSize(size):
			ratio = size
		case isWxHSize(size):
			ratio = ratioFromWxH(size)
		}
	}
	req.Resolution = resolution
	req.Ratio = ratio
}

func legacyVideoResolutionRatio(metadata map[string]interface{}) (string, string) {
	if metadata == nil {
		return "", ""
	}
	resolution, _ := metadata["resolution"].(string)
	ratio, _ := metadata["ratio"].(string)
	return strings.TrimSpace(resolution), strings.TrimSpace(ratio)
}

func isAspectRatioSize(size string) bool {
	if size == "" || strings.ContainsAny(size, "xX*") {
		return false
	}
	left, right, ok := strings.Cut(size, ":")
	if !ok || left == "" || right == "" || strings.Contains(right, ":") {
		return false
	}
	return isASCIIDigits(left) && isASCIIDigits(right)
}

func isASCIIDigits(s string) bool {
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}

func isWxHSize(size string) bool {
	parts := strings.FieldsFunc(size, func(r rune) bool { return r == 'x' || r == 'X' || r == '*' })
	if len(parts) != 2 {
		return false
	}
	w, errW := strconv.Atoi(strings.TrimSpace(parts[0]))
	h, errH := strconv.Atoi(strings.TrimSpace(parts[1]))
	return errW == nil && errH == nil && w > 0 && h > 0
}

func isResolutionToken(size string) bool {
	if size == "" || isAspectRatioSize(size) || isWxHSize(size) {
		return false
	}
	lower := strings.ToLower(size)
	return strings.Contains(lower, "p") || strings.Contains(lower, "k")
}

func resolutionTierFromWxH(size string) string {
	parts := strings.FieldsFunc(size, func(r rune) bool { return r == 'x' || r == 'X' || r == '*' })
	w, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
	h, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
	short := w
	if h < w {
		short = h
	}
	switch {
	case short >= 2160:
		return "4k"
	case short >= 1080:
		return "1080p"
	case short >= 720:
		return "720p"
	default:
		return "480p"
	}
}

func ratioFromWxH(size string) string {
	parts := strings.FieldsFunc(size, func(r rune) bool { return r == 'x' || r == 'X' || r == '*' })
	w, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
	h, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
	value := float64(w) / float64(h)
	best := "16:9"
	bestDiff := 1e9
	for _, item := range []struct {
		v     float64
		label string
	}{
		{16.0 / 9.0, "16:9"},
		{9.0 / 16.0, "9:16"},
		{1, "1:1"},
		{4.0 / 3.0, "4:3"},
		{3.0 / 4.0, "3:4"},
		{21.0 / 9.0, "21:9"},
	} {
		diff := value - item.v
		if diff < 0 {
			diff = -diff
		}
		if diff < bestDiff {
			bestDiff = diff
			best = item.label
		}
	}
	return best
}
