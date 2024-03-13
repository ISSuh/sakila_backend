package factory

import (
	"github.com/ISSuh/msago-sample/internal/controller/rest/handler"
	"github.com/ISSuh/msago-sample/internal/logger"
)

type Handlers struct {
	Store       *handler.Store
	Healthcheck *handler.Healthcheck
}

func NewHandlers(l logger.Logger, s *Services) (*Handlers, error) {
	h := &Handlers{
		Store:       handler.NewStoreHandler(l, s.Store),
		Healthcheck: handler.NewHealthcheck(l),
	}
	return h, nil
}
