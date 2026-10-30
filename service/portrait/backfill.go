package portrait

import (
	"github.com/QuantumNous/new-api/common"
	"github.com/QuantumNous/new-api/model"
	"github.com/QuantumNous/new-api/setting/portrait_setting"
)

// BackfillResult 是一次补建的统计。
type BackfillResult struct {
	Created int      `json:"created"` // 新补建的渠道素材数
	Skipped int      `json:"skipped"` // 跳过的素材数（审核失败或素材组缺失）
	Failed  int      `json:"failed"`  // 补建调用失败的素材数
	Errors  []string `json:"errors,omitempty"`
}

const backfillPageSize = 100
const backfillMaxErrors = 20

// BackfillMissingChannels 把缺少渠道绑定的素材补建到所有参与渠道。
// 主要用于线上历史素材（只在火山账号里，没有任何绑定）。幂等：已有绑定的渠道跳过，
// 可以重复执行，用于补上开启多写前后新建的素材。
func BackfillMissingChannels() (*BackfillResult, error) {
	result := &BackfillResult{}
	for _, alias := range fanoutAliasesFn() {
		if alias == portrait_setting.DefaultAlias {
			continue
		}
		cred, err := resolveCredFn(alias)
		if err != nil {
			return result, channelErr(alias, "凭据不可用", err)
		}
		backfillChannel(alias, cred, result)
	}
	return result, nil
}

func backfillChannel(alias string, cred portrait_setting.Credential, result *BackfillResult) {
	afterId := 0
	for {
		assets, err := model.GetPortraitAssetsMissingChannel(alias, afterId, backfillPageSize)
		if err != nil {
			result.addError("查询待补建素材失败: " + err.Error())
			return
		}
		if len(assets) == 0 {
			return
		}
		for _, asset := range assets {
			afterId = asset.Id
			backfillAsset(alias, cred, asset, result)
		}
	}
}

func (r *BackfillResult) addError(msg string) {
	r.Failed++
	if len(r.Errors) < backfillMaxErrors {
		r.Errors = append(r.Errors, msg)
	}
	common.SysLog("portrait backfill: " + msg)
}

func backfillAsset(alias string, cred portrait_setting.Credential, asset *model.PortraitAsset, result *BackfillResult) {
	// 审核失败的素材不再补建。
	if asset.Status == model.PortraitStatusFailed {
		result.Skipped++
		return
	}

	// 先保证 volc 绑定存在，历史素材的 volc 状态取自现有记录。
	if err := ensureVolcAssetBinding(asset); err != nil {
		result.addError("素材 " + asset.RemoteAssetId + " 写入火山绑定失败: " + err.Error())
		return
	}

	group, err := model.GetPortraitGroupByRemoteId(asset.UserId, asset.RemoteGroupId)
	if err != nil || group == nil {
		result.Skipped++
		common.SysLog("portrait backfill: 素材 " + asset.RemoteAssetId + " 的素材组 " + asset.RemoteGroupId + " 不存在，已跳过")
		return
	}
	refs, err := EnsureGroupBindings(group)
	if err != nil {
		result.addError("素材 " + asset.RemoteAssetId + " 补建素材组失败: " + err.Error())
		return
	}
	var groupId string
	for _, ref := range refs {
		if ref.Alias == alias {
			groupId = ref.UpstreamId
			break
		}
	}
	if groupId == "" {
		result.Skipped++
		return
	}

	created, err := createAssetFn(cred, groupId, asset.SourceUrl, asset.AssetType, asset.Name)
	if err != nil {
		result.addError("素材 " + asset.RemoteAssetId + " 在渠道 " + alias + " 补建失败: " + err.Error())
		return
	}
	binding := &model.PortraitAssetBinding{
		UserId:          asset.UserId,
		AssetId:         asset.RemoteAssetId,
		PortraitChannel: alias,
		UpstreamAssetId: created.AssetId,
		Status:          model.PortraitStatusSubmitted,
	}
	if err = model.CreatePortraitAssetBinding(binding); err != nil {
		result.addError("素材 " + asset.RemoteAssetId + " 保存渠道绑定失败: " + err.Error())
		return
	}
	if err = RecomputeAssetStatus(asset); err != nil {
		common.SysLog("portrait backfill: 重算素材 " + asset.RemoteAssetId + " 状态失败: " + err.Error())
	}
	result.Created++
}

// ensureVolcAssetBinding 给没有火山绑定的素材补一条，上游 ID 就是逻辑 ID。
func ensureVolcAssetBinding(asset *model.PortraitAsset) error {
	bindings, err := model.GetPortraitAssetBindings(asset.RemoteAssetId)
	if err != nil {
		return err
	}
	for _, b := range bindings {
		if b.PortraitChannel == portrait_setting.DefaultAlias {
			return nil
		}
	}
	return model.CreatePortraitAssetBinding(&model.PortraitAssetBinding{
		UserId:          asset.UserId,
		AssetId:         asset.RemoteAssetId,
		PortraitChannel: portrait_setting.DefaultAlias,
		UpstreamAssetId: asset.RemoteAssetId,
		Status:          asset.Status,
		ResolvedUrl:     asset.ResolvedUrl,
	})
}
