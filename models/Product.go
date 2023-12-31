package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Nombre    string
	Categoria string
	Precio    float64
	Stock     int
}
