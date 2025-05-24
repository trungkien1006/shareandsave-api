package handler

import (
	"final_project/internal/application/userapp"
	"final_project/internal/domain/filter"
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
// @Param page query int false "Current page" minimum(1) example(1)
// @Param limit query int false "Number record of page" minimum(1) example(10)
// @Param sort query string false "Sort column (vd: name, email)" example(name)
// @Param order query string false "Sort type: ASC hoặc DESC" enum(ASC,DESC) example(ASC)
// @Param filter query string false "Filter by name or email" example("{\"name\":\"John\", \"email\":\"john@gmail.com\"}")
// @Success 200 {object} userDTO.GetUserResponseWrapper
// @Failure 400 {object} enums.AppError
// @Router /users [get]
func (h *UserHandler) GetAllUser(c *gin.Context) {
	var req userDTO.GetUserRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			enums.NewAppError(http.StatusBadRequest, err.Error(), "ERR_VALIDATION"),
		)
		return
	}

	req.SetDefaults()

	var users []user.User

	var domain_req filter.FilterRequest

	domain_req.Page = req.Page
	domain_req.Limit = req.Limit
	domain_req.Sort = req.Sort
	domain_req.Order = req.Order
	domain_req.Filter = req.Filter

	totalPage, err := h.uc.GetAllUser(c.Request.Context(), &users, domain_req)

	if err != nil {
		c.JSON(
			http.StatusNotFound,
			enums.NewAppError(http.StatusNotFound, err.Error(), "ERR_USER_NOT_FOUND"),
		)
		return
	}

	usersDTORes := make([]userDTO.UserDTO, 0)

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
// @Description API get user by id
// @Tags users
// @Accept json
// @Produce json
// @Param userID path int true "ID user"
// @Success 200 {object} userDTO.GetUserByIDResponseWrapper
// @Failure 400 {object} enums.AppError
// @Router /users/{userID} [get]
func (h *UserHandler) GetUserByID(c *gin.Context) {
	var req userDTO.GetUserByIDRequest

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrValidate),
		)
		return
	}

	var user user.User

	if err := h.uc.GetUserByID(c.Request.Context(), &user, req.UserID); err != nil {
		c.JSON(
			http.StatusNotFound,
			enums.NewAppError(http.StatusNotFound, err.Error(), enums.ErrNotFound),
		)
		return
	}

	var userDTORes userDTO.UserDTO

	userDTORes = userDTO.ToUserDTO(user)

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Fetched user successfully",
		"data": userDTO.GetUserByIDResponse{
			User: userDTORes,
		},
	})
}

// @Summary Create user
// @Description API thêm người dùng
// @Tags users
// @Accept json
// @Produce json
// @Param request body userDTO.CreateUserRequest true "Create user info"
// @Success 201 {object} userDTO.CreateUserResponseWrapper "Created user successfully"
// @Failure 400 {object} enums.AppError
// @Router /users [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	var req userDTO.CreateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrValidate),
		)
		return
	}

	var user user.User

	user.Email = req.Email
	user.Password = req.Password
	user.Avatar = "" // nếu có field Avatar
	user.FullName = req.FullName
	user.PhoneNumber = req.PhoneNumber
	user.Address = req.Address
	user.Status = int(req.Status)
	user.GoodPoint = req.GoodPoint

	if err := h.uc.CreateUser(c.Request.Context(), &user); err != nil {
		c.JSON(
			http.StatusConflict,
			enums.NewAppError(http.StatusConflict, err.Error(), enums.ErrConflict),
		)
		return
	}

	var userDTORes userDTO.UserDTO

	userDTORes = userDTO.ToUserDTO(user)

	c.JSON(http.StatusCreated, gin.H{
		"code":    http.StatusCreated,
		"message": "Created user successfully",
		"data": userDTO.CreateUserResponse{
			User: userDTORes,
		},
	})
}

// @Summary Update user
// @Description API cập nhật người dùng
// @Tags users
// @Accept json
// @Produce json
// @Param request body userDTO.UpdateUserRequest true "Update user info"
// @Success 200 {object} userDTO.UpdateUserResponseWrapper "Updated user successfully"
// @Failure 400 {object} enums.AppError
// @Router /users [put]
func (h *UserHandler) UpdateUser(c *gin.Context) {
	var req userDTO.UpdateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrValidate),
		)
		return
	}

	var user user.User

	user.ID = req.ID
	user.Email = req.Email
	user.Password = req.Password
	user.Avatar = req.Avatar
	user.FullName = req.FullName
	user.PhoneNumber = req.PhoneNumber
	user.Address = req.Address
	user.Status = int(req.Status)
	user.GoodPoint = req.GoodPoint

	if err := h.uc.UpdateUser(c.Request.Context(), &user); err != nil {
		c.JSON(
			http.StatusConflict,
			enums.NewAppError(http.StatusConflict, err.Error(), enums.ErrConflict),
		)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Updated user successfully",
	})
}
