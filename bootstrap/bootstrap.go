package bootstrap

import (
	"api-solution/controllers"
	"api-solution/lib"
	"api-solution/middlewares"
	"api-solution/routes"
	"api-solution/services"
	"context"
	"fmt"

	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

var Module = fx.Options(
	controllers.Module,
	routes.Module,
	lib.Module,
	services.Module,
	middlewares.Module,
	fx.Invoke(bootstrap),
)

func bootstrap(
	lifecycle fx.Lifecycle,
	handler lib.RequestHandler,
	routes routes.Routes,
	env lib.Env,
	middlewares middlewares.Middlewares) {
	lifecycle.Append(fx.Hook{
		OnStart: func(c context.Context) error {
			fmt.Println("Starting Application")

			godotenv.Load()
			env.LoadEnv()

			go func() {
				middlewares.Setup()
				routes.Setup()
				handler.Gin.Run(env.ServerPort)
			}()
			return nil
		},
		OnStop: func(c context.Context) error {
			fmt.Println("Stopping App")
			return nil
		},
	})
}
