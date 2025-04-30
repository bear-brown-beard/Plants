package services

import (
	"go_plants/internal/models"

	"gorm.io/gorm"
)

// UserService определяет интерфейс для работы с пользователями
type UserService interface {
	CreateUser(db *gorm.DB, user *models.User) error
	LoginUser(db *gorm.DB, email, password string) (*models.User, error)
	GetAllUsers(db *gorm.DB) ([]models.User, error)
	GetUserByID(db *gorm.DB, id uint) (*models.User, error)
	UpdateUser(db *gorm.DB, id string, updated *models.User) error
	DeleteUser(db *gorm.DB, id string) error
}

// PlantService определяет интерфейс для работы с растениями
type PlantService interface {
	CreatePlant(db *gorm.DB, plant *models.Plant) error
	GetAllPlants(db *gorm.DB) ([]models.Plant, error)
	GetPlantByID(db *gorm.DB, id uint) (*models.Plant, error)
	UpdatePlant(db *gorm.DB, id uint, updated *models.Plant) error
	DeletePlant(db *gorm.DB, id uint) error
}
