package portrait

import (
	"fmt"
	"strings"
	"testing"

	"github.com/QuantumNous/new-api/model"
	"github.com/QuantumNous/new-api/setting/portrait_setting"
	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

// fakeUpstream 记录对各渠道账号的调用，并按预设返回结果。
type fakeUpstream struct {
	groupSeq      int
	assetSeq      int
	createdGroups map[string][]string // alias -> group ids
	createdAssets map[string][]string // alias -> asset ids
	assetStatus   map[string]string   // 上游素材 ID -> 状态
	assetError    map[string]string   // 上游素材 ID -> 失败信息
	failCreateOn  string              // 指定渠道的创建素材直接失败
	credLookups   []string
}

func newFakeUpstream() *fakeUpstream {
	return &fakeUpstream{
		createdGroups: map[string][]string{},
		createdAssets: map[string][]string{},
		assetStatus:   map[string]string{},
		assetError:    map[string]string{},
	}
}

func (f *fakeUpstream) install(t *testing.T, aliases ...string) {
	origCreateGroup, origCreateAsset, origGetAsset := createGroupFn, createAssetFn, getAssetFn
	origUpdateGroup, origUpdateAsset := updateGroupFn, updateAssetFn
	origResolve, origFanout := resolveCredFn, fanoutAliasesFn
	t.Cleanup(func() {
		createGroupFn, createAssetFn, getAssetFn = origCreateGroup, origCreateAsset, origGetAsset
		updateGroupFn, updateAssetFn = origUpdateGroup, origUpdateAsset
		resolveCredFn, fanoutAliasesFn = origResolve, origFanout
	})

	fanoutAliasesFn = func() []string { return aliases }
	resolveCredFn = func(alias string) (portrait_setting.Credential, error) {
		f.credLookups = append(f.credLookups, alias)
		return portrait_setting.Credential{Alias: alias, AK: "ak", SK: "sk"}, nil
	}
	createGroupFn = func(cred portrait_setting.Credential, name, description, groupType string) (*CreateAssetGroupResult, error) {
		f.groupSeq++
		id := fmt.Sprintf("%s-group-%d", cred.Alias, f.groupSeq)
		f.createdGroups[cred.Alias] = append(f.createdGroups[cred.Alias], id)
		return &CreateAssetGroupResult{Id: id}, nil
	}
	createAssetFn = func(cred portrait_setting.Credential, groupId, assetURL, assetType, name string) (*CreateAssetResult, error) {
		if cred.Alias == f.failCreateOn {
			return nil, fmt.Errorf("审核服务拒绝")
		}
		f.assetSeq++
		id := fmt.Sprintf("%s-asset-%d", cred.Alias, f.assetSeq)
		f.createdAssets[cred.Alias] = append(f.createdAssets[cred.Alias], id)
		return &CreateAssetResult{AssetId: id}, nil
	}
	getAssetFn = func(cred portrait_setting.Credential, assetId string) (*GetAssetResult, error) {
		return &GetAssetResult{Status: f.assetStatus[assetId], Error: f.assetError[assetId]}, nil
	}
	updateGroupFn = func(cred portrait_setting.Credential, groupId, name, description string) (*UpdateAssetGroupResult, error) {
		return &UpdateAssetGroupResult{Id: groupId}, nil
	}
	updateAssetFn = func(cred portrait_setting.Credential, assetId, name string) (*UpdateAssetResult, error) {
		return &UpdateAssetResult{Id: assetId}, nil
	}
}

func setupDB(t *testing.T) {
	t.Helper()
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	require.NoError(t, err)
	sqlDB, err := db.DB()
	require.NoError(t, err)
	sqlDB.SetMaxOpenConns(1)
	require.NoError(t, db.AutoMigrate(
		&model.PortraitGroup{}, &model.PortraitAsset{},
		&model.PortraitGroupBinding{}, &model.PortraitAssetBinding{},
	))
	orig := model.DB
	model.DB = db
	t.Cleanup(func() { model.DB = orig })
}

func TestCreateAssetAllChannelsWritesEveryChannel(t *testing.T) {
	setupDB(t)
	up := newFakeUpstream()
	up.install(t, "volc", "zj")

	group, err := CreateGroupAllChannels(1, "g", "", "AIGC")
	require.NoError(t, err)
	assert.Equal(t, "volc-group-1", group.RemoteGroupId, "逻辑 group_id 取主账号上游 ID")

	asset, err := CreateAssetAllChannels(1, group.RemoteGroupId, "http://x/a.png", "Image", "a")
	require.NoError(t, err)
	assert.Equal(t, "volc-asset-1", asset.RemoteAssetId, "逻辑 asset_id 取主账号上游 ID")
	assert.Equal(t, model.PortraitStatusSubmitted, asset.Status)

	bindings, err := model.GetPortraitAssetBindings(asset.RemoteAssetId)
	require.NoError(t, err)
	require.Len(t, bindings, 2)
	assert.Equal(t, "volc", bindings[0].PortraitChannel)
	assert.Equal(t, "zj", bindings[1].PortraitChannel)
	assert.Equal(t, "zj-asset-2", bindings[1].UpstreamAssetId)

	channels := AssetChannels(asset)
	require.Len(t, channels, 2)
}

func TestCreateAssetAllChannelsFailureReturnsChannelError(t *testing.T) {
	setupDB(t)
	up := newFakeUpstream()
	up.install(t, "volc", "zj")
	group, err := CreateGroupAllChannels(1, "g", "", "AIGC")
	require.NoError(t, err)

	up.failCreateOn = "zj"
	_, err = CreateAssetAllChannels(1, group.RemoteGroupId, "http://x/a.png", "Image", "a")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "zj")
	assert.Contains(t, err.Error(), "审核服务拒绝")

	var count int64
	require.NoError(t, model.DB.Model(&model.PortraitAsset{}).Count(&count).Error)
	assert.Zero(t, count, "失败时不保存半成品")
}

func TestCreateAssetLazilyBackfillsLegacyGroup(t *testing.T) {
	setupDB(t)
	up := newFakeUpstream()
	up.install(t, "volc", "zj")

	// 线上历史素材组：只有火山 ID，没有绑定。
	legacy := &model.PortraitGroup{UserId: 1, RemoteGroupId: "legacy-volc-group", Name: "old"}
	require.NoError(t, model.CreatePortraitGroup(legacy))

	asset, err := CreateAssetAllChannels(1, "legacy-volc-group", "http://x/a.png", "Image", "a")
	require.NoError(t, err)

	groupBindings, err := model.GetPortraitGroupBindings("legacy-volc-group")
	require.NoError(t, err)
	require.Len(t, groupBindings, 2)
	assert.Equal(t, "legacy-volc-group", groupBindings[0].UpstreamGroupId, "volc 绑定沿用逻辑 ID，不再建组")
	assert.Len(t, up.createdGroups["volc"], 0)
	assert.Len(t, up.createdGroups["zj"], 1, "zj 当场补建素材组")
	assert.Equal(t, "volc-asset-1", asset.RemoteAssetId)
}

func TestCreateAssetRejectsUnknownGroup(t *testing.T) {
	setupDB(t)
	newFakeUpstream().install(t, "volc", "zj")
	_, err := CreateAssetAllChannels(1, "nope", "http://x/a.png", "Image", "a")
	require.Error(t, err)
}

func TestStrictAggregateAndRefresh(t *testing.T) {
	setupDB(t)
	up := newFakeUpstream()
	up.install(t, "volc", "zj")
	group, err := CreateGroupAllChannels(1, "g", "", "AIGC")
	require.NoError(t, err)
	asset, err := CreateAssetAllChannels(1, group.RemoteGroupId, "http://x/a.png", "Image", "a")
	require.NoError(t, err)

	// volc 先通过，zj 仍在审核：整体不能算 Active。
	up.assetStatus["volc-asset-1"] = "Active"
	up.assetStatus["zj-asset-2"] = "Processing"
	assert.Empty(t, RefreshAsset(asset))
	assert.Equal(t, model.PortraitStatusProcessing, asset.Status)

	// 两边都通过才是 Active。
	up.assetStatus["zj-asset-2"] = "Active"
	assert.Empty(t, RefreshAsset(asset))
	assert.Equal(t, model.PortraitStatusActive, asset.Status)
}

func TestFailedChannelFailsAssetAndKeepsError(t *testing.T) {
	setupDB(t)
	up := newFakeUpstream()
	up.install(t, "volc", "zj")
	group, err := CreateGroupAllChannels(1, "g", "", "AIGC")
	require.NoError(t, err)
	asset, err := CreateAssetAllChannels(1, group.RemoteGroupId, "http://x/a.png", "Image", "a")
	require.NoError(t, err)

	up.assetStatus["volc-asset-1"] = "Active"
	up.assetStatus["zj-asset-2"] = "Failed"
	up.assetError["zj-asset-2"] = "face not detected"
	RefreshAsset(asset)

	assert.Equal(t, model.PortraitStatusFailed, asset.Status)
	var zj *ChannelStatus
	for _, ch := range AssetChannels(asset) {
		if ch.PortraitChannel == "zj" {
			c := ch
			zj = &c
		}
	}
	require.NotNil(t, zj)
	assert.Equal(t, "Failed", zj.Status)
	assert.Equal(t, "face not detected", zj.LastError, "失败信息要原样返回给调用方")
}

func TestPollerUsesEachBindingCredentialAndSkipsBadOnes(t *testing.T) {
	setupDB(t)
	up := newFakeUpstream()
	up.install(t, "volc", "zj")
	group, err := CreateGroupAllChannels(1, "g", "", "AIGC")
	require.NoError(t, err)
	asset, err := CreateAssetAllChannels(1, group.RemoteGroupId, "http://x/a.png", "Image", "a")
	require.NoError(t, err)

	// 另一条素材的 zj 凭据解析失败，不能拖垮其他素材。
	bad := &model.PortraitAsset{UserId: 1, RemoteGroupId: "g", RemoteAssetId: "ghost-1", AssetType: "Image", SourceUrl: "u", Status: "Submitted"}
	require.NoError(t, model.CreatePortraitAssetWithBindings(bad, []*model.PortraitAssetBinding{
		{UserId: 1, AssetId: "ghost-1", PortraitChannel: "broken", UpstreamAssetId: "x", Status: "Submitted"},
	}))
	baseResolve := resolveCredFn
	resolveCredFn = func(alias string) (portrait_setting.Credential, error) {
		if alias == "broken" {
			return portrait_setting.Credential{}, fmt.Errorf("人像库渠道标识不存在或未配置")
		}
		return baseResolve(alias)
	}

	up.assetStatus["volc-asset-1"] = "Active"
	up.assetStatus["zj-asset-2"] = "Active"
	pollPendingAssets()

	got, err := model.GetPortraitAssetByRemoteId(1, asset.RemoteAssetId)
	require.NoError(t, err)
	assert.Equal(t, model.PortraitStatusActive, got.Status)

	ghost, err := model.GetPortraitAssetByRemoteId(1, "ghost-1")
	require.NoError(t, err)
	assert.Equal(t, "Submitted", ghost.Status)
}

func TestPollerRefreshesLegacyAssetWithVolcCredential(t *testing.T) {
	setupDB(t)
	up := newFakeUpstream()
	up.install(t, "volc", "zj")
	legacy := &model.PortraitAsset{UserId: 1, RemoteGroupId: "g", RemoteAssetId: "legacy-asset", AssetType: "Image", SourceUrl: "u", Status: "Processing"}
	require.NoError(t, model.CreatePortraitAsset(legacy))

	up.assetStatus["legacy-asset"] = "Active"
	pollPendingAssets()

	got, err := model.GetPortraitAssetByRemoteId(1, "legacy-asset")
	require.NoError(t, err)
	assert.Equal(t, model.PortraitStatusActive, got.Status)
	for _, alias := range up.credLookups {
		assert.Equal(t, "volc", alias, "历史素材只走火山凭据")
	}
}

func TestBackfillAddsMissingChannelAndIsIdempotent(t *testing.T) {
	setupDB(t)
	up := newFakeUpstream()
	up.install(t, "volc", "zj")

	legacyGroup := &model.PortraitGroup{UserId: 1, RemoteGroupId: "legacy-group", Name: "old"}
	require.NoError(t, model.CreatePortraitGroup(legacyGroup))
	legacyAsset := &model.PortraitAsset{UserId: 1, RemoteGroupId: "legacy-group", RemoteAssetId: "legacy-asset", AssetType: "Image", SourceUrl: "http://x/a.png", Status: "Active"}
	require.NoError(t, model.CreatePortraitAsset(legacyAsset))
	failed := &model.PortraitAsset{UserId: 1, RemoteGroupId: "legacy-group", RemoteAssetId: "legacy-failed", AssetType: "Image", SourceUrl: "http://x/b.png", Status: "Failed"}
	require.NoError(t, model.CreatePortraitAsset(failed))

	res, err := BackfillMissingChannels()
	require.NoError(t, err)
	assert.Equal(t, 1, res.Created)
	assert.Equal(t, 1, res.Skipped, "审核失败的历史素材不补建")
	assert.Zero(t, res.Failed)

	bindings, err := model.GetPortraitAssetBindings("legacy-asset")
	require.NoError(t, err)
	require.Len(t, bindings, 2)
	assert.Equal(t, "volc", bindings[0].PortraitChannel)
	assert.Equal(t, "Active", bindings[0].Status, "volc 绑定沿用历史状态")
	assert.Equal(t, "legacy-asset", bindings[0].UpstreamAssetId)

	got, err := model.GetPortraitAssetByRemoteId(1, "legacy-asset")
	require.NoError(t, err)
	assert.Equal(t, model.PortraitStatusProcessing, got.Status, "zj 未就绪前汇总状态回到 Processing")

	// 第二次执行不再补建。
	res, err = BackfillMissingChannels()
	require.NoError(t, err)
	assert.Zero(t, res.Created)
}

func TestUpdateAllChannelsReportsFailedChannel(t *testing.T) {
	setupDB(t)
	up := newFakeUpstream()
	up.install(t, "volc", "zj")
	group, err := CreateGroupAllChannels(1, "g", "", "AIGC")
	require.NoError(t, err)

	updateGroupFn = func(cred portrait_setting.Credential, groupId, name, description string) (*UpdateAssetGroupResult, error) {
		if cred.Alias == "zj" {
			return nil, fmt.Errorf("boom")
		}
		return &UpdateAssetGroupResult{Id: groupId}, nil
	}
	err = UpdateGroupAllChannels(group, "new", "")
	require.Error(t, err)
	assert.True(t, strings.Contains(err.Error(), "zj") && strings.Contains(err.Error(), "boom"))
}
