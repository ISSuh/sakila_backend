package implementation

import (
	"github.com/ISSuh/msago-sample/internal/domain/model"
	"github.com/ISSuh/msago-sample/internal/logger"
	"github.com/ISSuh/msago-sample/internal/repository"
	"github.com/ISSuh/msago-sample/pkg/db"
	"github.com/davecgh/go-spew/spew"
)

type StoreRepositoryImpl struct {
	log logger.Logger
	db  *db.Database
}

func NewStoreRepository(l logger.Logger, d *db.Database) repository.StoreRepository {
	return &StoreRepositoryImpl{
		log: l,
		db:  d,
	}
}

func (r *StoreRepositoryImpl) StoreById(id int) (*model.Store, error) {
	e := r.db.Engine()

	store := new(model.Store)
	err := e.Preload("Address").Preload("ManagerStaff").Where("store_id=?", id).Find(&store).Error
	if err != nil {
		return nil, err
	}

	r.log.Infof("[StoreById]%+v", spew.Sdump(store))
	return store, nil
}
