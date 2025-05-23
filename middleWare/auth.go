package middleware

import (
	"librarymanagement/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AdminAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}

		
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := utils.VerifyToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

	
		if claims.Role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Admin access only"})
			c.Abort()
			return
		}

		c.Set("email", claims.Email)
		c.Set("role", claims.Role)
		c.Next()
	}
}
