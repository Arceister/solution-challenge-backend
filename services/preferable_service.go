package services

import (
	"api-solution/models"
	"api-solution/repository"
)

type PreferableService struct {
	repository repository.PreferableRepository
}

func NewPreferableService(repository repository.PreferableRepository) PreferableService {
	return PreferableService{
		repository: repository,
	}
}

func (s PreferableService) GetAllPreferables() ([]models.UserPreferable, error) {
	preferables, err := s.repository.GetAll()
	return preferables, err
}

func (s PreferableService) InsertPreferables(preferable models.UserPreferable) error {
	err := s.repository.InsertOne(preferable)
	return err
}
