package routes

import (
	"go_plants/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterAdminAPI(r *gin.Engine, db *gorm.DB) {
	handlers.RegisterAdminHandlers(r, db)
}
