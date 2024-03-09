package router

import (
	"github.com/ISSuh/msago-sample/internal/controller/api/handler"
	mw "github.com/ISSuh/msago-sample/internal/controller/api/middleware"

	"github.com/labstack/echo/v4"
)

const (
	Version = "v1"
)

type Router struct {
	m mw.Middleware
	h handler.Handler
}

func NewRouter(middleware mw.Middleware, handler handler.Handler) Router {
	return Router{
		m: middleware,
		h: handler,
	}
}

func (r *Router) SetRoute(e *echo.Echo) error {
	v := e.Group(Version)

	{
		item := v.Group("/item")
		item.GET(":itemId", func(echoCtx echo.Context) error {
			return nil
		})

	}
	return nil
}
