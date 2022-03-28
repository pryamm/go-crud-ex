package model

import "gorm.io/gorm"

//entity
type Role struct {
	gorm.Model
	ID   uint   `gorm:"not null"`
	Name string `gorm:"size:255;not null"`
}

//response
type RoleResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
