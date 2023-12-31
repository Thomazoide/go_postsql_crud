package models

import "gorm.io/gorm"

type Usuario struct {
	gorm.Model
	Nombre   string
	Email    string
	Password string
	Carro    Cart `gorm:"default:null"`
}
