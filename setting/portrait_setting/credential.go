package portrait_setting

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"sync"
)

// DefaultAlias 是火山官方人像库。空标识与历史数据都归到这里。
const DefaultAlias = "volc"

// Credential 是一次人像库调用使用的账号凭据。
// volc 始终读取 VOLC_PORTRAIT_*（含系统设置的运行时覆盖）；
// 其他标识在进程启动时从 PORTRAIT_<ALIAS>_* 读入，修改后需要重启。
type Credential struct {
	Alias   string
	AK      string
	SK      string
	Project string
	Region  string
}

var aliasPattern = regexp.MustCompile(`^[a-z0-9][a-z0-9_-]{0,31}$`)

var (
	registryMu sync.RWMutex
	registry   = map[string]Credential{}
)

func init() {
	reloadRegistry(os.Environ())
}

func reloadRegistry(environ []string) {
	registryMu.Lock()
	registry = loadRegistry(environ)
	registryMu.Unlock()
}

// loadRegistry 扫描 PORTRAIT_<ALIAS>_ACCESS_KEY_ID。
// 环境变量名里的下划线同时登记成连字符键，使 byte_plus 与 byte-plus 命中同一套凭据。
// volc 不进表，避免盖住可热更新的 VOLC_PORTRAIT_*。
func loadRegistry(environ []string) map[string]Credential {
	env := make(map[string]string, len(environ))
	for _, item := range environ {
		key, value, ok := strings.Cut(item, "=")
		if !ok {
			continue
		}
		env[key] = value
	}

	out := make(map[string]Credential)
	for key, accessKey := range env {
		middle, ok := portraitEnvMiddle(key)
		if !ok || strings.EqualFold(middle, "VOLC") {
			continue
		}
		region := strings.TrimSpace(env["PORTRAIT_"+middle+"_REGION"])
		if region == "" {
			region = "cn-beijing"
		}
		cred := Credential{
			AK:      strings.TrimSpace(accessKey),
			SK:      strings.TrimSpace(env["PORTRAIT_"+middle+"_SECRET_ACCESS_KEY"]),
			Project: strings.TrimSpace(env["PORTRAIT_"+middle+"_PROJECT_NAME"]),
			Region:  region,
		}
		alias := strings.ToLower(middle)
		cred.Alias = alias
		out[alias] = cred
		if hyphen := strings.ReplaceAll(alias, "_", "-"); hyphen != alias {
			hyphenCred := cred
			hyphenCred.Alias = hyphen
			out[hyphen] = hyphenCred
		}
	}
	return out
}

func portraitEnvMiddle(key string) (string, bool) {
	const prefix = "PORTRAIT_"
	const suffix = "_ACCESS_KEY_ID"
	if !strings.HasPrefix(key, prefix) || !strings.HasSuffix(key, suffix) {
		return "", false
	}
	middle := strings.TrimSuffix(strings.TrimPrefix(key, prefix), suffix)
	if middle == "" {
		return "", false
	}
	return middle, true
}

// Resolve 按渠道标识取凭据。空串和 volc 使用火山官方配置。
// 未知、非法或缺少 AK/SK 的标识返回错误，不回落到火山。
func Resolve(alias string) (Credential, error) {
	alias = strings.ToLower(strings.TrimSpace(alias))
	if alias == "" || alias == DefaultAlias {
		return volcCredential()
	}
	if !aliasPattern.MatchString(alias) {
		return Credential{}, fmt.Errorf("人像库渠道标识不合法")
	}

	registryMu.RLock()
	cred, ok := registry[alias]
	registryMu.RUnlock()
	if !ok || cred.AK == "" || cred.SK == "" {
		return Credential{}, fmt.Errorf("人像库渠道标识不存在或未配置")
	}
	cred.Alias = alias
	if cred.Region == "" {
		cred.Region = "cn-beijing"
	}
	return cred, nil
}

func volcCredential() (Credential, error) {
	if VolcPortraitAccessKeyId == "" || VolcPortraitSecretAccessKey == "" {
		return Credential{}, fmt.Errorf("火山引擎凭据未配置")
	}
	region := VolcPortraitRegion
	if region == "" {
		region = "cn-beijing"
	}
	return Credential{
		Alias:   DefaultAlias,
		AK:      VolcPortraitAccessKeyId,
		SK:      VolcPortraitSecretAccessKey,
		Project: VolcPortraitProjectName,
		Region:  region,
	}, nil
}
