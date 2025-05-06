package repositories

import (
	"context"
	"database/sql"
	"go_plants/internal/models"
)

type PlantRepository interface {
	Create(ctx context.Context, plant *models.Plant) error
	GetAll(ctx context.Context) ([]*models.Plant, error)
	GetByID(ctx context.Context, id uint, userID uint) (*models.Plant, error)
	Update(ctx context.Context, id uint, userID uint, plant *models.Plant) error
	Delete(ctx context.Context, id uint, userID uint) error
}

type plantRepository struct {
	db *sql.DB
}

func NewPlantRepository(db *sql.DB) PlantRepository {
	return &plantRepository{db: db}
}

func (r *plantRepository) Create(ctx context.Context, plant *models.Plant) error {
	query := `INSERT INTO plants (name, description, watering, repotting, breeding, user_id) 
		VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

	return r.db.QueryRowContext(ctx, query,
		plant.Name, plant.Description, plant.Watering, plant.Repotting,
		plant.Breeding, plant.UserID).Scan(&plant.ID)
}
func (r *plantRepository) GetAll(ctx context.Context) ([]*models.Plant, error) {
	var plants []*models.Plant
	query := `SELECT id, name, description, watering, repotting, breeding, user_id 
		FROM plants`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var plant models.Plant
		if err := rows.Scan(&plant.ID, &plant.Name, &plant.Description, &plant.Watering, &plant.Repotting, &plant.Breeding, &plant.UserID); err != nil {
			return nil, err
		}
		plants = append(plants, &plant)
	}

	return plants, nil
}
func (r *plantRepository) GetByID(ctx context.Context, id uint, userID uint) (*models.Plant, error) {
	var plant models.Plant
	query := `SELECT id, name, description, watering, repotting, breeding, user_id 
		FROM plants WHERE id = $1 AND user_id = $2`

	row := r.db.QueryRowContext(ctx, query, id, userID)
	if err := row.Scan(&plant.ID, &plant.Name, &plant.Description,
		&plant.Watering, &plant.Repotting, &plant.Breeding, &plant.UserID); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &plant, nil
}
func (r *plantRepository) Update(ctx context.Context, id uint, userID uint, plant *models.Plant) error {
	query := `UPDATE plants SET name = $1, description = $2, watering = $3, 
		repotting = $4, breeding = $5 WHERE id = $6 AND user_id = $7`

	_, err := r.db.ExecContext(ctx, query,
		plant.Name, plant.Description, plant.Watering,
		plant.Repotting, plant.Breeding, id, userID)
	return err
}
func (r *plantRepository) Delete(ctx context.Context, id uint, userID uint) error {
	query := `DELETE FROM plants WHERE id = $1 AND user_id = $2`
	_, err := r.db.ExecContext(ctx, query, id, userID)
	return err
}
