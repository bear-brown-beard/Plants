package admin

import (
	"net/http"
	"strconv"

	"go_plants/internal/models"
	"go_plants/internal/plant"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreatePlantHandler handles creating a new plant
func CreatePlantHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newPlant models.Plant
		if err := c.ShouldBindJSON(&newPlant); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid plant data"})
			return
		}

		if err := plant.CreatePlant(db, &newPlant); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create plant"})
			return
		}

		c.JSON(http.StatusCreated, newPlant)
	}
}

// GetPlantByIDHandler handles getting a plant by ID
func GetPlantByIDHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseUint(c.Param("id"), 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid plant ID"})
			return
		}

		plant, err := plant.GetPlantByID(db, uint(id))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Plant not found"})
			return
		}

		c.JSON(http.StatusOK, plant)
	}
}

// GetAllPlantsHandler handles getting all plants
func GetAllPlantsHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		plants, err := plant.GetAllPlants(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get plants"})
			return
		}

		c.JSON(http.StatusOK, plants)
	}
}

// UpdatePlantHandler handles updating a plant
func UpdatePlantHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseUint(c.Param("id"), 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid plant ID"})
			return
		}

		var updatedPlant models.Plant
		if err := c.ShouldBindJSON(&updatedPlant); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid plant data"})
			return
		}

		updatedPlant.ID = uint(id)
		if err := plant.UpdatePlant(db, &updatedPlant); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update plant"})
			return
		}

		c.JSON(http.StatusOK, updatedPlant)
	}
}

// DeletePlantHandler handles deleting a plant
func DeletePlantHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseUint(c.Param("id"), 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid plant ID"})
			return
		}

		if err := plant.DeletePlant(db, uint(id)); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete plant"})
			return
		}

		c.Status(http.StatusNoContent)
	}
}
