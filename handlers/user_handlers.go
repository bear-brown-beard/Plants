package handlers

import (
	"go_plants/auth"
	"go_plants/models"
	"go_plants/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Регистрация пользователя
func RegisterUserHandlers(r *gin.Engine, db *gorm.DB) {
	// Регистрация пользователя
	r.POST("/users", func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
			return
		}

		if err := services.CreateUser(db, &user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании пользователя"})
			return
		}

		c.JSON(http.StatusCreated, user)
	})

	r.GET("/users", func(c *gin.Context) {
		users, err := services.GetAllUsers(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении пользователей"})
			return
		}

		c.JSON(http.StatusOK, users)
	})

	// Получение пользователя по ID
	r.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		user, err := services.GetUserByID(db, id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Пользователь не найден"})
			return
		}

		c.JSON(http.StatusOK, user)
	})

	// Обновление пользователя (требуется авторизация)
	r.PUT("/users/:id", authMiddleware(), func(c *gin.Context) {
		id := c.Param("id")
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
			return
		}

		if err := services.UpdateUser(db, id, &user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении пользователя"})
			return
		}

		c.JSON(http.StatusOK, user)
	})

	// Удаление пользователя (требуется авторизация)
	r.DELETE("/users/:id", authMiddleware(), func(c *gin.Context) {
		id := c.Param("id")
		if err := services.DeleteUser(db, id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при удалении пользователя"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Пользователь успешно удален"})
	})
}

// Middleware для проверки авторизации
func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Токен не предоставлен"})
			return
		}

		userID, err := auth.ParseToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Неверный токен"})
			return
		}

		// Сохраняем user_id в контекст, чтобы можно было использовать дальше
		c.Set("user_id", userID)
		c.Next()
	}
}

func UpdateProfile(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Получаем user_id из токена
		userIDRaw, exists := c.Get("user_id")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Пользователь не авторизован"})
			return
		}
		userID := userIDRaw.(uint)

		// Ищем пользователя
		var user models.User
		if err := db.First(&user, userID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Пользователь не найден"})
			return
		}

		// Входные данные
		var input struct {
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
			City      string `json:"city"`
		}

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
			return
		}

		// Обновляем только те поля, которые переданы
		if input.FirstName != "" {
			user.FirstName = input.FirstName
		}
		if input.LastName != "" {
			user.LastName = input.LastName
		}
		if input.City != "" {
			user.City = input.City
		}

		if err := db.Save(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении профиля"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Профиль успешно обновлён",
			"user":    user,
		})
	}
}
