package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	products []Product
	total    float32
}
