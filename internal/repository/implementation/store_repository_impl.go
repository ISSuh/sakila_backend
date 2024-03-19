package implementation

import (
	"github.com/ISSuh/msago-sample/internal/domain/model"
	"github.com/ISSuh/msago-sample/internal/logger"
	"github.com/ISSuh/msago-sample/internal/repository"
	"github.com/ISSuh/msago-sample/pkg/db"
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
	if err := e.Model(&store).Preload("Address.City.Country").Preload("ManagerStaff").Find(&store, id).Error; err != nil {
		return nil, err
	}
	return store, nil
}

func (r *StoreRepositoryImpl) StoreAddressById(id int) (*model.Address, error) {
	e := r.db.Engine()

	store := new(model.Store)
	if err := e.Model(&store).Preload("Address.City.Country").Find(&store, id).Error; err != nil {
		return nil, err
	}
	return store.Address, nil
}

func (r *StoreRepositoryImpl) StoresOnCountry(country *model.Countries) ([]*model.Store, error) {
	// e := r.db.Engine()

	return nil, nil
}
