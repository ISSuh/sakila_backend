package factory

import (
	"github.com/ISSuh/sakila_backend/internal/logger"
	"github.com/ISSuh/sakila_backend/internal/repository"
	"github.com/ISSuh/sakila_backend/pkg/db"
)

type Repositories struct {
	Transaction *db.Transaction
	Store       repository.StoreRepository
	Country     repository.CountryRepository
	City        repository.CityRepository
	Address     repository.AddressRepository
	Staff       repository.StaffRepository
	Film        repository.FilmRepository
}

func NewRepositories(l logger.Logger, d *db.Database) (*Repositories, error) {
	r := &Repositories{
		Transaction: db.NewTransaction(d),
		Store:       repository.NewStoreRepository(l, d),
		Country:     repository.NewCountryRepository(l, d),
		City:        repository.NewCityRepository(l, d),
		Address:     repository.NewAddressRepository(l, d),
		Staff:       repository.NewStaffRepository(l, d),
		Film:        repository.NewFilmRepository(l, d),
	}
	return r, nil
}
