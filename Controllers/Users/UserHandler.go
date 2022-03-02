package users

import "github.com/gin-gonic/gin"

func UserGet(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "user get",
	})
}
