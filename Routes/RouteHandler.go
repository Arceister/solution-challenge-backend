package routes

import (
	apiEntry "api-solution/Controllers"

	"github.com/gin-gonic/gin"
)

func RouteHandler() *gin.Engine {
	router := gin.Default()

	defaultGroup := router.Group("/")
	{
		defaultGroup.GET("/", apiEntry.ApiEntryGet)
	}

	UserRoutes(router)

	return router
}
