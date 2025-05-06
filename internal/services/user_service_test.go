package services

import (
	"context"
	"testing"

	"go_plants/internal/models"
	"go_plants/internal/repositories/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUserService(t *testing.T) {
	// Create mock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create mock repository
	mockRepo := mocks.NewMockUserRepository(ctrl)
	userService := NewUserService(mockRepo)

	// Test Create
	t.Run("CreateUser", func(t *testing.T) {
		user := &models.User{
			ID:        1,
			FirstName: "testuser",
			Password:  "password",
		}

		mockRepo.EXPECT().Create(context.Background(), user).Return(nil)
		err := userService.Create(context.Background(), user)
		assert.NoError(t, err)
	})

	// Test GetByID
	t.Run("GetUserByID", func(t *testing.T) {
		mockUser := &models.User{
			ID:        1,
			FirstName: "testuser",
		}

		mockRepo.EXPECT().GetByID(context.Background(), uint(1)).Return(mockUser, nil)
		user, err := userService.GetByID(context.Background(), uint(1))
		assert.NoError(t, err)
		assert.Equal(t, mockUser, user)
	})

	// Test GetByAllUsers
	t.Run("GetAllUsers", func(t *testing.T) {
		users := []*models.User{
			{ID: 1, FirstName: "user1"},
			{ID: 2, FirstName: "user2"},
		}

		mockRepo.EXPECT().GetByAllUsers(context.Background()).Return(users, nil)
		result, err := userService.GetByAllUsers(context.Background())
		assert.NoError(t, err)
		assert.Equal(t, users, result)
	})

	// Test Update
	t.Run("UpdateUser", func(t *testing.T) {
		user := &models.User{
			ID:        1,
			FirstName: "updateduser",
		}

		mockRepo.EXPECT().Update(context.Background(), user).Return(nil)
		err := userService.Update(context.Background(), user)
		assert.NoError(t, err)
	})

	// Test Delete
	t.Run("DeleteUser", func(t *testing.T) {
		mockRepo.EXPECT().Delete(context.Background(), uint(1)).Return(nil)
		err := userService.Delete(context.Background(), uint(1))
		assert.NoError(t, err)
	})
}
