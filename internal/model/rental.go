package model

import "time"

type Rental struct {
	RentalId int `gorm:"primaryKey"`

	RentalDate time.Time
	ReturnDate time.Time

	Customer *Customer
	Staff    *Staff

	LastUpdate time.Time
}

func (Rental) TableName() string {
	return "rental"
}
