package routes

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewRoutes),
	fx.Provide(NewUserRoutes),
	fx.Provide(NewEntryRoutes),
	fx.Provide(NewAuthRoutes),
	fx.Provide(NewDonateRoutes),
	fx.Provide(NewPreferableRoutes),
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
	preferableRoutes PreferableRoutes,
) Routes {
	return Routes{
		userRoutes,
		entryRoutes,
		authRoutes,
		donateRoutes,
		preferableRoutes,
	}
}

func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
