package repositories

import (
	"context"
	"database/sql"
	"user_services/internal/models"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	GetByID(ctx context.Context, id uint) (*models.User, error)
	GetByAllUsers(ctx context.Context) ([]*models.User, error)
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, id uint) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, user *models.User) error {
	query := `INSERT INTO users (first_name, last_name, email, password, city) 
		VALUES ($1, $2, $3, $4, $5) RETURNING id`

	// Сохраняем пользователя с уже хешированным паролем
	return r.db.QueryRowContext(ctx, query,
		user.FirstName, user.LastName, user.Email, user.Password, user.City).
		Scan(&user.ID)
}
func (r *userRepository) GetByID(ctx context.Context, id uint) (*models.User, error) {
	var user models.User
	query := `SELECT id, first_name, last_name, email, password, city 
		FROM users WHERE id = $1`

	row := r.db.QueryRowContext(ctx, query, id)
	if err := row.Scan(&user.ID, &user.FirstName, &user.LastName,
		&user.Email, &user.Password, &user.City); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}
func (r *userRepository) GetByAllUsers(ctx context.Context) ([]*models.User, error) {
	query := `SELECT id, first_name, last_name, email, password, city 
		FROM users`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.FirstName,
			&user.LastName, &user.Email, &user.Password, &user.City); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}
func (r *userRepository) Update(ctx context.Context, user *models.User) error {
	query := `UPDATE users SET first_name = $1, last_name = $2, email = $3, 
		city = $4 WHERE id = $5`

	_, err := r.db.ExecContext(ctx, query,
		user.FirstName, user.LastName, user.Email, user.City, user.ID)
	return err
}
func (r *userRepository) Delete(ctx context.Context, id uint) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
