package repository

import (
	"fmt"

	"github.com/ISSuh/sakila_backend/internal/logger"
	"github.com/ISSuh/sakila_backend/internal/model"
	"github.com/ISSuh/sakila_backend/pkg/db"
)

type ActorRepository interface {
	ActorById(actorId int) (*model.Actor, error)
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
	err := e.Set("gorm:auto_preload", true).
		Joins("JOIN film_actor on film_actor.actor_id=actor.actor_id").
		Joins("JOIN film on film.film_id=film_actor.film_id").
		Where("actor.actor_id=?", actorId).
		Preload("Films").
		Find(&actor).Error
	if err != nil {
		return nil, err
	}

	fmt.Printf("[TEST] %v\n", *actor)

	return actor, nil
}
