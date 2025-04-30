package api

import (
	"go_plants/internal/auth"
	"go_plants/internal/models"
	"go_plants/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// AuthService определяет интерфейс для сервисов аутентификации и регистрации
type AuthService interface {
	LoginUser(db *gorm.DB) gin.HandlerFunc
	RegisterUser(db *gorm.DB) gin.HandlerFunc
	GetProfile(db *gorm.DB) gin.HandlerFunc
}

// AuthServiceImpl реализует интерфейс AuthService
type AuthServiceImpl struct {
	userService services.UserService
}

// NewAuthService создает новый экземпляр AuthServiceImpl
func NewAuthService(userService services.UserService) AuthService {
	return &AuthServiceImpl{
		userService: userService,
	}
}

func (s *AuthServiceImpl) LoginUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input struct {
			Email    string `json:"email" binding:"required"`
			Password string `json:"password" binding:"required"`
		}

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат входных данных"})
			return
		}

		user, err := s.userService.LoginUser(db, input.Email, input.Password)
		if err != nil {
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

func (s *AuthServiceImpl) RegisterUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input models.User
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
			return
		}

		// Сохраняем пользователя через сервис
		if err := s.userService.CreateUser(db, &input); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "Пользователь успешно зарегистрирован",
			"user":    input,
		})
	}
}

func (s *AuthServiceImpl) GetProfile(db *gorm.DB) gin.HandlerFunc {
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
		user, err := s.userService.GetUserByID(db, uint(userID))
		if err != nil {
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
