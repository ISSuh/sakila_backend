package factory

import (
	"github.com/ISSuh/msago-sample/internal/logger"
	"github.com/ISSuh/msago-sample/internal/service"
	"github.com/ISSuh/msago-sample/internal/service/implementation"
)

type Services struct {
	Store service.StoreService
}

func NewServices(l logger.Logger, r *Repositories) (*Services, error) {
	s := &Services{
		Store: implementation.NewStoreService(l, r.Store),
	}
	return s, nil
}
