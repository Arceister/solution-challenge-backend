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
	service    services.UserService
	jwtService services.JWTAuthService
}

func NewUserController(userService services.UserService, jwtService services.JWTAuthService) UserController {
	return UserController{
		service:    userService,
		jwtService: jwtService,
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

func (u UserController) GetUserByOwnId(c *gin.Context) {
	header := c.Request.Header.Get("Authorization")
	header = header[len("Bearer "):]
	valid, err := u.jwtService.Authorize(header)
	if err != nil && !valid {
		c.JSON(http.StatusForbidden, gin.H{
			"success": false,
			"message": "JWT Error",
			"error":   err.Error(),
		})
		c.Abort()
	} else {
		claims := u.jwtService.ExtractClaims(header)
		id := claims["id"].(float64)
		users, err := u.service.GetUserById(uint(id))
		if err != nil {
			fmt.Println(err)
		}
		c.JSON(200, gin.H{"data": users})
	}
}

func (u UserController) InsertUser(c *gin.Context) {
	user := models.User{}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	users, err := u.service.InsertUser(user)
	token := u.jwtService.CreateToken(users)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		c.JSON(200, gin.H{
			"message": "Success",
			"data":    users,
			"token":   token,
		})
	}
}

func (u UserController) UpdateUser(c *gin.Context) {
	user := models.User{}
	header := c.Request.Header.Get("Authorization")
	header = header[len("Bearer "):]
	valid, err := u.jwtService.Authorize(header)
	if err != nil && !valid {
		c.JSON(http.StatusForbidden, gin.H{
			"success": false,
			"message": "JWT Error",
			"error":   err.Error(),
		})
		c.Abort()
	}

	claims := u.jwtService.ExtractClaims(header)
	id := claims["id"].(float64)

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := u.service.UpdateUser(uint(id), user); err != nil {
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
