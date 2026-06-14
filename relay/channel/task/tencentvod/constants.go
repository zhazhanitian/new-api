package tencentvod

import "fmt"

const ChannelName = "TencentVOD"

// ModelList exposes all models supported by this channel.
// Add more Tencent VOD AIGC models here as needed (e.g. "kling-v1", "hunyuan-image").
var ModelList = []string{"gpt-image-2"}

// modelNameMap maps origin model name to Tencent VOD API ModelName field.
// Tencent VOD valid ModelName values include: OG, GEM, Kling, Hunyuan, SI, Vidu, Qwen, MJ, etc.
var modelNameMap = map[string]string{
	"gpt-image-2": "OG",
}

// ogQualityMap maps OpenAI-style quality values to Tencent VOD OG ModelVersion.
// OG (gpt-image-2) valid versions: image2_low, image2_medium, image2_high
var ogQualityMap = map[string]string{
	"low":      "image2_low",
	"standard": "image2_low",
	"medium":   "image2_medium",
	"hd":       "image2_high",
	"high":     "image2_high",
}

// sizeMap maps OpenAI-style size strings to Tencent VOD AspectRatio.
var sizeMap = map[string]string{
	"1024x1024": "1:1",
	"1024x1536": "2:3",
	"1536x1024": "3:2",
	"1024x1792": "9:16",
	"1792x1024": "16:9",
	"auto":      "1:1",
}

func modelToTencentName(originModel string) string {
	if v, ok := modelNameMap[originModel]; ok {
		return v
	}
	return originModel
}

func qualityToModelVersion(modelName, quality string) string {
	switch modelName {
	case "OG":
		if v, ok := ogQualityMap[quality]; ok {
			return v
		}
		return "image2_low"
	default:
		// Generic fallback: pass quality as-is, adaptor should map per model
		return quality
	}
}

func sizeToAspectRatio(size string) string {
	if v, ok := sizeMap[size]; ok {
		return v
	}
	return "1:1"
}

// buildOGExtInfo builds the ExtInfo JSON string for OG model size configuration.
// OG requires pixel dimensions passed via ExtInfo.AdditionalParameters.
func buildOGExtInfo(size string) string {
	if size == "" || size == "auto" {
		return ""
	}
	inner := fmt.Sprintf(`{"size":"%s"}`, size)
	return fmt.Sprintf(`{"AdditionalParameters":"%s"}`, escapeJSONString(inner))
}

func escapeJSONString(s string) string {
	result := make([]byte, 0, len(s)+10)
	for i := 0; i < len(s); i++ {
		if s[i] == '"' {
			result = append(result, '\\', '"')
		} else if s[i] == '\\' {
			result = append(result, '\\', '\\')
		} else {
			result = append(result, s[i])
		}
	}
	return string(result)
}
