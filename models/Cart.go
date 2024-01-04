package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	UserID    uint      `gorm:"unique"`
	Productos []Product `gorm:"foreignKey:ProductID"`
	Total     float64
}
