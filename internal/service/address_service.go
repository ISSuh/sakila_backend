package service

import (
	"github.com/ISSuh/msago-sample/internal/domain/model"
	"github.com/ISSuh/msago-sample/internal/logger"
	"github.com/ISSuh/msago-sample/internal/repository"
)

type AddressService interface {
	AddressById(addressId int) (*model.Address, error)
	AddressOnCity(cityId int) ([]*model.Address, error)
	AddressOnCites(cities []*model.Cities) ([]*model.Address, error)
}

type addressService struct {
	log               logger.Logger
	addressRepository repository.AddressRepository
}

func NewAddressService(l logger.Logger, r repository.AddressRepository) AddressService {
	return &addressService{
		addressRepository: r,
	}
}

func (s *addressService) AddressById(addressId int) (*model.Address, error) {
	return s.addressRepository.AddressById(addressId)
}

func (s *addressService) AddressOnCity(cityId int) ([]*model.Address, error) {
	cities := []*model.Cities{
		&model.Cities{CityId: cityId},
	}
	return s.addressRepository.AddressOnCities(cities)
}

func (s *addressService) AddressOnCites(cities []*model.Cities) ([]*model.Address, error) {
	return s.addressRepository.AddressOnCities(cities)
}
