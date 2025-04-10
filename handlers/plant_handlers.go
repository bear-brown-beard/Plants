package handlers

import (
	"go_plants/models"
	"go_plants/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterPlantHandlers(r *gin.Engine, db *gorm.DB) {
	r.POST("/plants", func(c *gin.Context) {
		var plant models.Plant
		if err := c.ShouldBindJSON(&plant); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
			return
		}

		if err := services.CreatePlant(db, &plant); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании растения"})
			return
		}

		c.JSON(http.StatusCreated, plant)
	})

	r.GET("/plants", func(c *gin.Context) {
		plants, err := services.GetAllPlants(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить список растений"})
			return
		}

		c.JSON(http.StatusOK, plants)
	})

	r.GET("/plants/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		plant, err := services.GetPlantByID(db, uint(id))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Растение не найдено"})
			return
		}

		c.JSON(http.StatusOK, plant)
	})

	r.PUT("/plants/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		var updated models.Plant
		if err := c.ShouldBindJSON(&updated); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
			return
		}

		if err := services.UpdatePlant(db, uint(id), &updated); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении растения"})
			return
		}

		c.JSON(http.StatusOK, updated)
	})

	r.DELETE("/plants/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		if err := services.DeletePlant(db, uint(id)); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при удалении растения"})
			return
		}

		c.Status(http.StatusNoContent)
	})
}
