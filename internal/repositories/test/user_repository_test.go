package repositories

import (
	"context"
	"database/sql"
	"go_plants/internal/models"
	"go_plants/internal/repositories"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func setupTestUsersDB(t *testing.T) *sql.DB {
	connStr := "user=plants_user dbname=plants_db sslmode=disable port=5432 host=localhost password=secret"
	db, err := sql.Open("postgres", connStr)
	assert.NoError(t, err)

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		first_name TEXT NOT NULL,
		last_name TEXT NOT NULL,
		email TEXT UNIQUE NOT NULL,
		password TEXT NOT NULL
	);`)
	assert.NoError(t, err)

	_, err = db.Exec("TRUNCATE TABLE users RESTART IDENTITY CASCADE;")
	assert.NoError(t, err)

	return db
}

func seedTestUsers(t *testing.T, repo repositories.UserRepository) []*models.User {
	users := []*models.User{
		{FirstName: "John", LastName: "Doe", Email: "john.doe@example.com", Password: "pass123"},
		{FirstName: "Jane", LastName: "Smith", Email: "jane.smith@example.com", Password: "pass456"},
	}
	for _, user := range users {
		err := repo.Create(context.Background(), user)
		assert.NoError(t, err)
	}
	return users
}

func TestUserRepository(t *testing.T) {
	db := setupTestUsersDB(t)
	defer db.Close()

	repo := repositories.NewUserRepository(db)
	users := seedTestUsers(t, repo)

	t.Run("CreateUser", func(t *testing.T) {
		user := &models.User{
			FirstName: "Alice",
			LastName:  "Brown",
			Email:     "alice.brown@example.com",
			Password:  "password123",
		}
		err := repo.Create(context.Background(), user)
		assert.NoError(t, err)
		var count int
		err = db.QueryRow("SELECT COUNT(*) FROM users WHERE email = $1", user.Email).Scan(&count)
		assert.NoError(t, err)
		assert.Equal(t, 1, count)
	})

	t.Run("GetUserByID", func(t *testing.T) {
		user := users[0]
		foundUser, err := repo.GetByID(context.Background(), user.ID)
		assert.NoError(t, err)
		assert.NotNil(t, foundUser)
		assert.Equal(t, user.Email, foundUser.Email)
	})

	t.Run("UpdateUser", func(t *testing.T) {
		user := users[1]
		user.FirstName = "Jane Updated"
		err := repo.Update(context.Background(), user)
		assert.NoError(t, err)
		updatedUser, err := repo.GetByID(context.Background(), user.ID)
		assert.NoError(t, err)
		assert.Equal(t, "Jane Updated", updatedUser.FirstName)
	})

	t.Run("DeleteUser", func(t *testing.T) {
		user := users[0]
		err := repo.Delete(context.Background(), user.ID)
		assert.NoError(t, err)
		var count int
		err = db.QueryRow("SELECT COUNT(*) FROM users WHERE id = $1", user.ID).Scan(&count)
		assert.NoError(t, err)
		assert.Equal(t, 0, count)
	})

	t.Run("GetAllUsers", func(t *testing.T) {
		result, err := repo.GetByAllUsers(context.Background())
		assert.NoError(t, err)
		assert.GreaterOrEqual(t, len(result), 1)
	})
}
