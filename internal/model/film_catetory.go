package model

import "time"

type FilmCategory struct {
	FilmId     int    `gorm:"primaryKey"`
	CategoryId string `gorm:"primaryKey"`
	LastUpdate time.Time
}

func (FilmCategory) TableName() string {
	return "film_category"
}
