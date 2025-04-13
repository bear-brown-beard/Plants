package routes

import (
	"go_plants/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterAuthRoutes регистрирует маршруты для аутентификации и регистрации
func RegisterAuthRoutes(r *gin.Engine, db *gorm.DB) {
	// Регистрируем маршрут для логина
	r.POST("/login", handlers.LoginUser(db))
	// Регистрируем маршрут для регистрации пользователя
	r.POST("/register", handlers.RegisterUser(db))
}
