package services

import "sidecar-template/pkg/models"

type UserService struct {
	// Add any dependencies, e.g., database client
}

func (us *UserService) GetUserByID(userID int) (*models.User, error) {
	// Return a dummy user for demonstration
	return &models.User{ID: userID, Name: "John Doe", Email: "john@example.com", FavoriteColor: "blue"}, nil
}
