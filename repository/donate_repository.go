package repository

import (
	"api-solution/lib"
	"api-solution/models"

	"gorm.io/gorm"
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

func (r DonateRepository) GetById(donateId uint) (donate models.Donate, err error) {
	return donate, r.db.DB.Where("id = ?", donateId).First(&donate).Error
}

func (r DonateRepository) TakeDonate(user models.User, donate models.Donate) error {
	r.db.DB.Model(&donate).Association("User")
	r.db.DB.Model(&user).Update("xp_points_pencari", gorm.Expr("xp_points_pencari + ?", 50))
	return r.db.DB.Model(&donate).Update("kuantitas", gorm.Expr("kuantitas - ?", 1)).Error
}

func (r DonateRepository) Save(user models.User, donate models.Donate) error {
	r.db.DB.Model(&donate).Association("User").Append(&user)
	r.db.DB.Model(&user).Update("xp_points", gorm.Expr("xp_points + ?", 100))
	return r.db.DB.Omit("User.*").Create(&donate).Error
}

func (r DonateRepository) Update(donate models.Donate) (models.Donate, error) {
	return donate, r.db.DB.Preload("User").Where("id = ?", donate.ID).Updates(&donate).Error
}

func (r DonateRepository) Delete(donateId uint, donate models.Donate) error {
	r.db.DB.Model(&donate).Association("User").Delete(&donate)
	return r.db.DB.Select("User").Delete(&donate).Error
}
