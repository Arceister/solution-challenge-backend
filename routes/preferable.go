package routes

import (
	"api-solution/controllers"
	"api-solution/lib"
)

type PreferableRoutes struct {
	handler              lib.RequestHandler
	preferableController controllers.PreferableController
}

func NewPreferableRoutes(
	handler lib.RequestHandler,
	preferableController controllers.PreferableController,
) PreferableRoutes {
	return PreferableRoutes{
		handler:              handler,
		preferableController: preferableController,
	}
}

func (s PreferableRoutes) Setup() {
	api := s.handler.Gin.Group("/api/prefer")
	{
		api.GET("/all", s.preferableController.GetAllPreferables)
		api.POST("/", s.preferableController.InsertPreferable)
	}
}
