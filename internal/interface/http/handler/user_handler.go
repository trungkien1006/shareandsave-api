package handler

import (
	"final-project/internal/application/userapp"
	"final-project/internal/domain/user"
	"final-project/internal/dto/userDTO"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	uc *userapp.UseCase
}

func NewUserHandler(uc *userapp.UseCase) *UserHandler {
	return &UserHandler{uc: uc}
}

// type CreateUserRequest struct {
// 	Name     string `json:"name"`
// 	Email    string `json:"email"`
// 	Password string `json:"password"`
// }

func (h *UserHandler) GetAllUser(c *gin.Context) {
	var req userDTO.GetUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user []user.User

	if err := h.uc.GetAllUser(c.Request.Context(), &user); err != nil {
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user created"})
}
