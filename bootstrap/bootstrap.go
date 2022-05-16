package bootstrap

import (
	"api-solution/controllers"
	"api-solution/lib"
	"api-solution/middlewares"
	"api-solution/models"
	"api-solution/repository"
	"api-solution/routes"
	"api-solution/services"
	"context"
	"fmt"

	"go.uber.org/fx"
)

var Module = fx.Options(
	controllers.Module,
	routes.Module,
	lib.Module,
	services.Module,
	middlewares.Module,
	models.Module,
	repository.Module,
	fx.Invoke(bootstrap),
)

func bootstrap(
	lifecycle fx.Lifecycle,
	handler lib.RequestHandler,
	routes routes.Routes,
	env lib.Env,
	middlewares middlewares.Middlewares,
	migrations models.Migrations,
	mongo models.MongoCall,
) {
	lifecycle.Append(fx.Hook{
		OnStart: func(c context.Context) error {
			fmt.Println("Starting Application")

			env.LoadEnv()

			go func() {
				migrations.Migrate()
				mongo.Mongo()
				middlewares.Setup()
				routes.Setup()
				handler.Gin.Run()
			}()
			return nil
		},
		OnStop: func(c context.Context) error {
			fmt.Println("Stopping App")
			return nil
		},
	})
}
