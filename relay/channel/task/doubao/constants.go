package doubao

var ModelList = []string{
	"doubao-seedance-1-0-pro-250528",
	"doubao-seedance-1-0-lite-t2v",
	"doubao-seedance-1-0-lite-i2v",
	"doubao-seedance-1-5-pro-251215",
	"doubao-seedance-2-0-260128",
	"doubao-seedance-2-0-fast-260128",
	// BytePlus (international) equivalents
	"dreamina-seedance-2-0-260128",
	"dreamina-seedance-2-0-fast-260128",
}

var ChannelName = "doubao-video"

// videoInputRatioMap 视频输入折扣比率（含视频单价 / 不含视频单价）。
// 管理员应将 ModelRatio 设置为"不含视频"的较高费率，
// 系统在检测到视频输入时自动乘以此折扣。
// doubao-* 为火山引擎（国内），dreamina-* 为 BytePlus（国际），底层同款模型，折扣比率相同。
var videoInputRatioMap = map[string]float64{
	"doubao-seedance-2-0-260128":          28.0 / 46.0, // ~0.6087
	"doubao-seedance-2-0-fast-260128":     22.0 / 37.0, // ~0.5946
	"dreamina-seedance-2-0-260128":        28.0 / 46.0, // ~0.6087 (BytePlus)
	"dreamina-seedance-2-0-fast-260128":   22.0 / 37.0, // ~0.5946 (BytePlus)
}

func GetVideoInputRatio(modelName string) (float64, bool) {
	r, ok := videoInputRatioMap[modelName]
	return r, ok
}

// resolutionRatioMap 分辨率溢价比率（1080p单价 / 480p/720p同场景单价）。
// 管理员应将 ModelRatio 设置为 480p/720p 的基础费率，系统在检测到 1080p 时自动乘以此比率。
// 比率来源：官方定价表 含视频1080P(0.031) / 含视频480P(0.028) = 31/28 ≈ 1.10714
// fast 模型不支持 1080p 输出，无需配置。
var resolutionRatioMap = map[string]float64{
	"doubao-seedance-2-0-260128:1080p":   31.0 / 28.0, // ≈1.10714，官方定价表标注值
	"dreamina-seedance-2-0-260128:1080p": 31.0 / 28.0, // ≈1.10714 (BytePlus 同款模型)
}

func GetResolutionRatio(modelName string, resolution string) (float64, bool) {
	if resolution == "" {
		return 0, false
	}
	r, ok := resolutionRatioMap[modelName+":"+resolution]
	return r, ok
}
