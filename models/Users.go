package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	nombre string `gorm:"not null"`
	rut    string `gorm:"not null;unique_index"`
	email  string `gorm:"not null;unique_index"`
	tipo   string `gorm:"not null"`
	cart   Cart
}
