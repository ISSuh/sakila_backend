package factory

import (
	"github.com/ISSuh/sakila_backend/internal/logger"
	"github.com/ISSuh/sakila_backend/internal/service"
)

type Services struct {
	Store   service.StoreService
	Country service.CountryService
	City    service.CityService
	Address service.AddressService
	Staff   service.StaffService
	Film    service.FilmService
	Actor   service.ActorService
}

func NewServices(l logger.Logger, r *Repositories) (*Services, error) {
	country := service.NewCountryService(l, r.Country)
	city := service.NewCityService(l, r.City)
	address := service.NewAddressService(l, r.Address)
	store := service.NewStoreService(l, r.Transaction, r.Store, city, address)
	staff := service.NewStaffService(l, r.Staff)
	film := service.NewFilmService(l, r.Transaction, r.Film)
	actor := service.NewActorService(l, r.Transaction, r.Actor)

	s := &Services{
		Country: country,
		City:    city,
		Address: address,
		Store:   store,
		Staff:   staff,
		Film:    film,
		Actor:   actor,
	}
	return s, nil
}
