package model

import "time"

type Cities struct {
	CityId    int        `gorm:"primaryKey" json:"city_id"`
	City      string     `json:"city,omitempty"`
	CountryId int        `json:"-"`
	Country   *Countries `gorm:"foreignKey:CountryId;references:CountryId" json:"country,omitempty"`

	LastUpdate time.Time `json:"-"`
}

func (Cities) TableName() string {
	return "city"
}
