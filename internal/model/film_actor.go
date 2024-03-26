package model

import "time"

type FilmActor struct {
	ActorId int
	Actor   []*Actor `gorm:"foreignKey:ActorId;references:ActorId"`
	FilmId  int
	Film    []*Film `gorm:"foreignKey:FilmId;references:FilmId"`

	LastUpdate time.Time
}

func (FilmActor) TableName() string {
	return "film_actor"
}
