package plant

import (
	"fmt"
	"go_plants/internal/models"
	"log"

	"gorm.io/gorm"
)

// CopyPlantToUser копирует растение из основной таблицы в коллекцию пользователя
func CopyPlantToUser(db *gorm.DB, userID uint, originalPlantID uint) error {
	// Проверяем существование растения
	originalPlant, err := GetPlantByID(db, originalPlantID)
	if err != nil {
		return err
	}

	// Проверяем, есть ли уже это растение в коллекции пользователя
	var existingPlant models.UserPlant
	result := db.Where("user_id = ? AND original_plant_id = ?", userID, originalPlantID).First(&existingPlant)
	if result.Error == nil {
		return nil // Растение уже есть в коллекции
	}

	// Создаем новое растение в коллекции пользователя
	userPlant := models.UserPlant{
		UserID:          userID,
		OriginalPlantID: originalPlantID,
		Name:            originalPlant.Name, // Копируем название из оригинального растения
	}

	// Создаем запись в базе данных
	if err := db.Create(&userPlant).Error; err != nil {
		log.Printf("Error copying plant to user: %v", err)
		return err
	}

	log.Printf("Successfully copied plant ID %d to user ID %d", originalPlantID, userID)
	return nil
}

// UpdateUserPlant обновляет растение в коллекции пользователя
func UpdateUserPlant(db *gorm.DB, userID uint, plantID uint, newName string) error {
	// Логируем входные параметры
	log.Printf("Updating plant name. UserID: %d, PlantID: %d, NewName: %s", userID, plantID, newName)

	// Получаем растение пользователя
	var userPlant models.UserPlant
	if err := db.Preload("OriginalPlant").
		Where("user_id = ? AND id = ?", userID, plantID).
		First(&userPlant).Error; err != nil {
		log.Printf("Error finding user plant: %v", err)
		return err
	}

	// Логируем текущее состояние
	log.Printf("Current plant state - ID: %d, Name: %s", userPlant.ID, userPlant.Name)

	// Обновляем название
	result := db.Model(&models.UserPlant{}).
		Where("user_id = ? AND id = ?", userID, plantID).
		Update("name", newName)

	if result.Error != nil {
		log.Printf("Error updating plant name: %v", result.Error)
		return result.Error
	}

	// Проверяем количество обновленных строк
	if result.RowsAffected == 0 {
		log.Printf("No rows were updated")
		return fmt.Errorf("no rows were updated")
	}

	// Получаем обновленное растение для проверки
	var updatedPlant models.UserPlant
	if err := db.Preload("OriginalPlant").
		Where("id = ?", plantID).
		First(&updatedPlant).Error; err != nil {
		log.Printf("Error verifying update: %v", err)
		return err
	}

	log.Printf("Plant updated successfully. New name: %s", updatedPlant.Name)
	return nil
}

// GetUserPlants получает все растения пользователя
func GetUserPlants(db *gorm.DB, userID uint) ([]models.UserPlant, error) {
	var userPlants []models.UserPlant
	if err := db.Preload("OriginalPlant").
		Where("user_id = ?", userID).
		Find(&userPlants).Error; err != nil {
		log.Printf("Error getting user plants: %v", err)
		return nil, err
	}
	return userPlants, nil
}
