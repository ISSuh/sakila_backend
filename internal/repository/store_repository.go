package repository

import "github.com/ISSuh/msago-sample/internal/domain/model"

type StoreRepository interface {
	StoreById(storeId int) (*model.Store, error)
	StoreAddressById(storeId int) (*model.Address, error)
	StoresOnCountry(country *model.Countries) ([]*model.Store, error)
}
