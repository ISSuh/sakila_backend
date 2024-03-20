package service

import (
	"github.com/ISSuh/msago-sample/internal/domain/model"
	"github.com/ISSuh/msago-sample/internal/logger"
	"github.com/ISSuh/msago-sample/internal/repository"
)

type CountryService interface {
	CountryById(countryId int) (*model.Countries, error)
}

type countryService struct {
	log               logger.Logger
	countryRepository repository.CountryRepository
}

func NewCountryService(l logger.Logger, r repository.CountryRepository) CountryService {
	return &countryService{
		countryRepository: r,
	}
}

func (s *countryService) CountryById(countryId int) (*model.Countries, error) {
	return s.countryRepository.CountryById(countryId)
}
