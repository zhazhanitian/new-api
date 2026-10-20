package model

import (
	"errors"

	"github.com/QuantumNous/new-api/common"

	"gorm.io/gorm"
)

type ModelTrial struct {
	Id          int    `json:"id"`
	ModelName   string `json:"model_name" gorm:"size:128;not null;uniqueIndex"`
	APIKind     string `json:"api_kind" gorm:"size:32;not null"`
	Enabled     bool   `json:"enabled" gorm:"default:true"`
	Nodes       string `json:"nodes" gorm:"type:text"`
	Examples    string `json:"examples" gorm:"type:text"`
	CreatedTime int64  `json:"created_time" gorm:"bigint"`
	UpdatedTime int64  `json:"updated_time" gorm:"bigint"`
}

func (m *ModelTrial) Insert() error {
	now := common.GetTimestamp()
	m.CreatedTime = now
	m.UpdatedTime = now
	if m.Nodes == "" {
		m.Nodes = "[]"
	}
	if m.Examples == "" {
		m.Examples = "[]"
	}
	enabled := m.Enabled
	if err := DB.Create(m).Error; err != nil {
		return err
	}
	return DB.Model(&ModelTrial{}).Where("id = ?", m.Id).Update("enabled", enabled).Error
}

func (m *ModelTrial) Update() error {
	m.UpdatedTime = common.GetTimestamp()
	return DB.Model(&ModelTrial{}).Where("id = ?", m.Id).Select(
		"model_name", "api_kind", "enabled", "nodes", "examples", "updated_time",
	).Updates(m).Error
}

func GetModelTrialByID(id int) (*ModelTrial, error) {
	var item ModelTrial
	if err := DB.First(&item, id).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func GetModelTrialByName(name string) (*ModelTrial, error) {
	var item ModelTrial
	err := DB.Where("model_name = ?", name).First(&item).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func ListModelTrials() ([]*ModelTrial, error) {
	var items []*ModelTrial
	err := DB.Order("id desc").Find(&items).Error
	return items, err
}

func ListEnabledModelTrials() ([]*ModelTrial, error) {
	var items []*ModelTrial
	err := DB.Where("enabled = ?", true).Order("id desc").Find(&items).Error
	return items, err
}

func DeleteModelTrial(id int) error {
	return DB.Delete(&ModelTrial{}, id).Error
}

type ModelTrialProof struct {
	Id          int    `json:"id"`
	UserId      int    `json:"user_id" gorm:"index:idx_trial_proof,priority:1"`
	ModelName   string `json:"model_name" gorm:"size:128;index:idx_trial_proof,priority:2"`
	ValuesHash  string `json:"values_hash" gorm:"size:64;index:idx_trial_proof,priority:3"`
	TaskID      string `json:"task_id" gorm:"size:191;index"`
	Verified    bool   `json:"verified"`
	CreatedTime int64  `json:"created_time" gorm:"bigint"`
}

func UpsertPendingTrialProof(userId int, modelName, hash string) (*ModelTrialProof, error) {
	now := common.GetTimestamp()
	_ = DB.Where("user_id = ? AND model_name = ? AND values_hash = ? AND verified = ?", userId, modelName, hash, false).
		Delete(&ModelTrialProof{}).Error
	item := &ModelTrialProof{
		UserId:      userId,
		ModelName:   modelName,
		ValuesHash:  hash,
		CreatedTime: now,
	}
	if err := DB.Create(item).Error; err != nil {
		return nil, err
	}
	return item, nil
}

func GetTrialProof(id, userId int) (*ModelTrialProof, error) {
	var item ModelTrialProof
	err := DB.Where("id = ? AND user_id = ?", id, userId).First(&item).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func MarkTrialProofVerified(id int) error {
	return DB.Model(&ModelTrialProof{}).Where("id = ?", id).Update("verified", true).Error
}

func AttachTrialProofTask(id int, taskID string) error {
	return DB.Model(&ModelTrialProof{}).Where("id = ?", id).Update("task_id", taskID).Error
}

func HasVerifiedTrialProof(userId int, modelName, hash string) (bool, error) {
	var count int64
	err := DB.Model(&ModelTrialProof{}).
		Where("user_id = ? AND model_name = ? AND values_hash = ? AND verified = ?", userId, modelName, hash, true).
		Count(&count).Error
	return count > 0, err
}
