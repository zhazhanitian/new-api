package doubao

var ModelList = []string{
	// ── Video generation (async, seedance) ──────────────────────────────
	"doubao-seedance-1-0-pro-250528",
	"doubao-seedance-1-0-lite-t2v",
	"doubao-seedance-1-0-lite-i2v",
	"doubao-seedance-1-5-pro-251215",
	"doubao-seedance-2-0-260128",
	"doubao-seedance-2-0-fast-260128",
	"doubao-seedance-2-0-mini-260615",
	// BytePlus (international) equivalents
	"dreamina-seedance-2-0-260128",
	"dreamina-seedance-2-0-fast-260128",

	// ── Image generation (background goroutine, seedream) ────────────────
	// These models call a synchronous Volcengine image API wrapped as async.
	"doubao-seedream-3-0-t2i-250415",
	"doubao-seedream-4-5-t2i-250505",
	"doubao-seedream-4-5-251128",
	"doubao-seedream-5-0-t2i-250804",
}

var ChannelName = "doubao-video"

// videoInputRatioMap 视频输入折扣比率（含视频单价 / 不含视频单价）。
// 管理员应将 ModelRatio 设置为"不含视频"的较高费率，
// 系统在检测到视频输入时自动乘以此折扣。
// doubao-* 为火山引擎（国内），dreamina-* 为 BytePlus（国际），底层同款模型，折扣比率相同。
var videoInputRatioMap = map[string]float64{
	"doubao-seedance-2-0-260128":          28.0 / 46.0, // ~0.6087
	"doubao-seedance-2-0-fast-260128":     22.0 / 37.0, // ~0.5946
	"doubao-seedance-2-0-mini-260615":     14.0 / 23.0, // ~0.6087
	"dreamina-seedance-2-0-260128":        28.0 / 46.0, // ~0.6087 (BytePlus)
	"dreamina-seedance-2-0-fast-260128":   22.0 / 37.0, // ~0.5946 (BytePlus)
}

func GetVideoInputRatio(modelName string) (float64, bool) {
	r, ok := videoInputRatioMap[modelName]
	return r, ok
}

// resolutionRatioMap 分辨率调价比率（各分辨率单价 / 480p基准单价）。
// 管理员应将 ModelRatio 设置为 480p/720p 的基础费率，系统在检测到其他分辨率时自动乘以此比率。
// 1080P 为溢价（>1），4K 为折扣（<1）。fast/mini 模型仅支持 480P/720P，无需配置。
var resolutionRatioMap = map[string]float64{
	"doubao-seedance-2-0-260128:1080p":   31.0 / 28.0, // ≈1.10714，1080P溢价（含视频1080P/含视频480P）
	"doubao-seedance-2-0-260128:4K":      26.0 / 46.0, // ≈0.5652，4K折扣（不含视频4K/不含视频480P）
	"dreamina-seedance-2-0-260128:1080p": 31.0 / 28.0, // ≈1.10714 (BytePlus 同款模型)
	"dreamina-seedance-2-0-260128:4K":    26.0 / 46.0, // ≈0.5652 (BytePlus 同款模型)
}

func GetResolutionRatio(modelName string, resolution string) (float64, bool) {
	if resolution == "" {
		return 0, false
	}
	r, ok := resolutionRatioMap[modelName+":"+resolution]
	return r, ok
}

// silentVideoRatioMap 无声视频折扣比率（无声单价 / 有声单价）。
// 管理员应将 ModelRatio 设置为有声视频费率，系统在检测到未开启音频时自动乘以此折扣。
var silentVideoRatioMap = map[string]float64{
	"doubao-seedance-1-5-pro-251215": 8.0 / 16.0, // 0.5，无声视频为有声视频半价
}

func GetSilentVideoRatio(modelName string) (float64, bool) {
	r, ok := silentVideoRatioMap[modelName]
	return r, ok
}
