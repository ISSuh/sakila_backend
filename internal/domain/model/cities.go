package model

import "time"

type Cities struct {
	CityId int `gorm:"primaryKey"`
	City   string
	// Country *Countries `gorm:"foreignKey:CountryId"`

	LastUpdate time.Time
}

func (Cities) TableName() string {
	return "city"
}
