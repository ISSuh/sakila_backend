package service

import (
	"github.com/ISSuh/sakila_backend/internal/logger"
	"github.com/ISSuh/sakila_backend/internal/model"
	"github.com/ISSuh/sakila_backend/internal/repository"
)

type FilmActorService interface {
	ActorByFilmId(filmId int) ([]*model.Actor, error)
	FilmByActorId(actorId int) ([]*model.Film, error)
}

type filmActorService struct {
	log                 logger.Logger
	filmActorRepository repository.FilmActorRepository
}

func NewFilmActorService(l logger.Logger, r repository.FilmActorRepository) FilmActorService {
	return &filmActorService{
		log:                 l,
		filmActorRepository: r,
	}
}

func (s *filmActorService) ActorByFilmId(actorId int) ([]*model.Actor, error) {
	return s.filmActorRepository.ActorByFilmId(actorId)
}

func (s *filmActorService) FilmByActorId(filmId int) ([]*model.Film, error) {
	return s.filmActorRepository.FilmByActorId(filmId)
}
