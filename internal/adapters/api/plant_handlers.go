package api

import (
	"go_plants/internal/models"
	"go_plants/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// PlantHandler определяет интерфейс для обработчиков растений
type PlantHandler interface {
	CreatePlant(db *gorm.DB) gin.HandlerFunc
	GetAllPlants(db *gorm.DB) gin.HandlerFunc
	GetPlant(db *gorm.DB) gin.HandlerFunc
	UpdatePlant(db *gorm.DB) gin.HandlerFunc
	DeletePlant(db *gorm.DB) gin.HandlerFunc
}

// PlantHandlerImpl реализует интерфейс PlantHandler
type PlantHandlerImpl struct {
	plantService services.PlantService
}

// NewPlantHandler создает новый экземпляр PlantHandlerImpl
func NewPlantHandler(plantService services.PlantService) PlantHandler {
	return &PlantHandlerImpl{
		plantService: plantService,
	}
}

// CreatePlant создает новое растение
func (h *PlantHandlerImpl) CreatePlant(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var plant models.Plant
		if err := c.ShouldBindJSON(&plant); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
			return
		}

		if err := h.plantService.CreatePlant(db, &plant); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании растения"})
			return
		}

		c.JSON(http.StatusCreated, plant)
	}
}

// GetAllPlants получает список всех растений
func (h *PlantHandlerImpl) GetAllPlants(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		plants, err := h.plantService.GetAllPlants(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить список растений"})
			return
		}

		c.JSON(http.StatusOK, plants)
	}
}

// GetPlant получает информацию о растении по ID
func (h *PlantHandlerImpl) GetPlant(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID растения"})
			return
		}

		plant, err := h.plantService.GetPlantByID(db, uint(id))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Растение не найдено"})
			return
		}

		c.JSON(http.StatusOK, plant)
	}
}

// UpdatePlant обновляет информацию о растении
func (h *PlantHandlerImpl) UpdatePlant(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID растения"})
			return
		}

		var plant models.Plant
		if err := c.ShouldBindJSON(&plant); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
			return
		}

		if err := h.plantService.UpdatePlant(db, uint(id), &plant); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении растения"})
			return
		}

		c.JSON(http.StatusOK, plant)
	}
}

// DeletePlant удаляет растение
func (h *PlantHandlerImpl) DeletePlant(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID растения"})
			return
		}

		if err := h.plantService.DeletePlant(db, uint(id)); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при удалении растения"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Растение успешно удалено"})
	}
}
