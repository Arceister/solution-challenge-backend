package middlewares

import (
	"api-solution/lib"

	"github.com/gin-contrib/cors"
)

type CorsMiddleware struct {
	handler lib.RequestHandler
	env     lib.Env
}

func NewCorsMiddleware(handler lib.RequestHandler, env lib.Env) CorsMiddleware {
	return CorsMiddleware{
		handler: handler,
		env:     env,
	}
}

func (m CorsMiddleware) Setup() {
	m.handler.Gin.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowHeaders:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "HEAD", "DELETE", "OPTIONS"},
		ExposeHeaders:    []string{"Content-Length"},
	}))
}
