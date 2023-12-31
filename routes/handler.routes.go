package routes

import "gorm.io/gorm"

type SvResponse struct {
	Mensaje string
	Cuerpo  any
}

type handler struct {
	DB *gorm.DB
}

func NewHandler(db *gorm.DB) handler {
	return handler{db}
}
