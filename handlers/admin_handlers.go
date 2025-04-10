package handlers

import (
	"go_plants/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterAdminHandlers(r *gin.Engine, db *gorm.DB) {
	r.DELETE("/admin/plants", func(c *gin.Context) {
		if err := services.DeleteAllPlants(db); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось удалить все растения"})
			return
		}
		c.Status(http.StatusNoContent)
	})

	r.GET("/admin/plants/count", func(c *gin.Context) {
		count, err := services.CountPlants(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить количество растений"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"count": count})
	})
}
