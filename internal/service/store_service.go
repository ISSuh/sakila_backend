package service

import "github.com/ISSuh/msago-sample/internal/domain/model"

type StoreService interface {
	StoreById(storeId int) (*model.Store, error)
	StoreAddressById(storeId int) (*model.Address, error)
	StoresOnCountry(cuntryId int) ([]*model.Store, error)
}
