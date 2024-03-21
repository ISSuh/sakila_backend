package factory

import (
	"github.com/ISSuh/monolith-sakila/internal/controller/rest/handler"
	"github.com/ISSuh/monolith-sakila/internal/logger"
)

type Handlers struct {
	Store       *handler.StoreHandler
	Healthcheck *handler.HealthcheckHandler
	Staff       *handler.StaffHandler
	Film        *handler.FilmHandler
}

func NewHandlers(l logger.Logger, s *Services) (*Handlers, error) {
	h := &Handlers{
		Store:       handler.NewStoreHandler(l, s.Store),
		Healthcheck: handler.NewHealthcheck(l),
		Staff:       handler.NewStaffHandler(l, s.Staff),
		Film:        handler.NewFilmHandler(l, s.Film),
	}
	return h, nil
}
