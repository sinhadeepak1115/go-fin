package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	db, err := gorm.Open(postgres.Open("postgresql://admin:password@localhost:5432/go-db"))
	if err != nil {
		panic(err)
	}
	DB = db
}
