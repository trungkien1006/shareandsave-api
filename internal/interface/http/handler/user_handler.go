package handler

import (
	"final_project/internal/application/userapp"
	"final_project/internal/domain/user"
	"final_project/internal/dto/userDTO"
	"final_project/internal/pkg/enums"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	uc *userapp.UseCase
}

func NewUserHandler(uc *userapp.UseCase) *UserHandler {
	return &UserHandler{uc: uc}
}

// @Summary Get user
// @Description API bao gồm cả lọc, phân trang và sắp xếp
// @Tags users
// @Accept json
// @Produce json
// @Param page query int false "Số trang hiện tại" minimum(1) example(1)
// @Param limit query int false "Số lượng mỗi trang" minimum(1) example(10)
// @Param sort query string false "Trường cần sắp xếp (vd: name, email)" example(name)
// @Param order query string false "Thứ tự sắp xếp: ASC hoặc DESC" Enums(ASC, DESC) example(ASC)
// @Param filter query string false "Lọc theo tên hoặc email" example("{"name": "John", "email": "john@gmail.com"}")
// @Success 200 {object} userDTO.GetUserResponseWrapper
// @Failure 400 {object} enums.AppError
// @Router /users [get]
func (h *UserHandler) GetAllUser(c *gin.Context) {
	var req userDTO.GetUserRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			enums.NewAppError("ERR_VALIDATION", err.Error(), http.StatusBadRequest),
		)
		return
	}

	var users []user.User

	totalPage, err := h.uc.GetAllUser(c.Request.Context(), &users, req)

	if err != nil {
		c.JSON(
			http.StatusNotFound,
			enums.NewAppError("ERR_USER_NOT_FOUND", err.Error(), http.StatusNotFound),
		)
		return
	}

	var usersDTORes []userDTO.UserDTO

	for _, user := range users {
		usersDTORes = append(usersDTORes, userDTO.ToUserDTO(user))
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Fetched user successfully",
		"data": userDTO.GetUserResponse{
			Users:     usersDTORes,
			TotalPage: totalPage,
		},
	})
}

// @Summary Get user by ID
// @Description API lấy ra user bằng id
// @Tags users
// @Accept json
// @Produce json
// @Param userID query int true "ID nhân viên" example(1)
// @Success 200 {object} userDTO.GetUserByIDResponseWrapper
// @Failure 400 {object} enums.AppError
// @Router /users/:userID [get]
func (h *UserHandler) GetUserByID(c *gin.Context) {
	var req userDTO.GetUserByIDRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			enums.NewAppError("ERR_VALIDATION", err.Error(), http.StatusBadRequest),
		)
		return
	}

	var user user.User

	if err := h.uc.GetUserByID(c.Request.Context(), &user, req.UserID); err != nil {
		c.JSON(
			http.StatusNotFound,
			enums.NewAppError("ERR_USER_NOT_FOUND", err.Error(), http.StatusNotFound),
		)
		return
	}

	var userDTORes []userDTO.UserDTO

	userDTORes = append(userDTORes, userDTO.ToUserDTO(user))

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Fetched user successfully",
		"data": userDTO.GetUserByIDResponse{
			User: userDTORes,
		},
	})
}
