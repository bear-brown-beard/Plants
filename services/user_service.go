package services

import (
	"errors"
	"go_plants/models"
	"log"

	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, user *models.User) error {
	// Логируем входящие данные
	log.Printf("Попытка создать пользователя: %+v", user)

	// Проверка на существующий email
	var existingUser models.User
	if err := db.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		// Если пользователь с таким email уже существует, возвращаем ошибку
		log.Println("Ошибка: пользователь с таким email уже существует")
		return errors.New("пользователь с таким email уже существует")
	}

	// Создание пользователя
	if err := db.Create(user).Error; err != nil {
		// Логируем ошибку, если создание не удалось
		log.Printf("Ошибка при создании пользователя: %v", err)
		return err
	}

	// Логируем успешное создание
	log.Println("Пользователь успешно создан")
	return nil
}

func GetAllUsers(db *gorm.DB) ([]models.User, error) {
	var users []models.User
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func GetUserByID(db *gorm.DB, id string) (*models.User, error) {
	var user models.User
	if err := db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateUser(db *gorm.DB, id string, updated *models.User) error {
	var user models.User
	if err := db.First(&user, id).Error; err != nil {
		return err
	}

	return db.Model(&user).Updates(updated).Error
}

func DeleteUser(db *gorm.DB, id string) error {
	return db.Delete(&models.User{}, id).Error
}
