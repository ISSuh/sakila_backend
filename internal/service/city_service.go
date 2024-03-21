package service

import (
	"github.com/ISSuh/monolith-sakila/internal/domain/model"
	"github.com/ISSuh/monolith-sakila/internal/logger"
	"github.com/ISSuh/monolith-sakila/internal/repository"
)

type CityService interface {
	CityById(cityId int) (*model.Cities, error)
	CitiesOnCountry(countryId int) ([]*model.Cities, error)
}

type cityService struct {
	log            logger.Logger
	cityRepository repository.CityRepository
}

func NewCityService(l logger.Logger, r repository.CityRepository) CityService {
	return &cityService{
		cityRepository: r,
	}
}

func (s *cityService) CityById(cityId int) (*model.Cities, error) {
	return s.cityRepository.CityById(cityId)
}

func (s *cityService) CitiesOnCountry(countryId int) ([]*model.Cities, error) {
	return s.cityRepository.CitiesOnCountry(countryId)
}
