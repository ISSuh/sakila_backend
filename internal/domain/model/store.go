package model

import "time"

type Store struct {
	StoreId int `gorm:"primaryKey"`

	ManagerStaff *Staff
	Address      Address

	LastUpdate time.Time
}

func (Store) TableName() string {
	return "store"
}
