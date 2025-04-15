package auth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware извлекает user_id из JWT и сохраняет в context
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Токен не передан"})
			c.Abort()
			return
		}

		fmt.Println("Received token:", token) // Добавьте логирование

		// Убираем "Bearer " из строки токена
		token = token[len("Bearer "):]

		userID, err := ParseToken(token)
		if err != nil {
			fmt.Println("Token parsing error:", err) // Логируем ошибку парсинга
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Некорректный токен"})
			c.Abort()
			return
		}

		// Сохраняем userID в контексте, чтобы его можно было использовать дальше
		c.Set("user_id", userID)
		c.Next()
	}
}
