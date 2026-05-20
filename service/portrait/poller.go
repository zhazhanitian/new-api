package portrait

import (
	"time"

	"github.com/QuantumNous/new-api/common"
	"github.com/QuantumNous/new-api/model"
	"github.com/QuantumNous/new-api/setting/portrait_setting"
)

// StartPortraitPoller 启动后台素材审核状态轮询协程。
// 每隔 portrait_setting.VolcPortraitPollIntervalSeconds 秒扫描一次所有非终态素材，
// 调用火山引擎 GetAsset 刷新状态并写回数据库。
func StartPortraitPoller() {
	go func() {
		for {
			interval := portrait_setting.VolcPortraitPollIntervalSeconds
			if interval <= 0 {
				interval = 300
			}
			time.Sleep(time.Duration(interval) * time.Second)
			pollPendingAssets()
		}
	}()
	common.SysLog("portrait poller started")
}

func pollPendingAssets() {
	assets, err := model.GetPendingPortraitAssets()
	if err != nil {
		common.SysLog("portrait poller: 查询待审核素材失败: " + err.Error())
		return
	}
	if len(assets) == 0 {
		return
	}
	common.SysLog("portrait poller: 开始刷新 " + itoa(len(assets)) + " 个待审核素材")

	for _, asset := range assets {
		result, err := GetAsset(asset.RemoteAssetId)
		if err != nil {
			common.SysLog("portrait poller: 查询素材 " + asset.RemoteAssetId + " 失败: " + err.Error())
			continue
		}
		changed := false
		if result.Status != "" && result.Status != asset.Status {
			asset.Status = result.Status
			changed = true
		}
		if result.ResolvedUrl != "" && result.ResolvedUrl != asset.ResolvedUrl {
			asset.ResolvedUrl = result.ResolvedUrl
			changed = true
		}
		if changed {
			if saveErr := model.SavePortraitAsset(asset); saveErr != nil {
				common.SysLog("portrait poller: 保存素材 " + asset.RemoteAssetId + " 失败: " + saveErr.Error())
			}
		}
	}
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	buf := make([]byte, 0, 10)
	for n > 0 {
		buf = append([]byte{byte('0' + n%10)}, buf...)
		n /= 10
	}
	return string(buf)
}
