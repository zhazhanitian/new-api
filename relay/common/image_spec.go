package common

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/QuantumNous/new-api/dto"
)

// ImageFields 是图片生成的四个对外字段，加上张数和输出格式。
// Size 只接受宽x高。Resolution 是分辨率档位。Ratio 是宽高比。
// Quality 只表示渲染精度：low、medium、high、xhigh、max。
type ImageFields struct {
	Size         string
	Resolution   string
	Ratio        string
	Quality      string
	OutputFormat string
	N            int
}

type pixelRule struct {
	multiple  int
	minPixels int
	maxPixels int
	minEdge   int
	maxEdge   int
	maxAspect float64
	enforce   bool
}

type imageSpec struct {
	names                  []string
	prefixes               []string
	allowSize              bool
	allowResolution        bool
	allowRatio             bool
	allowQuality           bool
	resolutions            []string
	ratios                 []string
	qualities              []string
	sizeExcludesResolution bool
	sizeExcludesRatio      bool
	pixels                 pixelRule
	maxN                   int
	maxRef                 int
	outputFormats          []string
	defaultQuality         string
	// resolutionAliases、formatAliases 把调用方写法收成规格表里的标准值，键不区分大小写。
	resolutionAliases map[string]string
	formatAliases     map[string]string
}

var (
	qualityFive      = []string{"low", "medium", "high", "xhigh", "max"}
	ggRatios         = []string{"1:1", "2:3", "3:2", "3:4", "4:3", "4:5", "5:4", "9:16", "16:9", "21:9"}
	gg31Ratios       = []string{"1:1", "2:3", "3:2", "3:4", "4:3", "4:5", "5:4", "9:16", "16:9", "21:9", "1:4", "4:1", "1:8", "8:1"}
	klingRatios      = []string{"1:1", "2:3", "3:2", "3:4", "4:3", "9:16", "16:9", "21:9"}
	omniRatios       = []string{"1:1", "2:3", "3:2", "3:4", "4:3", "9:16", "16:9", "21:9", "auto"}
	apiRatios        = []string{"1:1", "3:2", "2:3", "4:3", "3:4", "5:4", "4:5", "16:9", "9:16", "2:1", "1:2", "21:9", "9:21", "3:1", "1:3"}
	fluxRatios       = []string{"21:9", "2:1", "16:9", "3:2", "7:5", "4:3", "5:4", "1:1", "4:5", "3:4", "5:7", "2:3", "9:16", "1:2", "9:21", "auto"}
	maiRatios        = []string{"1:1", "4:3", "3:4", "3:2", "2:3", "16:9", "9:16", "2:1", "1:2", "21:9", "9:21", "4:1", "1:4", "auto"}
	qwenImage3Ratios = []string{"1:1", "4:3", "3:4", "16:9", "9:16", "3:2", "2:3"}
	museRatios       = []string{"1:1", "16:9", "9:16", "4:3", "3:4", "3:2", "2:3"}
	hyImageRatios    = []string{"1:1", "2:3", "3:2", "3:4", "4:3", "4:5", "5:4", "9:16", "16:9", "21:9", "auto"}
)

var (
	fluxResolutions       = []string{"768sq", "1k", "1.5k", "2k", "4k"}
	maiResolutions        = []string{"1K", "2K"}
	qwenImage3Resolutions = []string{"1K", "2K"}
	hyImageResolutions    = []string{"1K", "2K", "4K"}
)

func imageSpecs() []imageSpec {
	gptPixels := pixelRule{multiple: 16, minPixels: 655360, maxPixels: 8294400, enforce: true}
	apiPixels := pixelRule{multiple: 16, minPixels: 655360, maxPixels: 8294400, maxEdge: 3840, maxAspect: 3, enforce: true}
	return []imageSpec{
		{
			names:                  []string{"gpt-image-2", "OG-image-2"},
			allowSize:              true,
			allowResolution:        true,
			allowRatio:             true,
			allowQuality:           true,
			resolutions:            []string{"1K", "2K", "4K"},
			ratios:                 []string{"1:1", "3:2", "2:3", "3:4", "4:3", "16:9", "9:16", "21:9", "9:21"},
			qualities:              qualityFive,
			sizeExcludesResolution: true,
			sizeExcludesRatio:      true,
			pixels:                 gptPixels,
		},
		{
			names:           []string{"gemini-2.5-flash-image", "GG-2.5-flash-image", "gemini-3-pro-image", "GG-3-pro-image"},
			allowResolution: true,
			allowRatio:      true,
			resolutions:     []string{"1K", "2K", "4K"},
			ratios:          ggRatios,
		},
		{
			names:           []string{"gemini-3.1-flash-image", "GG-3.1-flash-image"},
			allowResolution: true,
			allowRatio:      true,
			resolutions:     []string{"720P", "1K", "2K", "4K"},
			ratios:          gg31Ratios,
		},
		{
			names:           []string{"kling-image-2.1", "kling-image-3.0"},
			allowResolution: true,
			allowRatio:      true,
			resolutions:     []string{"1k", "2k"},
			ratios:          klingRatios,
		},
		{
			names:           []string{"kling-image-3.0-omni", "kling-image-o1"},
			allowResolution: true,
			allowRatio:      true,
			resolutions:     []string{"1k", "2k", "4k"},
			ratios:          omniRatios,
		},
		{
			names: []string{"kling-image-scene"},
		},
		{
			names:           []string{"vidu-image-q2"},
			allowResolution: true,
			allowRatio:      true,
			resolutions:     []string{"1080p", "2K", "4K"},
			ratios:          klingRatios,
		},
		{
			names:     []string{"qwen-image-0925"},
			allowSize: true,
			pixels:    pixelRule{minPixels: 261632, maxPixels: 4194304, enforce: true},
		},
		{
			names:     []string{"hunyuan-image-3.0"},
			allowSize: true,
			pixels:    pixelRule{minEdge: 512, maxEdge: 2048, maxPixels: 1048576, enforce: true},
		},
		{
			names: []string{
				"doubao-seedream-5.0-pro",
				"doubao-seedream-5.0-flash",
				"doubao-seedream-5-0-pro-260628",
				"doubao-seedream-5-0-flash-260915",
			},
			allowSize:              true,
			allowResolution:        true,
			allowRatio:             true,
			resolutions:            []string{"1K", "1.5K", "2K"},
			ratios:                 seedreamRatios,
			sizeExcludesResolution: true,
			sizeExcludesRatio:      true,
		},
		{
			names:                  []string{"doubao-seedream-5-0-260128", "doubao-seedream-5-0-t2i-250804", "doubao-seedream-5.0-lite"},
			allowSize:              true,
			allowResolution:        true,
			allowRatio:             true,
			resolutions:            []string{"2K", "3K", "4K"},
			ratios:                 seedreamRatios,
			sizeExcludesResolution: true,
			sizeExcludesRatio:      true,
		},
		{
			names:                  []string{"doubao-seedream-4-5-251128", "doubao-seedream-4-5-t2i-250505"},
			allowSize:              true,
			allowResolution:        true,
			allowRatio:             true,
			resolutions:            []string{"2K", "4K"},
			ratios:                 seedreamRatios,
			sizeExcludesResolution: true,
			sizeExcludesRatio:      true,
		},
		{
			names:                  []string{"doubao-seedream-4-0-250828", "seedream-4-0-250828"},
			allowSize:              true,
			allowResolution:        true,
			allowRatio:             true,
			resolutions:            []string{"1K", "2K", "4K"},
			ratios:                 seedreamRatios,
			sizeExcludesResolution: true,
			sizeExcludesRatio:      true,
		},
		{
			names:     []string{"doubao-seedream-3-0-t2i-250415"},
			allowSize: true,
		},
		{
			names: []string{
				"gpt-image-2.5-flare", "gpt-image-2.5-sunburst", "gpt-image-2.6-flare", "gpt-image-2.6-sunburst",
				"OG-image-2.5-flare", "OG-image-2.5-sunburst", "OG-image-2.6-flare", "OG-image-2.6-sunburst",
			},
			allowSize:              true,
			allowResolution:        true,
			allowRatio:             true,
			allowQuality:           true,
			resolutions:            []string{"1K", "2K", "4K"},
			ratios:                 apiRatios,
			qualities:              qualityFive,
			sizeExcludesResolution: true,
			sizeExcludesRatio:      true,
			pixels:                 apiPixels,
			maxN:                   4,
			maxRef:                 16,
			outputFormats:          []string{"png", "jpeg", "webp"},
			defaultQuality:         "low",
		},
		{
			names:             []string{"flux-3-image"},
			allowResolution:   true,
			allowRatio:        true,
			resolutions:       fluxResolutions,
			ratios:            fluxRatios,
			resolutionAliases: map[string]string{"768": "768sq"},
			maxN:              1,
			maxRef:            10,
		},
		{
			names:             []string{"flux-3-image-klein"},
			allowResolution:   true,
			allowRatio:        true,
			resolutions:       fluxResolutions,
			ratios:            fluxRatios,
			resolutionAliases: map[string]string{"768": "768sq"},
			formatAliases:     map[string]string{"jpg": "jpeg"},
			maxN:              1,
			maxRef:            10,
			outputFormats:     []string{"jpeg", "png"},
		},
		{
			names:             []string{"mai-image-2.6", "mai-image-2.6-flash"},
			allowSize:         true,
			allowResolution:   true,
			allowRatio:        true,
			resolutions:       maiResolutions,
			ratios:            maiRatios,
			sizeExcludesRatio: true,
			pixels:            pixelRule{minEdge: 768, maxPixels: 2359296, maxAspect: 4, enforce: true},
			maxN:              1,
			maxRef:            5,
		},
		{
			names:                  []string{"qwen-image-3.0", "qwen-image-3.0-pro"},
			allowSize:              true,
			allowResolution:        true,
			allowRatio:             true,
			resolutions:            qwenImage3Resolutions,
			ratios:                 qwenImage3Ratios,
			sizeExcludesResolution: true,
			sizeExcludesRatio:      true,
			pixels:                 pixelRule{minEdge: 512, maxEdge: 2048, maxAspect: 8, enforce: true},
			maxN:                   6,
			maxRef:                 3,
		},
		{
			names:      []string{"muse-image-ext"},
			allowRatio: true,
			ratios:     museRatios,
			maxN:       1,
			maxRef:     9,
		},
		{
			names:                  []string{"hy-image-v3.5-preview"},
			allowSize:              true,
			allowResolution:        true,
			allowRatio:             true,
			resolutions:            hyImageResolutions,
			ratios:                 hyImageRatios,
			sizeExcludesResolution: true,
			sizeExcludesRatio:      true,
			pixels:                 pixelRule{minEdge: 256, maxEdge: 8192, maxPixels: 8294400, enforce: true},
			maxN:                   1,
			maxRef:                 5,
		},
	}
}

// seedreamRatios 是火山图片文档列出的常见宽高比。网关按这组比例把 resolution 映射成像素 size。
var seedreamRatios = []string{"1:1", "4:3", "3:4", "16:9", "9:16", "3:2", "2:3", "21:9"}

// NormalizeImageFields 按模型规格校验并归一化四个字段。
// strict 为 true 时，没登记的模型直接拒绝，用于 v2。
// strict 为 false 时，没登记的模型原样返回，用于 v1 上的 dall-e、gpt-image-1 等。
func NormalizeImageFields(model string, in ImageFields, strict bool) (ImageFields, error) {
	model = strings.TrimSpace(model)
	spec, ok := lookupImageSpec(model)
	if !ok {
		if strict {
			if model == "" {
				return in, fmt.Errorf("model is required")
			}
			return in, fmt.Errorf("model %s is not registered for image generation", model)
		}
		return in, nil
	}

	out := in
	out.Size = strings.TrimSpace(in.Size)
	out.Resolution = applyAlias(spec.resolutionAliases, strings.TrimSpace(in.Resolution))
	out.Ratio = strings.TrimSpace(in.Ratio)
	out.Quality = strings.TrimSpace(in.Quality)
	out.OutputFormat = applyAlias(spec.formatAliases, strings.TrimSpace(in.OutputFormat))

	if out.Size != "" {
		if !spec.allowSize {
			return in, fmt.Errorf("model %s does not accept size", model)
		}
		normalized, err := normalizeWxH(out.Size, spec.pixels)
		if err != nil {
			return in, err
		}
		out.Size = normalized
	}
	if out.Resolution != "" {
		if !spec.allowResolution {
			return in, fmt.Errorf("model %s does not accept resolution", model)
		}
		canon, ok := canonical(out.Resolution, spec.resolutions)
		if !ok {
			return in, fmt.Errorf("resolution %q is not supported for %s", out.Resolution, model)
		}
		out.Resolution = canon
	}
	if out.Ratio != "" {
		if !spec.allowRatio {
			return in, fmt.Errorf("model %s does not accept ratio", model)
		}
		if canon, ok := canonical(out.Ratio, spec.ratios); ok {
			out.Ratio = canon
		} else if !ratioAllowed(out.Ratio, spec.ratios) {
			return in, fmt.Errorf("ratio %q is not supported for %s", out.Ratio, model)
		}
	}
	if out.Quality != "" {
		if !spec.allowQuality {
			return in, fmt.Errorf("model %s does not accept quality", model)
		}
		canon, ok := canonical(out.Quality, spec.qualities)
		if !ok {
			return in, fmt.Errorf("quality %q is not supported for %s", out.Quality, model)
		}
		out.Quality = canon
	}
	if out.Size != "" && out.Resolution != "" && spec.sizeExcludesResolution {
		return in, fmt.Errorf("size and resolution cannot be used together")
	}
	if out.Size != "" && out.Ratio != "" && spec.sizeExcludesRatio {
		return in, fmt.Errorf("size and ratio cannot be used together")
	}
	if spec.maxN > 0 && out.N > spec.maxN {
		return in, fmt.Errorf("n must be between 1 and %d", spec.maxN)
	}
	if out.N < 0 {
		return in, fmt.Errorf("n must be positive")
	}
	if out.OutputFormat != "" && len(spec.outputFormats) > 0 {
		canon, ok := canonical(out.OutputFormat, spec.outputFormats)
		if !ok {
			return in, fmt.Errorf("output_format %q is not supported for %s", out.OutputFormat, model)
		}
		out.OutputFormat = canon
	}
	return out, nil
}

// ImageUpstreamDefaultQuality 返回网关发给上游时补上的 quality。
// 只对需要确定计费档位的模型返回非空，目前是 APIMart。
func ImageUpstreamDefaultQuality(model string) string {
	spec, ok := lookupImageSpec(model)
	if !ok {
		return ""
	}
	return spec.defaultQuality
}

// ImageRefLimit 返回参考图张数上限。0 表示不在这里截断。
func ImageRefLimit(model string) int {
	spec, ok := lookupImageSpec(model)
	if !ok {
		return 0
	}
	return spec.maxRef
}

// ApplyImageFieldRules 校验 dto.ImageRequest，并把归一化后的值写回去。
// resolution、ratio 放在 Extra 里，避免 ImageRequest 序列化时把它们带给不认识这两个字段的渠道。
func ApplyImageFieldRules(req *dto.ImageRequest, strict bool) error {
	if req == nil {
		return nil
	}
	n := 0
	if req.N != nil {
		n = int(*req.N)
	}
	out, err := NormalizeImageFields(req.Model, ImageFields{
		Size:         req.Size,
		Resolution:   ExtraString(req.Extra, "resolution"),
		Ratio:        ExtraString(req.Extra, "ratio"),
		Quality:      req.Quality,
		OutputFormat: rawJSONString(req.OutputFormat),
		N:            n,
	}, strict)
	if err != nil {
		return err
	}
	if _, ok := lookupImageSpec(req.Model); !ok {
		return nil
	}
	req.Size = out.Size
	req.Quality = out.Quality
	if out.OutputFormat != "" && rawJSONString(req.OutputFormat) != out.OutputFormat {
		encoded, mErr := json.Marshal(out.OutputFormat)
		if mErr == nil {
			req.OutputFormat = encoded
		}
	}
	setExtraString(req, "resolution", out.Resolution)
	setExtraString(req, "ratio", out.Ratio)
	return nil
}

// ExtraString 从 ImageRequest.Extra 里取出一个字符串字段。
func ExtraString(extra map[string]json.RawMessage, key string) string {
	if len(extra) == 0 {
		return ""
	}
	raw, ok := extra[key]
	if !ok || len(raw) == 0 {
		return ""
	}
	var s string
	if err := json.Unmarshal(raw, &s); err != nil {
		return ""
	}
	return strings.TrimSpace(s)
}

func setExtraString(req *dto.ImageRequest, key, value string) {
	if value == "" {
		if req.Extra != nil {
			delete(req.Extra, key)
		}
		return
	}
	if req.Extra == nil {
		req.Extra = map[string]json.RawMessage{}
	}
	encoded, err := json.Marshal(value)
	if err != nil {
		return
	}
	req.Extra[key] = encoded
}

func rawJSONString(raw json.RawMessage) string {
	if len(raw) == 0 {
		return ""
	}
	var s string
	if err := json.Unmarshal(raw, &s); err != nil {
		return ""
	}
	return strings.TrimSpace(s)
}

func lookupImageSpec(model string) (imageSpec, bool) {
	var prefixHit imageSpec
	prefixLen := 0
	foundPrefix := false
	for _, spec := range imageSpecs() {
		for _, name := range spec.names {
			if model == name {
				return spec, true
			}
		}
		for _, prefix := range spec.prefixes {
			if strings.HasPrefix(model, prefix) && len(prefix) > prefixLen {
				prefixHit = spec
				prefixLen = len(prefix)
				foundPrefix = true
			}
		}
	}
	if foundPrefix {
		return prefixHit, true
	}
	return imageSpec{}, false
}

func applyAlias(aliases map[string]string, value string) string {
	if value == "" || len(aliases) == 0 {
		return value
	}
	for key, alias := range aliases {
		if strings.EqualFold(value, key) {
			return alias
		}
	}
	return value
}

func canonical(value string, allowed []string) (string, bool) {
	for _, item := range allowed {
		if strings.EqualFold(value, item) {
			return item, true
		}
	}
	return "", false
}

func ratioAllowed(value string, allowed []string) bool {
	if len(allowed) == 1 && allowed[0] == "*" {
		parts := strings.Split(value, ":")
		if len(parts) != 2 {
			return false
		}
		w, errW := strconv.Atoi(parts[0])
		h, errH := strconv.Atoi(parts[1])
		return errW == nil && errH == nil && w > 0 && h > 0
	}
	_, ok := canonical(value, allowed)
	return ok
}

func normalizeWxH(size string, rule pixelRule) (string, error) {
	if strings.Contains(size, ":") {
		return "", fmt.Errorf("size must be widthxheight, for example 1536x1024")
	}
	parts := strings.Split(size, "x")
	if len(parts) == 1 {
		parts = strings.Split(size, "X")
	}
	if len(parts) != 2 {
		return "", fmt.Errorf("size must be widthxheight, for example 1536x1024")
	}
	w, errW := strconv.Atoi(strings.TrimSpace(parts[0]))
	h, errH := strconv.Atoi(strings.TrimSpace(parts[1]))
	if errW != nil || errH != nil || w <= 0 || h <= 0 {
		return "", fmt.Errorf("size must be widthxheight, for example 1536x1024")
	}
	if err := checkPixels(w, h, rule); err != nil {
		return "", err
	}
	return strconv.Itoa(w) + "x" + strconv.Itoa(h), nil
}

func checkPixels(w, h int, rule pixelRule) error {
	if !rule.enforce {
		return nil
	}
	if rule.multiple > 0 && (w%rule.multiple != 0 || h%rule.multiple != 0) {
		return fmt.Errorf("size width and height must both be multiples of %d", rule.multiple)
	}
	if rule.minEdge > 0 && (w < rule.minEdge || h < rule.minEdge) {
		return fmt.Errorf("size width and height must be at least %d", rule.minEdge)
	}
	if rule.maxEdge > 0 && (w > rule.maxEdge || h > rule.maxEdge) {
		return fmt.Errorf("size width and height must be at most %d", rule.maxEdge)
	}
	pixels := w * h
	if rule.minPixels > 0 && pixels < rule.minPixels {
		return fmt.Errorf("size total pixels %d is below %d", pixels, rule.minPixels)
	}
	if rule.maxPixels > 0 && pixels > rule.maxPixels {
		return fmt.Errorf("size total pixels %d is above %d", pixels, rule.maxPixels)
	}
	if rule.maxAspect > 0 {
		long, short := w, h
		if h > w {
			long, short = h, w
		}
		if short == 0 || float64(long)/float64(short) > rule.maxAspect+1e-6 {
			return fmt.Errorf("size long edge cannot exceed %.0f times the short edge", rule.maxAspect)
		}
	}
	return nil
}
