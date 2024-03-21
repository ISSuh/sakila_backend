package model

import "time"

type Store struct {
	StoreId        int      `gorm:"primaryKey" json:"store_id"`
	ManagerStaffId int      `json:"-"`
	ManagerStaff   *Staff   `gorm:"foreignKey:ManagerStaffId;references:StaffId" json:"manger_staff,omitempty"`
	AddressId      int      `json:"-"`
	Address        *Address `gorm:"foreignKey:AddressId;references:AddressId" json:"address,omitempty"`

	LastUpdate time.Time `json:"-"`
}

func (Store) TableName() string {
	return "store"
}
