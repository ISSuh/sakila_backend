package model

import (
	"time"
)

type Actor struct {
	ActorId    int          `gorm:"primaryKey" json:"actor_id"`
	FirstName  string       `json:"firt_name,omitempty"`
	LastName   string       `json:"last_name,omitempty"`
	FilmsActor []*FilmActor `json:"-"`
	Films      []*Film      `gorm:"-" json:"films,omitempty"`

	LastUpdate time.Time `json:"-"`
}

func (Actor) TableName() string {
	return "actor"
}
