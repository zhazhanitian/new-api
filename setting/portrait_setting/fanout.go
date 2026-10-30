package portrait_setting

import (
	"os"
	"strings"
	"sync"

	"github.com/QuantumNous/new-api/constant"
)

// FanoutEnvKey 指定参与多写的人像库账号，逗号分隔，例如 "volc,zj"。
// 不配置时只有 volc，行为与单账号一致。volc 始终存在且排第一，作为主账号：
// 对外的 group_id / asset_id 取主账号的上游 ID。
const FanoutEnvKey = "PORTRAIT_FANOUT_CHANNELS"

var (
	fanoutMu      sync.RWMutex
	fanoutAliases = parseFanoutAliases(os.Getenv(FanoutEnvKey))
)

// parseFanoutAliases 解析参与渠道列表：转小写、去重、丢弃非法标识，volc 固定排第一。
// 环境变量里的下划线和连字符写法归一成同一个标识（byte_plus 与 byte-plus 视为同一套）。
func parseFanoutAliases(raw string) []string {
	out := []string{DefaultAlias}
	seen := map[string]bool{DefaultAlias: true}
	for _, item := range strings.Split(raw, ",") {
		alias := strings.ToLower(strings.TrimSpace(item))
		if alias == "" || seen[alias] || !aliasPattern.MatchString(alias) {
			continue
		}
		seen[alias] = true
		out = append(out, alias)
	}
	return out
}

// FanoutAliases 返回参与多写的人像库账号标识，第一个永远是 volc。
func FanoutAliases() []string {
	fanoutMu.RLock()
	defer fanoutMu.RUnlock()
	out := make([]string, len(fanoutAliases))
	copy(out, fanoutAliases)
	return out
}

// setFanoutAliasesForTest 仅供测试覆盖参与渠道。
func setFanoutAliasesForTest(raw string) {
	fanoutMu.Lock()
	fanoutAliases = parseFanoutAliases(raw)
	fanoutMu.Unlock()
}

// AliasForChannelType 把视频渠道类型映射到人像库账号标识。
// 这是渠道类型到人像库账号的唯一对应表：不放进渠道配置，也不解析模型名。
// 未登记的渠道类型按火山处理。
func AliasForChannelType(channelType int) string {
	switch channelType {
	case constant.ChannelTypeZhongJiVideo:
		return "zj"
	default:
		return DefaultAlias
	}
}
