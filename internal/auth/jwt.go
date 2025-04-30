package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("секретный_ключ") // можно поменять, желательно вынести в переменные окружения

// Генерация JWT токена
func GenerateToken(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(72 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// Проверка и парсинг токена
func ParseToken(tokenString string) (uint, error) {
	fmt.Println("Parsing token:", tokenString) // Логируем перед парсингом

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Проверка метода подписи
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("неподдерживаемый метод подписи")
		}
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		fmt.Println("Token is invalid:", err) // Логируем ошибку
		return 0, errors.New("некорректный токен")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println("Error extracting claims") // Логируем ошибку
		return 0, errors.New("ошибка при извлечении claims")
	}

	userIDFloat, ok := claims["user_id"].(float64)
	if !ok {
		fmt.Println("Error converting user_id") // Логируем ошибку
		return 0, errors.New("user_id не найден в токене")
	}

	return uint(userIDFloat), nil
}
