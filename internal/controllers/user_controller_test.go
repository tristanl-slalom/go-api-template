package controllers

import (
	"go-api-template/pkg/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock for UserService
type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) GetUserByID(userID int) (*models.User, error) {
	args := m.Called(userID)
	return args.Get(0).(*models.User), args.Error(1)
}

func TestGetUser(t *testing.T) {
	mockService := new(MockUserService)
	userController := NewUserController(mockService)

	mockService.On("GetUserByID", 123).Return(&models.User{ID: 123, Name: "John Doe", Email: "john@example.com", FavoriteColor: "blue"}, nil)

	user, err := userController.GetUser(123)

	assert.NoError(t, err)
	assert.NotNil(t, user)
}
