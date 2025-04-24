package main

import (
	"log"
	"os"

	"go_plants/config"
	"go_plants/internal/admin"
	"go_plants/internal/user"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database
	db := config.InitDB()

	// Create router
	r := gin.Default()

	// Register routes
	user.RegisterUserRoutes(r, db)
	admin.RegisterAdminHandlers(r, db)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
