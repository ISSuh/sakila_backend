package repository

import (
	"github.com/ISSuh/sakila_backend/internal/logger"
	"github.com/ISSuh/sakila_backend/internal/model"
	"github.com/ISSuh/sakila_backend/pkg/db"
)

type ActorRepository interface {
	ActorById(actorId int) (*model.Actor, error)
	ActorWithFilmById(actorId int) (*model.Actor, error)
	ActorByLastName(lastName string) ([]*model.Actor, error)
}

type actorRepository struct {
	log logger.Logger
	db  *db.Database
}

func NewActorRepository(l logger.Logger, d *db.Database) ActorRepository {
	return &actorRepository{
		log: l,
		db:  d,
	}
}

func (r *actorRepository) ActorById(actorId int) (*model.Actor, error) {
	e := r.db.Engine()

	actor := new(model.Actor)
	if err := e.
		Model(actor).
		Where("actor_id=?", actorId).
		Find(actor).Error; err != nil {
		return nil, err
	}

	return actor, nil
}

func (r *actorRepository) ActorWithFilmById(actorId int) (*model.Actor, error) {
	e := r.db.Engine()

	actor := new(model.Actor)
	if err := e.
		Model(actor).
		Preload("FilmsActor.Film").
		Where("actor.actor_id=?", actorId).
		Find(actor).Error; err != nil {
		return nil, err
	}

	films := []*model.Film{}
	for _, filmActor := range actor.FilmsActor {
		films = append(films, filmActor.Film...)
	}

	actor.Films = films
	return actor, nil
}

func (r *actorRepository) ActorByLastName(lastName string) ([]*model.Actor, error) {
	e := r.db.Engine()

	actor := []*model.Actor{}
	if err := e.
		Model(&actor).
		Where("last_name=?", lastName).
		Find(&actor).Error; err != nil {
		return nil, err
	}

	return actor, nil
}
