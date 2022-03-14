package services

import (
	"api-solution/models"
	"api-solution/repository"
	"fmt"

	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) UserService {
	return UserService{
		repository: repository,
	}
}

func (s UserService) GetAllUser() ([]models.User, error) {
	users, err := s.repository.GetAll()
	return users, err
}

func (s UserService) GetUserById(userId uint) (models.User, error) {
	users, err := s.repository.GetUserById(userId)
	return users, err
}

func (s UserService) GetUserByEmail(email string) (models.User, error) {
	user, err := s.repository.GetEmail(email)
	return user, err
}

func (s UserService) InsertUser(user models.User) (models.User, error) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	user.Password = string(hashedPassword)
	users, err := s.repository.Save(user)
	return users, err
}

func (s UserService) UpdateUser(id uint, user models.User) error {
	userNew, err := s.GetUserById(id)
	if err != nil {
		return err
	}

	copier.Copy(&userNew, &user)

	userNew.ID = id

	_, err = s.repository.Update(userNew)
	return err
}

func (s UserService) DeleteUser(userId uint) error {
	_, err := s.repository.Delete(userId)
	return err
}

func (s UserService) CreateUser() {
	fmt.Println("Create user service called.")
}
