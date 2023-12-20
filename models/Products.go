package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	nombre string  `gorm:"not null"`
	precio float32 `gorm:"not null"`
}
