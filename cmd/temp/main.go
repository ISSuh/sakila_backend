package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type TestA struct {
	AId int `gorm:"primaryKey"`

	BFKId int
	B     *TestB `gorm:"foreignKey:BFKId;references:BId"`

	CFKId int
	C     *TestC `gorm:"foreignKey:CFFId;references:CId"`

	ValueA int
}

type TestB struct {
	BId    int `gorm:"primaryKey"`
	ValueB int
}

type TestC struct {
	CId    int `gorm:"primarykey"`
	DFKId  int
	D      *TestD `gorm:"foreignkey:DFKId;references:DId"`
	ValueC int
}

type TestD struct {
	DId    int `gorm:"primarykey"`
	ValueD int

	// CFKId int
	// C     *TestC `gorm:"foreignkey:CFKId;references:CId"`
}

func main() {
	dsn := "root:1@tcp(127.0.0.1:33551)/TESTDB?charset=utf8mb4&parseTime=True&loc=Local"
	dialector := mysql.Open(dsn)
	config := gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}
	db, err := gorm.Open(dialector, &config)
	if err != nil {
		fmt.Printf("%s", err)
	}

	db.AutoMigrate(
		&TestA{},
		&TestB{},
		&TestC{},
		&TestD{})

	// d := []TestD{}
	// for i := 0; i < 10; i++ {
	// 	d = append(d, TestD{ValueD: i})

	// }
	// if err := db.Create(d).Error; err != nil {
	// 	fmt.Println(err)
	// }

	// // for i := 0; i < 10; i++ {
	// c := TestC{
	// 	D: &TestD{
	// 		ValueD: 100,
	// 	},
	// 	ValueC: 1,
	// }
	// if err := db.Create(&c).Error; err != nil {
	// 	fmt.Println(err)
	// }
	// // }

}
