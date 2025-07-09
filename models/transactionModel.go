package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	Id          int     `json:"id" gorm:"primaryKey;autoIncrement"`
	Amount      float64 `json:"amount" gorm:"not null"`
	Type        string  `json:"type" gorm:"not null"`     // "income" or "expense"
	Category    string  `json:"category" gorm:"not null"` // e.g., "groceries", "salary"
	Description string  `json:"description" gorm:"not null"`
	CreatedAt   string  `json:"created_at" gorm:"not null"` // ISO 8601 format
}

// {
// "id": 1,
// "amount": 150.50,
// "type": "expense",
// "category": "groceries",
// "description": "Weekly shopping",
// "created_at": "2023-07-20T10:00:00Z"
// },
