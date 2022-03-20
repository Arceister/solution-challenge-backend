package routes

import "api-solution/lib"

type DonateRoutes struct {
	handler lib.RequestHandler
}

func (s DonateRoutes) Setup() {
	donateApi := s.handler.Gin.Group("/api/donate")
	{
		donateApi.GET("/")
	}
}
