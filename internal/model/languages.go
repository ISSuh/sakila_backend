package model

import "time"

type Languages struct {
	LanguageId int    `gorm:"primaryKey" json:"-"`
	Name       string `json:"name,omitempty"`

	LastUpdate time.Time `json:"-"`
}

func (Languages) TableName() string {
	return "language"
}
