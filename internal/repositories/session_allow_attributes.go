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

func (r *SessionAllowAttributesRepository) Update(sessionAllowAttributes *model.SessionAllowAttributes, sessionId int64) error {
	return r.db.Model(sessionAllowAttributes).Where("session_id = ?", sessionId).Updates(sessionAllowAttributes).Error
}

func (r *SessionAllowAttributesRepository) Delete(sessionAllowAttributes *model.SessionAllowAttributes) error {
	return r.db.Delete(sessionAllowAttributes).Error
}

func (r *SessionAllowAttributesRepository) GetByID(id int) (*model.SessionAllowAttributes, error) {
	var sessionAllowAttributes model.SessionAllowAttributes
	return &sessionAllowAttributes, r.db.First(&sessionAllowAttributes, id).Error
}

func (r *SessionAllowAttributesRepository) GetAll() ([]*model.SessionAllowAttributes, error) {
	var sessionAllowAttributes []*model.SessionAllowAttributes
	return sessionAllowAttributes, r.db.Find(&sessionAllowAttributes).Error
}

func (r *SessionAllowAttributesRepository) FindBySessionIdAndUserWithThisAttribute(sessionID int64,
	userWithThisAttribute int64) (*model.SessionAllowAttributes, error) {

	var sessionAllowAttributes model.SessionAllowAttributes

	if err := r.db.Where("session_id = ? AND user_with_this_attribute = ?", sessionID,
		userWithThisAttribute).First(&sessionAllowAttributes).Error; err != nil {
		return nil, err
	}

	return &sessionAllowAttributes, nil
}
