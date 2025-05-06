package services

import (
	"context"
	"go_plants/internal/models"
	"go_plants/mocks"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/golang/mock/gomock"
)

func TestUserService(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := mocks.NewMockUserService(ctrl)
	ctx := context.Background()

	// Шаблонные данные
	user1 := &models.User{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "johndoe@example.com",
		Password:  "securepassword123",
		City:      "New York",
	}
	user2 := &models.User{
		ID:        2,
		FirstName: "Jane",
		LastName:  "Smith",
		Email:     "janesmith@example.com",
		Password:  "anotherpassword456",
		City:      "Los Angeles",
	}
	users := []*models.User{user1, user2}

	// ===== Тестирование Create =====
	mockUserService.EXPECT().Create(ctx, user1).Return(nil)

	err := mockUserService.Create(ctx, user1)
	assert.NoError(t, err, "Create should not return error")

	// ===== Тестирование GetByID =====
	mockUserService.EXPECT().GetByID(ctx, uint(1)).Return(user1, nil)

	foundUser, err := mockUserService.GetByID(ctx, 1)
	assert.NoError(t, err)
	assert.Equal(t, "John", foundUser.FirstName)
	assert.Equal(t, "Doe", foundUser.LastName)
	assert.Equal(t, "johndoe@example.com", foundUser.Email)

	// ===== Тестирование GetByAllUsers =====
	mockUserService.EXPECT().GetByAllUsers(ctx).Return(users, nil).Times(1)

	resultUsers, err := mockUserService.GetByAllUsers(ctx)
	assert.NoError(t, err)
	assert.Len(t, resultUsers, 2)
	assert.Equal(t, "John", resultUsers[0].FirstName)
	assert.Equal(t, "Smith", resultUsers[1].LastName)
	assert.Equal(t, "janesmith@example.com", resultUsers[1].Email)
	assert.Equal(t, "anotherpassword456", resultUsers[1].Password)
	assert.Equal(t, "Los Angeles", resultUsers[1].City)

	// ===== Тестирование Delete =====
	mockUserService.EXPECT().Delete(ctx, uint(1)).Return(nil)

	err = mockUserService.Delete(ctx, 1)
	assert.NoError(t, err, "Delete should not return error")
}
