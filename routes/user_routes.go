package routes

import (
	"go_plants/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterUserAPI(r *gin.Engine, db *gorm.DB) {
	handlers.RegisterUserHandlers(r, db)
}
