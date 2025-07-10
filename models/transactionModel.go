package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	Id          int     `json:"id" gorm:"primaryKey;autoIncrement"`
	UserId      int     `json:"user_id" gorm:"not null"`
	Amount      float64 `json:"amount" gorm:"not null"`
	Type        string  `json:"type" gorm:"not null"`
	Category    string  `json:"category" gorm:"not null"`
	Description string  `json:"description" gorm:"not null"`
}
