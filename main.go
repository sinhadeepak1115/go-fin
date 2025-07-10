package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sinhadeepak1115/personal-finance/api"
	"github.com/sinhadeepak1115/personal-finance/config"
)

func main() {
	router := gin.Default()

	// Connecting the database
	config.ConnectDB()
	// Test Handler
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})
	// User Handler
	router.GET("/api", api.GetAllUser)
	router.POST("/api/auth/signup", api.SignupUser)
	router.POST("/api/auth/signin", api.SigninUser)

	// Transaction Handler
	router.POST("/api/transaction", api.PostTransaction)
	router.GET("/api/transaction", api.GetAllTransactions)
	router.Run(":3000")
}
