package service

import (
	"tenet-profile/internal/model"
	repository "tenet-profile/internal/repositories"
)

type TenetSessionAllowAttributesService struct {
	repo *repository.SessionAllowAttributesRepository
}

func NewSessionAllowAttributesService(repo *repository.SessionAllowAttributesRepository) *TenetSessionAllowAttributesService {
	return &TenetSessionAllowAttributesService{
		repo: repo,
	}
}
func (s *TenetSessionAllowAttributesService) Save(sessionAllowAttributes *model.SessionAllowAttributes) error {
	return s.repo.Create(sessionAllowAttributes)
}

func (s *TenetSessionAllowAttributesService) Update(sessionAllowAttributes *model.SessionAllowAttributes, sessionId int64) error {
	return s.repo.Update(sessionAllowAttributes, sessionId)
}

func (s *TenetSessionAllowAttributesService) Delete(sessionAllowAttributes *model.SessionAllowAttributes) error {
	return s.repo.Delete(sessionAllowAttributes)
}

func (s *TenetSessionAllowAttributesService) GetByID(id int) (*model.SessionAllowAttributes, error) {
	return s.repo.GetByID(id)
}

func (s *TenetSessionAllowAttributesService) GetAll() ([]*model.SessionAllowAttributes, error) {
	return s.repo.GetAll()
}
