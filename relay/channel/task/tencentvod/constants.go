package tencentvod

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	relaycommon "github.com/QuantumNous/new-api/relay/common"
)

const ChannelName = "TencentVOD"

// ModelList exposes all models supported by this channel.
// Add more Tencent VOD AIGC models here as needed.
var ModelList = []string{
	"gpt-image-2",
	// GG (Gemini) 系列生图模型（腾讯 VOD ModelName="GG"）
	// 系列          官方模型ID                    Nano Banana 名称
	// Gemini 2.5 Flash Image  gemini-2.5-flash-image  Nano Banana
	// Gemini 3 Pro Image      gemini-3-pro-image      Nano Banana Pro
	// Gemini 3.1 Flash Image  gemini-3.1-flash-image  Nano Banana 2
	"gemini-2.5-flash-image",
	"gemini-3-pro-image",
	"gemini-3.1-flash-image",
	// Kling 系列生图模型（腾讯 VOD ModelName="Kling"）
	// kling-image-scene 为扩图模式，需配合 metadata.scene_type="image_expand" 使用
	"kling-image-2.1",
	"kling-image-3.0",
	"kling-image-3.0-omni",
	"kling-image-o1",
	"kling-image-scene",
	// Vidu 系列生图模型（腾讯 VOD ModelName="Vidu"）
	"vidu-image-q2",
	// Qwen 系列生图模型（腾讯 VOD ModelName="Qwen"）
	// Qwen 0925 使用 ExtInfo 传自定义像素尺寸，不使用 OutputConfig.AspectRatio/Resolution
	"qwen-image-0925",
	// Hunyuan 系列生图模型（腾讯 VOD ModelName="Hunyuan"）
	// Hunyuan 3.0 使用 ExtInfo 传自定义像素尺寸；支持全景图模式 via extra_fields.scene_type="3d_panorama"
	"hunyuan-image-3.0",
}

// modelNameMap maps origin model name to Tencent VOD API ModelName field.
// Tencent VOD valid ModelName values include: OG, GG, Kling, Hunyuan, SI, Vidu, Qwen, MJ, etc.
var modelNameMap = map[string]string{
	"gpt-image-2": "OG",
	// GG 系列：三个子模型均映射到 ModelName="GG"，版本通过 ggVersionMap 区分
	"gemini-2.5-flash-image": "GG",
	"gemini-3-pro-image":     "GG",
	"gemini-3.1-flash-image": "GG",
	// Kling 系列：五个子模型均映射到 ModelName="Kling"，版本通过 klingVersionMap 区分
	"kling-image-2.1":      "Kling",
	"kling-image-3.0":      "Kling",
	"kling-image-3.0-omni": "Kling",
	"kling-image-o1":       "Kling",
	"kling-image-scene":    "Kling",
	// Vidu 系列
	"vidu-image-q2": "Vidu",
	// Qwen 系列
	"qwen-image-0925": "Qwen",
	// Hunyuan 系列
	"hunyuan-image-3.0": "Hunyuan",
}

// ggVersionMap maps our origin model name to Tencent VOD GG ModelVersion.
// GG (Gemini) valid versions: 2.5, 3.0, 3.1
var ggVersionMap = map[string]string{
	"gemini-2.5-flash-image": "2.5",
	"gemini-3-pro-image":     "3.0",
	"gemini-3.1-flash-image": "3.1",
}

// ggMaxRefImages maps GG ModelVersion to maximum number of reference images allowed.
// Reference: 腾讯 VOD API FileInfos.N 各模型支持最大参考图数量
//
//	GG 2.5：3张；GG 3.0：14张；GG 3.1：14张
var ggMaxRefImages = map[string]int{
	"2.5": 3,
	"3.0": 14,
	"3.1": 14,
}

// ggValidAspectRatios lists valid AspectRatio values per GG ModelVersion.
// Reference: 腾讯 VOD OutputConfig.AspectRatio 各 GG 版本可选值
//
//	GG 2.5 / GG 3.0: 1:1, 2:3, 3:2, 3:4, 4:3, 4:5, 5:4, 9:16, 16:9, 21:9
//	GG 3.1: 在上述基础上额外支持 1:4, 1:8, 4:1, 8:1（超长/超宽比例）
var ggValidAspectRatios = map[string][]string{
	"2.5": {"1:1", "2:3", "3:2", "3:4", "4:3", "4:5", "5:4", "9:16", "16:9", "21:9"},
	"3.0": {"1:1", "2:3", "3:2", "3:4", "4:3", "4:5", "5:4", "9:16", "16:9", "21:9"},
	// GG 3.1 额外支持极端宽高比 1:4, 1:8, 4:1, 8:1
	"3.1": {"1:1", "1:4", "1:8", "2:3", "3:2", "3:4", "4:1", "4:3", "4:5", "5:4", "8:1", "9:16", "16:9", "21:9"},
}

// ggValidResolutions lists valid Resolution values per GG ModelVersion.
// Reference: 腾讯 VOD OutputConfig.Resolution 各 GG 版本可选值
//
//	GG 2.5 / GG 3.0: 1K, 2K, 4K
//	GG 3.1: 720P, 1K, 2K, 4K（额外支持 720P）
var ggValidResolutions = map[string][]string{
	"2.5": {"1K", "2K", "4K"},
	"3.0": {"1K", "2K", "4K"},
	// GG 3.1 额外支持 720P 分辨率
	"3.1": {"720P", "1K", "2K", "4K"},
}

// ggQualityResolutionMap maps OpenAI-style quality values to GG Resolution tiers.
// 注意：GG 的 quality 映射到 Resolution（1K/2K/4K），与 OG 映射到 ModelVersion 不同。
// 用户也可以直接传入 "1K"/"2K"/"4K"/"720P" 作为 quality 值。
var ggQualityResolutionMap = map[string]string{
	// OpenAI 标准 quality 值映射
	"low":      "1K",
	"standard": "1K",
	"medium":   "2K",
	"hd":       "2K",
	"high":     "2K",
	// 直接传分辨率档位（大小写均可，getGGResolution 会先 ToUpper 归一化）
	"720p": "720P",
	"720P": "720P",
	"1k":   "1K",
	"1K":   "1K",
	"2k":   "2K",
	"2K":   "2K",
	"4k":   "4K",
	"4K":   "4K",
}

// ggAspectRatioFloatMap 用于像素尺寸到宽高比的最近邻查找（包含所有 GG 版本支持的宽高比）。
// key = 宽高比字符串, value = 对应的 float64 数值（宽/高）
var ggAspectRatioFloatMap = map[string]float64{
	"1:1":  1.0 / 1.0,
	"2:3":  2.0 / 3.0,
	"3:2":  3.0 / 2.0,
	"3:4":  3.0 / 4.0,
	"4:3":  4.0 / 3.0,
	"4:5":  4.0 / 5.0,
	"5:4":  5.0 / 4.0,
	"9:16": 9.0 / 16.0,
	"16:9": 16.0 / 9.0,
	"21:9": 21.0 / 9.0,
	// GG 3.1 额外值
	"1:4": 1.0 / 4.0,
	"1:8": 1.0 / 8.0,
	"4:1": 4.0 / 1.0,
	"8:1": 8.0 / 1.0,
}

// ──────────────────────────────────────────────────────────────────────────────
// Kling 系列生图模型（腾讯 VOD ModelName="Kling"）
// ──────────────────────────────────────────────────────────────────────────────

// klingVersionMap maps origin model name to Tencent VOD Kling ModelVersion.
// Kling valid versions: 2.1, 3.0, 3.0-Omni, O1, scene
var klingVersionMap = map[string]string{
	"kling-image-2.1":      "2.1",
	"kling-image-3.0":      "3.0",
	"kling-image-3.0-omni": "3.0-Omni",
	"kling-image-o1":       "O1",
	// scene 为 Kling 扩图模式，需配合 SceneType="image_expand" 使用；
	// 扩图参数通过 extra_fields.scene_type + extra_fields.expansion 传入；
	// 各比例取值范围 [0,2]，新图面积不超过原图 3 倍
	"kling-image-scene": "scene",
}

// klingMaxRefImages maps Kling ModelVersion to max reference images allowed.
// Reference: 腾讯 VOD API FileInfos.N 各 Kling 版本最大参考图数量
//
//	Kling 2.1：4张；Kling 3.0：1张；Kling 3.0-Omni：10张；Kling O1：10张
//	Kling scene：扩图模式以单张原图为基础，上限设为 1（扩图只需 1 张原图）
var klingMaxRefImages = map[string]int{
	"2.1":     4,
	"3.0":     1,
	"3.0-Omni": 10,
	"O1":      10,
	"scene":   1,
}

// klingValidAspectRatios lists valid AspectRatio values per Kling version.
// Reference: 腾讯 VOD OutputConfig.AspectRatio 各 Kling 版本可选值
//
//	Kling 2.1/3.0:     16:9, 9:16, 1:1, 4:3, 3:4, 3:2, 2:3, 21:9
//	Kling 3.0-Omni/O1: 在上述基础上额外支持 "auto"（模型自动选择最佳宽高比）
//	Kling scene:       扩图模式由 ExtInfo 扩图参数决定输出尺寸，不使用 AspectRatio
var klingValidAspectRatios = map[string][]string{
	"2.1": {"1:1", "2:3", "3:2", "3:4", "4:3", "9:16", "16:9", "21:9"},
	"3.0": {"1:1", "2:3", "3:2", "3:4", "4:3", "9:16", "16:9", "21:9"},
	// Kling 3.0-Omni 和 O1 额外支持 "auto"
	"3.0-Omni": {"1:1", "2:3", "3:2", "3:4", "4:3", "9:16", "16:9", "21:9", "auto"},
	"O1":       {"1:1", "2:3", "3:2", "3:4", "4:3", "9:16", "16:9", "21:9", "auto"},
	// scene 扩图模式不使用 AspectRatio
	"scene": nil,
}

// klingValidResolutions lists valid Resolution values per Kling version.
// Reference: 腾讯 VOD OutputConfig.Resolution 各 Kling 版本可选值
//
// 注意：Kling 系列 Resolution 使用小写（1k/2k/4k），与 GG 系列大写（1K/2K/4K）不同，
// 需与腾讯 VOD API 文档保持一致。
//
//	Kling 2.1/3.0:     1k, 2k
//	Kling 3.0-Omni/O1: 1k, 2k, 4k
//	Kling scene:       扩图模式不使用 Resolution
var klingValidResolutions = map[string][]string{
	"2.1":     {"1k", "2k"},
	"3.0":     {"1k", "2k"},
	"3.0-Omni": {"1k", "2k", "4k"},
	"O1":      {"1k", "2k", "4k"},
	"scene":   nil, // scene 扩图模式不使用 Resolution
}

// klingQualityResolutionMap maps OpenAI-style quality values to Kling Resolution tiers.
// 注意：Kling Resolution 使用小写（1k/2k/4k），与 GG 大写（1K/2K/4K）不同。
var klingQualityResolutionMap = map[string]string{
	"low":      "1k",
	"standard": "1k",
	"medium":   "1k",
	"hd":       "2k",
	"high":     "2k",
	// 直接传档位值（兼容大小写输入，统一输出小写）
	"1k": "1k",
	"1K": "1k",
	"2k": "2k",
	"2K": "2k",
	"4k": "4k",
	"4K": "4k",
}

// klingAspectRatioFloatMap 用于 Kling 像素尺寸到宽高比的最近邻查找。
// 注意：Kling 不支持 GG 系列的 4:5、5:4、1:4、1:8、4:1、8:1 等比例
var klingAspectRatioFloatMap = map[string]float64{
	"1:1":  1.0 / 1.0,
	"2:3":  2.0 / 3.0,
	"3:2":  3.0 / 2.0,
	"3:4":  3.0 / 4.0,
	"4:3":  4.0 / 3.0,
	"9:16": 9.0 / 16.0,
	"16:9": 16.0 / 9.0,
	"21:9": 21.0 / 9.0,
}

// klingExpansionParams Kling 扩图参数，仅 scene 版本（SceneType="image_expand"）时有效。
// Reference: 腾讯 VOD API ExtInfo.AdditionalParameters 扩图参数说明
//
//	各方向取值范围 [0, 2]，新图整体面积不得超过原图的 3 倍。
//	up_expansion_ratio:    向上扩充范围，基于原图高度的倍数（原图高 H，值为 r 则顶边距新图顶 H*r）
//	down_expansion_ratio:  向下扩充范围，基于原图高度的倍数
//	left_expansion_ratio:  向左扩充范围，基于原图宽度的倍数
//	right_expansion_ratio: 向右扩充范围，基于原图宽度的倍数
//
// 通过 extra_fields.expansion 传入
type klingExpansionParams struct {
	Up    float64 `json:"up"`
	Down  float64 `json:"down"`
	Left  float64 `json:"left"`
	Right float64 `json:"right"`
}

// ──────────────────────────────────────────────────────────────────────────────
// Vidu 系列生图模型（腾讯 VOD ModelName="Vidu"）
// ──────────────────────────────────────────────────────────────────────────────

// viduVersionMap maps origin model name to Tencent VOD Vidu ModelVersion.
var viduVersionMap = map[string]string{
	"vidu-image-q2": "q2",
}

// viduMaxRefImages maps Vidu ModelVersion to max reference images allowed.
// Reference: 腾讯 VOD API FileInfos.N
//
//	Vidu q2：7 张
var viduMaxRefImages = map[string]int{
	"q2": 7,
}

// viduValidAspectRatios lists valid AspectRatio values for Vidu models (all versions).
// Reference: 腾讯 VOD OutputConfig.AspectRatio Vidu q2 可选值
//
//	Vidu q2: 16:9、9:16、1:1、3:4、4:3、21:9、2:3、3:2
var viduValidAspectRatios = []string{"1:1", "2:3", "3:2", "3:4", "4:3", "9:16", "16:9", "21:9"}

// viduValidResolutions lists valid Resolution values for Vidu models (all versions).
// Reference: 腾讯 VOD OutputConfig.Resolution Vidu q2 可选值
//
// 注意：Vidu Resolution 大小写特殊："1080p"（小写 p）/ "2K"（大写 K）/ "4K"（大写 K），
// 与 GG 的全大写（1K/2K/4K）和 Kling 的全小写（1k/2k/4k）均不同，需严格匹配腾讯 API 文档。
//
//	Vidu q2: 1080p, 2K, 4K，默认 1080p
var viduValidResolutions = []string{"1080p", "2K", "4K"}

// viduQualityResolutionMap maps OpenAI-style quality values to Vidu Resolution tiers.
// 注意：Vidu Resolution 使用 "1080p"/"2K"/"4K"，大小写与 GG 和 Kling 均不同。
var viduQualityResolutionMap = map[string]string{
	"low":      "1080p",
	"standard": "1080p",
	"medium":   "2K",
	"hd":       "2K",
	"high":     "2K",
	// 直接传档位（兼容不同大小写输入）
	"1080p": "1080p",
	"1080P": "1080p",
	"2k":    "2K",
	"2K":    "2K",
	"4k":    "4K",
	"4K":    "4K",
}

// ──────────────────────────────────────────────────────────────────────────────
// Qwen 系列生图模型（腾讯 VOD ModelName="Qwen"）
// ──────────────────────────────────────────────────────────────────────────────

// qwenVersionMap maps origin model name to Tencent VOD Qwen ModelVersion.
var qwenVersionMap = map[string]string{
	"qwen-image-0925": "0925",
}

// qwenMaxRefImages maps Qwen ModelVersion to max reference images allowed.
// Reference: 腾讯 VOD API FileInfos.N
//
//	Qwen 0925：1 张（仅支持单张参考图）
var qwenMaxRefImages = map[string]int{
	"0925": 1,
}

// ──────────────────────────────────────────────────────────────────────────────
// Hunyuan 系列生图模型（腾讯 VOD ModelName="Hunyuan"）
// ──────────────────────────────────────────────────────────────────────────────

// hunyuanVersionMap maps origin model name to Tencent VOD Hunyuan ModelVersion.
var hunyuanVersionMap = map[string]string{
	"hunyuan-image-3.0": "3.0",
}

// hunyuanMaxRefImages maps Hunyuan ModelVersion to max reference images allowed.
// Reference: 腾讯 VOD API FileInfos.N
//
//	Hunyuan 3.0：3 张
var hunyuanMaxRefImages = map[string]int{
	"3.0": 3,
}

// ──────────────────────────────────────────────────────────────────────────────
// OG (gpt-image-2) 参数映射
// ──────────────────────────────────────────────────────────────────────────────

// ogQualityMap maps OpenAI-style quality values to Tencent VOD OG ModelVersion.
// OG (gpt-image-2) valid versions: image2_low, image2_medium, image2_high
var ogQualityMap = map[string]string{
	"low":      "image2_low",
	"standard": "image2_low",
	"medium":   "image2_medium",
	"hd":       "image2_high",
	"high":     "image2_high",
}

// sizeMap maps common OpenAI-style pixel size strings to Tencent VOD AspectRatio.
// 适用于 OG（ExtInfo 路径会绕过此表）以及 GG（无直接匹配时使用像素比最近邻算法）。
var sizeMap = map[string]string{
	"1024x1024": "1:1",
	"1024x1536": "2:3",
	"1536x1024": "3:2",
	"1024x1792": "9:16",
	"1792x1024": "16:9",
	"auto":      "1:1",
}

// ──────────────────────────────────────────────────────────────────────────────
// Helper functions
// ──────────────────────────────────────────────────────────────────────────────

// resolveModelName returns the effective model name for Tencent API lookups.
// When a channel model mapping is configured (e.g. "OG-image-2" → "gpt-image-2"),
// UpstreamModelName holds the mapped target; otherwise it falls back to OriginModelName.
// All model dispatch and version lookups should use this instead of OriginModelName directly.
func resolveModelName(info *relaycommon.RelayInfo) string {
	if info.UpstreamModelName != "" {
		return info.UpstreamModelName
	}
	return info.OriginModelName
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

// getGGVersion 根据原始模型名返回 GG ModelVersion（"2.5" / "3.0" / "3.1"）。
func getGGVersion(originModelName string) string {
	if v, ok := ggVersionMap[originModelName]; ok {
		return v
	}
	return "2.5" // 默认回退到最低版本
}

// getGGMaxRefImages 返回指定 GG 版本允许的最大参考图数量。
func getGGMaxRefImages(version string) int {
	if n, ok := ggMaxRefImages[version]; ok {
		return n
	}
	return 3 // 默认取最保守值
}

// getGGResolution 将 quality 字符串映射到 GG 的 Resolution 档位（"1K"/"2K"/"4K"/"720P"）。
//
// 特殊说明：
//   - "720P" 仅 GG 3.1 支持，若在 GG 2.5/3.0 使用会返回降级到 "1K" 并返回错误。
//   - 空字符串/未识别值均默认返回 "1K"。
func getGGResolution(version, quality string) (string, error) {
	resolution := "1K" // default
	if q := strings.ToUpper(strings.TrimSpace(quality)); q != "" {
		if r, ok := ggQualityResolutionMap[q]; ok {
			resolution = r
		} else if r, ok := ggQualityResolutionMap[strings.ToLower(quality)]; ok {
			resolution = r
		}
		// else: 未识别的 quality 保持默认 1K
	}

	// 校验 version 是否支持该 Resolution
	validResolutions, ok := ggValidResolutions[version]
	if !ok {
		return resolution, nil // 未知版本不校验
	}
	for _, r := range validResolutions {
		if r == resolution {
			return resolution, nil
		}
	}
	// 不支持的 Resolution（如 GG 2.5/3.0 传 720P）降级到 1K 并报错
	return "1K", fmt.Errorf("GG %s 不支持 Resolution=%s，已降级到 1K；GG 3.1 支持 720P", version, resolution)
}

// getGGAspectRatio 将 size 字符串映射到 GG 的 AspectRatio。
//
// 处理规则（优先级从高到低）：
//  1. 空/"auto" → "1:1"
//  2. size 本身就是合法的宽高比字符串（如 "16:9"）→ 直接使用并校验
//  3. size 是像素尺寸（如 "1024x1792"）→ 先查 sizeMap，未找到则用最近邻算法计算最接近的宽高比
//  4. 其他无法识别的格式 → 返回 "1:1" 并报错（非致命，调用方可记录日志继续）
//
// 注意：GG 各版本支持的宽高比不同，传入不支持的值会返回 error；
// GG 3.1 额外支持极端比例 1:4、1:8、4:1、8:1。
func getGGAspectRatio(version, size string) (string, error) {
	if size == "" || size == "auto" {
		return "1:1", nil
	}

	validRatios, hasValidList := ggValidAspectRatios[version]

	// Case 1: size 本身是宽高比格式（包含 ":" 但不包含 "x"）
	if strings.Contains(size, ":") && !strings.Contains(size, "x") {
		if hasValidList {
			for _, r := range validRatios {
				if r == size {
					return size, nil
				}
			}
			return "1:1", fmt.Errorf("GG %s 不支持 AspectRatio=%s，支持值: %v；已回退到 1:1",
				version, size, validRatios)
		}
		return size, nil
	}

	// Case 2: size 是像素尺寸格式
	if strings.Contains(size, "x") {
		// 先查快速映射表
		if ratio, ok := sizeMap[size]; ok {
			if err := validateGGAspectRatio(version, ratio, validRatios, hasValidList); err != nil {
				return "1:1", err
			}
			return ratio, nil
		}
		// 解析宽高，寻找最近邻宽高比
		parts := strings.SplitN(size, "x", 2)
		if len(parts) == 2 {
			w, errW := strconv.ParseFloat(strings.TrimSpace(parts[0]), 64)
			h, errH := strconv.ParseFloat(strings.TrimSpace(parts[1]), 64)
			if errW == nil && errH == nil && h > 0 {
				ratio := findClosestGGAspectRatio(w/h, validRatios, hasValidList)
				return ratio, nil
			}
		}
	}

	// Case 3: 无法识别
	return "1:1", fmt.Errorf("GG：无法识别 size=%s，已回退到 1:1", size)
}

// validateGGAspectRatio 检查宽高比是否在 validRatios 列表中。
func validateGGAspectRatio(version, ratio string, validRatios []string, hasValidList bool) error {
	if !hasValidList {
		return nil
	}
	for _, r := range validRatios {
		if r == ratio {
			return nil
		}
	}
	return fmt.Errorf("GG %s 不支持 AspectRatio=%s，支持值: %v；已回退到 1:1", version, ratio, validRatios)
}

// findClosestGGAspectRatio 从候选列表中找到与目标比值（宽/高）最接近的宽高比字符串。
// 若没有候选列表，则从全量 ggAspectRatioFloatMap 中查找。
func findClosestGGAspectRatio(targetRatio float64, validRatios []string, hasValidList bool) string {
	bestRatio := "1:1"
	bestDiff := math.MaxFloat64

	candidates := validRatios
	if !hasValidList {
		// 如果没有 version 限制，从全量中找
		for r := range ggAspectRatioFloatMap {
			candidates = append(candidates, r)
		}
	}

	for _, r := range candidates {
		if val, ok := ggAspectRatioFloatMap[r]; ok {
			diff := math.Abs(val - targetRatio)
			if diff < bestDiff {
				bestDiff = diff
				bestRatio = r
			}
		}
	}
	return bestRatio
}

// ──────────────────────────────────────────────────────────────────────────────
// Kling helper functions
// ──────────────────────────────────────────────────────────────────────────────

// getKlingVersion 根据原始模型名返回 Kling ModelVersion。
func getKlingVersion(originModelName string) string {
	if v, ok := klingVersionMap[originModelName]; ok {
		return v
	}
	return "2.1" // 默认回退到最低版本
}

// getKlingMaxRefImages 返回指定 Kling 版本允许的最大参考图数量。
func getKlingMaxRefImages(version string) int {
	if n, ok := klingMaxRefImages[version]; ok {
		return n
	}
	return 1 // 默认取最保守值
}

// getKlingAspectRatio 将 size 字符串映射到 Kling 的 AspectRatio。
//
// 处理规则：
//  1. scene 版本 → 返回 "" （扩图模式不使用 AspectRatio）
//  2. 空/"auto" → "auto"（仅 3.0-Omni/O1 支持）或 "1:1"（其他版本）
//  3. size 包含 ":" → 直接使用并校验
//  4. size 包含 "x" → 查 sizeMap，未找到则最近邻算法
//  5. 其他 → 返回 "1:1" 并报错
func getKlingAspectRatio(version, size string) (string, error) {
	// scene 扩图模式不使用 AspectRatio
	if version == "scene" {
		return "", nil
	}

	validRatios, hasValidList := klingValidAspectRatios[version]

	if size == "" {
		return "1:1", nil
	}
	if size == "auto" {
		// auto 仅 Kling 3.0-Omni 和 O1 支持
		if hasValidList {
			for _, r := range validRatios {
				if r == "auto" {
					return "auto", nil
				}
			}
		}
		return "1:1", nil // 其他版本 auto → 1:1
	}

	// size 是宽高比格式
	if strings.Contains(size, ":") && !strings.Contains(size, "x") {
		if hasValidList {
			for _, r := range validRatios {
				if r == size {
					return size, nil
				}
			}
			return "1:1", fmt.Errorf("Kling %s 不支持 AspectRatio=%s，支持值: %v；已回退到 1:1",
				version, size, validRatios)
		}
		return size, nil
	}

	// size 是像素尺寸格式
	if strings.Contains(size, "x") {
		if ratio, ok := sizeMap[size]; ok {
			if hasValidList {
				for _, r := range validRatios {
					if r == ratio {
						return ratio, nil
					}
				}
			}
			// sizeMap 命中但比例不在 Kling 支持列表，用最近邻
		}
		parts := strings.SplitN(size, "x", 2)
		if len(parts) == 2 {
			w, errW := strconv.ParseFloat(strings.TrimSpace(parts[0]), 64)
			h, errH := strconv.ParseFloat(strings.TrimSpace(parts[1]), 64)
			if errW == nil && errH == nil && h > 0 {
				return findClosestAspectRatio(w/h, validRatios, klingAspectRatioFloatMap), nil
			}
		}
	}

	return "1:1", fmt.Errorf("Kling：无法识别 size=%s，已回退到 1:1", size)
}

// getKlingResolution 将 quality 字符串映射到 Kling 的 Resolution 档位。
//
// 特殊说明：
//   - scene 版本不使用 Resolution（扩图模式）
//   - Kling Resolution 使用小写（1k/2k/4k），与 GG 大写不同
//   - 默认值为 "1k"（最低档）
func getKlingResolution(version, quality string) (string, error) {
	// scene 扩图模式不使用 Resolution
	if version == "scene" {
		return "", nil
	}

	resolution := "1k" // default
	if q := strings.TrimSpace(quality); q != "" {
		if r, ok := klingQualityResolutionMap[q]; ok {
			resolution = r
		} else if r, ok := klingQualityResolutionMap[strings.ToLower(q)]; ok {
			resolution = r
		}
		// else: 未识别的 quality 保持默认 1k
	}

	// 校验 version 是否支持该 Resolution
	validResolutions, ok := klingValidResolutions[version]
	if !ok || validResolutions == nil {
		return resolution, nil
	}
	for _, r := range validResolutions {
		if r == resolution {
			return resolution, nil
		}
	}
	// 不支持的 Resolution（如 Kling 2.1/3.0 传 4k）降级到 1k 并报错
	return "1k", fmt.Errorf("Kling %s 不支持 Resolution=%s，已降级到 1k；支持值: %v",
		version, resolution, validResolutions)
}

// buildKlingExpansionExtInfo builds ExtInfo JSON for Kling image expansion.
// Kling 扩图参数通过 ExtInfo.AdditionalParameters 传入。
//
// Reference: 腾讯 VOD API ExtInfo 扩图参数格式
// 特殊说明：仅 Kling scene 版本（SceneType="image_expand"）时使用；
// 各方向扩充比例均为 0 时返回空字符串（不设置 ExtInfo）
func buildKlingExpansionExtInfo(params *klingExpansionParams) string {
	if params == nil {
		return ""
	}
	if params.Up == 0 && params.Down == 0 && params.Left == 0 && params.Right == 0 {
		return ""
	}
	inner := fmt.Sprintf(
		`{"down_expansion_ratio":%g,"left_expansion_ratio":%g,"right_expansion_ratio":%g,"up_expansion_ratio":%g}`,
		params.Down, params.Left, params.Right, params.Up,
	)
	return fmt.Sprintf(`{"AdditionalParameters":"%s"}`, escapeJSONString(inner))
}

// buildKlingExpansionExtInfoFromMap builds Kling expansion ExtInfo from a raw interface map.
// 用于同步路径（adaptor.go）从 extra_fields.expansion 读取扩图参数时使用。
func buildKlingExpansionExtInfoFromMap(expansion map[string]interface{}) string {
	if len(expansion) == 0 {
		return ""
	}
	params := &klingExpansionParams{}
	if v, ok := expansion["up"].(float64); ok {
		params.Up = v
	}
	if v, ok := expansion["down"].(float64); ok {
		params.Down = v
	}
	if v, ok := expansion["left"].(float64); ok {
		params.Left = v
	}
	if v, ok := expansion["right"].(float64); ok {
		params.Right = v
	}
	return buildKlingExpansionExtInfo(params)
}

// ──────────────────────────────────────────────────────────────────────────────
// Vidu helper functions
// ──────────────────────────────────────────────────────────────────────────────

// getViduVersion 根据原始模型名返回 Vidu ModelVersion。
func getViduVersion(originModelName string) string {
	if v, ok := viduVersionMap[originModelName]; ok {
		return v
	}
	return "q2"
}

// getViduMaxRefImages 返回 Vidu 版本允许的最大参考图数量。
func getViduMaxRefImages(version string) int {
	if n, ok := viduMaxRefImages[version]; ok {
		return n
	}
	return 7
}

// getViduAspectRatio 将 size 字符串映射到 Vidu 的 AspectRatio（有效性校验）。
// Vidu 所有版本共享相同的宽高比支持列表。
func getViduAspectRatio(size string) (string, error) {
	if size == "" || size == "auto" {
		return "1:1", nil
	}

	// size 是宽高比格式
	if strings.Contains(size, ":") && !strings.Contains(size, "x") {
		for _, r := range viduValidAspectRatios {
			if r == size {
				return size, nil
			}
		}
		return "1:1", fmt.Errorf("Vidu 不支持 AspectRatio=%s，支持值: %v；已回退到 1:1",
			size, viduValidAspectRatios)
	}

	// size 是像素尺寸格式
	if strings.Contains(size, "x") {
		if ratio, ok := sizeMap[size]; ok {
			for _, r := range viduValidAspectRatios {
				if r == ratio {
					return ratio, nil
				}
			}
		}
		parts := strings.SplitN(size, "x", 2)
		if len(parts) == 2 {
			w, errW := strconv.ParseFloat(strings.TrimSpace(parts[0]), 64)
			h, errH := strconv.ParseFloat(strings.TrimSpace(parts[1]), 64)
			if errW == nil && errH == nil && h > 0 {
				return findClosestAspectRatio(w/h, viduValidAspectRatios, ggAspectRatioFloatMap), nil
			}
		}
	}

	return "1:1", fmt.Errorf("Vidu：无法识别 size=%s，已回退到 1:1", size)
}

// getViduResolution 将 quality 字符串映射到 Vidu 的 Resolution 档位。
// 注意：Vidu Resolution 格式为 "1080p"/"2K"/"4K"，与 GG/Kling 不同。
func getViduResolution(quality string) (string, error) {
	resolution := "1080p" // default
	if q := strings.TrimSpace(quality); q != "" {
		if r, ok := viduQualityResolutionMap[q]; ok {
			resolution = r
		} else if r, ok := viduQualityResolutionMap[strings.ToLower(q)]; ok {
			resolution = r
		}
	}
	for _, r := range viduValidResolutions {
		if r == resolution {
			return resolution, nil
		}
	}
	return "1080p", fmt.Errorf("Vidu 不支持 Resolution=%s，已降级到 1080p；支持值: %v",
		resolution, viduValidResolutions)
}

// ──────────────────────────────────────────────────────────────────────────────
// Qwen / Hunyuan helper functions
// ──────────────────────────────────────────────────────────────────────────────

// getQwenVersion 根据原始模型名返回 Qwen ModelVersion。
func getQwenVersion(originModelName string) string {
	if v, ok := qwenVersionMap[originModelName]; ok {
		return v
	}
	return "0925"
}

// getHunyuanVersion 根据原始模型名返回 Hunyuan ModelVersion。
func getHunyuanVersion(originModelName string) string {
	if v, ok := hunyuanVersionMap[originModelName]; ok {
		return v
	}
	return "3.0"
}

// ──────────────────────────────────────────────────────────────────────────────
// 通用辅助函数
// ──────────────────────────────────────────────────────────────────────────────

// getMaxRefImages 返回指定模型允许的最大参考图数量。
// 返回 -1 表示无上限限制（或未知模型）。
// 适用于所有 TencentVOD 模型的统一参考图数量校验。
func getMaxRefImages(tencentModelName, version string) int {
	switch tencentModelName {
	case "GG":
		return getGGMaxRefImages(version)
	case "Kling":
		return getKlingMaxRefImages(version)
	case "Vidu":
		return getViduMaxRefImages(version)
	case "Qwen":
		if n, ok := qwenMaxRefImages[version]; ok {
			return n
		}
		return 1
	case "Hunyuan":
		if n, ok := hunyuanMaxRefImages[version]; ok {
			return n
		}
		return 3
	}
	return -1
}

// findClosestAspectRatio 从候选列表中找到与目标比值最接近的宽高比字符串。
// 与 findClosestGGAspectRatio 功能相同，但接受自定义浮点映射表，供 Kling/Vidu 等模型使用。
//
// 注意：不要删除 findClosestGGAspectRatio，该函数仍被 getGGAspectRatio 使用。
func findClosestAspectRatio(targetRatio float64, validRatios []string, floatMap map[string]float64) string {
	bestRatio := "1:1"
	bestDiff := math.MaxFloat64
	for _, r := range validRatios {
		if val, ok := floatMap[r]; ok {
			diff := math.Abs(val - targetRatio)
			if diff < bestDiff {
				bestDiff = diff
				bestRatio = r
			}
		}
	}
	return bestRatio
}

func sizeToAspectRatio(size string) string {
	if v, ok := sizeMap[size]; ok {
		return v
	}
	return "1:1"
}

// buildOGExtInfo builds the ExtInfo JSON string for OG model size configuration.
// OG 支持通过 ExtInfo.AdditionalParameters.size 传入自定义像素尺寸（非宽高比）。
// 特殊说明：
//   - 总像素数必须在 [655,360, 8,294,400] 范围内，且需被 16 整除
//   - GG 系列不走此路径，GG 使用 OutputConfig.Resolution + AspectRatio
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
