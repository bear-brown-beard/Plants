package services

import (
	"errors"
	"go_plants/models"

	"gorm.io/gorm"
)

func CreatePlant(db *gorm.DB, plant *models.Plant) error {
	return db.Create(plant).Error
}

func GetAllPlants(db *gorm.DB) ([]models.Plant, error) {
	var plants []models.Plant
	err := db.Find(&plants).Error
	return plants, err
}

func GetPlantByID(db *gorm.DB, id uint) (*models.Plant, error) {
	var plant models.Plant
	result := db.First(&plant, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("Растение не найдено")
	}
	return &plant, result.Error
}

func UpdatePlant(db *gorm.DB, id uint, updated *models.Plant) error {
	var plant models.Plant
	if err := db.First(&plant, id).Error; err != nil {
		return err
	}
	return db.Model(&plant).Updates(updated).Error
}

func DeletePlant(db *gorm.DB, id uint) error {
	return db.Delete(&models.Plant{}, id).Error
}
