package admin

import (
	"go_plants/internal/user"
	"log"

	"gorm.io/gorm"
)

// CountUsers returns the total number of users in the database
func CountUsers(db *gorm.DB) (int64, error) {
	var count int64
	if err := db.Model(&user.User{}).Count(&count).Error; err != nil {
		log.Println("Error counting users:", err)
		return 0, err
	}
	return count, nil
}

// DeleteAllUsers deletes all users from the database
func DeleteAllUsers(db *gorm.DB) error {
	if err := db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&user.User{}).Error; err != nil {
		log.Println("Error deleting all users:", err)
		return err
	}
	log.Println("All users deleted successfully")
	return nil
}

// DeleteUserByID deletes a specific user by ID
func DeleteUserByID(db *gorm.DB, userID uint) error {
	if err := db.Delete(&user.User{}, userID).Error; err != nil {
		log.Println("Error deleting user:", err)
		return err
	}
	log.Println("User deleted successfully")
	return nil
}

// ResetUserAutoIncrement resets the user ID auto-increment counter
func ResetUserAutoIncrement(db *gorm.DB) error {
	if err := db.Exec("ALTER SEQUENCE users_id_seq RESTART WITH 1").Error; err != nil {
		log.Println("Error resetting user auto-increment:", err)
		return err
	}
	log.Println("User auto-increment reset successfully")
	return nil
} 