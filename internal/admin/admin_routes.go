package admin

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterAdminHandlers registers all admin-related routes
func RegisterAdminHandlers(r *gin.Engine, db *gorm.DB) {
	// Admin routes group
	admin := r.Group("/admin")
	{
		// User management routes
		users := admin.Group("/users")
		{
			// Get total user count
			users.GET("/count", GetUserCountHandler(db))
			
			// Reset user ID auto-increment
			users.PUT("/reset-id", ResetUserAutoIncrementHandler(db))
			
			// Delete all users
			users.DELETE("", DeleteAllUsersHandler(db))
			
			// Delete specific user by ID
			users.DELETE("/:id", DeleteUserByIDHandler(db))
		}

		// Plant management routes
		plants := admin.Group("/plants")
		{
			// Create a new plant
			plants.POST("", CreatePlantHandler(db))
			
			// Get all plants
			plants.GET("", GetAllPlantsHandler(db))
			
			// Get plant by ID
			plants.GET("/:id", GetPlantByIDHandler(db))
			
			// Update plant by ID
			plants.PUT("/:id", UpdatePlantHandler(db))
			
			// Delete plant by ID
			plants.DELETE("/:id", DeletePlantHandler(db))
		}
	}
}
