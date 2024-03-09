package model

import "time"

type FilmActor struct {
	Actor []Actor
	Film  []Film

	LastUpdate time.Time
}

func (FilmActor) TableName() string {
	return "city"
}
