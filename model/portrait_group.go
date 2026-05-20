package model

import "time"

type PortraitGroup struct {
	Id            int       `json:"id" gorm:"primaryKey;autoIncrement"`
	UserId        int       `json:"user_id" gorm:"index;not null"`
	RemoteGroupId string    `json:"remote_group_id" gorm:"type:varchar(191);index;not null"`
	Name          string    `json:"name" gorm:"type:varchar(255);not null"`
	ProjectName   string    `json:"project_name" gorm:"type:varchar(255)"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func GetPortraitGroupsByUserId(userId int) ([]*PortraitGroup, error) {
	var groups []*PortraitGroup
	err := DB.Where("user_id = ?", userId).Order("created_at desc").Find(&groups).Error
	return groups, err
}

func CreatePortraitGroup(group *PortraitGroup) error {
	return DB.Create(group).Error
}

func GetPortraitGroupByRemoteId(userId int, remoteGroupId string) (*PortraitGroup, error) {
	var group PortraitGroup
	err := DB.Where("user_id = ? AND remote_group_id = ?", userId, remoteGroupId).First(&group).Error
	return &group, err
}
