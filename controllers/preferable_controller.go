package controllers

import (
	"api-solution/models"
	"api-solution/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PreferableController struct {
	service    services.PreferableService
	jwtService services.JWTAuthService
}

func NewPreferableController(
	service services.PreferableService,
	jwtService services.JWTAuthService,
) PreferableController {
	return PreferableController{
		service:    service,
		jwtService: jwtService,
	}
}

func (p PreferableController) GetAllPreferables(c *gin.Context) {
	preferables, err := p.service.GetAllPreferables()
	if err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"error":   err.Error(),
		})
	}
	c.JSON(200, gin.H{
		"data": preferables,
	})
}

func (p PreferableController) InsertPreferable(c *gin.Context) {
	jsonPreferable := models.JsonReciever{}
	header := c.Request.Header.Get("Authorization")
	header = header[len("Bearer "):]

	if err := c.ShouldBindJSON(&jsonPreferable); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	valid, err := p.jwtService.Authorize(header)
	if err != nil && !valid {
		c.JSON(http.StatusForbidden, gin.H{
			"success": false,
			"message": "JWT Error",
			"error":   err.Error(),
		})
		c.Abort()
	}

	claims := p.jwtService.ExtractClaims(header)
	userId := claims["id"].(float64)

	userPreferable := models.UserPreferable{
		User_ID:     uint(userId),
		Preferable:  jsonPreferable.Preferable,
		Prohibition: jsonPreferable.Prohibition,
	}

	queryErr := p.service.InsertPreferables(userPreferable)

	if queryErr != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Inserted preferables!",
	})
}
