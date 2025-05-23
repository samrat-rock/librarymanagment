package routes

import (
	"github.com/gin-gonic/gin"
	"librarymanagement/controllers"
	"librarymanagement/middleWare"
)

func RegisterRoutes(r *gin.Engine) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
		auth.POST("/logout", controllers.Logout)
	}

	
	adminRoutes := r.Group("/admins", middleware.AdminAuthMiddleware())
	{
		adminRoutes.GET("/", controllers.GetAllAdmins)
		adminRoutes.PUT("/:id", controllers.UpdateAdmin)
		adminRoutes.DELETE("/:id", controllers.DeleteAdmin)
	}
}
