package api

import (
	"database/sql"

	"go_plants/internal/repositories"
	"go_plants/internal/services"

	"github.com/gin-gonic/gin"
)

func RegisterUserAPI(r *gin.Engine, db *sql.DB) {
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := NewUserHandler(userService)

	userRoutes := r.Group("/users") // Base route for users
	{
		userRoutes.GET("/:id", userHandler.GetUser)       // Get user by ID
		userRoutes.GET("/all", userHandler.GetByAllUsers) // Get all users
		userRoutes.POST("/creat", userHandler.CreateUser) // Create a new user
		userRoutes.PUT("/:id", userHandler.UpdateUser)    // Update user by ID
		userRoutes.DELETE("/:id", userHandler.DeleteUser) // Delete user by ID
	}
}
