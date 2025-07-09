package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sinhadeepak1115/personal-finance/models"
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

func GetAllUser(c *gin.Context) {
	var users []models.User
	result := models.DB.Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
		return
	}
	response := make([]gin.H, len(users))
	for i, user := range users {
		response[i] = gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"users": response,
		"count": len(users),
	})
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
