package model

import "time"

type Actor struct {
	ActorId   int    `gorm:"primaryKey;type:smallint unsigned"`
	FirstName string `gorm:"type:varchar(45)"`
	LastName  string `gorm:"type:varchar(45)"`

	LastUpdate time.Time
}

func (Actor) TableName() string {
	return "actor"
}
