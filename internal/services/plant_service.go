package services

import (
	"context"
	"go_plants/internal/models"
	"go_plants/internal/repositories"
)

type PlantService interface {
	Create(ctx context.Context, plant *models.Plant) error
	GetAll(ctx context.Context) ([]*models.Plant, error)
	GetByID(ctx context.Context, id uint, userID uint) (*models.Plant, error)
	Update(ctx context.Context, id uint, userID uint, plant *models.Plant) error
	Delete(ctx context.Context, id uint, userID uint) error
}

type plantService struct {
	plantRepository repositories.PlantRepository
}

func NewPlantService(plantRepo repositories.PlantRepository) PlantService {
	return &plantService{
		plantRepository: plantRepo,
	}
}

func (s *plantService) Create(ctx context.Context, plant *models.Plant) error {
	return s.plantRepository.Create(ctx, plant)
}
func (s *plantService) GetAll(ctx context.Context) ([]*models.Plant, error) {
	return s.plantRepository.GetAll(ctx)
}
func (s *plantService) GetByID(ctx context.Context, id uint, userID uint) (*models.Plant, error) {
	return s.plantRepository.GetByID(ctx, id, userID)
}
func (s *plantService) Update(ctx context.Context, id uint, userID uint, plant *models.Plant) error {
	return s.plantRepository.Update(ctx, id, userID, plant)
}
func (s *plantService) Delete(ctx context.Context, id uint, userID uint) error {
	return s.plantRepository.Delete(ctx, id, userID)
}
