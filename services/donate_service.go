package services

import (
	"api-solution/models"
	"api-solution/repository"
)

type DonateService struct {
	repository repository.DonateRepository
}

func NewDonateService(repository repository.DonateRepository) DonateService {
	return DonateService{
		repository: repository,
	}
}

func (s DonateService) InsertDonate(user models.User, donate models.Donate) error {
	err := s.repository.Save(user, donate)
	return err
}
