package model

import "time"

type Payment struct {
	PaymentId int `gorm:"primaryKey"`

	Customer *Customer
	Staff    *Staff
	Rental   *Rental

	Amount      int
	PaymentDate time.Time

	LastUpdate time.Time
}

func (Payment) TableName() string {
	return "payment"
}
