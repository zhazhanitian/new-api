package volcengine

import "strings"

// 像素对照表来自火山图片生成文档「模型实际映射的宽高像素参考值」。
// pro 与 flash 共用一张表。lite、4.5、4.0 的 2K/4K 相同，1K 和 3K 按各自文档分开。

var seedreamProPixels = map[string]map[string]string{
	"1K": {
		"1:1": "1024x1024", "4:3": "1152x864", "3:4": "864x1152", "16:9": "1424x800",
		"9:16": "800x1424", "3:2": "1248x832", "2:3": "832x1248", "21:9": "1568x672",
	},
	"1.5K": {
		"1:1": "1536x1536", "4:3": "1792x1344", "3:4": "1344x1792", "16:9": "2048x1152",
		"9:16": "1152x2048", "3:2": "1872x1248", "2:3": "1248x1872", "21:9": "2352x1008",
	},
	"2K": {
		"1:1": "2048x2048", "4:3": "2368x1776", "3:4": "1776x2368", "16:9": "2816x1584",
		"9:16": "1584x2816", "3:2": "2496x1664", "2:3": "1664x2496", "21:9": "3136x1344",
	},
}

var seedreamWidePixels = map[string]map[string]string{
	"2K": {
		"1:1": "2048x2048", "4:3": "2304x1728", "3:4": "1728x2304", "16:9": "2848x1600",
		"9:16": "1600x2848", "3:2": "2496x1664", "2:3": "1664x2496", "21:9": "3136x1344",
	},
	"4K": {
		"1:1": "4096x4096", "4:3": "4704x3520", "3:4": "3520x4704", "16:9": "5504x3040",
		"9:16": "3040x5504", "3:2": "4992x3328", "2:3": "3328x4992", "21:9": "6240x2656",
	},
}

var seedreamLitePixels = map[string]map[string]string{
	"2K": seedreamWidePixels["2K"],
	"3K": {
		"1:1": "3072x3072", "4:3": "3456x2592", "3:4": "2592x3456", "16:9": "4096x2304",
		"9:16": "2304x4096", "3:2": "3744x2496", "2:3": "2496x3744", "21:9": "4704x2016",
	},
	"4K": seedreamWidePixels["4K"],
}

var seedream40Pixels = map[string]map[string]string{
	"1K": {
		"1:1": "1024x1024", "4:3": "1152x864", "3:4": "864x1152", "16:9": "1280x720",
		"9:16": "720x1280", "3:2": "1248x832", "2:3": "832x1248", "21:9": "1512x648",
	},
	"2K": seedreamWidePixels["2K"],
	"4K": seedreamWidePixels["4K"],
}

// seedreamPixelSize 把 resolution 和 ratio 映射成上游 size。对不上官方表时返回 false。
func seedreamPixelSize(model, resolution, ratio string) (string, bool) {
	table := seedreamPixelTable(model)
	if table == nil {
		return "", false
	}
	ratios, ok := table[canonicalSeedreamResolution(resolution)]
	if !ok {
		return "", false
	}
	size, ok := ratios[strings.TrimSpace(ratio)]
	return size, ok && size != ""
}

func seedreamPixelTable(model string) map[string]map[string]string {
	name := strings.ToLower(strings.TrimSpace(model))
	switch {
	case strings.Contains(name, "5.0-pro"), strings.Contains(name, "5-0-pro"),
		strings.Contains(name, "5.0-flash"), strings.Contains(name, "5-0-flash"):
		return seedreamProPixels
	case strings.Contains(name, "5.0-lite"), strings.Contains(name, "5-0-260128"), strings.Contains(name, "5-0-t2i-250804"):
		return seedreamLitePixels
	case strings.Contains(name, "4-5-"), strings.Contains(name, "4.5"):
		return seedreamWidePixels
	case strings.Contains(name, "4-0-"), strings.Contains(name, "seedream-4-0"):
		return seedream40Pixels
	default:
		return nil
	}
}

func canonicalSeedreamResolution(resolution string) string {
	switch strings.ToLower(strings.TrimSpace(resolution)) {
	case "1k":
		return "1K"
	case "1.5k":
		return "1.5K"
	case "2k":
		return "2K"
	case "3k":
		return "3K"
	case "4k":
		return "4K"
	default:
		return strings.TrimSpace(resolution)
	}
}
