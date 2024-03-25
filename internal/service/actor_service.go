package service

import (
	"github.com/ISSuh/sakila_backend/internal/logger"
	"github.com/ISSuh/sakila_backend/internal/model"
	"github.com/ISSuh/sakila_backend/internal/repository"
	"github.com/ISSuh/sakila_backend/pkg/db"
)

type ActorService interface {
	ActorById(actorId int) (*model.Actor, error)
}

type actorService struct {
	*db.Transaction

	log             logger.Logger
	actorRepository repository.ActorRepository
}

func NewActorService(l logger.Logger, tx *db.Transaction, r repository.ActorRepository) ActorService {
	return &actorService{
		log:             l,
		Transaction:     tx,
		actorRepository: r,
	}
}

func (s *actorService) ActorById(id int) (*model.Actor, error) {
	return s.actorRepository.ActorById(id)
}
