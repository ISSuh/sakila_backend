package repository

import (
	"github.com/ISSuh/msago-sample/internal/domain/model"
	"github.com/ISSuh/msago-sample/internal/logger"
	"github.com/ISSuh/msago-sample/pkg/db"
)

const (
	sqlWhereInCityId = "city_id In ?"
)

type AddressRepository interface {
	AddressById(addressId int) (*model.Address, error)
	AddressOnCities(cities []*model.Cities) ([]*model.Address, error)
}

type addressRepository struct {
	log logger.Logger
	db  *db.Database
}

func NewAddressRepository(l logger.Logger, d *db.Database) AddressRepository {
	return &addressRepository{
		log: l,
		db:  d,
	}
}

func (r *addressRepository) AddressById(addressId int) (*model.Address, error) {
	e := r.db.Engine()

	address := new(model.Address)
	if err := e.Model(&address).Find(&address, addressId).Error; err != nil {
		return nil, err
	}
	return address, nil
}

func (r *addressRepository) AddressOnCities(cities []*model.Cities) ([]*model.Address, error) {
	e := r.db.Engine()

	cityIds := []int{}
	for _, city := range cities {
		cityIds = append(cityIds, city.CityId)
	}

	addresses := []*model.Address{}
	if err := e.Model(&addresses).Where(sqlWhereInCityId, cityIds).Find(&addresses).Error; err != nil {
		return nil, err
	}
	return addresses, nil
}
