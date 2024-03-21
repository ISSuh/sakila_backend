package factory

import (
	"github.com/ISSuh/monolith-sakila/internal/logger"
	"github.com/ISSuh/monolith-sakila/internal/repository"
	"github.com/ISSuh/monolith-sakila/pkg/db"
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
