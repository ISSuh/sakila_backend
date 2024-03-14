package model

import "time"

// type Location struct {
// 	X float32 `json:"x"`
// 	Y float32 `json:"y"`
// }

type Address struct {
	AddressId  int     `grom:"primaryKey;type:smallint unsigned"`
	Address    string  `gorm:"type:varchar(50)"`
	Address2   string  `gorm:"type:varchar(50)"`
	District   string  `gorm:"type:varchar(20)"`
	City       *Cities `gorm:"foreignKey:CityId"`
	PostalCode string  `gorm:"type:varchar(10)"`
	Phone      string  `gorm:"type:varchar(20)"`
	// Location   Location
	LastUpdate time.Time
}

func (Address) TableName() string {
	return "address"
}
