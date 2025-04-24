package config

import (
	"go_plants/internal/models"
	"go_plants/internal/user"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB initializes the database connection
func InitDB() *gorm.DB {
	// Подключение к базе данных с жестко заданными параметрами
	dsn := "host=localhost user=plants_user password=secret dbname=plants_db port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Ошибка подключения к базе данных", err)
	}

	// Автоматическая миграция схемы
	if err := db.AutoMigrate(&user.User{}, &models.Plant{}, &models.UserPlant{}); err != nil {
		log.Fatal("Ошибка миграции:", err)
	}

	DB = db
	log.Println("База данных успешно подключена")
	return db
}
