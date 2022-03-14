package controllers

import (
	"api-solution/models"
	"api-solution/services"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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
	payload := models.UserAuthentication{}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err := j.userService.GetUserByEmail(payload.Email)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	match := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))
	if match != nil {
		c.JSON(500, gin.H{
			"error": "Password not match!",
		})
	} else {
		token := j.service.CreateToken(user)
		c.JSON(200, gin.H{
			"message": "login success",
			"token":   token,
		})
	}
}
