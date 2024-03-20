package repository

import (
	"github.com/ISSuh/msago-sample/internal/domain/model"
	"github.com/ISSuh/msago-sample/internal/logger"
	"github.com/ISSuh/msago-sample/pkg/db"
)

type CountryRepository interface {
	CountryById(countryId int) (*model.Countries, error)
}

type countryRepository struct {
	log logger.Logger
	db  *db.Database
}

func NewCountryRepository(l logger.Logger, d *db.Database) CountryRepository {
	return &countryRepository{
		log: l,
		db:  d,
	}
}

func (r *countryRepository) CountryById(countryId int) (*model.Countries, error) {
	return nil, nil
}
