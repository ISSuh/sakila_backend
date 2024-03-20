package factory

import (
	"github.com/ISSuh/msago-sample/internal/logger"
	"github.com/ISSuh/msago-sample/internal/repository"
	"github.com/ISSuh/msago-sample/pkg/db"
)

type Repositories struct {
	Transaction *db.Transaction
	Store       repository.StoreRepository
	Country     repository.CountryRepository
	City        repository.CityRepository
	Address     repository.AddressRepository
}

func NewRepositories(l logger.Logger, d *db.Database) (*Repositories, error) {
	r := &Repositories{
		Transaction: db.NewTransaction(d),
		Store:       repository.NewStoreRepository(l, d),
		Country:     repository.NewCountryRepository(l, d),
		City:        repository.NewCityRepository(l, d),
		Address:     repository.NewAddressRepository(l, d),
	}
	return r, nil
}
