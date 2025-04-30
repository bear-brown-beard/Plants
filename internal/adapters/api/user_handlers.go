package api

import (
	//"go_plants/internal/auth"
	"go_plants/internal/models"
	"go_plants/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// UserHandler определяет интерфейс для обработчиков пользователей
type UserHandler interface {
	CreateUser(db *gorm.DB) gin.HandlerFunc
	GetUser(db *gorm.DB) gin.HandlerFunc
	UpdateUser(db *gorm.DB) gin.HandlerFunc
	DeleteUser(db *gorm.DB) gin.HandlerFunc
	UpdateProfile(db *gorm.DB) gin.HandlerFunc
}

// UserHandlerImpl реализует интерфейс UserHandler
type UserHandlerImpl struct {
	userService services.UserService
}

// NewUserHandler создает новый экземпляр UserHandlerImpl
func NewUserHandler(userService services.UserService) UserHandler {
	return &UserHandlerImpl{
		userService: userService,
	}
}

// CreateUser создает нового пользователя
func (h *UserHandlerImpl) CreateUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
			return
		}

		if err := h.userService.CreateUser(db, &user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании пользователя"})
			return
		}

		c.JSON(http.StatusCreated, user)
	}
}

// GetUser получает информацию о пользователе по ID
func (h *UserHandlerImpl) GetUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID пользователя"})
			return
		}

		user, err := h.userService.GetUserByID(db, uint(id))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Пользователь не найден"})
			return
		}

		c.JSON(http.StatusOK, user)
	}
}

// UpdateUser обновляет информацию о пользователе
func (h *UserHandlerImpl) UpdateUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID пользователя"})
			return
		}

		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
			return
		}

		if err := h.userService.UpdateUser(db, id, &user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении пользователя"})
			return
		}

		c.JSON(http.StatusOK, user)
	}
}

// DeleteUser удаляет пользователя
func (h *UserHandlerImpl) DeleteUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID пользователя"})
			return
		}

		if err := h.userService.DeleteUser(db, id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при удалении пользователя"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Пользователь успешно удален"})
	}
}

// UpdateProfile обновляет профиль текущего пользователя
func (h *UserHandlerImpl) UpdateProfile(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get("user_id")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Пользователь не авторизован"})
			return
		}

		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
			return
		}

		id := strconv.FormatUint(uint64(userID.(uint)), 10)
		if err := h.userService.UpdateUser(db, id, &user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении профиля"})
			return
		}

		c.JSON(http.StatusOK, user)
	}
}
