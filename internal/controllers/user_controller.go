package controllers

import (
	"sidecar-template/internal/services"
	"sidecar-template/pkg/models"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController(us *services.UserService) *UserController {
	return &UserController{userService: us}
}

func (uc *UserController) GetUser(userID int) (*models.User, error) {
	return uc.userService.GetUserByID(userID)
}
