package service

import (
	"tenet-profile/internal/model"
	repository "tenet-profile/internal/repositories"
)

type TenetProfileService struct {
	repo                       *repository.TenetProfileRepository
	sessionAllowAttributesRepo *repository.SessionAllowAttributesRepository
}

func NewTenetProfileService(repo *repository.TenetProfileRepository,
	sessionAllowAttributesRepo *repository.SessionAllowAttributesRepository) *TenetProfileService {
	return &TenetProfileService{
		repo:                       repo,
		sessionAllowAttributesRepo: sessionAllowAttributesRepo,
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

func (s *TenetProfileService) Update(profileDTO *model.ProfileDTO, profileId int64) (*model.Profile, error) {

	profile := profileDTO.ToEntity()

	updatedProfile, err := s.repo.UpdateTenetProfile(profile, profileId)
	if err != nil {
		return nil, err
	}

	return updatedProfile, nil
}

func (s *TenetProfileService) GetFiltered(sessionId int64, userIDParam int64) (map[string]interface{}, error) {

	profile, err := s.repo.GetTenetProfileByUserID(userIDParam)
	if err != nil {
		return nil, err
	}

	sessionAllowAttributes, err := s.sessionAllowAttributesRepo.FindBySessionIdAndUserWithThisAttribute(
		sessionId,
		userIDParam,
	)
	if err != nil {
		return nil, err
	}

	filteredProfile := profile.FilterByAttributes(sessionAllowAttributes.Attributes)

	return filteredProfile, nil

}
