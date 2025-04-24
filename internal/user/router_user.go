package user

import (
	"go_plants/internal/middleware"
	"go_plants/internal/plant"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterUserRoutes регистрирует все маршруты, связанные с пользователями
func RegisterUserRoutes(r *gin.Engine, db *gorm.DB) {
	// Публичные маршруты
	public := r.Group("/api")
	{
		// Маршруты аутентификации
		public.POST("/register", RegisterUser(db))
		public.POST("/login", LoginHandler(db))
	}

	// Защищенные маршруты
	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		// Маршруты профиля
		profile := protected.Group("/profile")
		{
			profile.GET("", GetProfileHandler(db))
			profile.PUT("", UpdateProfileHandler(db))
		}

		// Маршруты коллекции растений
		plants := protected.Group("/my-plants")
		{
			// Копировать растение в коллекцию
			plants.POST("/copy/:id", plant.CopyPlantToCollectionHandler(db))
			// ОБновление названия
			plants.PUT("/:id", plant.UpdateUserPlantHandler(db))
			// Получение всех растений пользователя
			plants.GET("", plant.GetUserPlantsHandler(db))
		}

		// Маршруты для просмотра растений из основной таблицы
		plantsCatalog := protected.Group("/plants")
		{
			// Получить список всех растений
			plantsCatalog.GET("", plant.GetPlantsHandler(db))

			// Получить растение по ID
			plantsCatalog.GET("/:id", plant.GetPlantByIDHandler(db))
		}

		// Маршруты управления пользователями
		protected.GET("/users", GetAllUsersHandler(db))
		protected.GET("/users/:id", GetUserByIDHandler(db))
	}
}
