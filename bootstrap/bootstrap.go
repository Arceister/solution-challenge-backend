package bootstrap

import (
	"api-solution/controllers"
	"api-solution/lib"
	"api-solution/routers"
	"api-solution/services"
	"context"
	"fmt"

	"go.uber.org/fx"
)

var Module = fx.Options(
	controllers.Module,
	routers.Module,
	lib.Module,
	services.Module,
	fx.Invoke(bootstrap),
)

func bootstrap(
	lifecycle fx.Lifecycle,
	handler lib.RequestHandler,
	routes routers.Routes) {
	lifecycle.Append(fx.Hook{
		OnStart: func(c context.Context) error {
			fmt.Println("Starting Application")
			go func() {
				routes.Setup()
				handler.Gin.Run(":8085")
			}()
			return nil
		},
		OnStop: func(c context.Context) error {
			fmt.Println("Stopping App")
			return nil
		},
	})
}
