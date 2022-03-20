package routes

import (
	"api-solution/controllers"
	"api-solution/lib"
)

type DonateRoutes struct {
	handler          lib.RequestHandler
	donateController controllers.DonateController
}

func (s DonateRoutes) Setup() {
	donateApi := s.handler.Gin.Group("/api/donate")
	{
		donateApi.GET("/", s.donateController.GetAllDonate)
		donateApi.GET("/:id", s.donateController.GetDonateById)
		donateApi.POST("/", s.donateController.InsertDonate)
	}
}

func NewDonateRoutes(
	handler lib.RequestHandler,
	donateController controllers.DonateController,
) DonateRoutes {
	return DonateRoutes{
		handler:          handler,
		donateController: donateController,
	}
}
