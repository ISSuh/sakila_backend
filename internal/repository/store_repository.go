package repository

import (
	"github.com/ISSuh/msago-sample/internal/domain/model"
	"github.com/ISSuh/msago-sample/internal/logger"
	"github.com/ISSuh/msago-sample/pkg/db"
)

const (
	sqlWhereInAddressId = "address_id In ?"
)

type StoreRepository interface {
	StoreById(storeId int) (*model.Store, error)
	StoreAddressById(storeId int) (*model.Address, error)
	StoresOnAddress(addresses []*model.Address) ([]*model.Store, error)
}

type storeRepository struct {
	log logger.Logger
	db  *db.Database
}

func NewStoreRepository(l logger.Logger, d *db.Database) StoreRepository {
	return &storeRepository{
		log: l,
		db:  d,
	}
}

func (r *storeRepository) StoreById(id int) (*model.Store, error) {
	e := r.db.Engine()

	store := new(model.Store)
	if err := e.Model(&store).Preload("Address.City.Country").Preload("ManagerStaff").Find(&store, id).Error; err != nil {
		return nil, err
	}
	return store, nil
}

func (r *storeRepository) StoreAddressById(id int) (*model.Address, error) {
	e := r.db.Engine()

	store := new(model.Store)
	if err := e.Model(&store).Preload("Address.City.Country").Find(&store, id).Error; err != nil {
		return nil, err
	}
	return store.Address, nil
}

func (r *storeRepository) StoresOnAddress(addresses []*model.Address) ([]*model.Store, error) {
	e := r.db.Engine()

	addressIds := []int{}
	for _, address := range addresses {
		addressIds = append(addressIds, address.AddressId)
	}

	stores := []*model.Store{}
	if err := e.Model(&stores).
		Preload("Address.City.Country").Preload("ManagerStaff").
		Where(sqlWhereInAddressId, addressIds).Find(&stores).Error; err != nil {
		return nil, err
	}
	return stores, nil
}
