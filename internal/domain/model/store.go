package model

import "time"

type Store struct {
	StoreId      int `gorm:"primaryKey"`
	StaffFKId    int
	ManagerStaff *Staff `gorm:"foreignKey:StaffFKId;references:StaffId"`
	AddressFKId  int
	Address      *Address `gorm:"foreignKey:AddressFKId;references:AddressId"`

	LastUpdate time.Time
}

func (Store) TableName() string {
	return "store"
}
