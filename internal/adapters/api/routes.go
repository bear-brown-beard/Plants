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

func RegisterPlantAPI(r *gin.Engine, db *sql.DB) {
	plantRepo := repositories.NewPlantRepository(db)
	plantService := services.NewPlantService(plantRepo)
	plantHandler := NewPlantHandler(plantService)

	plantRoutes := r.Group("/plants") // Base route for plants
	{
		plantRoutes.GET("/:id", plantHandler.GetPlant)       // Get a specific plant by ID
		plantRoutes.GET("/all", plantHandler.GetAllPlants)   // Get all plants
		plantRoutes.POST("/creat", plantHandler.CreatePlant) // Get all plants
		plantRoutes.DELETE("/:id", plantHandler.DeletePlant) // Delete a specific plant by ID

		plantRoutes.GET("/user/:user_id", plantHandler.GetPlantsUser)    // Get all plants for a user
		plantRoutes.GET("/user/:user_id/:id", plantHandler.GetPlantUser) // Get a specific plant by ID

	}
}
