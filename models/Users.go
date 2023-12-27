package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID     uint   `gorm:"not null;primaryKey"`
	Nombre string `gorm:"not null"`
	Rut    string `gorm:"not null;unique_index"`
	Email  string `gorm:"not null;unique_index"`
	Tipo   string `gorm:"not null"`
	Cart   Cart
}
