package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type Film struct {
	FilmId             int             `gorm:"primaryKey" json:"id"`
	Title              string          `json:"title"`
	Description        string          `json:"description,omitempty"`
	ReleaseYear        Year            `json:"release_year,omitempty"`
	LanguageId         int             `json:"-"`
	Language           *Languages      `gorm:"foreignKey:LanguageId;references:LanguageId" json:"laguage,omitempty"`
	OriginalLanguageId int             `json:"-"`
	OriginalLanguage   *Languages      `gorm:"foreignKey:OriginalLanguageId;references:LanguageId" json:"original_laguage,omitempty"`
	RentalDuration     int             `json:"rental_duration"`
	RentalRate         decimal.Decimal `json:"rental_rate"`
	Length             int             `json:"length,omitempty"`
	ReplacementCost    decimal.Decimal `json:"replacement_cost"`
	Rating             FilmRating      `json:"rating,omitempty"`
	SpecialFeatures    SpecialFeatures `json:"special_features,omitempty"`

	LastUpdate time.Time `json:"-"`
}

func (Film) TableName() string {
	return "film"
}
