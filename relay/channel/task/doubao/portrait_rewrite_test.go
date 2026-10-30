package doubao

import (
	"testing"

	"github.com/QuantumNous/new-api/constant"
	"github.com/QuantumNous/new-api/model"
	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func setupPortraitDB(t *testing.T) {
	t.Helper()
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	require.NoError(t, err)
	sqlDB, err := db.DB()
	require.NoError(t, err)
	sqlDB.SetMaxOpenConns(1)
	require.NoError(t, db.AutoMigrate(&model.PortraitAsset{}, &model.PortraitAssetBinding{}))
	orig := model.DB
	model.DB = db
	t.Cleanup(func() { model.DB = orig })

	require.NoError(t, model.CreatePortraitAssetWithBindings(
		&model.PortraitAsset{UserId: 1, RemoteGroupId: "g", RemoteAssetId: "volc-a", AssetType: "Image", SourceUrl: "u", Status: "Active"},
		[]*model.PortraitAssetBinding{
			{UserId: 1, AssetId: "volc-a", PortraitChannel: "volc", UpstreamAssetId: "volc-a", Status: "Active"},
			{UserId: 1, AssetId: "volc-a", PortraitChannel: "zj", UpstreamAssetId: "zj-a", Status: "Active"},
		}))
}

func portraitContent() []ContentItem {
	return []ContentItem{
		{Type: "text", Text: "hello"},
		{Type: "image_url", ImageURL: &MediaURL{URL: "asset://volc-a"}},
		{Type: "image_url", ImageURL: &MediaURL{URL: "https://example.com/a.png"}},
	}
}

func TestRewritePortraitAssetsPerChannel(t *testing.T) {
	setupPortraitDB(t)

	zj := portraitContent()
	require.NoError(t, rewritePortraitAssets(1, constant.ChannelTypeZhongJiVideo, zj))
	assert.Equal(t, "asset://zj-a", zj[1].ImageURL.URL, "中纪渠道改写成中纪账号的素材 ID")
	assert.Equal(t, "https://example.com/a.png", zj[2].ImageURL.URL)

	volc := portraitContent()
	require.NoError(t, rewritePortraitAssets(1, constant.ChannelTypeDoubaoVideo, volc))
	assert.Equal(t, "asset://volc-a", volc[1].ImageURL.URL)
}

func TestRewritePortraitAssetsRejectsForeignUser(t *testing.T) {
	setupPortraitDB(t)
	err := rewritePortraitAssets(2, constant.ChannelTypeZhongJiVideo, portraitContent())
	require.Error(t, err)
}
