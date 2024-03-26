package repository

import (
	"github.com/ISSuh/sakila_backend/internal/logger"
	"github.com/ISSuh/sakila_backend/internal/model"
	"github.com/ISSuh/sakila_backend/pkg/db"
)

type FilmActorRepository interface {
	ActorByFilmId(filmId int) ([]*model.Actor, error)
	FilmByActorId(actorId int) ([]*model.Film, error)
}

type filmActorRepository struct {
	log logger.Logger
	db  *db.Database
}

func NewFilmActorRepository(l logger.Logger, d *db.Database) FilmActorRepository {
	return &filmActorRepository{
		log: l,
		db:  d,
	}
}

func (r *filmActorRepository) ActorByFilmId(filmId int) ([]*model.Actor, error) {
	// e := r.db.Engine()

	// actors := []*model.Actor{}
	// if err := e.Model(&actors).Find(&actors, actorId).Error; err != nil {
	// 	return nil, err
	// }
	return nil, nil
}

func (r *filmActorRepository) FilmByActorId(actorId int) ([]*model.Film, error) {
	// e := r.db.Engine()
	return nil, nil
}
