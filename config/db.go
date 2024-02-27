package config

import (
	"github.com/afanasyevadina/go-test/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	db, err := gorm.Open(postgres.Open("host=localhost user=app password=!ChangeMe! dbname=app port=5432 sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = db
	DB.AutoMigrate(&models.Task{})
}
