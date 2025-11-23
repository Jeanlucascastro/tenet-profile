package repository

import (
	"tenet-profile/internal/model"

	"gorm.io/gorm"
)

type TenetProfileRepository struct {
	DB *gorm.DB
}

func NewTenetProfileRepository(db *gorm.DB) *TenetProfileRepository {
	return &TenetProfileRepository{
		DB: db,
	}
}

func (r *TenetProfileRepository) GetTenetProfileByID(id string) (*model.Profile, error) {
	var profile model.Profile
	if err := r.DB.First(&profile, id).Error; err != nil {
		return nil, err
	}
	return &profile, nil
}

func (r *TenetProfileRepository) GetTenetProfileByUserID(userID int64) (*model.Profile, error) {
	var profile model.Profile
	if err := r.DB.Where("user_id = ?", userID).First(&profile).Error; err != nil {
		return nil, err
	}
	return &profile, nil
}

func (r *TenetProfileRepository) CreateTenetProfile(profile *model.Profile) error {
	return r.DB.Create(profile).Error
}

func (r *TenetProfileRepository) UpdateTenetProfile(profile *model.Profile) error {
	return r.DB.Save(profile).Error
}
