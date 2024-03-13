package implementation

import (
	"github.com/ISSuh/msago-sample/internal/domain/model"
	"github.com/ISSuh/msago-sample/internal/logger"
	"github.com/ISSuh/msago-sample/internal/repository"
)

type StoreRepositoryImpl struct {
	log logger.Logger
}

func NewStoreRepository(l logger.Logger) repository.StoreRepository {
	return &StoreRepositoryImpl{
		log: l,
	}
}

func (r *StoreRepositoryImpl) StoreById(int) model.Store {
	return model.Store{}
}
