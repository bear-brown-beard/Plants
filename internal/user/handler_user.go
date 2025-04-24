package user

import (
	"go_plants/internal/middleware"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetProfileHandler handles getting user profile
func GetProfileHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := middleware.GetUserIDFromContext(c)
		if userID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
			return
		}

		user, err := GetUserByID(db, strconv.FormatUint(uint64(userID), 10))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		c.JSON(http.StatusOK, toUserResponse(*user))
	}
}

// UpdateProfileHandler handles updating user profile
func UpdateProfileHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := middleware.GetUserIDFromContext(c)
		if userID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
			return
		}

		var req UpdateProfileRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
			return
		}

		user := User{
			FirstName: req.FirstName,
			LastName:  req.LastName,
			City:      req.City,
		}

		if err := UpdateUser(db, strconv.FormatUint(uint64(userID), 10), &user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update profile"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Profile updated successfully",
			"user":    toUserResponse(user),
		})
	}
}

// GetAllUsersHandler handles getting all users
func GetAllUsersHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := GetAllUsers(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get users"})
			return
		}

		response := make([]UserResponse, len(users))
		for i, u := range users {
			response[i] = toUserResponse(u)
		}

		c.JSON(http.StatusOK, response)
	}
}

// GetUserByIDHandler handles getting a user by ID
func GetUserByIDHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		user, err := GetUserByID(db, id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		c.JSON(http.StatusOK, toUserResponse(*user))
	}
}

// toUserResponse converts a User model to a UserResponse
func toUserResponse(user User) UserResponse {
	return UserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		//Email:     user.Email,
		City: user.City,
	}
}
