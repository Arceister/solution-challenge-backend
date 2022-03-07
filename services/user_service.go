package services

import "fmt"

type UserService struct {
}

func NewUserService() UserService {
	return UserService{}
}

func (s UserService) CreateUser() {
	fmt.Println("Create user service called.")
}
