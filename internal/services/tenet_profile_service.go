package service

import (
	"strconv"
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

func (s *TenetProfileService) GetFiltered(sessionId int64, userIDParam int64) (map[string]interface{}, error) {

	userId := strconv.FormatInt(userIDParam, 10)

	profile, err := s.repo.GetTenetProfileByID(userId)
	if err != nil {
		return nil, err
	}

	//TODO: return null if fields are not allowed
	sessionAllowAttributes, err := s.repo.FindBySessionIdAndUserWithThisAttribute(
		sessionId,
		userIDParam,
	)
	if err != nil {
		return nil, err
	}

	filteredProfile := profile.FilterByAttributes(sessionAllowAttributes.Attributes)

	return filteredProfile, nil

}
