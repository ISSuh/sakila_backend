package repository

import (
	"github.com/ISSuh/msago-sample/internal/domain/model"
	"github.com/ISSuh/msago-sample/internal/logger"
	"github.com/ISSuh/msago-sample/pkg/db"
)

type CityRepository interface {
	CityById(cityId int) (*model.Cities, error)
	CitiesOnCountry(countryId int) ([]*model.Cities, error)
}

type cityRepository struct {
	log logger.Logger
	db  *db.Database
}

func NewCityRepository(l logger.Logger, d *db.Database) CityRepository {
	return &cityRepository{
		log: l,
		db:  d,
	}
}

func (r *cityRepository) CityById(cityId int) (*model.Cities, error) {
	e := r.db.Engine()

	city := new(model.Cities)
	if err := e.Model(&city).Find(&city, cityId).Error; err != nil {
		return nil, err
	}
	return city, nil
}

func (r *cityRepository) CitiesOnCountry(countryId int) ([]*model.Cities, error) {
	e := r.db.Engine()

	cities := make([]*model.Cities, 0)
	if err := e.Model(&cities).Where("country_id=?", countryId).Find(&cities).Error; err != nil {
		return nil, err
	}
	return cities, nil
}
