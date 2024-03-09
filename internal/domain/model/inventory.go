package model

import "time"

type Inventory struct {
	InventoryId int `gorm:"primaryKey"`

	Film  *Film
	Store *Store

	LastUpdate time.Time
}

func (Inventory) TableName() string {
	return "inventory"
}
