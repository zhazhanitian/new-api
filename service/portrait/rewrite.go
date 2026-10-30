package portrait

import (
	"fmt"
	"strings"

	"github.com/QuantumNous/new-api/model"
	"github.com/QuantumNous/new-api/setting/portrait_setting"
)

// AssetURLPrefix 是生视频请求里引用人像素材的 URL 前缀。
const AssetURLPrefix = "asset://"

var (
	getAssetRecordFn   = model.GetPortraitAssetByRemoteId
	getAssetBindingsFn = model.GetPortraitAssetBindings
)

// ParseAssetURL 取出 asset:// 后面的逻辑素材 ID。不是素材引用时 ok 为 false。
func ParseAssetURL(url string) (assetId string, ok bool) {
	url = strings.TrimSpace(url)
	if !strings.HasPrefix(url, AssetURLPrefix) {
		return "", false
	}
	id := strings.TrimSpace(strings.TrimPrefix(url, AssetURLPrefix))
	if id == "" {
		return "", false
	}
	return id, true
}

// ResolveUpstreamAssetID 把对外的逻辑素材 ID 换成指定渠道账号里的上游素材 ID。
// 只检查"当前选中渠道"的绑定：素材必须属于该用户，且该渠道的绑定为 Active。
// 没有任何绑定的线上历史素材视为仅火山：volc 渠道原样放行（要求本地状态为 Active），其他渠道报未就绪。
func ResolveUpstreamAssetID(userId int, alias, assetId string) (string, error) {
	alias = strings.ToLower(strings.TrimSpace(alias))
	if alias == "" {
		alias = portrait_setting.DefaultAlias
	}

	asset, err := getAssetRecordFn(userId, assetId)
	if err != nil || asset == nil {
		return "", fmt.Errorf("素材 %s 不存在或不属于当前用户", assetId)
	}
	bindings, err := getAssetBindingsFn(assetId)
	if err != nil {
		return "", fmt.Errorf("查询素材 %s 的渠道绑定失败: %w", assetId, err)
	}

	if len(bindings) == 0 {
		if alias != portrait_setting.DefaultAlias {
			return "", fmt.Errorf("素材 %s 尚未在渠道 %s 就绪", assetId, alias)
		}
		if asset.Status != model.PortraitStatusActive {
			return "", fmt.Errorf("素材 %s 状态为 %s，需为 Active 后才能使用", assetId, asset.Status)
		}
		return assetId, nil
	}

	for _, b := range bindings {
		if b.PortraitChannel != alias {
			continue
		}
		if b.Status != model.PortraitStatusActive {
			msg := fmt.Sprintf("素材 %s 在渠道 %s 的状态为 %s，需为 Active 后才能使用", assetId, alias, b.Status)
			if b.LastError != "" {
				msg += "：" + b.LastError
			}
			return "", fmt.Errorf("%s", msg)
		}
		return b.UpstreamAssetId, nil
	}
	return "", fmt.Errorf("素材 %s 尚未在渠道 %s 就绪", assetId, alias)
}

// RewriteAssetURL 把 asset://逻辑ID 改写成指定渠道的 asset://上游ID。
// 不是 asset:// 的 URL 原样返回。
func RewriteAssetURL(userId int, alias, url string) (string, error) {
	assetId, ok := ParseAssetURL(url)
	if !ok {
		return url, nil
	}
	upstreamId, err := ResolveUpstreamAssetID(userId, alias, assetId)
	if err != nil {
		return "", err
	}
	return AssetURLPrefix + upstreamId, nil
}
