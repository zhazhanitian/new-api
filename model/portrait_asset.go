package model

import "time"

type PortraitAsset struct {
	Id            int       `json:"id" gorm:"primaryKey;autoIncrement"`
	UserId        int       `json:"user_id" gorm:"index;not null"`
	RemoteGroupId string    `json:"remote_group_id" gorm:"type:varchar(191);index"`
	RemoteAssetId string    `json:"remote_asset_id" gorm:"type:varchar(191);index;not null"`
	Name          string    `json:"name" gorm:"type:varchar(255)"`
	AssetType     string    `json:"asset_type" gorm:"type:varchar(20);not null"` // Image / Video / Audio
	SourceUrl     string    `json:"source_url" gorm:"type:text;not null"`
	Status        string    `json:"status" gorm:"type:varchar(50)"` // Submitted / Processing / Active / Failed
	ResolvedUrl   string    `json:"resolved_url" gorm:"type:text"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
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

// GetPendingPortraitAssets 返回所有状态不是终态的素材（用于后台轮询）
// 终态：Active（审核通过）、Failed（失败）
func GetPendingPortraitAssets() ([]*PortraitAsset, error) {
	var assets []*PortraitAsset
	err := DB.Where("status NOT IN ?", []string{"Active", "Failed"}).
		Order("updated_at asc").
		Find(&assets).Error
	return assets, err
}
