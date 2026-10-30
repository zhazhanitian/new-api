package model

import (
	"time"

	"gorm.io/gorm"
)

type PortraitGroup struct {
	Id            int    `json:"id" gorm:"primaryKey;autoIncrement"`
	UserId        int    `json:"user_id" gorm:"index;not null"`
	RemoteGroupId string `json:"remote_group_id" gorm:"type:varchar(191);index;not null"`
	Name          string `json:"name" gorm:"type:varchar(255);not null"`
	// RemoteGroupId 是对外的逻辑 group_id，取主账号（volc）的上游 ID。
	// 各渠道账号里的上游 ID 见 PortraitGroupBinding。
	ProjectName string    `json:"project_name" gorm:"type:varchar(255)"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func GetPortraitGroupsByUserIdPaged(userId, startIdx, pageSize int) ([]*PortraitGroup, int64, error) {
	var groups []*PortraitGroup
	var total int64
	query := DB.Model(&PortraitGroup{}).Where("user_id = ?", userId)
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	err := query.Order("created_at desc").Offset(startIdx).Limit(pageSize).Find(&groups).Error
	return groups, total, err
}

func CreatePortraitGroup(group *PortraitGroup) error {
	return DB.Create(group).Error
}

// CreatePortraitGroupWithBindings 在同一事务里写入逻辑素材组和各渠道绑定。
func CreatePortraitGroupWithBindings(group *PortraitGroup, bindings []*PortraitGroupBinding) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(group).Error; err != nil {
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

func GetPortraitGroupByRemoteId(userId int, remoteGroupId string) (*PortraitGroup, error) {
	var group PortraitGroup
	err := DB.Where("user_id = ? AND remote_group_id = ?", userId, remoteGroupId).First(&group).Error
	return &group, err
}
