package model

import "time"

type Store struct {
	StoreId int `gorm:"primaryKey"`

	ManagerStaff *Staff   `gorm:"foreignKey:StaffId"`
	Address      *Address `gorm:"foreignKey:AddressId"`

	LastUpdate time.Time
}

func (Store) TableName() string {
	return "store"
}
