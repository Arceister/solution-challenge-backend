package controllers

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewUserController),
	fx.Provide(NewEntryController),
	fx.Provide(NewJWTAuthController),
	fx.Provide(NewDonateController),
)
