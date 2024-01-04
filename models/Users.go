package models

import "gorm.io/gorm"

type Usuario struct {
	gorm.Model
	Nombre   string
	Tipo     string
	Email    string `gorm:"unique"`
	Password string
	Carro    Cart `gorm:"foreignKey:UserID"`
}
