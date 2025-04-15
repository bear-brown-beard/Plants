package handlers

import (
	"go_plants/auth"
	"go_plants/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func LoginUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input struct {
			Email    string `json:"email" binding:"required"`
			Password string `json:"password" binding:"required"`
		}

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат входных данных"})
			return
		}

		var user models.User
		if err := db.Where("email = ?", input.Email).First(&user).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный email или пароль"})
			return
		}

		// Сравниваем пароли
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный email или пароль"})
			return
		}

		// Генерация JWT токена
		token, err := auth.GenerateToken(user.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при генерации токена"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"token": token,
			"user":  user,
		})
	}
}

func GetProfile(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Извлекаем user_id из контекста
		userIDInterface, exists := c.Get("user_id")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Пользователь не авторизован"})
			return
		}

		userID, ok := userIDInterface.(uint)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Ошибка преобразования user_id"})
			return
		}

		// Ищем пользователя в базе данных
		var user models.User
		if err := db.First(&user, userID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Пользователь не найден"})
			return
		}

		// Возвращаем информацию о пользователе
		c.JSON(http.StatusOK, gin.H{
			"id":         user.ID,
			"first_name": user.FirstName,
			"last_name":  user.LastName,
			"email":      user.Email,
			"city":       user.City,
		})
	}
}
