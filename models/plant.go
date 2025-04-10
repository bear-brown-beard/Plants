package models

import (
	"gorm.io/gorm"
)

// Структура для растения
type Plant struct {
	gorm.Model
	Name        string `gorm:"column:name;type:varchar(100);not null"`
	Description string `gorm:"column:description;type:text"`
	Watering    string `gorm:"column:watering;type:varchar(50)"`
	Repotting   string `gorm:"column:repotting;type:varchar(50)"`
	Breeding    string `gorm:"column:breeding;type:varchar(50)"`
}
