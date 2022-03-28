package model

import "gorm.io/gorm"

//entity
type Status struct {
	gorm.Model
	ID   uint   `gorm:"not null"`
	Name string `gorm:"size:255;not null"`
}

//response
type StatusResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
