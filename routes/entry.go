package routes

import (
	"api-solution/controllers"
	"api-solution/lib"
	"fmt"
)

type EntryRoutes struct {
	handler         lib.RequestHandler
	entryController controllers.EntryController
}

func (s EntryRoutes) Setup() {
	fmt.Println("Setting routes")
	entry := s.handler.Gin.Group("/")
	{
		entry.GET("/", s.entryController.Get)
		entry.POST("/")
	}
}

func NewEntryRoutes(handler lib.RequestHandler, entryController controllers.EntryController) EntryRoutes {
	return EntryRoutes{
		handler:         handler,
		entryController: entryController,
	}
}
