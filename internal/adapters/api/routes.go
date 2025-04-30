package api

import (
	"go_plants/internal/auth"
	"go_plants/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterUserAPI регистрирует маршруты для пользователей
func RegisterUserAPI(r *gin.Engine, db *gorm.DB) {
	userService := services.NewUserService()
	userHandler := NewUserHandler(userService)

	// Регистрируем маршруты для пользователей
	r.POST("/users", userHandler.CreateUser(db))
	r.GET("/users/:id", userHandler.GetUser(db))
	r.PUT("/users/:id", userHandler.UpdateUser(db))
	r.DELETE("/users/:id", userHandler.DeleteUser(db))
}

// RegisterPlantAPI регистрирует маршруты для растений
func RegisterPlantAPI(r *gin.Engine, db *gorm.DB) {
	plantService := services.NewPlantService()
	plantHandler := NewPlantHandler(plantService)

	// Регистрируем маршруты для растений
	r.POST("/plants", plantHandler.CreatePlant(db))
	r.GET("/plants", plantHandler.GetAllPlants(db))
	r.GET("/plants/:id", plantHandler.GetPlant(db))
	r.PUT("/plants/:id", plantHandler.UpdatePlant(db))
	r.DELETE("/plants/:id", plantHandler.DeletePlant(db))
}

// RegisterAuthRoutes регистрирует маршруты для аутентификации
func RegisterAuthRoutes(r *gin.Engine, db *gorm.DB) {
	userService := services.NewUserService()
	authService := NewAuthService(userService)

	// Регистрируем маршруты для аутентификации
	r.POST("/login", authService.LoginUser(db))
	r.POST("/register", authService.RegisterUser(db))
	r.GET("/profile", auth.AuthMiddleware(), authService.GetProfile(db))
}
