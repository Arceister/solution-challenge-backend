package repository

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewUserRepository),
	fx.Provide(NewDonateRepository),
	fx.Provide(NewPreferableRepository),
)
