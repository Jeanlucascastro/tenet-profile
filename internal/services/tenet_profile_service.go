package service

import (
	"tenet-profile/internal/model"
	repository "tenet-profile/internal/repositories"
)

type TenetProfileService struct {
	repo *repository.TenetProfileRepository
}

func NewTenetProfileService(repo *repository.TenetProfileRepository) *TenetProfileService {
	return &TenetProfileService{
		repo: repo,
	}
}

func (s *TenetProfileService) Save(profileDTO *model.ProfileDTO) (*model.Profile, error) {

	profile := profileDTO.ToEntity()

	return s.repo.CreateTenetProfile(profile)
}

func (s *TenetProfileService) GetAllByID(userIDParam int64) ([]model.Profile, error) {

	profile, err := s.repo.FindAllByUserID(userIDParam)
	if err != nil {
		return nil, err
	}

	return profile, nil
}
