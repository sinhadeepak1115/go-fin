package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sinhadeepak1115/personal-finance/api"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})
	router.POST("/api/auth/signup", api.SignupUser)
	// router.POST("/api/auth/signin", api.PostUser)
	router.Run(":3000")
}
