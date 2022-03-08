package controllers

import "github.com/gin-gonic/gin"

type EntryController struct {
	Get func(c *gin.Context)
}

func NewEntryController() EntryController {
	return EntryController{
		Get: getEntrypointController(),
	}
}

func getEntrypointController() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "api accessed!",
		})
	}
}
