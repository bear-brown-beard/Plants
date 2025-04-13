package models

import (
	"gorm.io/gorm"
)

// Структура для растения
type Plant struct {
	gorm.Model
	Name        string `gorm:"column:name;type:varchar(100);not null" json:"name"`
	Description string `gorm:"column:description;type:text" json:"description"`
	Watering    string `gorm:"column:watering;type:varchar(50)" json:"watering"`
	Repotting   string `gorm:"column:repotting;type:varchar(50)" json:"repotting"`
	Breeding    string `gorm:"column:breeding;type:varchar(50)" json:"breeding"`
	UserID      uint   `gorm:"not null" json:"user_id"`
}
