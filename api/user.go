package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserSignup struct {
	Email    string `json:"email" binding:"required,email"`
	Name     string `json:"name" binding:"required,nonzero"`
	Password string `jsin:"password" binding:"required,min=6"`
}

type UserSingin struct {
	Email string `json:"email" binding:"required,email"`
}

func SignupUser(c *gin.Context) {
	var req UserSignup
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}
	ressponse := gin.H{
		"received": gin.H{
			"email": req.Email,
			"name":  req.Name,
		},
		"message": "User registered successfully",
	}
	c.JSON(http.StatusCreated, ressponse)
}
