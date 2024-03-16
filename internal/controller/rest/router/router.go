package router

import (
	mw "github.com/ISSuh/msago-sample/internal/controller/rest/middleware"
	"github.com/ISSuh/msago-sample/internal/factory"

	"github.com/labstack/echo/v4"
)

const (
	Version = "v1"
)

func Route(e *echo.Echo, m *mw.Middleware, h *factory.Handlers) error {
	v := e.Group(Version)
	{
		item := v.Group("/item")
		item.GET(":itemId", func(echoCtx echo.Context) error {
			return nil
		})
	}

	{
		healthcheck := v.Group("/healthcheck")
		healthcheck.GET("", h.Healthcheck.Check())
	}

	{
		store := v.Group("/store")
		store.GET("/:id", h.Store.StoreById())
	}

	return nil
}
