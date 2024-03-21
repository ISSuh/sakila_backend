package model

import "time"

type Address struct {
	AddressId  int     `grom:"primaryKey" json:"address_id,omitempty"`
	Address    string  `json:"address,omitempty"`
	Address2   string  `json:"address2,omitempty"`
	District   string  `json:"district,omitempty"`
	CityId     int     `json:"-"`
	City       *Cities `gorm:"foreignKey:CityId;references:CityId" json:"city,omitempty"`
	PostalCode string  `json:"postal_code,omitempty"`
	Phone      string  `json:"phone,omitempty"`
	// Location   Location `json:"location,omitempty"`

	LastUpdate time.Time `json:"-"`
}

func (Address) TableName() string {
	return "address"
}
