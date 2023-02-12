package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"assignment1GO/models"
)

func Connect() {
	connection, err := gorm.Open(postgres.Open("postgres://postgres:12345@localhost:5432/postgres"), &gorm.Config{})
	if err != nil {
		panic("Could not connect to the db")
	}
	connection.AutoMigrate(&models.User{})
}
