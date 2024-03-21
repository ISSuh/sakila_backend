package model

import "time"

type Countries struct {
	CountryId int      `gorm:"primaryKey" json:"country_id"`
	Country   string   `json:"country,omitempty"`
	CityId    int      `json:"-"`
	City      []Cities `gorm:"foreignKey:CityId;references:CityId" json:"city,omitempty"`

	LastUpdate time.Time `json:"-"`
}

func (Countries) TableName() string {
	return "country"
}
