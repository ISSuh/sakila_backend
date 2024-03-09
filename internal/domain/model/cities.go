package model

import "time"

type Cities struct {
	CityId  int    `gorm:"primaryKey;type:smallint unsigned"`
	City    string `gorm:"type:varchar(50)"`
	Country Countries

	LastUpdate time.Time
}

func (Cities) TableName() string {
	return "city"
}