package models

import (
	"gorm.io/gorm"
)

// UserPlant represents a plant that a user has copied from the original plants table
type UserPlant struct {
	gorm.Model
	UserID          uint   `gorm:"not null" json:"user_id"`
	OriginalPlantID uint   `gorm:"not null" json:"original_plant_id"`
	OriginalPlant   Plant  `gorm:"foreignKey:OriginalPlantID"`
	Name            string `gorm:"size:255" json:"name"`
}

// Plant represents a plant in the system
type Plant struct {
	gorm.Model
	Name string `gorm:"size:255;not null" json:"name"`
}
