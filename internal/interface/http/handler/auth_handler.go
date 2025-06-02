package handler

import (
	"final_project/internal/application/authapp"
	authdto "final_project/internal/dto/authDTO"
	"final_project/internal/pkg/enums"
	"final_project/internal/shared/validator"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	uc *authapp.UseCase
}

func NewAuthHandler(uc *authapp.UseCase) *AuthHandler {
	return &AuthHandler{uc: uc}
}

// @Summary Login
// @Description Đăng nhập người dùng với email và mật khẩu mạnh
// @Tags auth
// @Accept json
// @Produce json
// @Param login body authdto.LoginRequest true "Dữ liệu đăng nhập"
// @Router /login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var req authdto.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrValidate))
		return
	}

	if err := validator.Validate.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrValidate))
		return
	}

}
