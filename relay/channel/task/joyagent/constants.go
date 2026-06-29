package joyagent

const ChannelName = "JoyAgent"

// ModelList 包含 HappyHorse 系列模型。
// 当前 JoyAgent 平台已上线：1.0-t2v / 1.0-i2v / 1.0-r2v / 1.0-video-edit
// 1.1 版本在阿里 DashScope 侧已发布，待 JoyAgent 平台同步上线后即可直接使用。
var ModelList = []string{
	"happyhorse-1.0-t2v",
	"happyhorse-1.1-t2v",
	"happyhorse-1.0-i2v",
	"happyhorse-1.1-i2v",
	"happyhorse-1.0-r2v",
	"happyhorse-1.1-r2v",
	"happyhorse-1.0-video-edit",
}

// queryToolName 结果查询工具名（已由 JoyAgent 接口文档确认）
const queryToolName = "tasks-query"

// resolutionMap 将 TaskSubmitReq.Size 映射为 JoyAgent parameters.resolution
var resolutionMap = map[string]string{
	"480p":      "480P",
	"720p":      "720P",
	"1080p":     "1080P",
	"1280x720":  "720P",
	"1920x1080": "1080P",
}

func sizeToResolution(size string) string {
	if v, ok := resolutionMap[size]; ok {
		return v
	}
	return ""
}
