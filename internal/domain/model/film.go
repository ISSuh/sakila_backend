package model

import "time"

type Film struct {
	FilmId int `gorm:"primaryKey"`

	Title           string `gorm:"type:varchar(50)"`
	Description     string
	ReleaseYear     time.Time
	Laguage         Languages
	OriginalLaguage Languages

	RentalDuration  int
	RentalRate      int
	Length          int
	ReplacementCost int

	Rating          int
	SpecialFeatures string

	LastUpdate time.Time
}

func (Film) TableName() string {
	return "film"
}
