package factory

import (
	"github.com/ISSuh/monolith-sakila/internal/logger"
	"github.com/ISSuh/monolith-sakila/internal/service"
)

type Services struct {
	Store   service.StoreService
	Country service.CountryService
	City    service.CityService
	Address service.AddressService
}

func NewServices(l logger.Logger, r *Repositories) (*Services, error) {
	country := service.NewCountryService(l, r.Country)
	city := service.NewCityService(l, r.City)
	address := service.NewAddressService(l, r.Address)
	store := service.NewStoreService(l, r.Transaction, r.Store, city, address)

	s := &Services{
		Country: country,
		City:    city,
		Address: address,
		Store:   store,
	}
	return s, nil
}
