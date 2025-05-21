package controllers

import (
	"net/http"
	"librarymanagement/config"
	"librarymanagement/models"
	"github.com/gin-gonic/gin"
)

func CreateTeacher(c *gin.Context) {
	var teacher models.Teacher
	if err := c.ShouldBindJSON(&teacher); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Create(&teacher).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create teacher"})
		return
	}
	c.JSON(http.StatusCreated, teacher)
}

func GetAllTeachers(c *gin.Context) {
	var teachers []models.Teacher
	if err := config.DB.Find(&teachers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve teachers"})
		return
	}
	c.JSON(http.StatusOK, teachers)
}

func GetTeacherByID(c *gin.Context) {
	id := c.Param("id")
	var teacher models.Teacher
	if err := config.DB.First(&teacher, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Teacher not found"})
		return
	}
	c.JSON(http.StatusOK, teacher)
}

func UpdateTeacher(c *gin.Context) {
	id := c.Param("id")
	var teacher models.Teacher
	if err := config.DB.First(&teacher, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Teacher not found"})
		return
	}
	if err := c.ShouldBindJSON(&teacher); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Save(&teacher).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update teacher"})
		return
	}
	c.JSON(http.StatusOK, teacher)
}

func DeleteTeacher(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.Teacher{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete teacher"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Teacher deleted"})
}
