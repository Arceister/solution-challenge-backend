package repository

import (
	"api-solution/lib"
	"api-solution/models"
)

type UserRepository struct {
	db lib.Database
}

func NewUserRepository(db lib.Database) UserRepository {
	return UserRepository{
		db: db,
	}
}

func (r UserRepository) GetAll() (users []models.User, err error) {
	return users, r.db.DB.Find(&users).Error
}

func (r UserRepository) GetUserById(userId uint) (users models.User, err error) {
	return users, r.db.DB.Where("id = ?", userId).First(&users).Error
}

func (r UserRepository) Save(user models.User) (models.User, error) {
	return user, r.db.DB.Create(&user).Error
}

func (r UserRepository) Update(user models.User) (models.User, error) {
	return user, r.db.DB.Updates(&user).Error
}
