package portrait_setting

import "os"

var (
	VolcPortraitAccessKeyId     = os.Getenv("VOLC_PORTRAIT_ACCESS_KEY_ID")
	VolcPortraitSecretAccessKey = os.Getenv("VOLC_PORTRAIT_SECRET_ACCESS_KEY")
	VolcPortraitProjectName     = os.Getenv("VOLC_PORTRAIT_PROJECT_NAME")
	VolcPortraitRegion          = getEnvOrDefault("VOLC_PORTRAIT_REGION", "cn-beijing")
	// 后台轮询非终态素材的间隔（秒），默认 300 秒（5 分钟）
	VolcPortraitPollIntervalSeconds = 300
)

func getEnvOrDefault(key, defaultVal string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultVal
}
