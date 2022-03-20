package repository

import (
	"api-solution/lib"
	"api-solution/models"
)

type DonateRepository struct {
	db lib.Database
}

func NewDonateRepository(db lib.Database) DonateRepository {
	return DonateRepository{
		db: db,
	}
}

func (r DonateRepository) GetAll() (donates []models.Donate, err error) {
	return donates, r.db.DB.Order("created_at DESC").Find(&donates).Error
}

func (r DonateRepository) Save(donate models.Donate) (models.Donate, error) {
	return donate, r.db.DB.Create(&donate).Error
}
