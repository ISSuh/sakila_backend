package model

import "time"

type Staff struct {
	StaffId int `gorm:"primaryKey"`

	FirstName string
	LastName  string
	Address   *Address `gorm:"foreignKey:AddressId"`
	// Picture   []byte
	Email string
	Store *Store `gorm:"foreignKey:StoreId"`

	Active   int
	Username string
	Password string

	LastUpdate time.Time
}

func (Staff) TableName() string {
	return "staff"
}
