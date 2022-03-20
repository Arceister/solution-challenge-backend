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

func (r DonateRepository) Save(user models.User, donate models.Donate) error {
	r.db.DB.Model(&donate).Association("User").Append([]models.User{user})
	return r.db.DB.Preload("User").Create(&donate).Error
}
