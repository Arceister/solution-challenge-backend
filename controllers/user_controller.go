package controllers

import (
	"api-solution/models"
	"api-solution/services"
	"fmt"
	"net/http"
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

func (u UserController) InsertUser(c *gin.Context) {
	user := models.User{}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := u.service.InsertUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		c.JSON(200, gin.H{
			"message": "Success",
		})
	}
}

func (u UserController) UpdateUser(c *gin.Context) {
	user := models.User{}
	idParam, _ := strconv.Atoi(c.Param("id"))

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := u.service.UpdateUser(uint(idParam), user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Update success!",
	})
}

func (u UserController) DeleteUser(c *gin.Context) {
	idParam, _ := strconv.Atoi(c.Param("id"))

	if err := u.service.DeleteUser(uint(idParam)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Delete success!",
	})
}
