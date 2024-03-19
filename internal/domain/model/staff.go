package model

import "time"

type Staff struct {
	StaffId   int      `gorm:"primaryKey" json:"staff_id"`
	FirstName string   `json:"first_name,omitempty"`
	LastName  string   `json:"last_name,omitempty"`
	AddressId int      `json:"address_id,omitempty"`
	Address   *Address `gorm:"foreignKey:AddressId;references:AddressId" json:"address,omitempty"`
	StoreId   int      `json:"-"`
	Store     *Store   `gorm:"foreignKey:StoreId;references:StoreId"`
	Email     string   `json:"email,omitempty"`
	Active    int      `json:"active,omitempty"`
	Username  string   `json:"username,omitempty"`
	Password  string   `json:"-"`
	Picture   []byte   `json:"-"`

	LastUpdate time.Time `json:"-"`
}

func (Staff) TableName() string {
	return "staff"
}
