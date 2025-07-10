package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sinhadeepak1115/personal-finance/config"
	"github.com/sinhadeepak1115/personal-finance/models"
)

type UserSignup struct {
	Email    string `json:"email" binding:"required,email"`
	Name     string `json:"name" `
	Password string `json:"password" binding:"required,min=6"`
}

type UserSingin struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// Get user
func GetAllUser(c *gin.Context) {
	var users []models.User
	result := config.DB.Find(&users)
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

// Register User
func SignupUser(c *gin.Context) {
	var req UserSignup
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}
	// Check the user existence
	var existingUser models.User
	if err := config.DB.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
		return
	}

	// Bycrypt Password

	// Create new User
	user := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	response := gin.H{
		"message": "User registered successfully",
		"received": gin.H{
			"email": req.Email,
			"name":  req.Name,
		},
	}
	c.JSON(http.StatusCreated, response)
}

// Login User
func SigninUser(c *gin.Context) {
	var req UserSingin
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}
	// Compare the user password with the database

	var user models.User
	if err := config.DB.Where("email = ? AND password = ?", req.Email, req.Password).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	response := gin.H{
		"message": "User logged in successfully",
		"user": gin.H{
			"email":    req.Email,
			"password": req.Password,
		},
	}
	c.JSON(http.StatusOK, response)
}
