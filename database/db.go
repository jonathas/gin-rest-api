package database

import (
	"log"

	"github.com/jonathas/gin-rest-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
	err error
)

func Connect() {
	dsn := "host=localhost user=root password=root dbname=students port=5432 sslmode=disable TimeZone=Europe/Prague"
	DB, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Panic("Failed to connect to database!")
	}

	DB.AutoMigrate(&models.Student{})
}
