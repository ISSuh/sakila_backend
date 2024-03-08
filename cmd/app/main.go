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

type Countries struct {
	CountryId  int      `gorm:"primaryKey;type:smallint unsigned"`
	Country    string   `gorm:"type:varchar(50)"`
	City       []Cities `gorm:"foreignKey:CityId"`
	LastUpdate time.Time
}

func (Countries) TableName() string {
	return "country"
}

type Cities struct {
	CityId int    `gorm:"primaryKey;type:smallint unsigned"`
	City   string `gorm:"type:varchar(50)"`
	// CountryId  incountry_idt
	Country    Countries `gorm:"foreignKey:CountryId"`
	LastUpdate time.Time
}

func (Cities) TableName() string {
	return "city"
}

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", " ")
	return string(s)
}

func main() {
	dsn := "root:1@tcp(127.0.0.1:33551)/temp?charset=utf8mb4&parseTime=True&loc=Local"
	dialector := mysql.Open(dsn)
	config := gorm.Config{}
	db, err := gorm.Open(dialector, &config)
	if err != nil {
		fmt.Printf("%s", err)
	}

	db.AutoMigrate(
		&Actor{},
		&Cities{},
	)

	// 	var a Actor
	// 	db.Where("actor_id=?", 1).Find(&a)

	// 	fmt.Printf("%s\n", prettyPrint(a))

	// 	var country Countries
	// 	err = db.Model(&Countries{}).Preload("City").Where("country_id=?", 82).Find(&country).Error

	// 	if err != nil {
	// 		fmt.Printf("%s", err.Error())
	// 	}

	// 	fmt.Printf("%s\n", prettyPrint(country))

	// 	var cities []Cities
	// 	err = db.Model(&Cities{}).Preload("Country").Where("country_id=?", 82).Find(&cities).Error

	// 	if err != nil {
	// 		fmt.Printf("%s", err.Error())
	// 	}

	// fmt.Printf("%s\n", prettyPrint(cities))
}
