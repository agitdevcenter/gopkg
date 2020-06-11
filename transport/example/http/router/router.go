package router

import (
	Handler "github.com/agitdevcenter/gopkg/transport/example/http/handler"
	Echo "github.com/labstack/echo/v4"
)

type Router struct {
	handler *Handler.Handler
}

func New(handler *Handler.Handler) *Router {
	return &Router{handler: handler}
}

func (r *Router) Route(echo *Echo.Echo) {
	health := echo.Group("/health")
	health.GET("", r.handler.Health.Ping)
	health.GET("/error", r.handler.Health.PingError)
}
