package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	Productos []Product `gorm:"default:null"`
	Total     float64
}
