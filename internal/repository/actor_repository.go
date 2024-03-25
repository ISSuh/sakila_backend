package repository

import (
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
	if err := e.Model(&actor).Find(&actor, actorId).Error; err != nil {
		return nil, err
	}
	return actor, nil
}
