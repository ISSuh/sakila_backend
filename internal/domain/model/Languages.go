package model

import "time"

type Languages struct {
	LanguageId int    `gorm:"primaryKey"`
	Name       string `gorm:"type:char(20)"`
	LastUpdate time.Time
}

func (Languages) TableName() string {
	return "language"
}
