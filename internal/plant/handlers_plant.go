package plant

import (
	"go_plants/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreatePlantHandler обрабатывает запрос на создание нового растения
func CreatePlantHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request CreatePlantRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		newPlant := models.Plant{
			Name: request.Name,
		}

		if err := CreatePlant(db, &newPlant); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create plant"})
			return
		}

		c.JSON(http.StatusCreated, ConvertToPlantResponse(&newPlant))
	}
}

// GetPlantsHandler обрабатывает запрос на получение списка всех растений
func GetPlantsHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		plants, err := GetPlants(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get plants"})
			return
		}

		response := make([]PlantResponse, len(plants))
		for i, p := range plants {
			response[i] = ConvertToPlantResponse(&p)
		}

		c.JSON(http.StatusOK, response)
	}
}

// GetPlantByIDHandler обрабатывает запрос на получение растения по ID
func GetPlantByIDHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		plantID, err := strconv.ParseUint(c.Param("id"), 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid plant ID"})
			return
		}

		plant, err := GetPlantByID(db, uint(plantID))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Plant not found"})
			return
		}

		c.JSON(http.StatusOK, ConvertToPlantResponse(plant))
	}
}

// UpdatePlantHandler обрабатывает запрос на обновление растения
func UpdatePlantHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		plantID, err := strconv.ParseUint(c.Param("id"), 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid plant ID"})
			return
		}

		var request UpdatePlantRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		plant, err := GetPlantByID(db, uint(plantID))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Plant not found"})
			return
		}

		plant.Name = request.Name
		if err := UpdatePlant(db, plant); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update plant"})
			return
		}

		c.JSON(http.StatusOK, ConvertToPlantResponse(plant))
	}
}

// DeletePlantHandler обрабатывает запрос на удаление растения
func DeletePlantHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		plantID, err := strconv.ParseUint(c.Param("id"), 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid plant ID"})
			return
		}

		if err := DeletePlant(db, uint(plantID)); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete plant"})
			return
		}

		c.Status(http.StatusNoContent)
	}
} 