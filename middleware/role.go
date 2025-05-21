package middleware

import (
	
	"net/http"
	

	"github.com/gin-gonic/gin"
)

func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists || role != "admin" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Admins only"})
			return
		}
		c.Next()
	}
}

func StudentOrTeacherOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists || (role != "student" && role != "teacher") {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Access restricted to students or teachers"})
			return
		}
		c.Next()
	}
}
