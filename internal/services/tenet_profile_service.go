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

func (s *TenetProfileService) Save(profileDTO *model.ProfileDTO) error {

	profile := profileDTO.ToEntity()

	return s.repo.CreateTenetProfile(profile)
}
