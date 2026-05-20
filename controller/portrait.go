package controller

import (
	"encoding/json"
	"net/http"
	"strings"

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

// PortraitCreateAssetRequest 创建素材请求体
type PortraitCreateAssetRequest struct {
	GroupId   string `json:"group_id"`
	URL       string `json:"url"`
	AssetType string `json:"asset_type"`
	Name      string `json:"name"`
}

func portraitError(c *gin.Context, status int, msg string) {
	c.JSON(status, gin.H{"success": false, "message": msg})
}

// PortraitCreateGroup POST /v1/portrait/groups
func PortraitCreateGroup(c *gin.Context) {
	userId := c.GetInt("id")
	if userId == 0 {
		portraitError(c, http.StatusUnauthorized, "未授权")
		return
	}

	var req PortraitCreateGroupRequest
	if err := json.NewDecoder(c.Request.Body).Decode(&req); err != nil {
		portraitError(c, http.StatusBadRequest, "请求体解析失败")
		return
	}
	req.Name = strings.TrimSpace(req.Name)
	if req.Name == "" {
		portraitError(c, http.StatusBadRequest, "素材组名称不能为空")
		return
	}
	if req.GroupType == "" {
		req.GroupType = "AIGC"
	}

	result, err := portraitSvc.CreateAssetGroup(req.Name, req.Description, req.GroupType)
	if err != nil {
		portraitError(c, http.StatusBadGateway, "调用火山引擎失败: "+err.Error())
		return
	}

	group := &model.PortraitGroup{
		UserId:        userId,
		RemoteGroupId: result.Id,
		Name:          req.Name,
	}
	if err = model.CreatePortraitGroup(group); err != nil {
		portraitError(c, http.StatusInternalServerError, "保存素材组记录失败: "+err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":  true,
		"group_id": result.Id,
		"name":     req.Name,
		"id":       group.Id,
	})
}

// PortraitListGroups GET /v1/portrait/groups
func PortraitListGroups(c *gin.Context) {
	userId := c.GetInt("id")
	if userId == 0 {
		portraitError(c, http.StatusUnauthorized, "未授权")
		return
	}

	groups, err := model.GetPortraitGroupsByUserId(userId)
	if err != nil {
		portraitError(c, http.StatusInternalServerError, "查询素材组失败: "+err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    groups,
	})
}

// PortraitCreateAsset POST /v1/portrait/assets
func PortraitCreateAsset(c *gin.Context) {
	userId := c.GetInt("id")
	if userId == 0 {
		portraitError(c, http.StatusUnauthorized, "未授权")
		return
	}

	var req PortraitCreateAssetRequest
	if err := json.NewDecoder(c.Request.Body).Decode(&req); err != nil {
		portraitError(c, http.StatusBadRequest, "请求体解析失败")
		return
	}

	req.GroupId = strings.TrimSpace(req.GroupId)
	req.URL = strings.TrimSpace(req.URL)
	req.AssetType = strings.TrimSpace(req.AssetType)

	if req.GroupId == "" {
		portraitError(c, http.StatusBadRequest, "group_id 不能为空")
		return
	}
	if req.URL == "" {
		portraitError(c, http.StatusBadRequest, "url 不能为空")
		return
	}
	validTypes := map[string]bool{"Image": true, "Video": true, "Audio": true}
	if !validTypes[req.AssetType] {
		portraitError(c, http.StatusBadRequest, "asset_type 须为 Image / Video / Audio")
		return
	}

	result, err := portraitSvc.CreateAsset(req.GroupId, req.URL, req.AssetType, req.Name)
	if err != nil {
		portraitError(c, http.StatusBadGateway, "调用火山引擎失败: "+err.Error())
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
		portraitError(c, http.StatusInternalServerError, "保存素材记录失败: "+err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":        true,
		"id":             asset.Id,
		"asset_id":       result.AssetId,
		"remote_group_id": req.GroupId,
		"status":         asset.Status,
	})
}

// PortraitListAssets GET /v1/portrait/assets
func PortraitListAssets(c *gin.Context) {
	userId := c.GetInt("id")
	if userId == 0 {
		portraitError(c, http.StatusUnauthorized, "未授权")
		return
	}

	groupId := c.Query("group_id")
	assets, err := model.GetPortraitAssetsByUserId(userId, groupId)
	if err != nil {
		portraitError(c, http.StatusInternalServerError, "查询素材列表失败: "+err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    assets,
	})
}

// PortraitGetAsset GET /v1/portrait/assets/:assetId
// 实时调用火山引擎刷新状态，并更新本地记录
func PortraitGetAsset(c *gin.Context) {
	userId := c.GetInt("id")
	if userId == 0 {
		portraitError(c, http.StatusUnauthorized, "未授权")
		return
	}

	remoteAssetId := c.Param("assetId")
	if remoteAssetId == "" {
		portraitError(c, http.StatusBadRequest, "assetId 不能为空")
		return
	}

	asset, err := model.GetPortraitAssetByRemoteId(userId, remoteAssetId)
	if err != nil {
		portraitError(c, http.StatusNotFound, "素材不存在")
		return
	}

	// 如果已是终态，直接返回本地数据，不再调火山引擎
	if asset.Status == "Approved" || asset.Status == "Rejected" || asset.Status == "Failed" {
		c.JSON(http.StatusOK, gin.H{
			"success":      true,
			"asset_id":     asset.RemoteAssetId,
			"status":       asset.Status,
			"resolved_url": asset.ResolvedUrl,
			"name":         asset.Name,
			"asset_type":   asset.AssetType,
			"source_url":   asset.SourceUrl,
			"created_at":   asset.CreatedAt,
			"updated_at":   asset.UpdatedAt,
		})
		return
	}

	// 非终态：实时查询火山引擎
	volcResult, err := portraitSvc.GetAsset(remoteAssetId)
	if err != nil {
		// 火山引擎调用失败仍返回本地缓存状态，并附带错误信息
		c.JSON(http.StatusOK, gin.H{
			"success":        true,
			"asset_id":       asset.RemoteAssetId,
			"status":         asset.Status,
			"resolved_url":   asset.ResolvedUrl,
			"name":           asset.Name,
			"asset_type":     asset.AssetType,
			"source_url":     asset.SourceUrl,
			"refresh_error":  err.Error(),
			"created_at":     asset.CreatedAt,
			"updated_at":     asset.UpdatedAt,
		})
		return
	}

	if volcResult.Status != "" {
		asset.Status = volcResult.Status
	}
	if volcResult.ResolvedUrl != "" {
		asset.ResolvedUrl = volcResult.ResolvedUrl
	}
	_ = model.SavePortraitAsset(asset)

	c.JSON(http.StatusOK, gin.H{
		"success":      true,
		"asset_id":     asset.RemoteAssetId,
		"status":       asset.Status,
		"resolved_url": asset.ResolvedUrl,
		"name":         asset.Name,
		"asset_type":   asset.AssetType,
		"source_url":   asset.SourceUrl,
		"created_at":   asset.CreatedAt,
		"updated_at":   asset.UpdatedAt,
	})
}
