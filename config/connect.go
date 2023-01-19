package config

import (
	"github.com/prashar32/rest-api-1/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=hotel password=hotel dbname=hotel port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.Hotel{})
	db.AutoMigrate(&models.Customer{})
	db.AutoMigrate(&models.Bookings{})
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Shop{})
	DB = db
}
