package routes

import (
	"github.com/gin-gonic/gin"
	"librarymanagement/controllers"
	
)

func RegisterRoutes(r *gin.Engine) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
	}

	api := r.Group("/api")
	// REMOVE middleware here
	{
		adminRoutes := api.Group("/admins")
		// REMOVE middleware here
		{
			adminRoutes.GET("/", controllers.GetAllAdmins)
			adminRoutes.PUT("/:id", controllers.UpdateAdmin)
			adminRoutes.DELETE("/:id", controllers.DeleteAdmin)
		}
	}
}

