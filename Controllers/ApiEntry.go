package controllers

import "github.com/gin-gonic/gin"

func ApiEntryGet(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "API ready to go!",
	})
}
