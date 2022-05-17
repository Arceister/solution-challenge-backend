package controllers

import (
	"api-solution/services"

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
