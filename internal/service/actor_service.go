package service

import (
	"github.com/ISSuh/sakila_backend/internal/logger"
	"github.com/ISSuh/sakila_backend/internal/model"
	"github.com/ISSuh/sakila_backend/internal/repository"
	"github.com/ISSuh/sakila_backend/pkg/db"
)

type ActorService interface {
	ActorById(actorId int, withFilm bool) (*model.Actor, error)
	FilmOnActorById(id int) ([]*model.Film, error)
	ActorByLastName(lastName string) ([]*model.Actor, error)
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

func (s *actorService) ActorById(id int, withFilm bool) (*model.Actor, error) {
	if withFilm {
		return s.actorRepository.ActorWithFilmById(id)
	}
	return s.actorRepository.ActorById(id)
}

func (s *actorService) FilmOnActorById(id int) ([]*model.Film, error) {
	actor, err := s.actorRepository.ActorWithFilmById(id)
	if err != nil {
		return nil, err
	}
	return actor.Films, nil
}

func (s *actorService) ActorByLastName(lastName string) ([]*model.Actor, error) {
	return s.actorRepository.ActorByLastName(lastName)
}
