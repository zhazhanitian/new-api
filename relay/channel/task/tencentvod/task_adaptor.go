package tencentvod

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/QuantumNous/new-api/common"
	"github.com/QuantumNous/new-api/constant"
	"github.com/QuantumNous/new-api/dto"
	"github.com/QuantumNous/new-api/model"
	"github.com/QuantumNous/new-api/relay/channel"
	taskcommon "github.com/QuantumNous/new-api/relay/channel/task/taskcommon"
	relaycommon "github.com/QuantumNous/new-api/relay/common"
	relayconstant "github.com/QuantumNous/new-api/relay/constant"
	"github.com/QuantumNous/new-api/service"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

// ── Request / Response structures ────────────────────────────────────────────

// outputConfig 对应腾讯 VOD API 的 AigcImageOutputConfig。
//
// 特殊说明（各字段模型支持情况）：
//   - AspectRatio：OG / GG / Kling 等均支持；SI 系列不支持（需在 prompt 里指定）
//   - Resolution：GG / OG / Kling / SI 等支持，具体档位见各模型文档；
//     GG 2.5/3.0 支持 1K/2K/4K，GG 3.1 额外支持 720P；
//     OG 支持 1K/2K/4K（OG 也可通过 ExtInfo 传自定义像素尺寸，两者不互斥）
//   - OutputImageCount：生成图片张数，仅 OG（1-8）和 Kling（1-9）支持；
//     其他模型（GG/Vidu/Qwen/Hunyuan 等）不支持此字段，传入无效；
//     由请求 n 字段映射而来，n > 上限时自动截断并记录日志。
type outputConfig struct {
	AspectRatio string `json:"AspectRatio,omitempty"`
	// Resolution 分辨率档位，GG 系列使用；取值：1K / 2K / 4K（GG 3.1 额外支持 720P）
	Resolution string `json:"Resolution,omitempty"`
	// OutputImageCount 生成图片张数，仅 OG（1-8）和 Kling（1-9）支持。
	// n=0 或 n=1 时不设此字段（使用 API 默认值 1 张）。
	OutputImageCount int `json:"OutputImageCount,omitempty"`
	// OutputFormat 指定输出图片文件格式，可选值：jpeg / png。
	// 若不指定则跟随模型默认值。通过请求的 output_format 字段传入。
	OutputFormat string `json:"OutputFormat,omitempty"`
}

// fileInfo 对应腾讯 VOD API 的 AigcImageTaskInputFileInfo
// Type 取值：File（VOD文件ID）| Url（外部可访问URL）| Base64（Base64字符串）
type fileInfo struct {
	Type   string `json:"Type,omitempty"`
	FileId string `json:"FileId,omitempty"`
	Url    string `json:"Url,omitempty"`
	Base64 string `json:"Base64,omitempty"`
}

type createTaskRequest struct {
	SubAppId       int64         `json:"SubAppId,omitempty"`
	ModelName      string        `json:"ModelName"`
	ModelVersion   string        `json:"ModelVersion"`
	Prompt         string        `json:"Prompt"`
	FileInfos      []fileInfo    `json:"FileInfos,omitempty"`
	OutputConfig   *outputConfig `json:"OutputConfig,omitempty"`
	Seed           int64         `json:"Seed,omitempty"`
	NegativePrompt string        `json:"NegativePrompt,omitempty"`
	// EnhancePrompt 自动优化提示词，取值 "Enabled" / "Disabled"。
	// 支持模型：GG 系列（通过 extra_fields.enhance_prompt 传入）；
	// 开启后腾讯 VOD 会自动优化传入的 Prompt 以提升生成质量。
	EnhancePrompt string `json:"EnhancePrompt,omitempty"`
	// SceneType 场景类型，取值示例：
	//   Kling scene 版本: "image_expand"（扩图）
	//   Hunyuan 3.0:      "3d_panorama"（全景图）
	// 通过 extra_fields.scene_type 传入
	SceneType string `json:"SceneType,omitempty"`
	ExtInfo   string `json:"ExtInfo,omitempty"`
}

type createTaskResponse struct {
	Response struct {
		TaskId    string `json:"TaskId"`
		RequestId string `json:"RequestId"`
		Error     *struct {
			Code    string `json:"Code"`
			Message string `json:"Message"`
		} `json:"Error,omitempty"`
	} `json:"Response"`
}

type describeTaskRequest struct {
	SubAppId int64  `json:"SubAppId,omitempty"`
	TaskId   string `json:"TaskId"`
}

// ── Video request structs ─────────────────────────────────────────────────────

// videoFileInfo corresponds to AigcVideoTaskInputFileInfo.
//
// Scope notes:
//   - Type / FileId / Url: [shared] present in both image and video tasks
//   - Category / Usage / ReferenceType / Text: [video-only]
//   - Base64: [image-only], not supported for video (omitted)
type videoFileInfo struct {
	Type          string `json:"Type,omitempty"`          // [shared] "Url" | "File"
	FileId        string `json:"FileId,omitempty"`        // [shared] VOD file ID (not exposed externally)
	Url           string `json:"Url,omitempty"`           // [shared] external URL
	Category      string `json:"Category,omitempty"`      // [video-only] "Image" | "Video"
	ReferenceType string `json:"ReferenceType,omitempty"` // [video-only] GV:"asset"/"style"; Kling+Video:"feature"/"base"; PixVerse:"subject"/"background"
	Usage         string `json:"Usage,omitempty"`         // [video-only] "FirstFrame" | "Reference" | "LastFrame"
	Text          string `json:"Text,omitempty"`          // [video-only] PixVerse multi-subject: name referenced in Prompt with @name
}

// videoOutputConfig corresponds to AigcVideoOutputConfig.
//
// Scope notes:
//   - Resolution / AspectRatio: [shared concept] but values/formats differ significantly from the image side
//   - Duration / AudioGeneration / EnhanceSwitch / FrameInterpolate: [video-only]
//   - OutputImageCount / OutputFormat: [image-only], not present here
type videoOutputConfig struct {
	Duration         float64 `json:"Duration,omitempty"`         // [video-only] seconds; per-model constraints apply
	Resolution       string  `json:"Resolution,omitempty"`       // [shared concept] differs from image side (see per-model table)
	AspectRatio      string  `json:"AspectRatio,omitempty"`      // [shared concept] Hailuo does not support it
	AudioGeneration  string  `json:"AudioGeneration,omitempty"`  // [video-only] "Enabled"/"Disabled"; GV/OS/Vidu/Kling only
	EnhanceSwitch    string  `json:"EnhanceSwitch,omitempty"`    // [video-only] "Enabled"/"Disabled"; video super-resolution
	FrameInterpolate string  `json:"FrameInterpolate,omitempty"` // [video-only] "Enabled"/"Disabled"; Vidu-exclusive
}

// createVideoTaskRequest corresponds to the CreateAigcVideoTask request body.
//
// Scope notes vs createTaskRequest (image):
//   - LastFrameUrl / InputRegion: [video-only]
//   - SceneType / ExtInfo: [shared] but semantics differ (see plan)
//   - All other top-level fields: [shared]
type createVideoTaskRequest struct {
	SubAppId       int64              `json:"SubAppId,omitempty"`
	ModelName      string             `json:"ModelName"`
	ModelVersion   string             `json:"ModelVersion"`
	Prompt         string             `json:"Prompt,omitempty"`
	NegativePrompt string             `json:"NegativePrompt,omitempty"`
	EnhancePrompt  string             `json:"EnhancePrompt,omitempty"`
	FileInfos      []videoFileInfo    `json:"FileInfos,omitempty"`
	LastFrameUrl   string             `json:"LastFrameUrl,omitempty"`  // [video-only]
	OutputConfig   *videoOutputConfig `json:"OutputConfig,omitempty"`
	Seed           int64              `json:"Seed,omitempty"`
	SceneType      string             `json:"SceneType,omitempty"`
	ExtInfo        string             `json:"ExtInfo,omitempty"`
	InputRegion    string             `json:"InputRegion,omitempty"`   // [video-only] "Mainland"/"Oversea"
}

// videoTaskMetadata holds video-specific fields decoded from TaskSubmitReq.Metadata.
//
// Scope notes:
//   - seed / negative_prompt / enhance_prompt / scene_type / ext_info: [shared] with image side
//   - All remaining fields: [video-only]
type videoTaskMetadata struct {
	// [shared with image side]
	Seed           int64  `json:"seed"`
	NegativePrompt string `json:"negative_prompt"`
	EnhancePrompt  string `json:"enhance_prompt"`
	SceneType      string `json:"scene_type"`
	ExtInfo        string `json:"ext_info"`

	// [video-only]
	AudioGeneration  string `json:"audio_generation"`   // "Enabled"/"Disabled"
	EnhanceSwitch    string `json:"enhance_switch"`     // "Enabled"/"Disabled"
	FrameInterpolate string `json:"frame_interpolate"`  // "Enabled"/"Disabled"; Vidu only
	LastFrame        string `json:"last_frame"`         // tail-frame image URL
	Resolution       string `json:"resolution"`         // fallback resolution (overridden by size)
	InputRegion      string `json:"input_region"`       // "Mainland"/"Oversea"
	ReferenceType    string `json:"reference_type"`
	InputUsage       string `json:"input_usage"`        // force "Reference" mode for single image
	VideoURL         string `json:"video_url"`          // video reference/edit input
	// PixVerse multi-subject: JSON array [{"url":"...","type":"subject","name":"小猫"}]
	PixverseSubjects string `json:"pixverse_subjects"`
}

// pixverseSubjectItem is one entry in the pixverse_subjects JSON array.
type pixverseSubjectItem struct {
	URL  string `json:"url"`
	Type string `json:"type"` // "subject" or "background"
	Name string `json:"name"`
}

// ── Video response struct ─────────────────────────────────────────────────────

// aigcVideoTask is the sub-struct inside DescribeTaskDetail for video tasks.
type aigcVideoTask struct {
	Status     string `json:"Status"`
	ErrCode    int    `json:"ErrCode"`
	ErrCodeExt string `json:"ErrCodeExt"`
	ErrMsg     string `json:"Message"`
	// Progress is a 0-100 integer from the Tencent API; use directly for accurate progress reporting.
	Progress int `json:"Progress"`
	Output   *struct {
		FileInfos []struct {
			FileUrl string `json:"FileUrl"`
			// UsageType discriminates video body ("") from auxiliary files:
			//   ""         → video file body (take this one)
			//   "scene_url"  → dynamic shot segment
			//   "point_url"  → NeRF point cloud
			//   "mesh_url"   → 3-D mesh
			//   "image_url"  → cover/thumbnail
			UsageType string `json:"UsageType"`
			// MetaData carries output properties returned by Tencent upon completion.
			// Duration is the actual video length in seconds with up to 8 decimal places of precision,
			// used for post-completion billing reconciliation (多退少补).
			MetaData *struct {
				Duration float64 `json:"Duration"`
			} `json:"MetaData,omitempty"`
		} `json:"FileInfos"`
	} `json:"Output,omitempty"`
}

type describeTaskResponse struct {
	Response struct {
		RequestId string `json:"RequestId"`
		Status    string `json:"Status"`
		Error     *struct {
			Code    string `json:"Code"`
			Message string `json:"Message"`
		} `json:"Error,omitempty"`
		AigcImageTask *struct {
			Status     string `json:"Status"`
			ErrCode    int    `json:"ErrCode"`
			ErrCodeExt string `json:"ErrCodeExt"`
			ErrMsg     string `json:"Message"`
			// Input 包含提交时的请求参数，原样回显在 DescribeTaskDetail 响应中。
			// 用于在轮询阶段读取实际发出的 OutputImageCount，以判断是否真正请求了多图，
			// 避免对不支持多图的模型（n 被忽略）误触发"少图退款"逻辑。
			Input *struct {
				OutputConfig *struct {
					// OutputImageCount 实际发往腾讯 API 的多图数量。
					// 0 或 1 表示未请求多图；> 1 表示主动请求了 n 张。
					OutputImageCount int `json:"OutputImageCount"`
				} `json:"OutputConfig,omitempty"`
			} `json:"Input,omitempty"`
			Output *struct {
				FileInfos []struct {
					FileUrl string `json:"FileUrl"`
				} `json:"FileInfos"`
			} `json:"Output,omitempty"`
		} `json:"AigcImageTask,omitempty"`
		// AigcVideoTask is non-nil for video tasks; AigcImageTask and AigcVideoTask are
		// mutually exclusive — DescribeTaskDetail only returns one of them.
		AigcVideoTask *aigcVideoTask `json:"AigcVideoTask,omitempty"`
	} `json:"Response"`
}

// metadata fields clients may include inside the "metadata" object
type taskMetadata struct {
	Seed           int64  `json:"seed"`
	NegativePrompt string `json:"negative_prompt"`
	// EnhancePrompt 自动优化提示词，取值 "Enabled" / "Disabled"。
	// 仅 GG 系列支持，通过 extra_fields.enhance_prompt 传入。
	EnhancePrompt string `json:"enhance_prompt"`
	// SceneType 场景类型，通过 extra_fields.scene_type 传入。
	// 支持模型：
	//   Kling scene：取值 "image_expand"（扩图模式），配合 KlingExpansion 参数使用
	//   Hunyuan 3.0：取值 "3d_panorama"（全景图模式）
	SceneType string `json:"scene_type"`
	// KlingExpansion Kling 扩图参数，仅当 SceneType="image_expand" 时有效。
	// 通过 extra_fields.expansion 传入，各方向比例取值 [0, 2]，新图面积不超过原图 3 倍。
	// 示例：{"up": 0.1, "down": 0.2, "left": 0.3, "right": 0.4}
	KlingExpansion *klingExpansionParams `json:"expansion,omitempty"`
}

// ── TaskAdaptor ───────────────────────────────────────────────────────────────

type TaskAdaptor struct {
	taskcommon.BaseBilling
	subAppId  int64
	secretId  string
	secretKey string
	baseURL   string
}

func (a *TaskAdaptor) Init(info *relaycommon.RelayInfo) {
	a.baseURL = info.ChannelBaseUrl
	parts := strings.Split(info.ApiKey, "|")
	if len(parts) == 3 {
		var n int64
		fmt.Sscan(strings.TrimSpace(parts[0]), &n)
		a.subAppId = n
		a.secretId = strings.TrimSpace(parts[1])
		a.secretKey = strings.TrimSpace(parts[2])
	}
}

func (a *TaskAdaptor) ValidateRequestAndSetAction(c *gin.Context, info *relaycommon.RelayInfo) *dto.TaskError {
	if isVideoModel(info.OriginModelName) {
		return relaycommon.ValidateBasicTaskRequest(c, info, constant.TaskActionGenerate)
	}
	return relaycommon.ValidateBasicTaskRequest(c, info, constant.TaskActionImageGenerate)
}

func (a *TaskAdaptor) BuildRequestURL(info *relaycommon.RelayInfo) (string, error) {
	return a.baseURL, nil
}

func (a *TaskAdaptor) BuildRequestHeader(c *gin.Context, req *http.Request, info *relaycommon.RelayInfo) error {
	bodyBytes, err := io.ReadAll(req.Body)
	if err != nil {
		return err
	}
	// Reset body so the HTTP stack can send it after we read it for signing.
	req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	// action must exactly match the value used in the TC3 signature; mismatch → AuthFailure.
	action := "CreateAigcImageTask"
	if isVideoModel(info.OriginModelName) {
		action = "CreateAigcVideoTask"
	}

	ts := time.Now().Unix()
	auth := buildAuthorization(a.secretId, a.secretKey, action, string(bodyBytes), ts)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Host", tencentVODHost)
	req.Header.Set("X-TC-Action", action)
	req.Header.Set("X-TC-Version", "2018-07-17")
	req.Header.Set("X-TC-Timestamp", fmt.Sprintf("%d", ts))
	req.Header.Set("Authorization", auth)
	return nil
}

func (a *TaskAdaptor) BuildRequestBody(c *gin.Context, info *relaycommon.RelayInfo) (io.Reader, error) {
	if a.subAppId == 0 {
		return nil, fmt.Errorf("TencentVOD: SubAppId 未配置，请将渠道 API Key 格式设置为 subAppId|secretId|secretKey")
	}
	if isVideoModel(info.OriginModelName) {
		return a.buildVideoRequestBody(c, info)
	}

	taskReq, err := relaycommon.GetTaskRequest(c)
	if err != nil {
		return nil, err
	}

	var meta taskMetadata
	_ = taskcommon.UnmarshalMetadata(taskReq.Metadata, &meta)

	// quality comes from top-level field; fall back to metadata["quality"] for compatibility
	quality := taskReq.Quality
	if quality == "" {
		quality, _ = taskReq.Metadata["quality"].(string)
	}
	tencentModelName := modelToTencentName(info.OriginModelName)

	// 各模型的 ModelVersion 来源不同，不能统一通过 quality 推断：
	// - OG: quality 映射到精度档位（image2_low/medium/high），由 qualityToModelVersion 处理
	// - GG/Kling/Vidu/Qwen/Hunyuan: 版本固定由原始模型名决定，quality 另独立映射到 Resolution
	var modelVersion string
	if tencentModelName == "GG" {
		modelVersion = getGGVersion(info.OriginModelName)
	} else if tencentModelName == "Kling" {
		modelVersion = getKlingVersion(info.OriginModelName)
	} else if tencentModelName == "Vidu" {
		modelVersion = getViduVersion(info.OriginModelName)
	} else if tencentModelName == "Qwen" {
		modelVersion = getQwenVersion(info.OriginModelName)
	} else if tencentModelName == "Hunyuan" {
		modelVersion = getHunyuanVersion(info.OriginModelName)
	} else {
		modelVersion = qualityToModelVersion(tencentModelName, quality)
	}

	body := createTaskRequest{
		SubAppId:       a.subAppId,
		ModelName:      tencentModelName,
		ModelVersion:   modelVersion,
		Prompt:         taskReq.Prompt,
		Seed:           meta.Seed,
		NegativePrompt: meta.NegativePrompt,
		EnhancePrompt:  meta.EnhancePrompt,
	}

	// ── 尺寸/分辨率处理（各模型路径不同）────────────────────────────────────
	// OG:      ExtInfo.AdditionalParameters.size 传自定义像素尺寸（[655,360–8,294,400] 像素）
	// GG:      OutputConfig.Resolution（大写 1K/2K/4K）+ AspectRatio（宽高比）
	// Kling:   OutputConfig.Resolution（小写 1k/2k/4k）+ AspectRatio
	//          scene 版本：跳过 OutputConfig，仅设 SceneType 和扩图 ExtInfo
	// Vidu:    OutputConfig.Resolution（1080p/2K/4K）+ AspectRatio
	// Qwen:    ExtInfo.AdditionalParameters.size 传自定义像素尺寸（[512×512–2048×2048] 像素）
	// Hunyuan: ExtInfo.AdditionalParameters.size 传自定义像素尺寸（各维度 [512,2048]，乘积≤1M）
	// 其他:    OutputConfig.AspectRatio（fallback）
	if tencentModelName == "OG" {
		if extInfo := buildOGExtInfo(taskReq.Size); extInfo != "" {
			body.ExtInfo = extInfo
		}
	} else if tencentModelName == "GG" {
		aspectRatio, arErr := getGGAspectRatio(modelVersion, taskReq.Size)
		if arErr != nil {
			common.SysLog(fmt.Sprintf("TencentVOD GG BuildRequestBody: aspect ratio fallback: %v", arErr))
		}
		resolution, resErr := getGGResolution(modelVersion, quality)
		if resErr != nil {
			common.SysLog(fmt.Sprintf("TencentVOD GG BuildRequestBody: resolution fallback: %v", resErr))
		}
		body.OutputConfig = &outputConfig{AspectRatio: aspectRatio, Resolution: resolution}
	} else if tencentModelName == "Kling" {
		aspectRatio, arErr := getKlingAspectRatio(modelVersion, taskReq.Size)
		if arErr != nil {
			common.SysLog(fmt.Sprintf("TencentVOD Kling BuildRequestBody: aspect ratio fallback: %v", arErr))
		}
		resolution, resErr := getKlingResolution(modelVersion, quality)
		if resErr != nil {
			common.SysLog(fmt.Sprintf("TencentVOD Kling BuildRequestBody: resolution fallback: %v", resErr))
		}
		// scene 版本：getKlingAspectRatio/getKlingResolution 均返回 ""，跳过 OutputConfig
		if aspectRatio != "" || resolution != "" {
			body.OutputConfig = &outputConfig{AspectRatio: aspectRatio, Resolution: resolution}
		}
		// SceneType：通过 extra_fields.scene_type 传入（extra_fields 自动合并进内部 Metadata）
		// 特殊说明：kling-image-scene 通常需设 scene_type="image_expand"；不自动注入，由调用方负责
		if meta.SceneType != "" {
			body.SceneType = meta.SceneType
		}
		// Kling 扩图参数：通过 extra_fields.expansion 传入
		if meta.KlingExpansion != nil {
			body.ExtInfo = buildKlingExpansionExtInfo(meta.KlingExpansion)
		}
	} else if tencentModelName == "Vidu" {
		aspectRatio, arErr := getViduAspectRatio(taskReq.Size)
		if arErr != nil {
			common.SysLog(fmt.Sprintf("TencentVOD Vidu BuildRequestBody: aspect ratio fallback: %v", arErr))
		}
		resolution, resErr := getViduResolution(quality)
		if resErr != nil {
			common.SysLog(fmt.Sprintf("TencentVOD Vidu BuildRequestBody: resolution fallback: %v", resErr))
		}
		body.OutputConfig = &outputConfig{AspectRatio: aspectRatio, Resolution: resolution}
	} else if tencentModelName == "Qwen" {
		// Qwen 不支持 OutputConfig.AspectRatio/Resolution，使用 ExtInfo 传自定义像素尺寸
		// 特殊说明：Qwen 0925 合法总像素范围 [512×512=261,632, 2048×2048=4,194,304]
		if extInfo := buildOGExtInfo(taskReq.Size); extInfo != "" {
			body.ExtInfo = extInfo
		}
	} else if tencentModelName == "Hunyuan" {
		// Hunyuan 不支持 OutputConfig.AspectRatio/Resolution，使用 ExtInfo 传自定义像素尺寸
		// 特殊说明：Hunyuan 3.0 宽高均在 [512, 2048] 像素范围内，宽高乘积 ≤ 1024×1024
		if extInfo := buildOGExtInfo(taskReq.Size); extInfo != "" {
			body.ExtInfo = extInfo
		}
		// SceneType：通过 extra_fields.scene_type 传入；Hunyuan 支持 "3d_panorama"（全景图）
		if meta.SceneType != "" {
			body.SceneType = meta.SceneType
		}
	} else {
		body.OutputConfig = &outputConfig{AspectRatio: sizeToAspectRatio(taskReq.Size)}
	}

	// OutputImageCount：将请求的 n 映射到 OutputConfig.OutputImageCount。
	// 仅 OG（上限 8）和 Kling（上限 9）支持；其他模型忽略 n 字段。
	// n <= 1 时不设此字段，使用 API 默认值（1 张）。
	if taskReq.N > 1 {
		maxCount := 0
		switch tencentModelName {
		case "OG":
			maxCount = 8
		case "Kling":
			maxCount = 9
		}
		if maxCount > 0 {
			count := taskReq.N
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

	// OutputFormat：将请求的 output_format 映射到 OutputConfig.OutputFormat。
	// 可选值：jpeg / png；不传时跟随模型默认值。
	if taskReq.OutputFormat != "" {
		if body.OutputConfig == nil {
			body.OutputConfig = &outputConfig{}
		}
		body.OutputConfig.OutputFormat = taskReq.OutputFormat
	}

	// 参考图：从 taskReq.Images 读取（单图 image 字段已在 ValidateBasicTaskRequest 里自动合并进 Images）
	// 各模型均有参考图数量上限，超出时截断并记录日志：
	//   GG 2.5: 3张；GG 3.0/3.1: 14张；Kling 2.1: 4张；Kling 3.0: 1张；
	//   Kling 3.0-Omni/O1: 10张；Kling scene: 1张；Vidu q2: 7张；
	//   Qwen 0925: 1张；Hunyuan 3.0: 3张
	refImages := taskReq.Images
	maxRef := getMaxRefImages(tencentModelName, modelVersion)
	if maxRef > 0 && len(refImages) > maxRef {
		common.SysLog(fmt.Sprintf(
			"TencentVOD %s %s: 参考图数量 %d 超出上限 %d，已截断至 %d 张",
			tencentModelName, modelVersion, len(refImages), maxRef, maxRef,
		))
		refImages = refImages[:maxRef]
	}
	for _, url := range refImages {
		if strings.TrimSpace(url) != "" {
			body.FileInfos = append(body.FileInfos, fileInfo{Type: "Url", Url: url})
		}
	}

	data, err := common.Marshal(body)
	if err != nil {
		return nil, err
	}

	// 打印发往腾讯 VOD 的请求体，便于排查参数映射问题（n/OutputImageCount/OutputFormat 等）
	common.SysLog(fmt.Sprintf("[TencentVOD] CreateAigcImageTask request body: %s", string(data)))

	return bytes.NewReader(data), nil
}

// buildVideoRequestBody constructs the CreateAigcVideoTask JSON request body.
func (a *TaskAdaptor) buildVideoRequestBody(c *gin.Context, info *relaycommon.RelayInfo) (io.Reader, error) {
	taskReq, err := relaycommon.GetTaskRequest(c)
	if err != nil {
		return nil, err
	}

	var meta videoTaskMetadata
	_ = taskcommon.UnmarshalMetadata(taskReq.Metadata, &meta)

	tencentModelName := getVideoModelName(info.OriginModelName)
	modelVersion := getVideoModelVersion(info.OriginModelName)

	body := createVideoTaskRequest{
		SubAppId:       a.subAppId,
		ModelName:      tencentModelName,
		ModelVersion:   modelVersion,
		Prompt:         taskReq.Prompt,
		Seed:           meta.Seed,
		NegativePrompt: meta.NegativePrompt,
		EnhancePrompt:  meta.EnhancePrompt,
	}

	// ── ExtInfo ────────────────────────────────────────────────────────────────
	if meta.ExtInfo != "" {
		body.ExtInfo = meta.ExtInfo
	}

	// ── SceneType (Kling / Vidu only; warn and drop for unsupported models) ───
	if meta.SceneType != "" {
		if sceneTypeSupported(tencentModelName) {
			body.SceneType = meta.SceneType
		} else {
			common.SysLog(fmt.Sprintf("[TencentVOD] video BuildRequestBody: model %s does not support SceneType=%s, ignoring",
				tencentModelName, meta.SceneType))
		}
	}

	// ── InputRegion ────────────────────────────────────────────────────────────
	if meta.InputRegion != "" {
		body.InputRegion = meta.InputRegion
	}

	// ── OutputConfig ───────────────────────────────────────────────────────────
	cfg := &videoOutputConfig{}
	hasConfig := false

	// Duration: prefer taskReq.Duration (int); fall back to taskReq.Seconds (string).
	requestedDur := float64(taskReq.Duration)
	if requestedDur == 0 && taskReq.Seconds != "" {
		if v, parseErr := strconv.Atoi(taskReq.Seconds); parseErr == nil {
			requestedDur = float64(v)
		}
	}
	dur := getVideoDuration(tencentModelName, requestedDur)
	if dur > 0 {
		cfg.Duration = dur
		hasConfig = true
	}
	// 记录预扣计费时长：表达式里用的有效时长（0→各模型默认值），用于任务完成后多退少补。
	info.BilledDurationSec = getBillingDuration(tencentModelName, dur, requestedDur)

	// Resolution: size > metadata.resolution
	sizeHint := taskReq.Size
	if sizeHint == "" {
		sizeHint = meta.Resolution
	}
	res := getVideoResolution(tencentModelName, sizeHint)
	if res != "" {
		cfg.Resolution = res
		hasConfig = true
	}

	// AspectRatio: only when there are no input images (i2v suppresses AspectRatio for Kling)
	// We resolve this after building FileInfos; store temporarily.
	willHaveFileInfos := len(taskReq.Images) > 0 || meta.VideoURL != "" || meta.PixverseSubjects != ""

	var aspectRatioForKling string
	if tencentModelName == "Kling" && willHaveFileInfos {
		// Kling i2v: aspect ratio is determined by the first frame image, do not send AspectRatio.
		aspectRatioForKling = ""
	} else {
		ar := getVideoAspectRatio(tencentModelName, modelVersion, taskReq.Size)
		if ar != "" {
			cfg.AspectRatio = ar
			hasConfig = true
		}
		aspectRatioForKling = ar
	}
	_ = aspectRatioForKling // used above; suppress unused warning

	// AudioGeneration
	if meta.AudioGeneration != "" && audioGenerationSupported(tencentModelName) {
		cfg.AudioGeneration = meta.AudioGeneration
		hasConfig = true
	}

	// EnhanceSwitch
	if meta.EnhanceSwitch != "" {
		cfg.EnhanceSwitch = meta.EnhanceSwitch
		hasConfig = true
	}

	// FrameInterpolate (Vidu-exclusive)
	if meta.FrameInterpolate != "" && frameInterpolateSupported(tencentModelName) {
		cfg.FrameInterpolate = meta.FrameInterpolate
		hasConfig = true
	}

	if hasConfig {
		body.OutputConfig = cfg
	}

	// ── FileInfos ──────────────────────────────────────────────────────────────
	// PixVerse multi-subject: metadata.pixverse_subjects takes priority over taskReq.Images.
	if meta.PixverseSubjects != "" && tencentModelName == "PixVerse" {
		var subjects []pixverseSubjectItem
		if jsonErr := common.Unmarshal([]byte(meta.PixverseSubjects), &subjects); jsonErr == nil {
			for _, s := range subjects {
				if s.URL != "" {
					body.FileInfos = append(body.FileInfos, videoFileInfo{
						Type:          "Url",
						Category:      "Image",
						Url:           s.URL,
						Usage:         "Reference",
						ReferenceType: s.Type,
						Text:          s.Name,
					})
				}
			}
		} else {
			common.SysLog(fmt.Sprintf("[TencentVOD] video BuildRequestBody: failed to parse pixverse_subjects: %v", jsonErr))
		}
	} else {
		images := taskReq.Images
		switch {
		case len(images) == 0:
			// t2v mode: no FileInfos

		case len(images) == 1 && meta.LastFrame != "":
			// First+last frame mode
			if strings.TrimSpace(images[0]) != "" {
				body.FileInfos = append(body.FileInfos, videoFileInfo{
					Type:     "Url",
					Category: "Image",
					Url:      images[0],
					Usage:    "FirstFrame",
				})
			}
			body.LastFrameUrl = meta.LastFrame

		case len(images) == 1 && strings.EqualFold(meta.InputUsage, "Reference"):
			// Reference image mode (single image, forced by caller)
			if strings.TrimSpace(images[0]) != "" {
				fi := videoFileInfo{
					Type:     "Url",
					Category: "Image",
					Url:      images[0],
					Usage:    "Reference",
				}
				if meta.ReferenceType != "" {
					fi.ReferenceType = meta.ReferenceType
				}
				body.FileInfos = append(body.FileInfos, fi)
			}

		case len(images) == 1:
			// Default i2v first-frame mode
			if strings.TrimSpace(images[0]) != "" {
				body.FileInfos = append(body.FileInfos, videoFileInfo{
					Type:     "Url",
					Category: "Image",
					Url:      images[0],
					Usage:    "FirstFrame",
				})
			}

		default:
			// Multiple images: first → FirstFrame, rest → Reference
			if strings.TrimSpace(images[0]) != "" {
				body.FileInfos = append(body.FileInfos, videoFileInfo{
					Type:     "Url",
					Category: "Image",
					Url:      images[0],
					Usage:    "FirstFrame",
				})
			}
			for _, img := range images[1:] {
				if strings.TrimSpace(img) != "" {
					fi := videoFileInfo{
						Type:     "Url",
						Category: "Image",
						Url:      img,
						Usage:    "Reference",
					}
					if meta.ReferenceType != "" {
						fi.ReferenceType = meta.ReferenceType
					}
					body.FileInfos = append(body.FileInfos, fi)
				}
			}
		}
	}

	// Append video reference/edit input (Kling / Vidu only)
	if meta.VideoURL != "" {
		body.FileInfos = append(body.FileInfos, videoFileInfo{
			Type:     "Url",
			Category: "Video",
			Url:      meta.VideoURL,
			Usage:    "FirstFrame",
		})
	}

	data, err := common.Marshal(body)
	if err != nil {
		return nil, err
	}

	common.SysLog(fmt.Sprintf("[TencentVOD] CreateAigcVideoTask request body: %s", string(data)))

	return bytes.NewReader(data), nil
}

func (a *TaskAdaptor) DoRequest(c *gin.Context, info *relaycommon.RelayInfo, requestBody io.Reader) (*http.Response, error) {
	return channel.DoTaskApiRequest(a, c, info, requestBody)
}

func (a *TaskAdaptor) DoResponse(c *gin.Context, resp *http.Response, info *relaycommon.RelayInfo) (taskID string, taskData []byte, taskErr *dto.TaskError) {
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", nil, service.TaskErrorWrapper(err, "read_response_failed", http.StatusInternalServerError)
	}
	defer resp.Body.Close()

	// 打印腾讯 VOD 原始响应，便于排查任务提交结果
	action := "CreateAigcImageTask"
	if isVideoModel(info.OriginModelName) {
		action = "CreateAigcVideoTask"
	}
	common.SysLog(fmt.Sprintf("[TencentVOD] %s response (status=%d): %s", action, resp.StatusCode, string(bodyBytes)))

	var createResp createTaskResponse
	if err := common.Unmarshal(bodyBytes, &createResp); err != nil {
		return "", nil, service.TaskErrorWrapper(
			errors.Wrapf(err, "body: %s", bodyBytes), "unmarshal_failed", http.StatusInternalServerError)
	}
	if createResp.Response.Error != nil {
		return "", nil, service.TaskErrorWrapper(
			fmt.Errorf("%s: %s", createResp.Response.Error.Code, createResp.Response.Error.Message),
			createResp.Response.Error.Code, http.StatusInternalServerError)
	}

	ov := dto.NewOpenAIVideo()
	ov.ID = info.PublicTaskID
	ov.TaskID = info.PublicTaskID
	ov.CreatedAt = time.Now().Unix()
	ov.Model = info.OriginModelName

	// v2 图片任务接口（RelayModeImageTaskSubmit）不在此处写响应，
	// 由控制器统一写 ImageTaskDto 格式，保持新旧接口响应结构一致。
	if c.GetInt("relay_mode") != relayconstant.RelayModeImageTaskSubmit {
		c.JSON(http.StatusOK, ov)
	}

	return createResp.Response.TaskId, bodyBytes, nil
}

// FetchTask calls DescribeTaskDetail to poll task status.
func (a *TaskAdaptor) FetchTask(baseUrl, key string, body map[string]any, proxy string) (*http.Response, error) {
	taskID, _ := body["task_id"].(string)

	// Re-parse credentials from the stored key
	parts := strings.Split(key, "|")
	var subAppId int64
	var secretId, secretKey string
	if len(parts) == 3 {
		fmt.Sscan(strings.TrimSpace(parts[0]), &subAppId)
		secretId = strings.TrimSpace(parts[1])
		secretKey = strings.TrimSpace(parts[2])
	}

	reqBody := describeTaskRequest{TaskId: taskID}
	if subAppId != 0 {
		reqBody.SubAppId = subAppId
	}
	payload, err := common.Marshal(reqBody)
	if err != nil {
		return nil, errors.Wrap(err, "marshal fetch payload failed")
	}

	ts := time.Now().Unix()
	auth := buildAuthorization(secretId, secretKey, "DescribeTaskDetail", string(payload), ts)

	// 打印查询请求体，便于确认 TaskId 是否正确
	common.SysLog(fmt.Sprintf("[TencentVOD] DescribeTaskDetail request body: %s", string(payload)))

	req, err := http.NewRequest(http.MethodPost, baseUrl, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Host", tencentVODHost)
	req.Header.Set("X-TC-Action", "DescribeTaskDetail")
	req.Header.Set("X-TC-Version", "2018-07-17")
	req.Header.Set("X-TC-Timestamp", fmt.Sprintf("%d", ts))
	req.Header.Set("Authorization", auth)

	client, err := service.GetHttpClientWithProxy(proxy)
	if err != nil {
		return nil, fmt.Errorf("get http client failed: %w", err)
	}
	return client.Do(req)
}

func (a *TaskAdaptor) ParseTaskResult(respBody []byte) (*relaycommon.TaskInfo, error) {
	// 打印腾讯 VOD 轮询原始响应，便于排查任务结果（图片数量、URL 等）
	common.SysLog(fmt.Sprintf("[TencentVOD] DescribeTaskDetail response: %s", string(respBody)))

	var descResp describeTaskResponse
	if err := common.Unmarshal(respBody, &descResp); err != nil {
		return nil, errors.Wrap(err, "unmarshal describe response failed")
	}

	result := &relaycommon.TaskInfo{}

	if descResp.Response.Error != nil {
		result.Code = 1
		result.Reason = fmt.Sprintf("%s: %s", descResp.Response.Error.Code, descResp.Response.Error.Message)
		result.Status = model.TaskStatusFailure
		result.Progress = "100%"
		return result, nil
	}

	// Dispatch to the video branch when AigcVideoTask is present in the response.
	if descResp.Response.AigcVideoTask != nil {
		return parseVideoTaskResult(descResp.Response.AigcVideoTask)
	}

	// ── Image task branch (existing logic) ───────────────────────────────────

	// Use AigcImageTask.Status when available; fall back to top-level Status
	taskStatus := descResp.Response.Status
	if descResp.Response.AigcImageTask != nil && descResp.Response.AigcImageTask.Status != "" {
		taskStatus = descResp.Response.AigcImageTask.Status
	}

	switch taskStatus {
	case "PROCESSING", "WAITING":
		result.Status = model.TaskStatusInProgress
		result.Progress = "50%"
	case "SUCCESS", "FINISH":
		// Tencent returns Status=FINISH for both success and failure; check ErrCode to distinguish.
		common.SysLog(fmt.Sprintf("TencentVOD ParseTaskResult: status=%s errCode=%d errCodeExt=%s",
			taskStatus,
			func() int {
				if descResp.Response.AigcImageTask != nil {
					return descResp.Response.AigcImageTask.ErrCode
				}
				return -1
			}(),
			func() string {
				if descResp.Response.AigcImageTask != nil {
					return descResp.Response.AigcImageTask.ErrCodeExt
				}
				return "nil"
			}()))
		if descResp.Response.AigcImageTask != nil &&
			(descResp.Response.AigcImageTask.ErrCode != 0 || descResp.Response.AigcImageTask.ErrCodeExt != "") {
			result.Status = model.TaskStatusFailure
			result.Progress = "100%"
			result.Reason = descResp.Response.AigcImageTask.ErrMsg
			if result.Reason == "" {
				result.Reason = descResp.Response.AigcImageTask.ErrCodeExt
			}
		} else {
			result.Status = model.TaskStatusSuccess
			result.Progress = "100%"
			if descResp.Response.AigcImageTask != nil &&
				descResp.Response.AigcImageTask.Output != nil &&
				len(descResp.Response.AigcImageTask.Output.FileInfos) > 0 {
				fileInfos := descResp.Response.AigcImageTask.Output.FileInfos
				result.Url = fileInfos[0].FileUrl

				// ActualImages 只在实际发出的请求中 OutputImageCount > 1 时才设置，
				// 用于"少图比例退款"逻辑（settleTaskBillingOnComplete 步骤 4）。
				// 判断依据是腾讯 DescribeTaskDetail 响应里回显的 Input.OutputConfig.OutputImageCount：
				//   > 1：代码主动请求了多图，计费表达式含 * n，实际数与请求数不符时才退款。
				//   0/1：未请求多图（模型不支持或 n=1），不设置 ActualImages，避免误触发退款。
				// 这样不需要写死任何模型名，新模型支持多图后只要代码里设了 OutputImageCount 即可自动生效。
				// 从腾讯 DescribeTaskDetail 回显的 Input.OutputConfig.OutputImageCount 中读取
				// 实际发往 API 的多图请求数，作为预扣和结算的共同基准（ExpectedImages）。
				// 这与计费表达式里的 * n 使用相同的值，保证预扣和结算完全对齐：
				//   预扣：expression 用 param("n") = 请求 n（= OutputImageCount，两者应一致）
				//   结算：quota * ActualImages / ExpectedImages（OutputImageCount）
				// 若 OutputImageCount <= 1（模型不支持多图或未请求多图），
				// 则不设置 ActualImages / ExpectedImages，结算逻辑不触发。
				requestedOutputImageCount := 0
				if descResp.Response.AigcImageTask.Input != nil &&
					descResp.Response.AigcImageTask.Input.OutputConfig != nil {
					requestedOutputImageCount = descResp.Response.AigcImageTask.Input.OutputConfig.OutputImageCount
				}
				if requestedOutputImageCount > 1 {
					result.ActualImages = len(fileInfos)
					result.ExpectedImages = requestedOutputImageCount
				}

				// 多图任务：收集所有 URL，task_polling.go 轮询路径会将其写入 task.Data（ImageResponse 格式），
				// TaskModel2ImageTaskDto 优先从 task.Data 读取，确保所有图片都返回给调用方。
				if len(fileInfos) > 1 {
					for _, fi := range fileInfos {
						if fi.FileUrl != "" {
							result.Urls = append(result.Urls, fi.FileUrl)
						}
					}
				}
			}
		}
	case "FAIL":
		result.Status = model.TaskStatusFailure
		result.Progress = "100%"
		if descResp.Response.AigcImageTask != nil {
			result.Reason = descResp.Response.AigcImageTask.ErrMsg
			if result.Reason == "" {
				result.Reason = descResp.Response.AigcImageTask.ErrCodeExt
			}
		}
	default:
		result.Status = model.TaskStatusQueued
		result.Progress = "10%"
	}
	return result, nil
}

// parseVideoTaskResult converts an aigcVideoTask poll response into a TaskInfo.
//
// Status mapping:
//
//	PROCESSING → InProgress  (use actual Progress integer)
//	FINISH + ErrCode==0 → Success
//	FINISH + ErrCode!=0 → Failure
//	others → Queued
//
// Video URL selection: take the first FileInfos entry where UsageType == "" and FileUrl != "".
// Other UsageType values (scene_url / point_url / mesh_url / image_url) are auxiliary outputs.
func parseVideoTaskResult(task *aigcVideoTask) (*relaycommon.TaskInfo, error) {
	result := &relaycommon.TaskInfo{}

	switch task.Status {
	case "PROCESSING", "WAITING":
		result.Status = model.TaskStatusInProgress
		// 轮询队列过滤条件：progress != "100%"（见 GetAllUnFinishSyncTasks）。
		// 腾讯在视频生成完毕、输出文件就绪之前会短暂返回 Status=PROCESSING + Progress=100，
		// 此时 Output.FileInfos 为空，视频尚未可用。若直接写 "100%" 会导致任务被踢出轮询队列，
		// 永远拿不到 FileInfos。因此限制最高 99%，只有真正进入 SUCCESS/FINISH/FAIL 才写 100%。
		p := task.Progress
		if p >= 100 {
			p = 99
		}
		result.Progress = fmt.Sprintf("%d%%", p)
	case "SUCCESS", "FINISH":
		if task.ErrCode != 0 || task.ErrCodeExt != "" {
			result.Status = model.TaskStatusFailure
			result.Progress = "100%"
			result.Reason = task.ErrMsg
			if result.Reason == "" {
				result.Reason = task.ErrCodeExt
			}
		} else {
			result.Status = model.TaskStatusSuccess
			result.Progress = "100%"
			if task.Output != nil {
				for _, fi := range task.Output.FileInfos {
					// UsageType "" → video body; skip auxiliary file types
					if fi.UsageType == "" && fi.FileUrl != "" {
						result.Url = fi.FileUrl
						break
					}
				}
			}
		}
	case "FAIL":
		result.Status = model.TaskStatusFailure
		result.Progress = "100%"
		result.Reason = task.ErrMsg
		if result.Reason == "" {
			result.Reason = task.ErrCodeExt
		}
	default:
		result.Status = model.TaskStatusQueued
		result.Progress = "10%"
	}
	return result, nil
}

// AdjustBillingOnComplete implements duration-based 多退少补 for video tasks.
//
// Tencent bills with 8-decimal-place precision on the actual output duration,
// so a 5s request that yields a 5.042s video should be charged for 5.042s, not 5s.
//
// Logic:
//  1. Skip if PerCallBilling or no BilledDurationSec stored at submit time.
//  2. Extract actual duration from AigcVideoTask.Output.FileInfos[0].MetaData.Duration.
//  3. final_quota = task.Quota × (actualDuration / billedDuration), rounded to int.
//
// Returns 0 to keep the pre-charged amount unchanged when reconciliation is not applicable.
func (a *TaskAdaptor) AdjustBillingOnComplete(task *model.Task, _ *relaycommon.TaskInfo) int {
	bc := task.PrivateData.BillingContext
	if bc == nil || bc.BilledDurationSec <= 0 {
		return 0
	}

	// Parse the raw task data (DescribeTaskDetail response stored at poll completion).
	var descResp describeTaskResponse
	if err := common.Unmarshal(task.Data, &descResp); err != nil {
		common.SysLog(fmt.Sprintf("[TencentVOD] AdjustBillingOnComplete: unmarshal task.Data failed for task %s: %v", task.TaskID, err))
		return 0
	}

	videoTask := descResp.Response.AigcVideoTask
	if videoTask == nil || videoTask.Output == nil {
		return 0
	}

	// Find the actual output duration from the first video-body FileInfo.
	var actualDuration float64
	for _, fi := range videoTask.Output.FileInfos {
		if fi.UsageType == "" && fi.MetaData != nil && fi.MetaData.Duration > 0 {
			actualDuration = fi.MetaData.Duration
			break
		}
	}
	if actualDuration <= 0 {
		common.SysLog(fmt.Sprintf("[TencentVOD] AdjustBillingOnComplete: no MetaData.Duration for task %s, keeping pre-charged quota", task.TaskID))
		return 0
	}

	billedDur := bc.BilledDurationSec
	finalQuota := int(float64(task.Quota) * actualDuration / billedDur)

	common.SysLog(fmt.Sprintf("[TencentVOD] AdjustBillingOnComplete: task=%s billedDur=%.3f actualDur=%.8f preQuota=%d finalQuota=%d",
		task.TaskID, billedDur, actualDuration, task.Quota, finalQuota))

	return finalQuota
}

func (a *TaskAdaptor) ConvertToOpenAIVideo(originTask *model.Task) ([]byte, error) {
	var descResp describeTaskResponse
	if err := common.Unmarshal(originTask.Data, &descResp); err != nil {
		return nil, errors.Wrap(err, "unmarshal task data failed")
	}

	ov := dto.NewOpenAIVideo()
	ov.ID = originTask.TaskID
	ov.Status = originTask.Status.ToVideoStatus()
	ov.SetProgressStr(originTask.Progress)
	ov.CreatedAt = originTask.CreatedAt
	ov.CompletedAt = originTask.UpdatedAt
	ov.Model = originTask.Properties.OriginModelName

	if descResp.Response.Error != nil {
		ov.Error = &dto.OpenAIVideoError{
			Message: descResp.Response.Error.Message,
			Code:    descResp.Response.Error.Code,
		}
	}

	// Video task branch: extract URL from AigcVideoTask output.
	if descResp.Response.AigcVideoTask != nil {
		task := descResp.Response.AigcVideoTask
		if task.Output != nil {
			for _, fi := range task.Output.FileInfos {
				// UsageType "" → video body; skip auxiliary outputs.
				if fi.UsageType == "" && fi.FileUrl != "" {
					ov.SetMetadata("url", fi.FileUrl)
					break
				}
			}
		}
		return common.Marshal(ov)
	}

	// Image task branch (existing logic).
	if descResp.Response.AigcImageTask != nil &&
		descResp.Response.AigcImageTask.Output != nil &&
		len(descResp.Response.AigcImageTask.Output.FileInfos) > 0 {
		ov.SetMetadata("url", descResp.Response.AigcImageTask.Output.FileInfos[0].FileUrl)
	}

	return common.Marshal(ov)
}

func (a *TaskAdaptor) GetModelList() []string {
	all := make([]string, 0, len(ModelList)+len(VideoModelList))
	all = append(all, ModelList...)
	all = append(all, VideoModelList...)
	return all
}
func (a *TaskAdaptor) GetChannelName() string { return ChannelName }
