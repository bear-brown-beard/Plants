package admin

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetUserCountHandler handles getting the total count of users
func GetUserCountHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		count, err := CountUsers(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user count"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"count": count})
	}
}

// ResetUserAutoIncrementHandler handles resetting the user ID auto-increment
func ResetUserAutoIncrementHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := ResetUserAutoIncrement(db); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to reset user auto-increment"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "User auto-increment reset successfully"})
	}
}

// DeleteAllUsersHandler handles deleting all users
func DeleteAllUsersHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := DeleteAllUsers(db); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete all users"})
			return
		}
		c.Status(http.StatusNoContent)
	}
}

// DeleteUserByIDHandler handles deleting a specific user by ID
func DeleteUserByIDHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, err := strconv.ParseUint(c.Param("id"), 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
			return
		}

		if err := DeleteUserByID(db, uint(userID)); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
			return
		}
		c.Status(http.StatusNoContent)
	}
} 