package config

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(mysql.Open("root:FleetEasy123@tcp(localhost:3306)/atm"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		panic("Please check DB Connection")
	}
	DB = connection
	log.Println("DB Connected")
}
