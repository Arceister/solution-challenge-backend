package services

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewUserService),
	fx.Provide(NewEntryService),
	fx.Provide(NewJWTAuthService),
	fx.Provide(NewDonateService),
	fx.Provide(NewPreferableService),
)
