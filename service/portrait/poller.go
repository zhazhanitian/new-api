package portrait

import (
	"strconv"
	"time"

	"github.com/QuantumNous/new-api/common"
	"github.com/QuantumNous/new-api/model"
	"github.com/QuantumNous/new-api/setting/portrait_setting"
)

// StartPortraitPoller 启动后台素材审核状态轮询协程。
// 每隔 portrait_setting.VolcPortraitPollIntervalSeconds 秒扫描一次所有非终态绑定，
// 按每条绑定自己的渠道取凭据刷新状态，再重算对应逻辑素材的汇总状态。
// 没有绑定的线上历史素材仍按火山凭据刷新。
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
	cache := newCredCache()
	pollLegacyAssets(cache)
	pollPendingBindings(cache)
}

func pollLegacyAssets(cache *credCache) {
	assets, err := model.GetLegacyPendingPortraitAssets()
	if err != nil {
		common.SysLog("portrait poller: 查询待审核历史素材失败: " + err.Error())
		return
	}
	if len(assets) == 0 {
		return
	}
	common.SysLog("portrait poller: 开始刷新 " + strconv.Itoa(len(assets)) + " 个待审核历史素材")
	for _, asset := range assets {
		if msg := refreshLegacyAsset(asset, cache); msg != "" {
			common.SysLog("portrait poller: 素材 " + asset.RemoteAssetId + " 刷新失败: " + msg)
		}
	}
}

func pollPendingBindings(cache *credCache) {
	bindings, err := model.GetPendingPortraitAssetBindings()
	if err != nil {
		common.SysLog("portrait poller: 查询待审核素材绑定失败: " + err.Error())
		return
	}
	if len(bindings) == 0 {
		return
	}
	common.SysLog("portrait poller: 开始刷新 " + strconv.Itoa(len(bindings)) + " 个待审核素材绑定")

	// 逻辑素材 -> 本轮是否有绑定被刷新，刷新后统一重算汇总状态。
	touched := map[string]int{}
	for _, b := range bindings {
		cred, credErr := cache.get(b.PortraitChannel)
		if credErr != nil {
			common.SysLog("portrait poller: 素材 " + b.AssetId + " 在渠道 " + b.PortraitChannel + " 凭据不可用: " + credErr.Error())
			continue
		}
		if refreshErr := refreshBinding(b, cred); refreshErr != nil {
			common.SysLog("portrait poller: 查询素材 " + b.AssetId + " 在渠道 " + b.PortraitChannel + " 失败: " + refreshErr.Error())
			continue
		}
		touched[b.AssetId] = b.UserId
	}

	for assetId, userId := range touched {
		asset, getErr := model.GetPortraitAssetByRemoteId(userId, assetId)
		if getErr != nil {
			common.SysLog("portrait poller: 读取素材 " + assetId + " 失败: " + getErr.Error())
			continue
		}
		if recomputeErr := RecomputeAssetStatus(asset); recomputeErr != nil {
			common.SysLog("portrait poller: 保存素材 " + assetId + " 失败: " + recomputeErr.Error())
		}
	}
}
