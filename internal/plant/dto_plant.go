package plant

import (
	"go_plants/internal/models"
)

// CreatePlantRequest представляет запрос на создание растения
type CreatePlantRequest struct {
	Name string `json:"name" binding:"required"`
}

// UpdatePlantRequest представляет запрос на обновление растения
type UpdatePlantRequest struct {
	Name string `json:"name" binding:"required"`
}

// UpdateUserPlantRequest представляет запрос на обновление растения пользователя
type UpdateUserPlantRequest struct {
	Name string `json:"name" binding:"required"`
}

// PlantResponse представляет ответ с информацией о растении
type PlantResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	//CreatedAt string `json:"created_at"`
	//UpdatedAt string `json:"updated_at"`
}

// UserPlantResponse представляет ответ с информацией о растении пользователя
type UserPlantResponse struct {
	ID            uint          `json:"id"`
	UserID        uint          `json:"user_id"`
	Name          string        `json:"name"`
	OriginalPlant PlantResponse `json:"original_plant"`
	//CreatedAt     string        `json:"created_at"`
	//UpdatedAt     string        `json:"updated_at"`
}

// ConvertToPlantResponse конвертирует модель Plant в PlantResponse
func ConvertToPlantResponse(plant *models.Plant) PlantResponse {
	return PlantResponse{
		ID:   plant.ID,
		Name: plant.Name,
		//CreatedAt: plant.CreatedAt.String(),
		//UpdatedAt: plant.UpdatedAt.String(),
	}
}

// ConvertToUserPlantResponse конвертирует модель UserPlant в UserPlantResponse
func ConvertToUserPlantResponse(userPlant *models.UserPlant) UserPlantResponse {
	response := UserPlantResponse{
		ID:     userPlant.ID,
		UserID: userPlant.UserID,
		Name:   userPlant.Name,
		//CreatedAt:     userPlant.CreatedAt.String(),
		//UpdatedAt:     userPlant.UpdatedAt.String(),
	}

	// Если OriginalPlant загружен, добавляем его в ответ
	if userPlant.OriginalPlant.ID != 0 {
		response.OriginalPlant = ConvertToPlantResponse(&userPlant.OriginalPlant)
	}
	return response
}
