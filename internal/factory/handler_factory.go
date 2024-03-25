package factory

import (
	"github.com/ISSuh/sakila_backend/internal/controller/rest/handler"
	"github.com/ISSuh/sakila_backend/internal/logger"
)

type Handlers struct {
	Store       *handler.StoreHandler
	Healthcheck *handler.HealthcheckHandler
	Staff       *handler.StaffHandler
	Film        *handler.FilmHandler
	Actor       *handler.ActorHandler
}

func NewHandlers(l logger.Logger, s *Services) (*Handlers, error) {
	h := &Handlers{
		Healthcheck: handler.NewHealthcheck(l),
		Store:       handler.NewStoreHandler(l, s.Store),
		Staff:       handler.NewStaffHandler(l, s.Staff),
		Film:        handler.NewFilmHandler(l, s.Film),
		Actor:       handler.NewActorHandler(l, s.Actor),
	}
	return h, nil
}
