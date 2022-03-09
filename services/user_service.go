package services

import (
	"api-solution/models"
	"api-solution/repository"
	"fmt"
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

func (s UserService) CreateUser() {
	fmt.Println("Create user service called.")
}
