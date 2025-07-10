package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sinhadeepak1115/personal-finance/models"
)

type TransactionRequest struct {
	Amount      float64 `json:"amount" binding:"required"`
	Type        string  `json:"type" binding:"required,oneof=income expense"`
	Category    string  `json:"category" binding:"required"`
	Description string  `json:"description" binding:"required"`
}

// {
// "amount": 75.25,
// "type": "expense",
// "category": "dining",
// "description": "Dinner with friends"
// }

func PostTransaction(c *gin.Context) {
	var req TransactionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request data"})
		return
	}
	var existingTransaction models.Transaction
}
