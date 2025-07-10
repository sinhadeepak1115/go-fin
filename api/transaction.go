package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sinhadeepak1115/personal-finance/config"
	"github.com/sinhadeepak1115/personal-finance/models"
)

type TransactionRequest struct {
	Amount      float64 `json:"amount" binding:"required"`
	UserId      int     `json:"user_id" binding:"required"`
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
func GetAllTransactions(c *gin.Context) {
	var transactions []models.Transaction
	result := config.DB.Find(&transactions)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve transactions"})
	}
	response := make([]gin.H, len(transactions))
	for i, transaction := range transactions {
		response[i] = gin.H{
			"Amount":      transaction.Amount,
			"UserId":      transaction.UserId,
			"Type":        transaction.Type,
			"Category":    transaction.Category,
			"Description": transaction.Description,
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"users": response,
		"count": len(response),
	})
}

func PostTransaction(c *gin.Context) {
	var req TransactionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request data"})
		return
	}
	transaction := models.Transaction{
		Amount:      req.Amount,
		UserId:      req.UserId,
		Type:        req.Type,
		Category:    req.Category,
		Description: req.Description,
	}
	if err := config.DB.Create(&transaction).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to create transaction"})
		return
	}
	response := gin.H{
		"message": "Transaction created successfully",
		"user": gin.H{
			"Amount":      transaction.Amount,
			"UserId":      transaction.UserId,
			"Type":        transaction.Type,
			"Category":    transaction.Category,
			"Description": transaction.Description,
		},
	}
	c.JSON(201, response)
}
