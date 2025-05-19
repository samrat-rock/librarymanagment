package routes

import (
	"github.com/gin-gonic/gin"
	"librarymanagement/controllers"
	"librarymanagement/middleware"
)

func RegisterRoutes(r *gin.Engine) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
	}

	api := r.Group("/api")
	api.Use(middlewares.JWTAuthMiddleware())
	{
		// api.GET("/books", controllers.GetBooks)
		// api.POST("/books", middlewares.AdminOnlyMiddleware(), controllers.CreateBook)

	}
}
