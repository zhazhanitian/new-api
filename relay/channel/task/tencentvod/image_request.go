package tencentvod

import (
	"fmt"
	"strings"

	"github.com/QuantumNous/new-api/common"
)

// imageBuildInput 是 v1 同步和 v2 异步共用的生图请求。
// Size 只表示宽x高。Resolution、Ratio 只在没有 Size 时写入 OutputConfig。
type imageBuildInput struct {
	SubAppId     int64
	LookupModel  string
	Prompt       string
	Size         string
	Resolution   string
	Ratio        string
	Quality      string
	N            int
	OutputFormat string
	Images       []string
	Extra        map[string]interface{}
}

// buildImageCreateRequest 构造腾讯 CreateAigcImageTask 请求。
// v1 ConvertImageRequest 和 v2 BuildRequestBody 都调用它，避免两套映射。
func buildImageCreateRequest(in imageBuildInput) createTaskRequest {
	tencentModelName := modelToTencentName(in.LookupModel)
	modelVersion := imageModelVersion(tencentModelName, in.LookupModel, in.Quality)

	body := createTaskRequest{
		SubAppId:     in.SubAppId,
		ModelName:    tencentModelName,
		ModelVersion: modelVersion,
		Prompt:       in.Prompt,
	}

	switch tencentModelName {
	case "OG":
		// 有像素尺寸时只写 ExtInfo。有分辨率或宽高比时写 OutputConfig，不写 ExtInfo。
		if in.Size != "" {
			body.ExtInfo = buildOGExtInfo(in.Size)
		} else if in.Resolution != "" || in.Ratio != "" {
			cfg := &outputConfig{}
			if in.Resolution != "" {
				cfg.Resolution = pickToken(in.Resolution, []string{"1K", "2K", "4K"}, "1K", "OG")
			}
			if in.Ratio != "" {
				cfg.AspectRatio = in.Ratio
			}
			body.OutputConfig = cfg
		}
	case "GG":
		body.OutputConfig = &outputConfig{
			AspectRatio: pickToken(in.Ratio, ggValidAspectRatios[modelVersion], "1:1", "GG"),
			Resolution:  pickToken(in.Resolution, ggValidResolutions[modelVersion], "1K", "GG"),
		}
	case "Kling":
		if modelVersion != "scene" {
			body.OutputConfig = &outputConfig{
				AspectRatio: pickToken(in.Ratio, klingValidAspectRatios[modelVersion], "1:1", "Kling"),
				Resolution:  pickToken(in.Resolution, klingValidResolutions[modelVersion], "1k", "Kling"),
			}
		}
	case "Vidu":
		body.OutputConfig = &outputConfig{
			AspectRatio: pickToken(in.Ratio, viduValidAspectRatios, "1:1", "Vidu"),
			Resolution:  pickToken(in.Resolution, viduValidResolutions, "1080p", "Vidu"),
		}
	case "Qwen", "Hunyuan":
		if extInfo := buildOGExtInfo(in.Size); extInfo != "" {
			body.ExtInfo = extInfo
		}
	default:
		if in.Ratio != "" {
			body.OutputConfig = &outputConfig{AspectRatio: in.Ratio}
		}
	}

	applyImageCountAndFormat(&body, tencentModelName, in.N, in.OutputFormat)
	applyImageExtra(&body, tencentModelName, in.Extra)
	body.FileInfos = imageFileInfos(tencentModelName, modelVersion, in.Images)
	return body
}

func imageModelVersion(tencentModelName, lookupModel, quality string) string {
	switch tencentModelName {
	case "GG":
		return getGGVersion(lookupModel)
	case "Kling":
		return getKlingVersion(lookupModel)
	case "Vidu":
		return getViduVersion(lookupModel)
	case "Qwen":
		return getQwenVersion(lookupModel)
	case "Hunyuan":
		return getHunyuanVersion(lookupModel)
	default:
		// OG：low/medium/high 映射到 image2_*。xhigh、max 和未传都落到 image2_low。
		return qualityToModelVersion(tencentModelName, quality)
	}
}

func pickToken(value string, allowed []string, fallback, model string) string {
	value = strings.TrimSpace(value)
	if value == "" || len(allowed) == 0 {
		return fallback
	}
	for _, item := range allowed {
		if strings.EqualFold(value, item) {
			return item
		}
	}
	common.SysLog(fmt.Sprintf("TencentVOD %s: %s is not supported, fallback to %s", model, value, fallback))
	return fallback
}

func applyImageCountAndFormat(body *createTaskRequest, tencentModelName string, n int, outputFormat string) {
	if n > 1 {
		maxCount := 0
		switch tencentModelName {
		case "OG":
			maxCount = 8
		case "Kling":
			maxCount = 9
		}
		if maxCount > 0 {
			count := n
			if count > maxCount {
				common.SysLog(fmt.Sprintf(
					"TencentVOD %s: OutputImageCount %d 超出上限 %d，已截断至 %d 张",
					tencentModelName, count, maxCount, maxCount,
				))
				count = maxCount
			}
			if body.OutputConfig == nil {
				body.OutputConfig = &outputConfig{}
			}
			body.OutputConfig.OutputImageCount = count
		}
	}
	if outputFormat != "" {
		if body.OutputConfig == nil {
			body.OutputConfig = &outputConfig{}
		}
		body.OutputConfig.OutputFormat = outputFormat
	}
}

func applyImageExtra(body *createTaskRequest, tencentModelName string, extra map[string]interface{}) {
	if len(extra) == 0 {
		return
	}
	if v, ok := extra["negative_prompt"].(string); ok {
		body.NegativePrompt = v
	}
	if seed := extraInt64(extra, "seed"); seed != 0 {
		body.Seed = seed
	}
	if v, ok := extra["enhance_prompt"].(string); ok && tencentModelName == "GG" {
		body.EnhancePrompt = v
	}
	if v, ok := extra["scene_type"].(string); ok && (tencentModelName == "Kling" || tencentModelName == "Hunyuan") {
		body.SceneType = v
	}
	if tencentModelName == "Kling" {
		if expansion, ok := extra["expansion"].(map[string]interface{}); ok {
			if extInfo := buildKlingExpansionExtInfoFromMap(expansion); extInfo != "" {
				body.ExtInfo = extInfo
			}
		}
	}
}

func imageFileInfos(tencentModelName, modelVersion string, images []string) []fileInfo {
	maxRef := getMaxRefImages(tencentModelName, modelVersion)
	if maxRef > 0 && len(images) > maxRef {
		common.SysLog(fmt.Sprintf(
			"TencentVOD %s %s: 参考图数量 %d 超出上限 %d，已截断至 %d 张",
			tencentModelName, modelVersion, len(images), maxRef, maxRef,
		))
		images = images[:maxRef]
	}
	var files []fileInfo
	for _, url := range images {
		if strings.TrimSpace(url) != "" {
			files = append(files, fileInfo{Type: "Url", Url: url})
		}
	}
	return files
}

func extraInt64(extra map[string]interface{}, key string) int64 {
	switch v := extra[key].(type) {
	case float64:
		return int64(v)
	case int:
		return int64(v)
	case int64:
		return v
	default:
		return 0
	}
}
