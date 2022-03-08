package routes

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewRoutes),
	fx.Provide(NewUserRoutes),
	fx.Provide(NewEntryRoutes),
)

type Routes []Route

type Route interface {
	Setup()
}

func NewRoutes(
	userRoutes UserRoutes,
	entryRoutes EntryRoutes,
) Routes {
	return Routes{
		userRoutes,
		entryRoutes,
	}
}

func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
