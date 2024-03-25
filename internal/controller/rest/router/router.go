package router

import (
	mw "github.com/ISSuh/sakila_backend/internal/controller/rest/middleware"
	"github.com/ISSuh/sakila_backend/internal/factory"
	"github.com/gin-gonic/gin"
)

const (
	Version1 = "v1"
)

func Route(e *gin.Engine, h *factory.Handlers) error {
	e.Use(mw.WrapContext())
	// e.Use(mw.ResponseMiddleware())

	v1 := e.Group(Version1)

	{
		healthcheck := v1.Group("/healthcheck")
		healthcheck.GET("", h.Healthcheck.Check())
	}

	{
		store := v1.Group("/store")
		store.GET("/:id", h.Store.StoreById())
		store.GET("/:id/address", h.Store.StoreAddressById())
		store.GET("/country/:countryId", h.Store.StoreOnCountry())
	}

	{
		store := v1.Group("/staff")
		store.GET("", h.Staff.Staffs())
		store.GET("/:id", h.Staff.StaffById())
	}

	{
		store := v1.Group("/film")
		store.GET("", h.Film.Films())
		store.GET("/:id", h.Film.FilmById())
	}

	return nil
}
