package routes

import (
	"github.com/gin-gonic/gin"
	"librarymanagement/controllers"
)

func RegisterRoutes(r *gin.Engine) {
	// Public routes
	auth := r.Group("/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
	}

	// Protected routes (You can add middleware later)
	api := r.Group("/api")
	{
		// Admin routes
		adminRoutes := api.Group("/admins")
		{
			adminRoutes.GET("/", controllers.GetAllAdmins)
			adminRoutes.PUT("/:id", controllers.UpdateAdmin)
			adminRoutes.DELETE("/:id", controllers.DeleteAdmin)
		}

		// Student routes
		studentRoutes := api.Group("/students")
		{
			studentRoutes.POST("/", controllers.CreateStudent)
			studentRoutes.GET("/", controllers.GetAllStudents)
			studentRoutes.GET("/:id", controllers.GetStudentByID)
			studentRoutes.PUT("/:id", controllers.UpdateStudent)
			studentRoutes.DELETE("/:id", controllers.DeleteStudent)
		}

		// Teacher routes
		teacherRoutes := api.Group("/teachers")
		{
			teacherRoutes.POST("/", controllers.CreateTeacher)
			teacherRoutes.GET("/", controllers.GetAllTeachers)
			teacherRoutes.GET("/:id", controllers.GetTeacherByID)
			teacherRoutes.PUT("/:id", controllers.UpdateTeacher)
			teacherRoutes.DELETE("/:id", controllers.DeleteTeacher)
		}

		// Book routes
		bookRoutes := api.Group("/books")
		{
			bookRoutes.POST("/", controllers.CreateBook)
			bookRoutes.GET("/", controllers.GetAllBooks)
			bookRoutes.GET("/:id", controllers.GetBookByID)
			bookRoutes.PUT("/:id", controllers.UpdateBook)
			bookRoutes.DELETE("/:id", controllers.DeleteBook)
		}
	}
}
