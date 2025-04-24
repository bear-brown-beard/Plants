package plant

import (
	"go_plants/internal/models"
	"log"

	"gorm.io/gorm"
)

// CreatePlant создает новое растение в базе данных
func CreatePlant(db *gorm.DB, plant *models.Plant) error {
	if err := db.Create(plant).Error; err != nil {
		log.Println("Error creating plant:", err)
		return err
	}
	return nil
}

// GetPlants возвращает список всех растений
func GetPlants(db *gorm.DB) ([]models.Plant, error) {
	var plants []models.Plant
	if err := db.Find(&plants).Error; err != nil {
		log.Println("Error getting plants:", err)
		return nil, err
	}
	return plants, nil
}

// GetPlantByID возвращает растение по ID
func GetPlantByID(db *gorm.DB, id uint) (*models.Plant, error) {
	var plant models.Plant
	if err := db.First(&plant, id).Error; err != nil {
		log.Println("Error getting plant by ID:", err)
		return nil, err
	}
	return &plant, nil
}

// GetAllPlants retrieves all plants from the database
func GetAllPlants(db *gorm.DB) ([]models.Plant, error) {
	var plants []models.Plant
	if err := db.Find(&plants).Error; err != nil {
		log.Println("Error getting plants:", err)
		return nil, err
	}
	return plants, nil
}

// UpdatePlant обновляет растение в базе данных
func UpdatePlant(db *gorm.DB, plant *models.Plant) error {
	if err := db.Save(plant).Error; err != nil {
		log.Println("Error updating plant:", err)
		return err
	}
	return nil
}

// DeletePlant удаляет растение из базы данных
func DeletePlant(db *gorm.DB, id uint) error {
	if err := db.Delete(&models.Plant{}, id).Error; err != nil {
		log.Println("Error deleting plant:", err)
		return err
	}
	return nil
}

// GetPublicPlantByID retrieves a plant by ID for public access
func GetPublicPlantByID(db *gorm.DB, id uint) (*models.Plant, error) {
	return GetPlantByID(db, id)
}

// GetAllPublicPlants retrieves all plants for public access
func GetAllPublicPlants(db *gorm.DB) ([]models.Plant, error) {
	return GetAllPlants(db)
}
