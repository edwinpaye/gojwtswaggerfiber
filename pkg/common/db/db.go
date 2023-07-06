package db

import (
	"log"

	"github.com/edwinpaye/gots/pkg/common/models"
	"github.com/edwinpaye/gots/pkg/cars"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(url string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Book{})
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&cars.Car{})

	return db
}
