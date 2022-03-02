package routes

import (
	userController "api-solution/Controllers/Users"

	"github.com/gin-gonic/gin"
)

func UserRoutes(route *gin.Engine) {

	userGroup := route.Group("/user")
	{
		userGroup.GET("/", userController.UserGet)
	}

}
