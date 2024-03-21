package service

import (
	"github.com/ISSuh/monolith-sakila/internal/common"
	"github.com/ISSuh/monolith-sakila/internal/logger"
	"github.com/ISSuh/monolith-sakila/internal/model"
	"github.com/ISSuh/monolith-sakila/internal/repository"
)

type StaffService interface {
	StaffById(staffId int) (*model.Staff, error)
	Staffs(page common.Pagnation) ([]*model.Staff, error)
}

type staffService struct {
	log             logger.Logger
	staffRepository repository.StaffRepository
}

func NewStaffService(l logger.Logger, r repository.StaffRepository) StaffService {
	return &staffService{
		log:             l,
		staffRepository: r,
	}
}

func (s *staffService) StaffById(staffId int) (*model.Staff, error) {
	return s.staffRepository.StaffById(staffId)
}

func (s *staffService) Staffs(page common.Pagnation) ([]*model.Staff, error) {
	return s.staffRepository.Staffs(page)
}
