package service

import (
	"errors"
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"tenet-profile/internal/model"
	repository "tenet-profile/internal/repositories"

	"github.com/google/uuid"
)

type profileImageService interface {
	UpdatePicture(profileID uint64, file *multipart.FileHeader) (string, error)
}

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

func (s *TenetProfileService) UpdatePicture(profileID uint64, file *multipart.FileHeader) (string, error) {

	if file.Size > 5<<20 {
		return "", errors.New("Image is too large")
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	if ext != ".png" && ext != ".jpg" && ext != ".jpeg" && ext != ".JPEG" {
		return "", errors.New("format invalid")
	}

	dir := "./uploads/profiles"
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return "", err
	}

	filename := fmt.Sprintf("%s%s", uuid.NewString(), ext)
	path := filepath.Join(dir, filename)

	if err := saveFile(file, path); err != nil {
		return "", err
	}

	url := fmt.Sprintf("/static/profiles/%s", filename)

	if err := s.repo.UpdatePicture(profileID, url); err != nil {
		return "", err
	}

	return url, nil

}

func saveFile(file *multipart.FileHeader, path string) error {

	src, err := file.Open()
	if err != nil {
		return err
	}

	defer src.Close()

	dst, err := os.Create(path)
	if err != nil {
		return err
	}

	defer dst.Close()

	_, err = dst.ReadFrom(src)
	return err

}
