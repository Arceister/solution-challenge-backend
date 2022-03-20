package routes

import (
	"api-solution/controllers"
	"api-solution/lib"
	"fmt"
)

type AuthRoutes struct {
	handler           lib.RequestHandler
	jwtAuthController controllers.JWTAuthController
}

func (s AuthRoutes) Setup() {
	fmt.Println("Setting routes")
	authApi := s.handler.Gin.Group("/api/auth")
	{
		authApi.POST("/login", s.jwtAuthController.SignIn)
		authApi.GET("/", s.jwtAuthController.AuthorizeToken)
	}
}

func NewAuthRoutes(
	handler lib.RequestHandler,
	jwtAuthController controllers.JWTAuthController,
) AuthRoutes {
	return AuthRoutes{
		handler:           handler,
		jwtAuthController: jwtAuthController,
	}
}
