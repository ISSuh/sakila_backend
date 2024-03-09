package model

import "time"

type Staff struct {
	StaffId int `gorm:"primaryKey"`

	FirstName string
	LastName  string
	Address   *Address
	Picture   []byte
	Email     string
	Store     Store

	Active   int
	Username string
	Password string

	LastUpdate time.Time
}

func (Staff) TableName() string {
	return "staff"
}
