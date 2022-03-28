package main

import (
	"backend-c-payment-monitoring/config"
	"backend-c-payment-monitoring/model"
	"log"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestMigration(t *testing.T) {
	//setup configuration
	db := config.DB

	db.AutoMigrate(&model.Role{}, &model.User{}, &model.Status{}, &model.Payment{}, &model.Authentication{})
}

func TestRoleSeeder(t *testing.T) {
	//setup configuration
	db := config.DB

	var roles = []model.Role{
		{Name: "Admin"},
		{Name: "Unit Kerja (Customer)"},
		{Name: "General Support"},
		{Name: "Accounting"},
	}

	err := db.Create(&roles).Error

	if err != nil {
		log.Println("Role Seed Failed")
	}

	log.Println("Role Seed Success")
}

func TestUserSeeder(t *testing.T) {
	//setup configuration
	db := config.DB

	passwordHash, err := bcrypt.GenerateFromPassword([]byte("12345678"), bcrypt.MinCost)

	var user = model.User{
		Name:     "Admin",
		Username: "admin",
		Password: string(passwordHash),
		RoleID:   1,
	}

	err = db.Create(&user).Error

	if err != nil {
		log.Println("User Seed Failed")
	}

	log.Println("User Seed Success")

}

func TestStatusSeeder(t *testing.T) {
	//setup configuration
	db := config.DB

	var status = []model.Status{
		{Name: "Menunggu Konfirmasi"},
		{Name: "Reject by GS"},
		{Name: "Diteruskan ke Accounting"},
		{Name: "Rejected by AC"},
		{Name: "Disetujui AC"},
	}

	err := db.Create(&status).Error

	if err != nil {
		log.Println("Status Seed Failed")
	}

	log.Println("Status Seed Success")
}
