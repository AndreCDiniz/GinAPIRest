package database

import (
	"log"

	"github.com/AndreCDiniz/GinAPIRest/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectToDatabase() {
	connectionString := "host=localhost user=root password=root dbname=root port=5432"
	DB, err = gorm.Open(postgres.Open(connectionString))
	if err != nil {
		log.Panic("Erro to connect to database")
	}
	DB.AutoMigrate(&models.Pizza{})
}
