package handler

import (
	"final_project/internal/application/authapp"
	"final_project/internal/domain/auth"
	"final_project/internal/domain/user"
	authdto "final_project/internal/dto/authDTO"
	userdto "final_project/internal/dto/userDTO"
	"final_project/internal/pkg/enums"
	"final_project/internal/pkg/helpers"
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

// @Summary Admin Login
// @Description Đăng nhập admin với email và mật khẩu mạnh
// @Security BearerAuth
// @Tags auth
// @Accept json
// @Produce json
// @Param login body authdto.LoginRequest true "Dữ liệu đăng nhập"
// @Success 200 {object} authdto.LoginResponseWrapper
// @Failure 400 {object} enums.AppError
// @Failure 401 {object} enums.AppError
// @Router /login [post]
func (h *AuthHandler) AdminLogin(c *gin.Context) {
	var (
		req             authdto.LoginRequest
		domainAuthLogin auth.AuthLogin
		JWT             string
		refreshToken    string
		domainUser      user.User
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrValidate))
		return
	}

	if err := validator.Validate.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrValidate))
		return
	}

	domainAuthLogin = authdto.AuthDTOToDomain(req)

	if err := h.uc.Login(c.Request.Context(), domainAuthLogin, &JWT, &refreshToken, &domainUser, true); err != nil {
		c.JSON(http.StatusUnauthorized, enums.NewAppError(http.StatusUnauthorized, err.Error(), enums.ErrUnauthorized))
		return
	}

	userDTO := userdto.DomainCommonUserToDTO(domainUser)

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Login successfully",
		"data": authdto.LoginResponse{
			JWT:          JWT,
			RefreshToken: refreshToken,
			User:         userDTO,
		},
	})
}

// @Summary Client Login
// @Description Đăng nhập client với email và mật khẩu mạnh
// @Security BearerAuth
// @Tags auth
// @Accept json
// @Produce json
// @Param login body authdto.LoginRequest true "Dữ liệu đăng nhập"
// @Success 200 {object} authdto.LoginResponseWrapper
// @Failure 400 {object} enums.AppError
// @Failure 401 {object} enums.AppError
// @Router /client/login [post]
func (h *AuthHandler) UserLogin(c *gin.Context) {
	var (
		req             authdto.LoginRequest
		domainAuthLogin auth.AuthLogin
		JWT             string
		refreshToken    string
		domainUser      user.User
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrValidate))
		return
	}

	if err := validator.Validate.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrValidate))
		return
	}

	domainAuthLogin = authdto.AuthDTOToDomain(req)

	if err := h.uc.Login(c.Request.Context(), domainAuthLogin, &JWT, &refreshToken, &domainUser, false); err != nil {
		c.JSON(http.StatusUnauthorized, enums.NewAppError(http.StatusUnauthorized, err.Error(), enums.ErrUnauthorized))
		return
	}

	userDTO := userdto.DomainCommonUserToDTO(domainUser)

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Login successfully",
		"data": authdto.LoginResponse{
			JWT:          JWT,
			RefreshToken: refreshToken,
			User:         userDTO,
		},
	})
}

// @Summary Refresh Token
// @Description Lấy access token
// @Tags auth
// @Accept json
// @Produce json
// @Param request body authdto.GetAccessTokenRequest true "Refresh Token"
// @Success 200 {object} authdto.GetAccessTokenResponseWrapper
// @Failure 400 {object} enums.AppError
// @Failure 401 {object} enums.AppError
// @Router /refresh-token [post]
func (h *AuthHandler) GetAccessToken(c *gin.Context) {
	var (
		req authdto.GetAccessTokenRequest
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrValidate),
		)
		return
	}

	if err := helpers.CheckJWT(c.Request.Context(), req.RefreshToken); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"error":   enums.ErrUnauthorized,
			"message": err.Error(),
		})

		return
	}

	JWTSub := helpers.GetTokenSubject(req.RefreshToken)

	jwt := helpers.GenerateToken(JWTSub)

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Get access token successfully",
		"data": authdto.GetAccessTokenResponse{
			JWT: jwt,
		},
	})
}

// @Summary Admin Logout
// @Description Đăng xuất dành cho admin
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {object} authdto.LogoutResponseWrapper
// @Failure 400 {object} enums.AppError
// @Failure 401 {object} enums.AppError
// @Router /logout [post]
func (h *AuthHandler) AdminLogout(c *gin.Context) {
	userID, err := helpers.GetUintFromContext(c, "userID")
	if err != nil {
		c.JSON(http.StatusBadRequest, enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrBadRequest))
		return
	}

	device, err := helpers.GetStringFromContext(c, "device")
	if err != nil {
		c.JSON(http.StatusBadRequest, enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrBadRequest))
		return
	}

	if err := h.uc.Logout(c.Request.Context(), userID, device, true); err != nil {
		c.JSON(http.StatusBadRequest, enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrBadRequest))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Logout successfully",
	})
}

// @Summary Client Logout
// @Description Đăng xuất dành cho client
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {object} authdto.LogoutResponseWrapper
// @Failure 400 {object} enums.AppError
// @Failure 401 {object} enums.AppError
// @Router /logout [post]
func (h *AuthHandler) ClientLogout(c *gin.Context) {
	userID, err := helpers.GetUintFromContext(c, "userID")
	if err != nil {
		c.JSON(http.StatusBadRequest, enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrBadRequest))
		return
	}

	device, err := helpers.GetStringFromContext(c, "device")
	if err != nil {
		c.JSON(http.StatusBadRequest, enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrBadRequest))
		return
	}

	if err := h.uc.Logout(c.Request.Context(), userID, device, false); err != nil {
		c.JSON(http.StatusBadRequest, enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrBadRequest))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Logout successfully",
	})
}
