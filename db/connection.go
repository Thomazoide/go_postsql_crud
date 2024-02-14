package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DSN string = "host=localhost user=thomas password=Thom1232 dbname=database port=5432"

var DB *gorm.DB

func ConectarDB() *gorm.DB {
	var error error
	DB, error := gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if error != nil {
		log.Println("ERROR AL CONECTAR BBDD...")
		log.Fatal(error)
		return nil
	}
	log.Println("Conectado a DB...")
	return DB
}
