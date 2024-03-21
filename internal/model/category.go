package model

import "time"

type Category struct {
	CategoryId int    `gorm:"primaryKey"`
	Name       string `gorm:"type:varchar(25)"`
	LastUpdate time.Time
}

func (Category) TableName() string {
	return "category"
}
