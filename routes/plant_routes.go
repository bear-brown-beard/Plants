package routes

import (
	"go_plants/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterPlantAPI(r *gin.Engine, db *gorm.DB) {
	handlers.RegisterPlantHandlers(r, db)
}
