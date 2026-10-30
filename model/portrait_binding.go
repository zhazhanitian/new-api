package model

import (
	"strings"
	"time"
)

// 人像素材状态。除 Active、Failed 外都视为非终态。
const (
	PortraitStatusSubmitted  = "Submitted"
	PortraitStatusProcessing = "Processing"
	PortraitStatusActive     = "Active"
	PortraitStatusFailed     = "Failed"
)

// PortraitGroupBinding 记录一个逻辑素材组在某个人像库账号里的上游素材组 ID。
// 逻辑 group_id 取主账号（volc）的上游 ID，因此线上历史素材组没有绑定记录。
type PortraitGroupBinding struct {
	Id              int       `json:"id" gorm:"primaryKey;autoIncrement"`
	UserId          int       `json:"user_id" gorm:"index;not null"`
	GroupId         string    `json:"group_id" gorm:"type:varchar(191);not null;uniqueIndex:idx_pgb_group_channel"`
	PortraitChannel string    `json:"portrait_channel" gorm:"type:varchar(32);not null;uniqueIndex:idx_pgb_group_channel"`
	UpstreamGroupId string    `json:"upstream_group_id" gorm:"type:varchar(191);not null"`
	CreatedAt       time.Time `json:"created_at"`
}

// PortraitAssetBinding 记录一个逻辑素材在某个人像库账号里的上游素材 ID 和审核状态。
type PortraitAssetBinding struct {
	Id              int       `json:"id" gorm:"primaryKey;autoIncrement"`
	UserId          int       `json:"user_id" gorm:"index;not null"`
	AssetId         string    `json:"asset_id" gorm:"type:varchar(191);not null;uniqueIndex:idx_pab_asset_channel"`
	PortraitChannel string    `json:"portrait_channel" gorm:"type:varchar(32);not null;uniqueIndex:idx_pab_asset_channel"`
	UpstreamAssetId string    `json:"upstream_asset_id" gorm:"type:varchar(191);not null"`
	Status          string    `json:"status" gorm:"type:varchar(50);index"`
	ResolvedUrl     string    `json:"resolved_url" gorm:"type:text"`
	LastError       string    `json:"last_error" gorm:"type:text"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// AggregatePortraitStatus 汇总各渠道状态得到逻辑素材状态（严格模式）：
// 任一 Failed 则 Failed；全部 Active 才 Active；其余为非终态，
// 全部还是 Submitted 时保持 Submitted，否则 Processing。没有任何绑定时返回空串。
func AggregatePortraitStatus(statuses []string) string {
	if len(statuses) == 0 {
		return ""
	}
	allActive := true
	allSubmitted := true
	for _, s := range statuses {
		switch s {
		case PortraitStatusFailed:
			return PortraitStatusFailed
		case PortraitStatusActive:
			allSubmitted = false
		case PortraitStatusSubmitted, "":
			allActive = false
		default:
			allActive = false
			allSubmitted = false
		}
	}
	switch {
	case allActive:
		return PortraitStatusActive
	case allSubmitted:
		return PortraitStatusSubmitted
	default:
		return PortraitStatusProcessing
	}
}

func CreatePortraitGroupBinding(b *PortraitGroupBinding) error {
	return DB.Create(b).Error
}

func GetPortraitGroupBindings(groupId string) ([]*PortraitGroupBinding, error) {
	var list []*PortraitGroupBinding
	err := DB.Where("group_id = ?", groupId).Order("id asc").Find(&list).Error
	return list, err
}

func CreatePortraitAssetBinding(b *PortraitAssetBinding) error {
	return DB.Create(b).Error
}

func SavePortraitAssetBinding(b *PortraitAssetBinding) error {
	return DB.Save(b).Error
}

func GetPortraitAssetBindings(assetId string) ([]*PortraitAssetBinding, error) {
	var list []*PortraitAssetBinding
	err := DB.Where("asset_id = ?", assetId).Order("id asc").Find(&list).Error
	return list, err
}

// GetPendingPortraitAssetBindings 返回所有非终态绑定，供后台轮询。
func GetPendingPortraitAssetBindings() ([]*PortraitAssetBinding, error) {
	var list []*PortraitAssetBinding
	err := DB.Where("status NOT IN ?", []string{PortraitStatusActive, PortraitStatusFailed}).
		Order("updated_at asc").
		Find(&list).Error
	return list, err
}

// GetLegacyPendingPortraitAssets 返回没有任何绑定记录的非终态素材（线上历史素材，仅火山）。
func GetLegacyPendingPortraitAssets() ([]*PortraitAsset, error) {
	var assets []*PortraitAsset
	err := DB.Where("status NOT IN ?", []string{PortraitStatusActive, PortraitStatusFailed}).
		Where("remote_asset_id NOT IN (SELECT asset_id FROM portrait_asset_bindings)").
		Order("updated_at asc").
		Find(&assets).Error
	return assets, err
}

// GetPortraitAssetsMissingChannel 返回缺少指定渠道绑定的素材（含没有任何绑定的历史素材），用于补建。
func GetPortraitAssetsMissingChannel(channel string, afterId, limit int) ([]*PortraitAsset, error) {
	var assets []*PortraitAsset
	err := DB.Where("id > ?", afterId).
		Where("remote_asset_id NOT IN (SELECT asset_id FROM portrait_asset_bindings WHERE portrait_channel = ?)", strings.ToLower(channel)).
		Order("id asc").
		Limit(limit).
		Find(&assets).Error
	return assets, err
}
