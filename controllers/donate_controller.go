package controllers

import (
	"api-solution/models"
	"api-solution/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DonateController struct {
	service     services.DonateService
	jwtService  services.JWTAuthService
	userService services.UserService
}

func NewDonateController(
	service services.DonateService,
	jwtService services.JWTAuthService,
	userService services.UserService,
) DonateController {
	return DonateController{
		service:     service,
		jwtService:  jwtService,
		userService: userService,
	}
}

func (d DonateController) InsertDonate(c *gin.Context) {
	donate := models.Donate{}
	header := c.Request.Header.Get("Authorization")
	header = header[len("Bearer "):]

	if err := c.ShouldBindJSON(&donate); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	valid, err := d.jwtService.Authorize(header)
	if err != nil && !valid {
		c.JSON(http.StatusForbidden, gin.H{
			"success": false,
			"message": "JWT Error",
			"error":   err.Error(),
		})
		c.Abort()
	}

	claims := d.jwtService.ExtractClaims(header)
	userId := claims["id"].(float64)

	user, _ := d.userService.GetUserById(uint(userId))

	if err := d.service.InsertDonate(user, donate); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Insert New Donation Success!",
	})
}
