package config

import (
	"go_plants/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// Подключение к базе данных с жестко заданными параметрами
	dsn := "host=localhost user=plants_user password=secret dbname=plants_db port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Ошибка подключения к базе данных", err)
	}

	// Автоматическая миграция для модели Plant
	if err := db.AutoMigrate(&models.Plant{}); err != nil {
		log.Fatal("Ошибка миграции:", err)
	}

	DB = db
	log.Println("База данных успешно подключена")
}
