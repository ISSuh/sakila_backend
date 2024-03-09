package model

import "time"

type Countries struct {
	CountryId  int      `gorm:"primaryKey;type:smallint unsigned"`
	Country    string   `gorm:"type:varchar(50)"`
	City       []Cities `gorm:"foreignKey:CountryId"`
	LastUpdate time.Time
}

func (Countries) TableName() string {
	return "country"
}
