package controllers

import (
	"librarymanagement/config"
	"librarymanagement/models"
	

	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt/v4"
	"os"
	"time"
)

func RegisterStudent(c *gin.Context) {
	var student models.StudentRegister

	
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(student.Password), 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	
	newStudent := models.Student{
		FirstName:   student.FirstName,
		LastName:    student.LastName,
		Email:       student.Email,
		Password:    string(hashedPassword),
		Phone:       student.Phone,
		ClassNumber: student.ClassNumber,
		RollNo:      student.RollNo,
	}


	if err := config.DB.Create(&newStudent).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register student"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Student registered successfully"})
}





func LoginStudent(c *gin.Context) {
	var creds models.StudentLogin

	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var student models.Student
	if err := config.DB.Where("email = ?", creds.Email).First(&student).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(student.Password), []byte(creds.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    student.ID,
		"email": student.Email,
		"role":  "student",
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	})

	secret := os.Getenv("JWT_SECRET")
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}



func GetOwnProfile(c *gin.Context) {
	email, _ := c.Get("email")

	var student models.Student
	if err := config.DB.Where("email = ?", email).First(&student).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}

	student.Password = "" 
	c.JSON(http.StatusOK, student)
}

func UpdateOwnProfile(c *gin.Context) {
	email, _ := c.Get("email")

	var input models.StudentUpdate
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var student models.Student
	if err := config.DB.Where("email = ?", email).First(&student).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}

	student.FirstName = input.FirstName
	student.LastName = input.LastName
	student.Phone = input.Phone
	student.ClassNumber = input.ClassNumber
	student.RollNo = input.RollNo

	if err := config.DB.Save(&student).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update student profile"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully"})
}