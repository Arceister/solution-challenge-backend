package routers

import (
	"api-solution/controllers"
	"api-solution/lib"
	"fmt"
)

type UserRoutes struct {
	handler        lib.RequestHandler
	userController controllers.UserController
}

func (s UserRoutes) Setup() {
	fmt.Println("Setting routes")
	api := s.handler.Gin.Group("/api")
	{
		api.GET("/user", s.userController.Get)
	}
}

func NewUserRoutes(handler lib.RequestHandler, userController controllers.UserController) UserRoutes {
	return UserRoutes{
		handler:        handler,
		userController: userController,
	}
}
