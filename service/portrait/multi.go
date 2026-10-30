package portrait

import (
	"fmt"
	"strings"

	"github.com/QuantumNous/new-api/model"
	"github.com/QuantumNous/new-api/setting/portrait_setting"
)

// 上游调用做成变量，便于测试替换。
var (
	createGroupFn   = CreateAssetGroup
	createAssetFn   = CreateAsset
	getAssetFn      = GetAsset
	updateGroupFn   = UpdateAssetGroup
	updateAssetFn   = UpdateAsset
	resolveCredFn   = portrait_setting.Resolve
	fanoutAliasesFn = portrait_setting.FanoutAliases
)

// channelRef 是逻辑对象在某个人像库账号里的上游 ID。
type channelRef struct {
	Alias      string
	UpstreamId string
}

// ChannelStatus 是素材在单个渠道上的状态，用于响应里的 channels 字段。
type ChannelStatus struct {
	PortraitChannel string `json:"portrait_channel"`
	Status          string `json:"status"`
	LastError       string `json:"last_error,omitempty"`
}

func channelErr(alias, action string, err error) error {
	return fmt.Errorf("人像库渠道 %s %s: %w", alias, action, err)
}

// resolveCreds 先把所有渠道的凭据取齐，任一不可用就直接失败，避免只建了一半。
func resolveCreds(aliases []string) ([]portrait_setting.Credential, error) {
	creds := make([]portrait_setting.Credential, 0, len(aliases))
	for _, alias := range aliases {
		cred, err := resolveCredFn(alias)
		if err != nil {
			return nil, channelErr(alias, "凭据不可用", err)
		}
		creds = append(creds, cred)
	}
	return creds, nil
}

// CreateGroupAllChannels 在所有参与渠道各建一个素材组。
// 对外的 group_id 取主账号（volc）的上游 ID。任一渠道失败即整体失败，已建成的不清理。
func CreateGroupAllChannels(userId int, name, description, groupType string) (*model.PortraitGroup, error) {
	creds, err := resolveCreds(fanoutAliasesFn())
	if err != nil {
		return nil, err
	}
	ids := make([]string, len(creds))
	for i, cred := range creds {
		result, createErr := createGroupFn(cred, name, description, groupType)
		if createErr != nil {
			return nil, channelErr(cred.Alias, "创建素材组失败", createErr)
		}
		ids[i] = result.Id
	}

	group := &model.PortraitGroup{
		UserId:        userId,
		RemoteGroupId: ids[0],
		Name:          name,
	}
	bindings := make([]*model.PortraitGroupBinding, len(creds))
	for i, cred := range creds {
		bindings[i] = &model.PortraitGroupBinding{
			UserId:          userId,
			GroupId:         ids[0],
			PortraitChannel: cred.Alias,
			UpstreamGroupId: ids[i],
		}
	}
	if err = model.CreatePortraitGroupWithBindings(group, bindings); err != nil {
		return nil, fmt.Errorf("保存素材组记录失败: %w", err)
	}
	return group, nil
}

// groupRefs 返回素材组在各渠道的上游 ID。线上历史素材组没有绑定，按"逻辑 ID 就是火山 ID"补上 volc。
func groupRefs(group *model.PortraitGroup) ([]channelRef, error) {
	bindings, err := model.GetPortraitGroupBindings(group.RemoteGroupId)
	if err != nil {
		return nil, err
	}
	refs := make([]channelRef, 0, len(bindings)+1)
	hasVolc := false
	for _, b := range bindings {
		if b.PortraitChannel == portrait_setting.DefaultAlias {
			hasVolc = true
		}
		refs = append(refs, channelRef{Alias: b.PortraitChannel, UpstreamId: b.UpstreamGroupId})
	}
	if !hasVolc {
		refs = append([]channelRef{{Alias: portrait_setting.DefaultAlias, UpstreamId: group.RemoteGroupId}}, refs...)
	}
	return refs, nil
}

// EnsureGroupBindings 保证素材组在每个参与渠道都有绑定，缺的当场补建（懒补建）。
// 线上历史素材组只在火山账号里：volc 补一条绑定，其他渠道调接口新建素材组。
func EnsureGroupBindings(group *model.PortraitGroup) ([]channelRef, error) {
	existing, err := model.GetPortraitGroupBindings(group.RemoteGroupId)
	if err != nil {
		return nil, err
	}
	byAlias := make(map[string]*model.PortraitGroupBinding, len(existing))
	for _, b := range existing {
		byAlias[b.PortraitChannel] = b
	}

	aliases := fanoutAliasesFn()
	refs := make([]channelRef, 0, len(aliases))
	for _, alias := range aliases {
		if b, ok := byAlias[alias]; ok {
			refs = append(refs, channelRef{Alias: alias, UpstreamId: b.UpstreamGroupId})
			continue
		}
		upstreamId := group.RemoteGroupId
		if alias != portrait_setting.DefaultAlias {
			cred, credErr := resolveCredFn(alias)
			if credErr != nil {
				return nil, channelErr(alias, "凭据不可用", credErr)
			}
			result, createErr := createGroupFn(cred, group.Name, "", "AIGC")
			if createErr != nil {
				return nil, channelErr(alias, "补建素材组失败", createErr)
			}
			upstreamId = result.Id
		}
		binding := &model.PortraitGroupBinding{
			UserId:          group.UserId,
			GroupId:         group.RemoteGroupId,
			PortraitChannel: alias,
			UpstreamGroupId: upstreamId,
		}
		if err = model.CreatePortraitGroupBinding(binding); err != nil {
			return nil, fmt.Errorf("保存素材组绑定失败: %w", err)
		}
		refs = append(refs, channelRef{Alias: alias, UpstreamId: upstreamId})
	}
	return refs, nil
}

// CreateAssetAllChannels 在所有参与渠道各注册一份素材。
// 对外的 asset_id 取主账号（volc）的上游 ID，各渠道的上游 ID 和状态记在绑定里。
func CreateAssetAllChannels(userId int, groupId, assetURL, assetType, name string) (*model.PortraitAsset, error) {
	group, err := model.GetPortraitGroupByRemoteId(userId, groupId)
	if err != nil || group == nil {
		return nil, fmt.Errorf("素材组不存在")
	}
	refs, err := EnsureGroupBindings(group)
	if err != nil {
		return nil, err
	}
	aliases := make([]string, len(refs))
	for i, ref := range refs {
		aliases[i] = ref.Alias
	}
	creds, err := resolveCreds(aliases)
	if err != nil {
		return nil, err
	}

	assetIds := make([]string, len(refs))
	for i, ref := range refs {
		result, createErr := createAssetFn(creds[i], ref.UpstreamId, assetURL, assetType, name)
		if createErr != nil {
			return nil, channelErr(ref.Alias, "创建素材失败", createErr)
		}
		assetIds[i] = result.AssetId
	}

	asset := &model.PortraitAsset{
		UserId:        userId,
		RemoteGroupId: group.RemoteGroupId,
		RemoteAssetId: assetIds[0],
		Name:          name,
		AssetType:     assetType,
		SourceUrl:     assetURL,
		Status:        model.PortraitStatusSubmitted,
	}
	bindings := make([]*model.PortraitAssetBinding, len(refs))
	for i, ref := range refs {
		bindings[i] = &model.PortraitAssetBinding{
			UserId:          userId,
			AssetId:         assetIds[0],
			PortraitChannel: ref.Alias,
			UpstreamAssetId: assetIds[i],
			Status:          model.PortraitStatusSubmitted,
		}
	}
	if err = model.CreatePortraitAssetWithBindings(asset, bindings); err != nil {
		return nil, fmt.Errorf("保存素材记录失败: %w", err)
	}
	return asset, nil
}

// UpdateGroupAllChannels 对素材组在所有渠道里的上游素材组各更新一次。
// 部分失败时返回的错误会列出失败的渠道，本地记录不更新。
func UpdateGroupAllChannels(group *model.PortraitGroup, name, description string) error {
	refs, err := groupRefs(group)
	if err != nil {
		return err
	}
	var failed []string
	for _, ref := range refs {
		cred, credErr := resolveCredFn(ref.Alias)
		if credErr != nil {
			failed = append(failed, channelErr(ref.Alias, "凭据不可用", credErr).Error())
			continue
		}
		if _, updateErr := updateGroupFn(cred, ref.UpstreamId, name, description); updateErr != nil {
			failed = append(failed, channelErr(ref.Alias, "更新素材组失败", updateErr).Error())
		}
	}
	if len(failed) > 0 {
		return fmt.Errorf("%s", strings.Join(failed, "; "))
	}
	return nil
}

// assetRefs 返回素材在各渠道的上游 ID。线上历史素材没有绑定，按"逻辑 ID 就是火山 ID"补上 volc。
func assetRefs(asset *model.PortraitAsset) ([]channelRef, error) {
	bindings, err := model.GetPortraitAssetBindings(asset.RemoteAssetId)
	if err != nil {
		return nil, err
	}
	if len(bindings) == 0 {
		return []channelRef{{Alias: portrait_setting.DefaultAlias, UpstreamId: asset.RemoteAssetId}}, nil
	}
	refs := make([]channelRef, 0, len(bindings))
	for _, b := range bindings {
		refs = append(refs, channelRef{Alias: b.PortraitChannel, UpstreamId: b.UpstreamAssetId})
	}
	return refs, nil
}

// UpdateAssetAllChannels 对素材在所有渠道里的上游素材各更新一次名称。
func UpdateAssetAllChannels(asset *model.PortraitAsset, name string) error {
	refs, err := assetRefs(asset)
	if err != nil {
		return err
	}
	var failed []string
	for _, ref := range refs {
		cred, credErr := resolveCredFn(ref.Alias)
		if credErr != nil {
			failed = append(failed, channelErr(ref.Alias, "凭据不可用", credErr).Error())
			continue
		}
		if _, updateErr := updateAssetFn(cred, ref.UpstreamId, name); updateErr != nil {
			failed = append(failed, channelErr(ref.Alias, "更新素材失败", updateErr).Error())
		}
	}
	if len(failed) > 0 {
		return fmt.Errorf("%s", strings.Join(failed, "; "))
	}
	return nil
}

// AssetChannels 返回素材在各渠道的状态。没有绑定的历史素材只有 volc 一条。
func AssetChannels(asset *model.PortraitAsset) []ChannelStatus {
	bindings, err := model.GetPortraitAssetBindings(asset.RemoteAssetId)
	if err != nil || len(bindings) == 0 {
		return []ChannelStatus{{PortraitChannel: portrait_setting.DefaultAlias, Status: asset.Status}}
	}
	out := make([]ChannelStatus, 0, len(bindings))
	for _, b := range bindings {
		out = append(out, ChannelStatus{PortraitChannel: b.PortraitChannel, Status: b.Status, LastError: b.LastError})
	}
	return out
}

func isTerminalPortraitStatus(status string) bool {
	return status == model.PortraitStatusActive || status == model.PortraitStatusFailed
}

// refreshBinding 向对应渠道账号查询一次素材审核状态并写回绑定。
func refreshBinding(b *model.PortraitAssetBinding, cred portrait_setting.Credential) error {
	result, err := getAssetFn(cred, b.UpstreamAssetId)
	if err != nil {
		return err
	}
	changed := false
	if result.Status != "" && result.Status != b.Status {
		b.Status = result.Status
		changed = true
	}
	if result.ResolvedUrl != "" && result.ResolvedUrl != b.ResolvedUrl {
		b.ResolvedUrl = result.ResolvedUrl
		changed = true
	}
	if result.Error != "" && result.Error != b.LastError {
		b.LastError = result.Error
		changed = true
	}
	if b.Status == model.PortraitStatusActive && b.LastError != "" {
		b.LastError = ""
		changed = true
	}
	if !changed {
		return nil
	}
	return model.SavePortraitAssetBinding(b)
}

// RecomputeAssetStatus 按所有渠道绑定重算逻辑素材的汇总状态（严格模式），变化时写回。
// 没有绑定的历史素材不处理。
func RecomputeAssetStatus(asset *model.PortraitAsset) error {
	bindings, err := model.GetPortraitAssetBindings(asset.RemoteAssetId)
	if err != nil || len(bindings) == 0 {
		return err
	}
	statuses := make([]string, len(bindings))
	resolved := ""
	for i, b := range bindings {
		statuses[i] = b.Status
		if b.ResolvedUrl == "" {
			continue
		}
		if b.PortraitChannel == portrait_setting.DefaultAlias || resolved == "" {
			resolved = b.ResolvedUrl
		}
	}
	status := model.AggregatePortraitStatus(statuses)
	if status == asset.Status && (resolved == "" || resolved == asset.ResolvedUrl) {
		return nil
	}
	asset.Status = status
	if resolved != "" {
		asset.ResolvedUrl = resolved
	}
	return model.SavePortraitAsset(asset)
}

// credCache 在一轮刷新里按渠道标识缓存凭据，解析失败也缓存，避免重复报错。
type credCache struct {
	creds map[string]portrait_setting.Credential
	errs  map[string]error
}

func newCredCache() *credCache {
	return &credCache{
		creds: map[string]portrait_setting.Credential{},
		errs:  map[string]error{},
	}
}

func (c *credCache) get(alias string) (portrait_setting.Credential, error) {
	if cred, ok := c.creds[alias]; ok {
		return cred, nil
	}
	if err, failed := c.errs[alias]; failed {
		return portrait_setting.Credential{}, err
	}
	cred, err := resolveCredFn(alias)
	if err != nil {
		c.errs[alias] = err
		return portrait_setting.Credential{}, err
	}
	c.creds[alias] = cred
	return cred, nil
}

// RefreshAsset 实时刷新素材的上游审核状态并重算汇总状态，返回刷新过程中的错误信息（可空）。
func RefreshAsset(asset *model.PortraitAsset) string {
	bindings, err := model.GetPortraitAssetBindings(asset.RemoteAssetId)
	if err != nil {
		return err.Error()
	}
	if len(bindings) == 0 {
		return refreshLegacyAsset(asset, nil)
	}

	cache := newCredCache()
	var errs []string
	for _, b := range bindings {
		if isTerminalPortraitStatus(b.Status) {
			continue
		}
		cred, credErr := cache.get(b.PortraitChannel)
		if credErr != nil {
			errs = append(errs, channelErr(b.PortraitChannel, "凭据不可用", credErr).Error())
			continue
		}
		if refreshErr := refreshBinding(b, cred); refreshErr != nil {
			errs = append(errs, channelErr(b.PortraitChannel, "查询素材失败", refreshErr).Error())
		}
	}
	if err = RecomputeAssetStatus(asset); err != nil {
		errs = append(errs, err.Error())
	}
	return strings.Join(errs, "; ")
}

// refreshLegacyAsset 刷新没有任何绑定的线上历史素材，只走火山凭据。
func refreshLegacyAsset(asset *model.PortraitAsset, cache *credCache) string {
	if isTerminalPortraitStatus(asset.Status) {
		return ""
	}
	if cache == nil {
		cache = newCredCache()
	}
	cred, err := cache.get(portrait_setting.DefaultAlias)
	if err != nil {
		return channelErr(portrait_setting.DefaultAlias, "凭据不可用", err).Error()
	}
	result, err := getAssetFn(cred, asset.RemoteAssetId)
	if err != nil {
		return channelErr(portrait_setting.DefaultAlias, "查询素材失败", err).Error()
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
			return saveErr.Error()
		}
	}
	return ""
}
