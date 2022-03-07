package routers

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewRoutes),
	fx.Provide(NewUserRoutes),
)

type Routes []Route

type Route interface {
	Setup()
}

func NewRoutes(
	userRoutes UserRoutes,
) Routes {
	return Routes{
		userRoutes,
	}
}

func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
