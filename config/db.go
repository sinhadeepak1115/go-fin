package config

import (
	"github.com/sinhadeepak1115/personal-finance/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	db, err := gorm.Open(postgres.Open("postgresql://admin:password@localhost:5432/go-db"))
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.User{})
	DB = db
}
