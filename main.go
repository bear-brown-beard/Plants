package main

import (
	"go_plants/internal/adapters/api"
	"go_plants/internal/config"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()
	db := config.DB

	r := gin.Default()

	// Регистрируем маршруты
	api.RegisterUserAPI(r, db)
	api.RegisterPlantAPI(r, db)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Ошибка при запуске сервера: ", err)
	}
}
