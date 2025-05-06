package services

import (
	"context"
	"errors"
	"testing"

	"go_plants/internal/models"
	"go_plants/internal/repositories/mocks"
	"go_plants/internal/services"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestPlantService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockPlantRepository(ctrl)
	plantService := services.NewPlantService(mockRepo)

	t.Run("CreatePlant_Success", func(t *testing.T) {
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
		assert.NotNil(t, plant)
		assert.Equal(t, "Test Plant", plant.Name)
		assert.Equal(t, "daily", plant.Watering)
	})

	t.Run("GetAllPlants_Success", func(t *testing.T) {
		plants := []*models.Plant{
			{ID: 1, Name: "Plant1", UserID: 1, Watering: "weekly", Description: "First test plant"},
			{ID: 2, Name: "Plant2", UserID: 1, Watering: "daily", Description: "Second test plant"},
		}

		mockRepo.EXPECT().GetAll(context.Background()).Return(plants, nil)
		result, err := plantService.GetAll(context.Background())

		assert.NoError(t, err)
		assert.Len(t, result, 2)
		assert.Equal(t, "Plant1", result[0].Name)
		assert.Equal(t, "Plant2", result[1].Name)
	})

	t.Run("GetPlantByID_Success", func(t *testing.T) {
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
		assert.NotNil(t, plant)
		assert.Equal(t, "Test Plant", plant.Name)
	})

	t.Run("GetPlantByID_NotFound", func(t *testing.T) {
		mockRepo.EXPECT().GetByID(context.Background(), uint(1), uint(99)).Return(nil, errors.New("plant not found"))
		plant, err := plantService.GetByID(context.Background(), uint(1), uint(99))

		assert.Error(t, err)
		assert.Nil(t, plant)
		assert.EqualError(t, err, "plant not found")
	})

	t.Run("UpdatePlant_Success", func(t *testing.T) {
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
		assert.Equal(t, "Updated Plant", plant.Name)
		assert.Equal(t, "weekly", plant.Watering)
	})

	t.Run("DeletePlant_Success", func(t *testing.T) {
		mockRepo.EXPECT().Delete(context.Background(), uint(1), uint(1)).Return(nil)
		err := plantService.Delete(context.Background(), uint(1), uint(1))

		assert.NoError(t, err)
	})

	t.Run("DeletePlant_NotFound", func(t *testing.T) {
		mockRepo.EXPECT().Delete(context.Background(), uint(1), uint(99)).Return(errors.New("plant not found"))
		err := plantService.Delete(context.Background(), uint(1), uint(99))

		assert.Error(t, err)
		assert.EqualError(t, err, "plant not found")
	})
}

func TestPlantServiceErrors(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockPlantRepository(ctrl)
	plantService := services.NewPlantService(mockRepo)

	// Test: Plant not found
	t.Run("GetPlantByID_PlantNotFound", func(t *testing.T) {
		mockRepo.EXPECT().GetByID(context.Background(), uint(1), uint(1)).Return(nil, errors.New("plant not found"))
		plant, err := plantService.GetByID(context.Background(), uint(1), uint(1))
		assert.Error(t, err)
		assert.Nil(t, plant)
		assert.Equal(t, "plant not found", err.Error())
	})

	// Test: Create Plant with empty name
	t.Run("CreatePlant_EmptyName", func(t *testing.T) {
		plant := &models.Plant{ID: 1, Name: ""}

		// Ожидаем, что Create будет вызван с пустым именем и вернет ошибку
		mockRepo.EXPECT().Create(context.Background(), plant).Return(errors.New("plant name cannot be empty")).Times(1)

		err := plantService.Create(context.Background(), plant)
		assert.Error(t, err)
		assert.Equal(t, "plant name cannot be empty", err.Error())
	})

	// Test: Update Plant with very large description
	t.Run("UpdatePlant_LargeDescription", func(t *testing.T) {
		plant := &models.Plant{ID: 1, Name: "Test Plant", Description: string(make([]byte, 1000)), UserID: 1}
		mockRepo.EXPECT().Update(context.Background(), uint(1), uint(1), plant).Return(nil)
		err := plantService.Update(context.Background(), uint(1), uint(1), plant)
		assert.NoError(t, err)
	})
}

func TestPlantServiceConcurrency(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockPlantRepository(ctrl)
	plantService := services.NewPlantService(mockRepo)

	plant := &models.Plant{ID: 1, Name: "Test Plant", UserID: 1}

	// Test: Two updates in parallel
	t.Run("UpdatePlant_Concurrency", func(t *testing.T) {
		mockRepo.EXPECT().Update(context.Background(), uint(1), uint(1), plant).Return(nil).Times(2)

		errChan := make(chan error, 2)

		// Simulating two concurrent updates
		go func() {
			errChan <- plantService.Update(context.Background(), uint(1), uint(1), plant)
		}()
		go func() {
			errChan <- plantService.Update(context.Background(), uint(1), uint(1), plant)
		}()

		err1 := <-errChan
		err2 := <-errChan

		// Ensure no errors happened in both operations
		assert.NoError(t, err1)
		assert.NoError(t, err2)
	})
}
