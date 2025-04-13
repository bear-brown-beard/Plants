package services

import (
	"errors"
	"go_plants/models"
	"log"
	"regexp"

	"golang.org/x/crypto/bcrypt"

	"gorm.io/gorm"
)

// Создание пользователя
func CreateUser(db *gorm.DB, user *models.User) error {
	// Хешируем пароль
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Ошибка при хешировании пароля:", err)
		return err
	}
	user.Password = string(hashedPassword)

	// Проверка на существующий email
	var existingUser models.User
	if err := db.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		log.Println("Ошибка: пользователь с таким email уже существует")
		return errors.New("пользователь с таким email уже существует")
	}

	if user.Password == "" {
		return errors.New("пароль не может быть пустым")
	}

	var passwordRegex = regexp.MustCompile(`^(?=.*[A-Za-z])(?=.*\d)[A-Za-z\d]{8,}$`)
	if !passwordRegex.MatchString(user.Password) {
		return errors.New("пароль должен содержать минимум 8 символов, включая буквы и цифры")
	}

	// Создание пользователя
	if err := db.Create(user).Error; err != nil {
		return err
	}
	log.Println("Пользователь успешно создан")
	return nil
}

// Логин пользователя
func LoginUser(db *gorm.DB, email, password string) (*models.User, error) {
	var user models.User
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, errors.New("пользователь не найден")
	}

	// Проверка пароля
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("неверный пароль")
	}

	return &user, nil
}

// Получение всех пользователей
func GetAllUsers(db *gorm.DB) ([]models.User, error) {
	var users []models.User
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// Получение пользователя по ID
func GetUserByID(db *gorm.DB, id string) (*models.User, error) {
	var user models.User
	if err := db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// Обновление пользователя
func UpdateUser(db *gorm.DB, id string, updated *models.User) error {
	var user models.User
	if err := db.First(&user, id).Error; err != nil {
		return err
	}

	return db.Model(&user).Updates(updated).Error
}

// Удаление пользователя
func DeleteUser(db *gorm.DB, id string) error {
	return db.Delete(&models.User{}, id).Error
}
