package handlers

import (
	"go_plants/models"
	"net/http"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input models.User
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
			return
		}

		// Проверка, существует ли уже пользователь с таким email
		var existing models.User
		if err := db.Where("email = ?", input.Email).First(&existing).Error; err == nil {
			c.JSON(http.StatusConflict, gin.H{"error": "Пользователь с таким email уже существует"})
			return
		}

		// Хэшируем пароль
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при шифровании пароля"})
			return
		}
		input.Password = string(hashedPassword)

		// Сохраняем пользователя
		if err := db.Create(&input).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании пользователя"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "Пользователь успешно зарегистрирован",
			"user":    input,
		})
	}
}
