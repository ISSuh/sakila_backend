package main

import (
	"encoding/json"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Actor struct {
	ActorId    int    `gorm:"primaryKey;type:smallint unsigned"`
	FirstName  string `gorm:"type:varchar(45)"`
	LastName   string `gorm:"type:varchar(45)"`
	LastUpdate time.Time
}

func (Actor) TableName() string {
	return "actor"
}

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", " ")
	return string(s)
}

func main() {
	dsn := "root:1@tcp(127.0.0.1:33551)/sakila?charset=utf8mb4&parseTime=True&loc=Local"
	dialector := mysql.Open(dsn)
	config := gorm.Config{}
	db, err := gorm.Open(dialector, &config)
	if err != nil {
		fmt.Printf("%s", err)
	}

	var a Actor
	db.Where("actor_id=?", 1).Find(&a)

	fmt.Printf("%s", prettyPrint(a))
}
