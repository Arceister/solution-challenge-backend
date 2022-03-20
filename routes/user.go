package routes

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
		api.GET("/users", s.userController.GetUser)
		api.GET("/user", s.userController.GetUserByOwnId)
		api.GET("/user/:id", s.userController.GetUserById)
		api.POST("/user/register", s.userController.InsertUser)
		api.PUT("/user", s.userController.UpdateUser)
		api.DELETE("/user/:id", s.userController.DeleteUser)
	}
}

func NewUserRoutes(handler lib.RequestHandler, userController controllers.UserController) UserRoutes {
	return UserRoutes{
		handler:        handler,
		userController: userController,
	}
}
