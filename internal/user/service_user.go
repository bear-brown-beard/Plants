package user

import (
	"errors"
	"log"
	"unicode"

	"golang.org/x/crypto/bcrypt"

	"gorm.io/gorm"
)

// Создание пользователя
func CreateUser(db *gorm.DB, user *User) error {
	// Проверка пароля
	if user.Password == "" {
		return errors.New("пароль не может быть пустым")
	}
	if !isPasswordValid(user.Password) {
		return errors.New("пароль должен содержать минимум 8 символов, включая хотя бы одну букву и одну цифру")
	}

	// Хешируем пароль
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Ошибка при хешировании пароля:", err)
		return err
	}
	user.Password = string(hashedPassword)

	// Проверка на существующий email
	var existingUser User
	if err := db.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		log.Println("Ошибка: пользователь с таким email уже существует")
		return errors.New("пользователь с таким email уже существует")
	}

	// Создание пользователя
	if err := db.Create(user).Error; err != nil {
		return err
	}

	log.Println("Пользователь успешно создан")
	return nil
}

// Логин пользователя
func AuthenticateUser(db *gorm.DB, email, password string) (*User, error) {
	var user User
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
func GetAllUsers(db *gorm.DB) ([]User, error) {
	var users []User
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// Получение пользователя по ID
func GetUserByID(db *gorm.DB, id string) (*User, error) {
	var user User
	if err := db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// Обновление пользователя
func UpdateUser(db *gorm.DB, id string, updated *User) error {
	var user User
	if err := db.First(&user, id).Error; err != nil {
		return err
	}

	return db.Model(&user).Updates(updated).Error
}

// Удаление пользователя
func DeleteUser(db *gorm.DB, id string) error {
	return db.Delete(&User{}, id).Error
}

func isPasswordValid(password string) bool {
	if len(password) < 8 {
		return false
	}

	hasLetter := false
	hasDigit := false

	for _, c := range password {
		if unicode.IsLetter(c) {
			hasLetter = true
		} else if unicode.IsDigit(c) {
			hasDigit = true
		}
	}

	return hasLetter && hasDigit
}
