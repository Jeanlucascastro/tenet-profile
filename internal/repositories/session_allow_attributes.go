package repository

import (
	"tenet-profile/internal/model"

	"gorm.io/gorm"
)

type SessionAllowAttributesRepository struct {
	db *gorm.DB
}

func NewSessionAllowAttributesRepository(db *gorm.DB) *SessionAllowAttributesRepository {
	return &SessionAllowAttributesRepository{db: db}
}

func (r *SessionAllowAttributesRepository) Create(sessionAllowAttributes *model.SessionAllowAttributes) error {
	return r.db.Create(sessionAllowAttributes).Error
}

func (r *SessionAllowAttributesRepository) Update(sessionAllowAttributes *model.SessionAllowAttributes) error {
	return r.db.Save(sessionAllowAttributes).Error
}

func (r *SessionAllowAttributesRepository) Delete(sessionAllowAttributes *model.SessionAllowAttributes) error {
	return r.db.Delete(sessionAllowAttributes).Error
}
