package services

import (
	"api-solution/models"
	"api-solution/repository"

	"github.com/jinzhu/copier"
)

type DonateService struct {
	repository repository.DonateRepository
}

func NewDonateService(repository repository.DonateRepository) DonateService {
	return DonateService{
		repository: repository,
	}
}

func (s DonateService) GetDonateAll() ([]models.Donate, error) {
	donates, err := s.repository.GetAll()
	return donates, err
}

func (s DonateService) GetDonateById(donateId uint) (models.Donate, error) {
	donate, err := s.repository.GetById(donateId)
	return donate, err
}

func (s DonateService) InsertDonate(user models.User, donate models.Donate) error {
	err := s.repository.Save(user, donate)
	return err
}

func (s DonateService) TakeDonation(
	quantity uint,
	user models.User,
	donate models.Donate,
	donatur models.User) (uint, error) {
	quantity, err := s.repository.TakeDonate(quantity, user, donate, donatur)
	return quantity, err
}

func (s DonateService) UpdateDonate(donateId uint, donate models.Donate) error {
	donateNew, err := s.GetDonateById(donateId)
	if err != nil {
		return err
	}

	copier.Copy(&donateNew, &donate)

	donateNew.ID = donateId

	_, err = s.repository.Update(donateNew)
	return err
}

func (s DonateService) DeleteDonate(donateId uint) error {
	donateNew, err := s.GetDonateById(donateId)
	if err != nil {
		return err
	}

	deleteErr := s.repository.Delete(donateNew)
	return deleteErr
}

func (s DonateService) GetUserByDonateId(donateId uint) (uint, error) {
	userDonationId, err := s.repository.GetUserDonationBased(donateId)
	return userDonationId, err
}
