package repository

import (
	"github.com/ISSuh/sakila_backend/internal/logger"
	"github.com/ISSuh/sakila_backend/internal/model"
	"github.com/ISSuh/sakila_backend/pkg/db"
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
