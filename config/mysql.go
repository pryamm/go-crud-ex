package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func NewMysqlDatabase(configuration Config) {
	var err error
	host := configuration.Get("DB_HOST")
	username := configuration.Get("DB_USERNAME")
	password := configuration.Get("DB_PASSWORD")
	port := configuration.Get("DB_PORT")
	database := configuration.Get("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, database)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
}
