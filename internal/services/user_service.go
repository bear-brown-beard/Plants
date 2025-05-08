package services

import (
	"context"
	"user_services/internal/models"
	"user_services/internal/repositories"

	"golang.org/x/crypto/bcrypt"
)

// UserService интерфейс для работы с пользователями
// Он определяет методы для создания, получения, обновления и удаления пользователей
// а также для получения всех пользователей
type UserService interface {
	Create(ctx context.Context, user *models.User) error
	GetByID(ctx context.Context, id uint) (*models.User, error)
	GetByAllUsers(ctx context.Context) ([]*models.User, error)
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, id uint) error
}

// userService структура которая реализует интерфейс UserService
// и содержит ссылку на репозиторий пользователей
type userService struct {
	userRepository repositories.UserRepository
}

// NewUserService конструктор для создания нового сервиса пользователя
func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

func (s *userService) Create(ctx context.Context, user *models.User) error {
	// Hash the password before saving to database
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	return s.userRepository.Create(ctx, user)
}
func (s *userService) GetByID(ctx context.Context, id uint) (*models.User, error) {
	return s.userRepository.GetByID(ctx, id)
}
func (s *userService) GetByAllUsers(ctx context.Context) ([]*models.User, error) {
	return s.userRepository.GetByAllUsers(ctx)
}
func (s *userService) Update(ctx context.Context, user *models.User) error {
	return s.userRepository.Update(ctx, user)
}
func (s *userService) Delete(ctx context.Context, id uint) error {
	return s.userRepository.Delete(ctx, id)
}
