package controllers

import (
	"librarymanagement/config"
	"librarymanagement/models"
	"librarymanagement/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var input models.AdminRegister
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Role == "admin" {
		var count int64
		config.DB.Model(&models.Admin{}).Where("role = ?", "admin").Count(&count)
		if count >= 2 {
			c.JSON(http.StatusForbidden, gin.H{"error": "Only two admins allowed"})
			return
		}
	}

	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	admin := models.Admin{
		Username: input.Username,
		Email:    input.Email,
		Password: hashedPassword,
		Role:     "admin",
	}

	if err := config.DB.Create(&admin).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not register admin"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Admin registered successfully"})
}

func Login(c *gin.Context) {
	var input models.AdminLogin
	var dbUser models.Admin

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Where("email = ?", input.Email).First(&dbUser).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	if !utils.CheckPasswordHash(input.Password, dbUser.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	token, _ := utils.GenerateToken(dbUser.ID, dbUser.Email, dbUser.Role)
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
	})
}
