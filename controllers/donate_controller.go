package controllers

import (
	"api-solution/models"
	"api-solution/services"
	"net/http"
	"strconv"

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

func (d DonateController) GetAllDonate(c *gin.Context) {
	users, err := d.service.GetDonateAll()
	if err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"error":   err.Error(),
		})
	}
	c.JSON(200, gin.H{"data": users})
}

func (d DonateController) GetDonateById(c *gin.Context) {
	idParam, _ := strconv.Atoi(c.Param("id"))
	users, err := d.service.GetDonateById(uint(idParam))
	if err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"error":   err.Error(),
		})
	}
	c.JSON(200, gin.H{"data": users})
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

func (d DonateController) UpdateDonate(c *gin.Context) {
	donate := models.Donate{}
	idParam, _ := strconv.Atoi(c.Param("id"))

	if err := c.ShouldBindJSON(&donate); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := d.service.UpdateDonate(uint(idParam), donate); err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Update Success!",
	})
}

func (d DonateController) DeleteDonate(c *gin.Context) {
	idParam, _ := strconv.Atoi(c.Param("id"))

	if err := d.service.DeleteDonate(uint(idParam)); err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Delete Donation!",
	})
}
