package controllers

import (
	"api-solution/services"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service services.UserService
}

func NewUserController(userService services.UserService) UserController {
	return UserController{
		service: userService,
	}
}

func (u UserController) GetUser(c *gin.Context) {
	users, err := u.service.GetAllUser()
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(200, gin.H{"data": users})
}

func (u UserController) GetUserById(c *gin.Context) {
	idParam, _ := strconv.Atoi(c.Param("id"))
	users, err := u.service.GetUserById(uint(idParam))
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(200, gin.H{"data": users})
}
