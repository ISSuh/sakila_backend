package model

import "time"

type Customer struct {
	CustomerId int `gorm:"primaryKey"`

	LastUpdate time.Time
}

func (Customer) TableName() string {
	return "customer"
}
