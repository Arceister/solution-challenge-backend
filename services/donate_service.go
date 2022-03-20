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

func (s DonateService) GetDonateById(donateId uint) (models.Donate, error) {
	donate, err := s.repository.GetById(donateId)
	return donate, err
}

func (s DonateService) InsertDonate(user models.User, donate models.Donate) error {
	err := s.repository.Save(user, donate)
	return err
}
