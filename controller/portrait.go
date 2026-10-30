package controller

import (
	"strings"

	"github.com/QuantumNous/new-api/common"
	"github.com/QuantumNous/new-api/model"
	portraitSvc "github.com/QuantumNous/new-api/service/portrait"
	"github.com/gin-gonic/gin"
)

// PortraitCreateGroupRequest 创建素材组请求体
type PortraitCreateGroupRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	GroupType   string `json:"group_type"`
}

// PortraitUpdateGroupRequest 更新素材组请求体
type PortraitUpdateGroupRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// PortraitCreateAssetRequest 创建素材请求体
type PortraitCreateAssetRequest struct {
	GroupId   string `json:"group_id"`
	URL       string `json:"url"`
	AssetType string `json:"asset_type"`
	Name      string `json:"name"`
}

// PortraitUpdateAssetRequest 更新素材请求体（当前仅支持更新 Name）
type PortraitUpdateAssetRequest struct {
	Name string `json:"name"`
}

// PortraitCreateGroup POST /v1/portrait/groups
func PortraitCreateGroup(c *gin.Context) {
	userId := c.GetInt("id")
	if userId == 0 {
		common.ApiErrorMsg(c, "未授权")
		return
	}

	var req PortraitCreateGroupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		common.ApiErrorMsg(c, "请求体解析失败")
		return
	}
	req.Name = strings.TrimSpace(req.Name)
	if req.Name == "" {
		common.ApiErrorMsg(c, "素材组名称不能为空")
		return
	}
	if req.GroupType == "" {
		req.GroupType = "AIGC"
	}

	group, err := portraitSvc.CreateGroupAllChannels(userId, req.Name, req.Description, req.GroupType)
	if err != nil {
		common.ApiErrorMsg(c, err.Error())
		return
	}

	common.ApiSuccess(c, gin.H{
		"group_id": group.RemoteGroupId,
		"name":     req.Name,
		"id":       group.Id,
	})
}

// PortraitListGroups GET /v1/portrait/groups
func PortraitListGroups(c *gin.Context) {
	userId := c.GetInt("id")
	if userId == 0 {
		common.ApiErrorMsg(c, "未授权")
		return
	}

	pageInfo := common.GetPageQuery(c)
	groups, total, err := model.GetPortraitGroupsByUserIdPaged(userId, pageInfo.GetStartIdx(), pageInfo.GetPageSize())
	if err != nil {
		common.ApiError(c, err)
		return
	}
	pageInfo.SetTotal(int(total))
	pageInfo.SetItems(groups)
	common.ApiSuccess(c, pageInfo)
}

// PortraitUpdateGroup PUT /v1/portrait/groups/:groupId
func PortraitUpdateGroup(c *gin.Context) {
	userId := c.GetInt("id")
	if userId == 0 {
		common.ApiErrorMsg(c, "未授权")
		return
	}

	remoteGroupId := c.Param("groupId")
	if remoteGroupId == "" {
		common.ApiErrorMsg(c, "groupId 不能为空")
		return
	}

	group, err := model.GetPortraitGroupByRemoteId(userId, remoteGroupId)
	if err != nil {
		common.ApiErrorMsg(c, "素材组不存在")
		return
	}

	var req PortraitUpdateGroupRequest
	if err = c.ShouldBindJSON(&req); err != nil {
		common.ApiErrorMsg(c, "请求体解析失败")
		return
	}
	req.Name = strings.TrimSpace(req.Name)
	req.Description = strings.TrimSpace(req.Description)
	if req.Name == "" && req.Description == "" {
		common.ApiErrorMsg(c, "name 或 description 至少传一个")
		return
	}

	if err = portraitSvc.UpdateGroupAllChannels(group, req.Name, req.Description); err != nil {
		common.ApiErrorMsg(c, err.Error())
		return
	}

	if req.Name != "" {
		group.Name = req.Name
		_ = model.DB.Save(group).Error
	}

	common.ApiSuccess(c, gin.H{
		"group_id": remoteGroupId,
		"name":     group.Name,
	})
}

// PortraitCreateAsset POST /v1/portrait/assets
func PortraitCreateAsset(c *gin.Context) {
	userId := c.GetInt("id")
	if userId == 0 {
		common.ApiErrorMsg(c, "未授权")
		return
	}

	var req PortraitCreateAssetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		common.ApiErrorMsg(c, "请求体解析失败")
		return
	}

	req.GroupId = strings.TrimSpace(req.GroupId)
	req.URL = strings.TrimSpace(req.URL)
	req.AssetType = strings.TrimSpace(req.AssetType)

	if req.GroupId == "" {
		common.ApiErrorMsg(c, "group_id 不能为空")
		return
	}
	if req.URL == "" {
		common.ApiErrorMsg(c, "url 不能为空")
		return
	}
	validTypes := map[string]bool{"Image": true, "Video": true, "Audio": true}
	if !validTypes[req.AssetType] {
		common.ApiErrorMsg(c, "asset_type 须为 Image / Video / Audio")
		return
	}

	asset, err := portraitSvc.CreateAssetAllChannels(userId, req.GroupId, req.URL, req.AssetType, req.Name)
	if err != nil {
		common.ApiErrorMsg(c, err.Error())
		return
	}

	common.ApiSuccess(c, gin.H{
		"id":              asset.Id,
		"asset_id":        asset.RemoteAssetId,
		"remote_group_id": asset.RemoteGroupId,
		"status":          asset.Status,
		"channels":        portraitSvc.AssetChannels(asset),
	})
}

// PortraitListAssets GET /v1/portrait/assets
func PortraitListAssets(c *gin.Context) {
	userId := c.GetInt("id")
	if userId == 0 {
		common.ApiErrorMsg(c, "未授权")
		return
	}

	groupId := c.Query("group_id")
	pageInfo := common.GetPageQuery(c)
	assets, total, err := model.GetPortraitAssetsByUserIdPaged(userId, groupId, pageInfo.GetStartIdx(), pageInfo.GetPageSize())
	if err != nil {
		common.ApiError(c, err)
		return
	}
	views := make([]portraitAssetView, 0, len(assets))
	for _, asset := range assets {
		views = append(views, portraitAssetView{PortraitAsset: asset, Channels: portraitSvc.AssetChannels(asset)})
	}
	pageInfo.SetTotal(int(total))
	pageInfo.SetItems(views)
	common.ApiSuccess(c, pageInfo)
}

// portraitAssetView 在素材记录上附带各渠道状态。
type portraitAssetView struct {
	*model.PortraitAsset
	Channels []portraitSvc.ChannelStatus `json:"channels"`
}

// PortraitBackfill POST /api/portrait/backfill（仅管理员）
// 把缺少渠道绑定的素材补建到所有参与渠道，幂等，可重复执行。
func PortraitBackfill(c *gin.Context) {
	result, err := portraitSvc.BackfillMissingChannels()
	if err != nil {
		common.ApiErrorMsg(c, err.Error())
		return
	}
	common.ApiSuccess(c, result)
}

// PortraitGetAsset GET /v1/portrait/assets/:assetId
// 实时刷新各渠道上游状态，并更新本地记录
func PortraitGetAsset(c *gin.Context) {
	userId := c.GetInt("id")
	if userId == 0 {
		common.ApiErrorMsg(c, "未授权")
		return
	}

	remoteAssetId := c.Param("assetId")
	if remoteAssetId == "" {
		common.ApiErrorMsg(c, "assetId 不能为空")
		return
	}

	asset, err := model.GetPortraitAssetByRemoteId(userId, remoteAssetId)
	if err != nil {
		common.ApiErrorMsg(c, "素材不存在")
		return
	}

	buildAssetResponse := func(refreshErr string) gin.H {
		h := gin.H{
			"asset_id":     asset.RemoteAssetId,
			"status":       asset.Status,
			"resolved_url": asset.ResolvedUrl,
			"name":         asset.Name,
			"asset_type":   asset.AssetType,
			"source_url":   asset.SourceUrl,
			"channels":     portraitSvc.AssetChannels(asset),
			"created_at":   asset.CreatedAt,
			"updated_at":   asset.UpdatedAt,
		}
		if refreshErr != "" {
			h["refresh_error"] = refreshErr
		}
		return h
	}

	// 终态：Active（所有渠道都通过）、Failed（任一渠道失败），不再调上游
	if asset.Status == model.PortraitStatusActive || asset.Status == model.PortraitStatusFailed {
		common.ApiSuccess(c, buildAssetResponse(""))
		return
	}

	// 非终态：实时刷新各渠道状态并重算汇总
	refreshErr := portraitSvc.RefreshAsset(asset)
	common.ApiSuccess(c, buildAssetResponse(refreshErr))
}

// PortraitUpdateAsset PUT /v1/portrait/assets/:assetId
func PortraitUpdateAsset(c *gin.Context) {
	userId := c.GetInt("id")
	if userId == 0 {
		common.ApiErrorMsg(c, "未授权")
		return
	}

	remoteAssetId := c.Param("assetId")
	if remoteAssetId == "" {
		common.ApiErrorMsg(c, "assetId 不能为空")
		return
	}

	asset, err := model.GetPortraitAssetByRemoteId(userId, remoteAssetId)
	if err != nil {
		common.ApiErrorMsg(c, "素材不存在")
		return
	}

	var req PortraitUpdateAssetRequest
	if err = c.ShouldBindJSON(&req); err != nil {
		common.ApiErrorMsg(c, "请求体解析失败")
		return
	}
	req.Name = strings.TrimSpace(req.Name)
	if req.Name == "" {
		common.ApiErrorMsg(c, "name 不能为空")
		return
	}

	if err = portraitSvc.UpdateAssetAllChannels(asset, req.Name); err != nil {
		common.ApiErrorMsg(c, err.Error())
		return
	}

	asset.Name = req.Name
	_ = model.SavePortraitAsset(asset)

	common.ApiSuccess(c, gin.H{
		"asset_id": remoteAssetId,
		"name":     asset.Name,
	})
}
