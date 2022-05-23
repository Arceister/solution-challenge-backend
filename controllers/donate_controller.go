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

func successResponse(
	c *gin.Context,
	code int,
	message string,
	data interface{},
) (int, interface{}) {
	response := gin.H{
		"success": true,
		"message": message,
		"data":    data,
	}
	return code, response
}

func failedResponse(
	c *gin.Context,
	code int,
	message string,
	err error,
	data interface{},
) (int, interface{}) {
	response := gin.H{
		"success": false,
		"message": message,
		"error":   err.Error(),
		"data":    data,
	}
	return code, response
}

func (d DonateController) GetAllDonate(c *gin.Context) {
	users, err := d.service.GetDonateAll()
	if err != nil {
		c.JSON(failedResponse(c, 500, "Error occured.", err, nil))
	}
	c.JSON(successResponse(c, 200, "Get all donate success!", users))
}

func (d DonateController) GetDonateById(c *gin.Context) {
	idParam, _ := strconv.Atoi(c.Param("id"))
	users, err := d.service.GetDonateById(uint(idParam))
	if err != nil {
		c.JSON(failedResponse(c, 500, "Error occured.", err, nil))
	}
	c.JSON(200, gin.H{"data": users})
}

func (d DonateController) InsertDonate(c *gin.Context) {
	donate := models.Donate{}
	header := c.Request.Header.Get("Authorization")
	header = header[len("Bearer "):]

	if err := c.ShouldBindJSON(&donate); err != nil {
		c.JSON(failedResponse(c, 500, "Error occured.", err, nil))
		return
	}

	valid, err := d.jwtService.Authorize(header)
	if err != nil && !valid {
		c.JSON(failedResponse(c, 403, "JWT Error.", err, nil))
		c.Abort()
	}

	claims := d.jwtService.ExtractClaims(header)
	userId := claims["id"].(float64)

	user, _ := d.userService.GetUserById(uint(userId))

	if err := d.service.InsertDonate(user, donate); err != nil {
		c.JSON(failedResponse(c, 500, "Error occured.", err, nil))
		return
	}

	c.JSON(200, gin.H{
		"message": "Insert New Donation Success!",
	})
}

func (d DonateController) TakeDonation(c *gin.Context) {
	header := c.Request.Header.Get("Authorization")
	header = header[len("Bearer "):]

	idParam, _ := strconv.Atoi(c.Param("id"))
	quantityQuery := c.DefaultQuery("quantity", "1")
	quantityParam, _ := strconv.ParseUint(quantityQuery, 10, 64)

	donate, err := d.service.GetDonateById(uint(idParam))
	if err != nil {
		c.JSON(failedResponse(c, 500, "Error occured.", err, nil))
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

	userDonation, err := d.service.GetUserByDonateId(uint(idParam))
	if err != nil {
		c.JSON(failedResponse(c, 500, "Error occured.", err, nil))
		return
	}

	user, _ := d.userService.GetUserById(uint(userId))
	donatur, _ := d.userService.GetUserById(uint(userDonation))

	donateNew, err := d.service.TakeDonation(uint(quantityParam), user, donate, donatur)
	if err != nil {
		c.JSON(failedResponse(c, 500, "Error occured.", err, nil))
		return
	}

	if donateNew == 0 {
		deleteDonateError := d.service.DeleteDonate(donate.ID)
		if deleteDonateError != nil {
			c.JSON(failedResponse(c, 500, "Error occured.", err, nil))
			return
		}
	}

	c.JSON(successResponse(c, 200, "Donation taken!", nil))
}

func (d DonateController) UpdateDonate(c *gin.Context) {
	donate := models.Donate{}
	idParam, _ := strconv.Atoi(c.Param("id"))

	if err := c.ShouldBindJSON(&donate); err != nil {
		c.JSON(failedResponse(c, 500, "Error occured.", err, nil))
		return
	}

	if err := d.service.UpdateDonate(uint(idParam), donate); err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(successResponse(c, 200, "Update success!", nil))
}

func (d DonateController) DeleteDonate(c *gin.Context) {
	idParam, _ := strconv.Atoi(c.Param("id"))

	if err := d.service.DeleteDonate(uint(idParam)); err != nil {
		c.JSON(failedResponse(c, 500, "Error occured.", err, nil))
		return
	}

	c.JSON(200, gin.H{
		"message": "Delete Donation!",
	})
}
