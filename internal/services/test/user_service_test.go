package services

import (
	"context"
	"errors"
	"testing"

	"user_services/internal/models"
	"user_services/internal/repositories/mocks"
	"user_services/internal/services"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUserService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepository(ctrl)
	userService := services.NewUserService(mockRepo)

	t.Run("CreateUser_Success", func(t *testing.T) {
		user := &models.User{
			ID:        1,
			FirstName: "testuser",
			Password:  "password",
		}

		mockRepo.EXPECT().Create(context.Background(), user).Return(nil)
		err := userService.Create(context.Background(), user)

		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, "testuser", user.FirstName)
	})

	t.Run("GetUserByID_Success", func(t *testing.T) {
		mockUser := &models.User{
			ID:        1,
			FirstName: "testuser",
		}

		mockRepo.EXPECT().GetByID(context.Background(), uint(1)).Return(mockUser, nil)
		user, err := userService.GetByID(context.Background(), uint(1))

		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, mockUser.ID, user.ID)
		assert.Equal(t, "testuser", user.FirstName)
	})

	t.Run("GetUserByID_NotFound", func(t *testing.T) {
		mockRepo.EXPECT().GetByID(context.Background(), uint(99)).Return(nil, errors.New("user not found"))
		user, err := userService.GetByID(context.Background(), uint(99))

		assert.Error(t, err)
		assert.Nil(t, user)
		assert.EqualError(t, err, "user not found")
	})

	t.Run("GetAllUsers_Success", func(t *testing.T) {
		users := []*models.User{
			{ID: 1, FirstName: "user1"},
			{ID: 2, FirstName: "user2"},
		}

		mockRepo.EXPECT().GetByAllUsers(context.Background()).Return(users, nil)
		result, err := userService.GetByAllUsers(context.Background())

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Len(t, result, 2)
		assert.Equal(t, "user1", result[0].FirstName)
	})

	t.Run("UpdateUser_Success", func(t *testing.T) {
		user := &models.User{
			ID:        1,
			FirstName: "updateduser",
		}

		mockRepo.EXPECT().Update(context.Background(), user).Return(nil)
		err := userService.Update(context.Background(), user)

		assert.NoError(t, err)
		assert.Equal(t, "updateduser", user.FirstName)
	})

	t.Run("DeleteUser_Success", func(t *testing.T) {
		mockRepo.EXPECT().Delete(context.Background(), uint(1)).Return(nil)
		err := userService.Delete(context.Background(), uint(1))

		assert.NoError(t, err)
	})
}

func TestUserServiceErrors(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepository(ctrl)
	userService := services.NewUserService(mockRepo)

	// Test: User not found
	t.Run("GetUserByID_UserNotFound", func(t *testing.T) {
		mockRepo.EXPECT().GetByID(context.Background(), uint(1)).Return(nil, errors.New("user not found"))
		user, err := userService.GetByID(context.Background(), uint(1))
		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Equal(t, "user not found", err.Error())
	})

	// Test: Create User with empty name
	t.Run("CreateUser_EmptyName", func(t *testing.T) {
		user := &models.User{ID: 1, FirstName: ""}
		mockRepo.EXPECT().Create(context.Background(), user).Return(errors.New("first name cannot be empty"))

		err := userService.Create(context.Background(), user)
		assert.Error(t, err)
		assert.Equal(t, "first name cannot be empty", err.Error())
	})
	// Test: Update User with very large name
	t.Run("UpdateUser_LargeName", func(t *testing.T) {
		user := &models.User{ID: 1, FirstName: string(make([]byte, 1000))}
		mockRepo.EXPECT().Update(context.Background(), user).Return(nil)
		err := userService.Update(context.Background(), user)
		assert.NoError(t, err)
	})
}

func TestUserServiceConcurrency(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepository(ctrl)
	userService := services.NewUserService(mockRepo)

	user := &models.User{ID: 1, FirstName: "Test"}

	// Test: Two updates in parallel
	t.Run("UpdateUser_Concurrency", func(t *testing.T) {
		mockRepo.EXPECT().Update(context.Background(), user).Return(nil).Times(2)

		errChan := make(chan error, 2)

		// Simulating two concurrent updates
		go func() {
			errChan <- userService.Update(context.Background(), user)
		}()
		go func() {
			errChan <- userService.Update(context.Background(), user)
		}()

		err1 := <-errChan
		err2 := <-errChan

		// Ensure no errors happened in both operations
		assert.NoError(t, err1)
		assert.NoError(t, err2)
	})
}
