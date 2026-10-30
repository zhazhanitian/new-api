package model

import (
	"time"

	"gorm.io/gorm"
)

type PortraitAsset struct {
	Id            int    `json:"id" gorm:"primaryKey;autoIncrement"`
	UserId        int    `json:"user_id" gorm:"index;not null"`
	RemoteGroupId string `json:"remote_group_id" gorm:"type:varchar(191);index"`
	RemoteAssetId string `json:"remote_asset_id" gorm:"type:varchar(191);index;not null"`
	// RemoteAssetId 是对外的逻辑 asset_id，取主账号（volc）的上游 ID。
	// 各渠道账号里的上游 ID 和状态见 PortraitAssetBinding。
	Name      string `json:"name" gorm:"type:varchar(255)"`
	AssetType string `json:"asset_type" gorm:"type:varchar(20);not null"` // Image / Video / Audio
	SourceUrl string `json:"source_url" gorm:"type:text;not null"`
	// Status 是汇总状态：有绑定时由所有渠道绑定汇总得出（全部 Active 才 Active），
	// 没有绑定的线上历史素材沿用火山的状态。
	Status      string    `json:"status" gorm:"type:varchar(50)"` // Submitted / Processing / Active / Failed
	ResolvedUrl string    `json:"resolved_url" gorm:"type:text"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func GetPortraitAssetsByUserIdPaged(userId int, groupId string, startIdx, pageSize int) ([]*PortraitAsset, int64, error) {
	var assets []*PortraitAsset
	var total int64
	query := DB.Model(&PortraitAsset{}).Where("user_id = ?", userId)
	if groupId != "" {
		query = query.Where("remote_group_id = ?", groupId)
	}
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	err := query.Order("created_at desc").Offset(startIdx).Limit(pageSize).Find(&assets).Error
	return assets, total, err
}

func GetPortraitAssetByRemoteId(userId int, remoteAssetId string) (*PortraitAsset, error) {
	var asset PortraitAsset
	err := DB.Where("user_id = ? AND remote_asset_id = ?", userId, remoteAssetId).First(&asset).Error
	return &asset, err
}

func CreatePortraitAsset(asset *PortraitAsset) error {
	return DB.Create(asset).Error
}

func SavePortraitAsset(asset *PortraitAsset) error {
	return DB.Save(asset).Error
}

// CreatePortraitAssetWithBindings 在同一事务里写入逻辑素材和各渠道绑定。
func CreatePortraitAssetWithBindings(asset *PortraitAsset, bindings []*PortraitAssetBinding) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(asset).Error; err != nil {
			return err
		}
		for _, b := range bindings {
			if err := tx.Create(b).Error; err != nil {
				return err
			}
		}
		return nil
	})
}
