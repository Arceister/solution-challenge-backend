package middlewares

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewCorsMiddleware),
	fx.Provide(NewMiddleware),
)

type Middlewares []MiddlewareSingle

type MiddlewareSingle interface {
	Setup()
}

func NewMiddleware(
	corsMiddleware CorsMiddleware,
) Middlewares {
	return Middlewares{
		corsMiddleware,
	}
}

func (m Middlewares) Setup() {
	for _, middleware := range m {
		middleware.Setup()
	}
}
