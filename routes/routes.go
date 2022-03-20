package routes

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewRoutes),
	fx.Provide(NewUserRoutes),
	fx.Provide(NewEntryRoutes),
	fx.Provide(NewAuthRoutes),
	fx.Provide(NewDonateRoutes),
)

type Routes []Route

type Route interface {
	Setup()
}

func NewRoutes(
	userRoutes UserRoutes,
	entryRoutes EntryRoutes,
	authRoutes AuthRoutes,
	donateRoutes DonateRoutes,
) Routes {
	return Routes{
		userRoutes,
		entryRoutes,
		authRoutes,
		donateRoutes,
	}
}

func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
