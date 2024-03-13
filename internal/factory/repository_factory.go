package factory

import (
	"github.com/ISSuh/msago-sample/internal/logger"
	"github.com/ISSuh/msago-sample/internal/repository"
	"github.com/ISSuh/msago-sample/internal/repository/implementation"
)

type Repositories struct {
	Store repository.StoreRepository
}

func NewRepositories(l logger.Logger) (*Repositories, error) {
	r := &Repositories{
		Store: implementation.NewStoreRepository(l),
	}
	return r, nil
}
