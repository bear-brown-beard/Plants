package api

import (
	"go_plants/internal/models"
	"go_plants/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PlantHandler struct {
	plantService services.PlantService
}

func NewPlantHandler(plantService services.PlantService) *PlantHandler {
	return &PlantHandler{
		plantService: plantService,
	}
}

func (h *PlantHandler) CreatePlant(c *gin.Context) {
	var plant models.Plant
	if err := c.ShouldBindJSON(&plant); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.plantService.Create(c.Request.Context(), &plant); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Plant created successfully"})
}
func (h *PlantHandler) GetAllPlants(c *gin.Context) {
	// Fetch all plants from the database
	plants, err := h.plantService.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, plants)
}
func (h *PlantHandler) GetPlant(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid plant ID"})
		return
	}

	plant, err := h.plantService.GetByID(c.Request.Context(), uint(id), 0)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if plant == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Plant not found"})
		return
	}

	c.JSON(http.StatusOK, plant)
}
func (h *PlantHandler) DeletePlant(c *gin.Context) {
	idParam := c.Param("id")
	userIDParam := c.Param("user_id")

	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid plant ID"})
		return
	}

	userID, err := strconv.ParseUint(userIDParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user_id"})
		return
	}

	if err := h.plantService.Delete(c.Request.Context(), uint(id), uint(userID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Plant deleted successfully"})
}

func (h *PlantHandler) GetPlantsUser(c *gin.Context) {
	userIDParam := c.Param("user_id")
	userID, err := strconv.ParseUint(userIDParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user_id"})
		return
	}

	plants, err := h.plantService.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Filter plants by user ID
	filteredPlants := make([]*models.Plant, 0)
	for _, plant := range plants {
		if plant.UserID == uint(userID) {
			filteredPlants = append(filteredPlants, plant)
		}
	}

	c.JSON(http.StatusOK, filteredPlants)
}
func (h *PlantHandler) GetPlantUser(c *gin.Context) {
	idParam := c.Param("id")
	userIDParam := c.Param("user_id")

	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid plant ID"})
		return
	}

	userID, err := strconv.ParseUint(userIDParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user_id"})
		return
	}

	plant, err := h.plantService.GetByID(c.Request.Context(), uint(id), uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if plant == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Plant not found"})
		return
	}

	c.JSON(http.StatusOK, plant)
}
