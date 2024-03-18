package model

import "time"

type Cities struct {
	CityId      int `gorm:"primaryKey"`
	City        string
	CountryFKId int
	Country     *Countries `gorm:"foreignKey:CountryFKId;references:CountryId"`

	LastUpdate time.Time
}

func (Cities) TableName() string {
	return "city"
}
