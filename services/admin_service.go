package services

import (
	"go_plants/models"

	"gorm.io/gorm"
)

// AllowGlobalUpdate - необходим для удаления всех записей , иначе GORM не даст этого сделать, одна из мер защиты
func DeleteAllPlants(db *gorm.DB) error {
	return db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&models.Plant{}).Error
}

func CountPlants(db *gorm.DB) (int64, error) {
	var count int64
	err := db.Model(&models.Plant{}).Count(&count).Error
	return count, err
}
