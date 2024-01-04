package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ProductID uint `gorm:"unique"`
	Nombre    string
	Categoria string
	Precio    float64
	Stock     int
}
