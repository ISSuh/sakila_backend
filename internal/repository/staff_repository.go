package repository

import (
	"github.com/ISSuh/monolith-sakila/internal/common"
	"github.com/ISSuh/monolith-sakila/internal/logger"
	"github.com/ISSuh/monolith-sakila/internal/model"
	"github.com/ISSuh/monolith-sakila/pkg/db"
)

type StaffRepository interface {
	StaffById(staffId int) (*model.Staff, error)
	Staffs(page common.Pagnation) ([]*model.Staff, error)
}

type staffRepository struct {
	log logger.Logger
	db  *db.Database
}

func NewStaffRepository(l logger.Logger, d *db.Database) StaffRepository {
	return &staffRepository{
		log: l,
		db:  d,
	}
}

func (r *staffRepository) StaffById(staffId int) (*model.Staff, error) {
	e := r.db.Engine()

	staff := new(model.Staff)
	if err := e.Model(&staff).Preload("Address.City.Country").Preload("Store").Find(&staff, staffId).Error; err != nil {
		return nil, err
	}
	return staff, nil
}

func (r *staffRepository) Staffs(page common.Pagnation) ([]*model.Staff, error) {
	e := r.db.Engine()

	staffs := []*model.Staff{}
	if err := e.Model(&staffs).Find(&staffs).Error; err != nil {
		return nil, err
	}
	return staffs, nil
}
