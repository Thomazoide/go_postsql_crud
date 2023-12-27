package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	ID       uint `gorm:"primaryKey; not null"`
	Products []Product
	Total    float32
}
