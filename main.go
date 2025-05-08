package main

import (
	"user_services/internal/adapters/api"
	"user_services/internal/config"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()
	db := config.DB

	r := gin.Default()

	api.RegisterUserAPI(r, db)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Ошибка при запуске сервера: ", err)
	}
}
