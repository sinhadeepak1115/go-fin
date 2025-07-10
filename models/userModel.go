package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id           int           `json:"id" gorm:"primaryKey;autoIncrement"`
	Name         string        `json:"name" gorm:"not null"`
	Email        string        `json:"email" gorm:"not null;unique"`
	Password     string        `json:"password" gorm:"not null"`
	Transactions []Transaction `json:"transactions" gorm:"foreignKey:UserId"`
}
