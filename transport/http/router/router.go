package router

import (
	Echo "github.com/labstack/echo/v4"
)

type Router interface {
	Route(echo *Echo.Echo)
}
