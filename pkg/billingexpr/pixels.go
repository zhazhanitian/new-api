package billingexpr

import (
	"fmt"
	"strconv"
	"strings"
)

// pixelsOf parses an image size value into total pixels (width * height).
// Used by the billing expression function pixels().
//
// Accepted forms:
//   - "WIDTHxHEIGHT" with separator x / X / * / ×, optional spaces (e.g. "2048x2048")
//   - named tiers: "1K"/"1k" → 1024×1024, "2K"/"2k" → 2048×2048, "4K"/"4k" → 3840×2160
//   - empty / "auto" / unparseable → 0
func pixelsOf(v interface{}) float64 {
	if v == nil {
		return 0
	}
	s := strings.TrimSpace(fmt.Sprint(v))
	if s == "" || strings.EqualFold(s, "auto") {
		return 0
	}

	switch strings.ToUpper(s) {
	case "1K":
		return 1024 * 1024
	case "2K":
		return 2048 * 2048
	case "4K":
		return 3840 * 2160
	}

	normalized := strings.NewReplacer("×", "x", "*", "x", "X", "x", " ", "").Replace(s)
	parts := strings.SplitN(normalized, "x", 2)
	if len(parts) != 2 {
		return 0
	}
	w, errW := strconv.ParseFloat(parts[0], 64)
	h, errH := strconv.ParseFloat(parts[1], 64)
	if errW != nil || errH != nil || w <= 0 || h <= 0 {
		return 0
	}
	return w * h
}
