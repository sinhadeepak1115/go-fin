package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserSignup struct {
	Email    string `json:"email" binding:"required,email"`
	Name     string `json:"name" `
	Password string `jsin:"password" binding:"required,min=6"`
}

type UserSingin struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `jsin:"password" binding:"required,min=6"`
}

func SignupUser(c *gin.Context) {
	var req UserSignup
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}
	response := gin.H{
		"received": gin.H{
			"email": req.Email,
			"name":  req.Name,
		},
		"message": "User registered successfully",
	}
	c.JSON(http.StatusCreated, response)
}

func SigninUser(c *gin.Context) {
	var req UserSingin
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}
	response := gin.H{
		"received": gin.H{
			"email":    req.Email,
			"password": req.Password,
		},
		"message": "User registered successfully",
	}
	c.JSON(http.StatusOK, response)
}
