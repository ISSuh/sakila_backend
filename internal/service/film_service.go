package service

import (
	"github.com/ISSuh/monolith-sakila/internal/common"
	"github.com/ISSuh/monolith-sakila/internal/logger"
	"github.com/ISSuh/monolith-sakila/internal/model"
	"github.com/ISSuh/monolith-sakila/internal/repository"
	"github.com/ISSuh/monolith-sakila/pkg/db"
)

type FilmService interface {
	FilmById(storeId int) (*model.Film, error)
	Films(page *common.Pagnation) (*common.Pagnation, error)
}

type filmService struct {
	*db.Transaction

	log            logger.Logger
	filmRepository repository.FilmRepository
}

func NewFilmService(l logger.Logger, tx *db.Transaction, r repository.FilmRepository) FilmService {
	return &filmService{
		log:            l,
		Transaction:    tx,
		filmRepository: r,
	}
}

func (s *filmService) FilmById(id int) (*model.Film, error) {
	return s.filmRepository.FilmById(id)
}

func (s *filmService) Films(page *common.Pagnation) (*common.Pagnation, error) {
	return s.filmRepository.Films(page)
}
