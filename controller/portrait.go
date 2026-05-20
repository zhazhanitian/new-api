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

	result, err := portraitSvc.CreateAssetGroup(req.Name, req.Description, req.GroupType)
	if err != nil {
		common.ApiErrorMsg(c, "调用火山引擎失败: "+err.Error())
		return
	}

	group := &model.PortraitGroup{
		UserId:        userId,
		RemoteGroupId: result.Id,
		Name:          req.Name,
	}
	if err = model.CreatePortraitGroup(group); err != nil {
		common.ApiErrorMsg(c, "保存素材组记录失败: "+err.Error())
		return
	}

	common.ApiSuccess(c, gin.H{
		"group_id": result.Id,
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

	if _, err = portraitSvc.UpdateAssetGroup(remoteGroupId, req.Name, req.Description); err != nil {
		common.ApiErrorMsg(c, "调用火山引擎失败: "+err.Error())
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

	result, err := portraitSvc.CreateAsset(req.GroupId, req.URL, req.AssetType, req.Name)
	if err != nil {
		common.ApiErrorMsg(c, "调用火山引擎失败: "+err.Error())
		return
	}

	asset := &model.PortraitAsset{
		UserId:        userId,
		RemoteGroupId: req.GroupId,
		RemoteAssetId: result.AssetId,
		Name:          req.Name,
		AssetType:     req.AssetType,
		SourceUrl:     req.URL,
		Status:        "Submitted",
	}
	if err = model.CreatePortraitAsset(asset); err != nil {
		common.ApiErrorMsg(c, "保存素材记录失败: "+err.Error())
		return
	}

	common.ApiSuccess(c, gin.H{
		"id":              asset.Id,
		"asset_id":        result.AssetId,
		"remote_group_id": req.GroupId,
		"status":          asset.Status,
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
	pageInfo.SetTotal(int(total))
	pageInfo.SetItems(assets)
	common.ApiSuccess(c, pageInfo)
}

// PortraitGetAsset GET /v1/portrait/assets/:assetId
// 实时调用火山引擎刷新状态，并更新本地记录
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
			"created_at":   asset.CreatedAt,
			"updated_at":   asset.UpdatedAt,
		}
		if refreshErr != "" {
			h["refresh_error"] = refreshErr
		}
		return h
	}

	// 终态：Active（通过）、Failed（失败），不再调火山
	if asset.Status == "Active" || asset.Status == "Failed" {
		common.ApiSuccess(c, buildAssetResponse(""))
		return
	}

	// 非终态：实时查询火山引擎
	volcResult, err := portraitSvc.GetAsset(remoteAssetId)
	if err != nil {
		common.ApiSuccess(c, buildAssetResponse(err.Error()))
		return
	}

	if volcResult.Status != "" {
		asset.Status = volcResult.Status
	}
	if volcResult.ResolvedUrl != "" {
		asset.ResolvedUrl = volcResult.ResolvedUrl
	}
	_ = model.SavePortraitAsset(asset)

	common.ApiSuccess(c, buildAssetResponse(""))
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

	if _, err = portraitSvc.UpdateAsset(remoteAssetId, req.Name); err != nil {
		common.ApiErrorMsg(c, "调用火山引擎失败: "+err.Error())
		return
	}

	asset.Name = req.Name
	_ = model.SavePortraitAsset(asset)

	common.ApiSuccess(c, gin.H{
		"asset_id": remoteAssetId,
		"name":     asset.Name,
	})
}
