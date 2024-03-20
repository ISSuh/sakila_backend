package service

import (
	"github.com/ISSuh/msago-sample/internal/domain/model"
	"github.com/ISSuh/msago-sample/internal/logger"
	"github.com/ISSuh/msago-sample/internal/repository"
	"github.com/ISSuh/msago-sample/pkg/db"
)

type StoreService interface {
	StoreById(storeId int) (*model.Store, error)
	StoreAddressById(storeId int) (*model.Address, error)
	StoresOnCountry(cuntryId int) ([]*model.Store, error)
}

type storeService struct {
	*db.Transaction

	log               logger.Logger
	serviceRepository repository.StoreRepository

	cityService    CityService
	addressService AddressService
}

func NewStoreService(l logger.Logger, tx *db.Transaction, r repository.StoreRepository, cityService CityService, addressService AddressService) StoreService {
	return &storeService{
		Transaction:       tx,
		serviceRepository: r,
		cityService:       cityService,
		addressService:    addressService,
	}
}

func (s *storeService) StoreById(id int) (*model.Store, error) {
	return s.serviceRepository.StoreById(id)
}

func (s *storeService) StoreAddressById(id int) (*model.Address, error) {
	return s.serviceRepository.StoreAddressById(id)
}

func (s *storeService) StoresOnCountry(countryId int) ([]*model.Store, error) {
	cities, err := s.cityService.CitiesOnCountry(countryId)
	if err != nil {
		return nil, err
	}

	addresses, err := s.addressService.AddressOnCites(cities)
	if err != nil {
		return nil, err
	}
	return s.serviceRepository.StoresOnAddress(addresses)
}
