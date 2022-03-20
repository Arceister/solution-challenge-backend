package services

import "api-solution/repository"

type DonateService struct {
	repository repository.DonateRepository
}

func NewDonateService(repository repository.DonateRepository) DonateService {
	return DonateService{
		repository: repository,
	}
}
