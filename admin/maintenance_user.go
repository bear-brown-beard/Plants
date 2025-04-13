package admin

import (
	"log"

	"gorm.io/gorm"
)

// Очистка таблицы пользователей
func ClearUsers(db *gorm.DB) error {
	// Очищаем таблицу пользователей
	err := db.Exec("DELETE FROM users").Error
	if err != nil {
		log.Println("Ошибка очистки таблицы пользователей:", err)
		return err
	}
	log.Println("Таблица пользователей очищена!")
	return nil
}

// Сброс автоинкремента для пользователей
func ResetUserAutoIncrement(db *gorm.DB) error {
	// Сбрасываем автоинкремент для пользователей
	err := db.Exec("ALTER SEQUENCE users_id_seq RESTART WITH 1").Error
	if err != nil {
		log.Println("Ошибка сброса автоинкремента для пользователей:", err)
		return err
	}
	log.Println("Автоинкремент пользователей сброшен на 1!")
	return nil
}
