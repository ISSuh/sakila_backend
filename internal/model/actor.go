package model

import (
	"time"
)

type Actor struct {
	ActorId   int     `gorm:"primaryKey" json:"actor_id"`
	FirstName string  `json:"firt_name,omitempty"`
	LastName  string  `json:"last_name,omitempty"`
	Films     []*Film `gorm:"many2many:film_actor;joinForeignKey:actor_id;" json:"films,omitempty"`

	LastUpdate time.Time `json:"_"`
}

func (Actor) TableName() string {
	return "actor"
}
