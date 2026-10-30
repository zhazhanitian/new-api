package portrait

import (
	"testing"

	"github.com/QuantumNous/new-api/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func seedRewriteData(t *testing.T) {
	t.Helper()
	setupDB(t)

	// 多写素材：volc 与 zj 都已 Active。
	require.NoError(t, model.CreatePortraitAssetWithBindings(
		&model.PortraitAsset{UserId: 1, RemoteGroupId: "g", RemoteAssetId: "volc-a", AssetType: "Image", SourceUrl: "u", Status: "Active"},
		[]*model.PortraitAssetBinding{
			{UserId: 1, AssetId: "volc-a", PortraitChannel: "volc", UpstreamAssetId: "volc-a", Status: "Active"},
			{UserId: 1, AssetId: "volc-a", PortraitChannel: "zj", UpstreamAssetId: "zj-a", Status: "Active"},
		}))
	// zj 仍在审核。
	require.NoError(t, model.CreatePortraitAssetWithBindings(
		&model.PortraitAsset{UserId: 1, RemoteGroupId: "g", RemoteAssetId: "volc-b", AssetType: "Image", SourceUrl: "u", Status: "Processing"},
		[]*model.PortraitAssetBinding{
			{UserId: 1, AssetId: "volc-b", PortraitChannel: "volc", UpstreamAssetId: "volc-b", Status: "Active"},
			{UserId: 1, AssetId: "volc-b", PortraitChannel: "zj", UpstreamAssetId: "zj-b", Status: "Failed", LastError: "face not detected"},
		}))
	// 线上历史素材：没有绑定。
	require.NoError(t, model.CreatePortraitAsset(
		&model.PortraitAsset{UserId: 1, RemoteGroupId: "g", RemoteAssetId: "legacy-active", AssetType: "Image", SourceUrl: "u", Status: "Active"}))
	require.NoError(t, model.CreatePortraitAsset(
		&model.PortraitAsset{UserId: 1, RemoteGroupId: "g", RemoteAssetId: "legacy-processing", AssetType: "Image", SourceUrl: "u", Status: "Processing"}))
}

func TestRewriteAssetURL(t *testing.T) {
	seedRewriteData(t)

	cases := []struct {
		name    string
		userId  int
		alias   string
		url     string
		want    string
		wantErr string
	}{
		{"非素材 URL 原样返回", 1, "zj", "https://x/a.png", "https://x/a.png", ""},
		{"volc 渠道取火山 ID", 1, "volc", "asset://volc-a", "asset://volc-a", ""},
		{"zj 渠道改写成中纪 ID", 1, "zj", "asset://volc-a", "asset://zj-a", ""},
		{"空渠道按 volc", 1, "", "asset://volc-a", "asset://volc-a", ""},
		{"zj 绑定失败返回渠道和原因", 1, "zj", "asset://volc-b", "", "face not detected"},
		{"同一素材在 volc 渠道仍可用", 1, "volc", "asset://volc-b", "asset://volc-b", ""},
		{"非本人素材", 2, "volc", "asset://volc-a", "", "不存在或不属于当前用户"},
		{"素材不存在", 1, "volc", "asset://nope", "", "不存在或不属于当前用户"},
		{"历史素材 volc 放行", 1, "volc", "asset://legacy-active", "asset://legacy-active", ""},
		{"历史素材 zj 未就绪", 1, "zj", "asset://legacy-active", "", "尚未在渠道 zj 就绪"},
		{"历史素材未 Active", 1, "volc", "asset://legacy-processing", "", "需为 Active"},
		{"未配置的渠道未就绪", 1, "other", "asset://volc-a", "", "尚未在渠道 other 就绪"},
	}
	for _, tc := range cases {
		got, err := RewriteAssetURL(tc.userId, tc.alias, tc.url)
		if tc.wantErr != "" {
			require.Error(t, err, tc.name)
			assert.Contains(t, err.Error(), tc.wantErr, tc.name)
			continue
		}
		require.NoError(t, err, tc.name)
		assert.Equal(t, tc.want, got, tc.name)
	}
}

func TestParseAssetURL(t *testing.T) {
	id, ok := ParseAssetURL(" asset://abc ")
	assert.True(t, ok)
	assert.Equal(t, "abc", id)
	for _, bad := range []string{"", "https://x", "asset://", "asset://  "} {
		_, ok := ParseAssetURL(bad)
		assert.False(t, ok, bad)
	}
}
