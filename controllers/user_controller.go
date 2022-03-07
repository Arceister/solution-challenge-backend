package controllers

import "github.com/gin-gonic/gin"

type UserController struct {
	Get func(c *gin.Context)
}

func NewUserController() UserController {
	return UserController{
		Get: getUsersController(),
	}
}

func getUsersController() func(*gin.Context) {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "user get",
		})
	}
}
