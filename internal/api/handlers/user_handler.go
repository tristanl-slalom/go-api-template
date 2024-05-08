package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go-api-template/internal/controllers"
)

type UserHandler struct {
	userController *controllers.UserController
}

func NewUserHandler(uc *controllers.UserController) *UserHandler {
	return &UserHandler{userController: uc}
}

// GetUser godoc
// @Summary Retrieve a user
// @Description get user by ID
// @ID get-user
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} models.User
// @Router /users/{id} [get]
func (uh *UserHandler) GetUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	user, err := uh.userController.GetUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error getting user"})
		return
	}

	c.JSON(http.StatusOK, user)
}
