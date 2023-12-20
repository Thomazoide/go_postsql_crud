package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConnect() {
	var error error
	DSN := "host=localhost port=5432 user=thomas password=Thom1232 dbname=bbdd sslmode=disable"
	DB, error = gorm.Open((postgres.Open(DSN)), &gorm.Config{})
	if error != nil {
		log.Fatal(error)
	} else {
		log.Println("Base de datos conectada...")
	}
}
