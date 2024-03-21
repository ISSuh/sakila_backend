package repository

import (
	"github.com/ISSuh/monolith-sakila/internal/domain/model"
	"github.com/ISSuh/monolith-sakila/internal/logger"
	"github.com/ISSuh/monolith-sakila/pkg/db"
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
