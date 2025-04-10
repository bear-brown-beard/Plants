package main

import (
	"go_plants/config"
	"go_plants/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()
	db := config.DB

	r := gin.Default()

	routes.RegisterAdminAPI(r, db)
	routes.RegisterPlantAPI(r, db)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Ошибка при запуске сервера: ", err)
	}
}
