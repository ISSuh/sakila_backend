package implementation

import (
	"github.com/ISSuh/msago-sample/internal/domain/model"
	"github.com/ISSuh/msago-sample/internal/logger"
	"github.com/ISSuh/msago-sample/internal/repository"
	"github.com/ISSuh/msago-sample/internal/service"
)

type StoreServiceImpl struct {
	log               logger.Logger
	serviceRepository repository.StoreRepository
}

func NewStoreService(l logger.Logger, sr repository.StoreRepository) service.StoreService {
	return &StoreServiceImpl{
		serviceRepository: sr,
	}
}

func (s *StoreServiceImpl) StoreById(id int) model.Store {
	return s.serviceRepository.StoreById(id)
}
