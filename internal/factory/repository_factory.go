package factory

import (
	"github.com/ISSuh/msago-sample/internal/logger"
	"github.com/ISSuh/msago-sample/internal/repository"
	"github.com/ISSuh/msago-sample/internal/repository/implementation"
	"github.com/ISSuh/msago-sample/pkg/db"
)

type Repositories struct {
	Store repository.StoreRepository
}

func NewRepositories(l logger.Logger, d *db.Database) (*Repositories, error) {
	r := &Repositories{
		Store: implementation.NewStoreRepository(l, d),
	}
	return r, nil
}
