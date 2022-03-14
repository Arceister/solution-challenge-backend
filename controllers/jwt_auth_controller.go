package controllers

import (
	"api-solution/services"

	"github.com/gin-gonic/gin"
)

type JWTAuthController struct {
	service     services.JWTAuthService
	userService services.UserService
}

func NewJWTAuthController(
	service services.JWTAuthService,
	userService services.UserService,
) JWTAuthController {
	return JWTAuthController{
		service:     service,
		userService: userService,
	}
}

func (j JWTAuthController) SignIn(c *gin.Context) {
	user, _ := j.userService.GetUserById(uint(1))
	token := j.service.CreateToken(user)
	c.JSON(200, gin.H{
		"message": "logged in",
		"token":   token,
	})
}

func (j JWTAuthController) Register(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "register route",
	})
}
