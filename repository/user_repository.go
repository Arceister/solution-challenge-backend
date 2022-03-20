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

func (r UserRepository) GetEmail(email string) (users models.User, err error) {
	return users, r.db.DB.Where("email = ?", email).First(&users).Error
}

func (r UserRepository) Save(user models.User) (models.User, error) {
	return user, r.db.DB.Preload("Donate").Create(&user).Error
}

func (r UserRepository) Update(user models.User) (models.User, error) {
	return user, r.db.DB.Preload("Donate").Where("id = ?", &user.ID).Updates(&user).Error
}

func (r UserRepository) Delete(userId uint) (users models.User, err error) {
	return users, r.db.DB.Where("id = ?", userId).Delete(&users).Error
}
