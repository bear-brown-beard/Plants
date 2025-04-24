package user

import (
	"net/http"
	"strconv"

	"go_plants/internal/plant"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetPublicPlantByIDHandler handles getting a plant by ID for public access
func GetPublicPlantByIDHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseUint(c.Param("id"), 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid plant ID"})
			return
		}

		plant, err := plant.GetPublicPlantByID(db, uint(id))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Plant not found"})
			return
		}

		c.JSON(http.StatusOK, plant)
	}
}

// GetAllPublicPlantsHandler handles getting all plants for public access
func GetAllPublicPlantsHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		plants, err := plant.GetAllPublicPlants(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get plants"})
			return
		}

		c.JSON(http.StatusOK, plants)
	}
} 