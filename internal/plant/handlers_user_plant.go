package plant

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CopyPlantToCollectionHandler handles copying a plant to user's collection
func CopyPlantToCollectionHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Получаем ID пользователя из контекста
		userID := c.GetUint("user_id")
		if userID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		// Получаем ID растения из параметров URL
		plantID, err := strconv.ParseUint(c.Param("id"), 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid plant ID"})
			return
		}

		// Копируем растение в коллекцию пользователя
		if err := CopyPlantToUser(db, userID, uint(plantID)); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to copy plant"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "The plant has been added"})
	}
}

// UpdateUserPlantHandler обрабатывает запрос на обновление растения в коллекции пользователя
func UpdateUserPlantHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Получаем ID пользователя из контекста
		userID := c.GetUint("user_id")
		if userID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		// Получаем ID растения из параметров URL
		plantID, err := strconv.ParseUint(c.Param("id"), 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid plant ID"})
			return
		}

		// Получаем новое название из тела запроса
		var request UpdateUserPlantRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		// Обновляем растение в коллекции пользователя
		if err := UpdateUserPlant(db, userID, uint(plantID), request.Name); err != nil {
			if err.Error() == "plant not found or not owned by user" {
				c.JSON(http.StatusNotFound, gin.H{"error": "Plant not found or not owned by user"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update plant"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "The data has been updated successfully"})
	}
}

// GetUserPlantsHandler обрабатывает запрос на получение всех растений пользователя
func GetUserPlantsHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Получаем ID пользователя из контекста
		userID := c.GetUint("user_id")
		if userID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		// Получаем растения пользователя
		plants, err := GetUserPlants(db, userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user plants"})
			return
		}

		// Конвертируем в ответ
		response := make([]UserPlantResponse, len(plants))
		for i, p := range plants {
			response[i] = ConvertToUserPlantResponse(&p)
		}

		c.JSON(http.StatusOK, response)
	}
}
