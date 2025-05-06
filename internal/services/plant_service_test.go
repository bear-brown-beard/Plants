package services

import (
	"context"
	"testing"

	"go_plants/internal/models"
	"go_plants/internal/repositories/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestPlantService(t *testing.T) {
	// Create mock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create mock repository
	mockRepo := mocks.NewMockPlantRepository(ctrl)
	plantService := NewPlantService(mockRepo)

	// Test Create
	t.Run("CreatePlant", func(t *testing.T) {
		plant := &models.Plant{
			ID:          1,
			Name:        "Test Plant",
			UserID:      1,
			Watering:    "daily",
			Description: "A test plant for testing purposes",
			Repotting:   "annually",
			Breeding:    "stem cuttings",
		}

		mockRepo.EXPECT().Create(context.Background(), plant).Return(nil)
		err := plantService.Create(context.Background(), plant)
		assert.NoError(t, err)
	})

	// Test GetAll
	t.Run("GetAllPlants", func(t *testing.T) {
		plants := []*models.Plant{
			{ID: 1, Name: "Plant1", UserID: 1, Watering: "weekly", Description: "First test plant"},
			{ID: 2, Name: "Plant2", UserID: 1, Watering: "daily", Description: "Second test plant"},
		}

		mockRepo.EXPECT().GetAll(context.Background()).Return(plants, nil)
		result, err := plantService.GetAll(context.Background())
		assert.NoError(t, err)
		assert.Equal(t, plants, result)
	})

	// Test GetByID
	t.Run("GetPlantByID", func(t *testing.T) {
		mockPlant := &models.Plant{
			ID:          1,
			Name:        "Test Plant",
			UserID:      1,
			Watering:    "daily",
			Description: "A test plant for testing purposes",
			Repotting:   "annually",
			Breeding:    "stem cuttings",
		}

		mockRepo.EXPECT().GetByID(context.Background(), uint(1), uint(1)).Return(mockPlant, nil)
		plant, err := plantService.GetByID(context.Background(), uint(1), uint(1))
		assert.NoError(t, err)
		assert.Equal(t, mockPlant, plant)
	})

	// Test Update
	t.Run("UpdatePlant", func(t *testing.T) {
		plant := &models.Plant{
			ID:          1,
			Name:        "Updated Plant",
			UserID:      1,
			Watering:    "weekly",
			Description: "Updated description",
			Repotting:   "biannually",
			Breeding:    "leaf cuttings",
		}

		mockRepo.EXPECT().Update(context.Background(), uint(1), uint(1), plant).Return(nil)
		err := plantService.Update(context.Background(), uint(1), uint(1), plant)
		assert.NoError(t, err)
	})

	// Test Delete
	t.Run("DeletePlant", func(t *testing.T) {
		mockRepo.EXPECT().Delete(context.Background(), uint(1), uint(1)).Return(nil)
		err := plantService.Delete(context.Background(), uint(1), uint(1))
		assert.NoError(t, err)
	})
}
