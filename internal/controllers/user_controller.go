package controllers

import (
	"go-api-template/internal/services"
	"go-api-template/pkg/models"
)

type UserController struct {
	userService services.IUserService
}

func NewUserController(us services.IUserService) *UserController {
	return &UserController{userService: us}
}

func (uc *UserController) GetUser(userID int) (*models.User, error) {
	return uc.userService.GetUserByID(userID)
}
