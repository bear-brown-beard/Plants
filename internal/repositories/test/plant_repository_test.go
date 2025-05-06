package repositories

import (
	"context"
	"database/sql"
	"testing"

	"go_plants/internal/models"
	"go_plants/internal/repositories"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func setupTestPlantsDB(t *testing.T) *sql.DB {
	connStr := "user=plants_user dbname=plants_db sslmode=disable port=5432 host=localhost password=secret"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		t.Fatalf("Failed to connect to test DB: %v", err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		first_name TEXT NOT NULL,
		last_name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL,
		city TEXT
	);`)
	if err != nil {
		t.Fatalf("Failed to create users table: %v", err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS plants (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		description TEXT,
		watering TEXT,
		repotting TEXT,
		breeding TEXT,
		user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE
	);`)
	if err != nil {
		t.Fatalf("Failed to create plants table: %v", err)
	}

	_, err = db.Exec("TRUNCATE TABLE plants RESTART IDENTITY CASCADE;")
	assert.NoError(t, err)
	_, err = db.Exec("TRUNCATE TABLE users RESTART IDENTITY CASCADE;")
	assert.NoError(t, err)

	_, err = db.Exec(`INSERT INTO users (first_name, last_name, email, password) 
		VALUES ('Test', 'User', 'test@example.com', 'password')`)
	assert.NoError(t, err)

	return db
}

func getTestUserID(t *testing.T, db *sql.DB) uint {
	var id uint
	err := db.QueryRow("SELECT id FROM users WHERE email = 'test@example.com'").Scan(&id)
	assert.NoError(t, err)
	return id
}

func seedTestPlants(t *testing.T, repo repositories.PlantRepository, userID uint) []*models.Plant {
	plants := []*models.Plant{
		{Name: "Ficus", Description: "Indoor tree", Watering: "Weekly", Repotting: "Yearly", Breeding: "Cuttings", UserID: userID},
		{Name: "Aloe", Description: "Medicinal plant", Watering: "Bi-weekly", Repotting: "Every 2 years", Breeding: "Offshoots", UserID: userID},
	}
	for _, p := range plants {
		err := repo.Create(context.Background(), p)
		assert.NoError(t, err)
	}
	return plants
}

func TestPlantRepository(t *testing.T) {
	db := setupTestPlantsDB(t)
	defer db.Close()
	userID := getTestUserID(t, db)
	repo := repositories.NewPlantRepository(db)
	seedTestPlants(t, repo, userID)

	t.Run("GetAllPlants", func(t *testing.T) {
		plants, err := repo.GetAll(context.Background())
		assert.NoError(t, err)
		assert.GreaterOrEqual(t, len(plants), 2) // Ожидаем, что должно быть хотя бы два растения
	})

	t.Run("GetPlantByID", func(t *testing.T) {
		plants, err := repo.GetAll(context.Background())
		assert.NoError(t, err)
		assert.GreaterOrEqual(t, len(plants), 2)

		plant, err := repo.GetByID(context.Background(), plants[0].ID, userID)
		assert.NoError(t, err)
		assert.NotNil(t, plant)
		assert.Equal(t, plants[0].ID, plant.ID)
	})

	t.Run("CreatePlant", func(t *testing.T) {
		newPlant := &models.Plant{
			Name:        "Cactus",
			Description: "Desert plant",
			Watering:    "Once a month",
			Repotting:   "Every 3 years",
			Breeding:    "Seeds",
			UserID:      userID,
		}
		err := repo.Create(context.Background(), newPlant)
		assert.NoError(t, err)
		assert.NotZero(t, newPlant.ID) // ID должен быть установлен
	})

	t.Run("UpdatePlant", func(t *testing.T) {
		plants, err := repo.GetAll(context.Background())
		assert.NoError(t, err)
		plants[0].Name = "Updated Ficus"
		err = repo.Update(context.Background(), plants[0].ID, userID, plants[0])
		assert.NoError(t, err)
		updatedPlant, err := repo.GetByID(context.Background(), plants[0].ID, userID)
		assert.NoError(t, err)
		assert.Equal(t, "Updated Ficus", updatedPlant.Name)
	})

	t.Run("DeletePlant", func(t *testing.T) {
		plants, err := repo.GetAll(context.Background())
		assert.NoError(t, err)
		err = repo.Delete(context.Background(), plants[0].ID, userID)
		assert.NoError(t, err)
		deletedPlant, err := repo.GetByID(context.Background(), plants[0].ID, userID)
		assert.NoError(t, err)
		assert.Nil(t, deletedPlant) // Растение должно быть удалено
	})
}
