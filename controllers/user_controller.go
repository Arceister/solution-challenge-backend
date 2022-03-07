package controllers

import (
	"api-solution/services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	Get func(c *gin.Context)
}

func NewUserController(userService services.UserService) UserController {
	return UserController{
		Get: getUsersController(userService),
	}
}

func getUsersController(userService services.UserService) func(*gin.Context) {
	return func(c *gin.Context) {
		userService.CreateUser()
		c.JSON(200, gin.H{
			"message": "user get",
		})
	}
}
