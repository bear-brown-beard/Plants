package admin

import (
	"log"

	"gorm.io/gorm"
)

// Очистка таблицы растений
func ClearPlants(db *gorm.DB) error {
	// Очищаем таблицу растений
	err := db.Exec("DELETE FROM plants").Error
	if err != nil {
		log.Println("Ошибка очистки таблицы растений:", err)
		return err
	}
	log.Println("Таблица растений очищена!")
	return nil
}

// Сброс автоинкремента
func ResetAutoIncrement(db *gorm.DB) error {
	// Сбрасываем автоинкремент
	err := db.Exec("ALTER SEQUENCE plants_id_seq RESTART WITH 1").Error
	if err != nil {
		log.Println("Ошибка сброса автоинкремента:", err)
		return err
	}
	log.Println("Автоинкремент сброшен на 1!")
	return nil
}
