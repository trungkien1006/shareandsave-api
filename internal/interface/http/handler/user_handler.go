package handler

import (
	"final_project/internal/application/userapp"
	"final_project/internal/domain/filter"
	"final_project/internal/domain/user"
	userdto "final_project/internal/dto/userDTO"
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

// @Summary Get users
// @Description API bao gồm cả lọc, phân trang và sắp xếp
// @Tags users
// @Accept json
// @Produce json
// @Param page query int false "Current page" minimum(1) example(1)
// @Param limit query int false "Number record of page" minimum(1) example(10)
// @Param sort query string false "Sort column (vd: fullName, email)" example(name)
// @Param order query string false "Sort type: ASC hoặc DESC" enum(ASC,DESC) example(ASC)
// @Param   searchBy   query    string  false  "Trường lọc (vd: email, fullName)"
// @Param   searchValue   query    string  false  "Giá trị lọc (vd:abc@gmail.com, John Doe)"
// @Success 200 {object} userdto.GetUserResponseWrapper
// @Failure 400 {object} enums.AppError
// @Router /users [get]
func (h *UserHandler) GetAllUser(c *gin.Context) {
	var req userdto.GetUserRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrValidate),
		)
		return
	}

	req.SetDefault()

	var users []user.User

	var domainReq filter.FilterRequest

	domainReq.Page = req.Page
	domainReq.Limit = req.Limit
	domainReq.Sort = req.Sort
	domainReq.Order = req.Order
	domainReq.SearchBy = req.SearchBy
	domainReq.SearchValue = req.SearchValue

	totalPage, err := h.uc.GetAllUser(c.Request.Context(), &users, domainReq)

	if err != nil {
		c.JSON(
			http.StatusNotFound,
			enums.NewAppError(http.StatusNotFound, err.Error(), "ERR_USER_NOT_FOUND"),
		)
		return
	}

	usersDTORes := make([]userdto.UserDTO, 0)

	for _, user := range users {
		usersDTORes = append(usersDTORes, userdto.DomainToDTO(user))
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Fetched user successfully",
		"data": userdto.GetUserResponse{
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
// @Success 200 {object} userdto.GetUserByIDResponseWrapper
// @Failure 400 {object} enums.AppError
// @Router /users/{userID} [get]
func (h *UserHandler) GetUserByID(c *gin.Context) {
	var req userdto.GetUserByIDRequest

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

	var userDTORes userdto.UserDTO

	userDTORes = userdto.DomainToDTO(user)

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Fetched user successfully",
		"data": userdto.GetUserByIDResponse{
			User: userDTORes,
		},
	})
}

// @Summary Create user
// @Description API thêm người dùng
// @Tags users
// @Accept json
// @Produce json
// @Param request body userdto.CreateUserRequest true "Create user info"
// @Success 201 {object} userdto.CreateUserResponseWrapper "Created user successfully"
// @Failure 400 {object} enums.AppError
// @Router /users [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	var req userdto.CreateUserRequest

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
	user.Status = int8(req.Status)
	user.GoodPoint = req.GoodPoint

	if err := h.uc.CreateUser(c.Request.Context(), &user); err != nil {
		c.JSON(
			http.StatusConflict,
			enums.NewAppError(http.StatusConflict, err.Error(), enums.ErrConflict),
		)
		return
	}

	var userDTORes userdto.UserDTO

	userDTORes = userdto.DomainToDTO(user)

	c.JSON(http.StatusCreated, gin.H{
		"code":    http.StatusCreated,
		"message": "Created user successfully",
		"data": userdto.CreateUserResponse{
			User: userDTORes,
		},
	})
}

// @Summary Update user
// @Description API cập nhật người dùng
// @Tags users
// @Accept json
// @Produce json
// @Param request body userdto.UpdateUserRequest true "Update user info"
// @Success 200 {object} userdto.UpdateUserResponseWrapper "Updated user successfully"
// @Failure 400 {object} enums.AppError
// @Router /users [put]
func (h *UserHandler) UpdateUser(c *gin.Context) {
	var req userdto.UpdateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrValidate),
		)
		return
	}

	var user user.User

	user.ID = req.ID
	user.Avatar = req.Avatar
	user.FullName = req.FullName
	user.PhoneNumber = req.PhoneNumber
	user.Address = req.Address
	user.Status = int8(req.Status)
	user.GoodPoint = req.GoodPoint
	user.Major = req.Major

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

// @Summary Delete user
// @Description API delete user by id
// @Tags users
// @Accept json
// @Produce json
// @Param userID path int true "ID user"
// @Success 200 {object} userdto.DeleteUserResponseWrapper
// @Failure 400 {object} enums.AppError
// @Router /users/{userID} [delete]
func (h *UserHandler) DeleteUser(c *gin.Context) {
	var req userdto.DeleteUserRequest

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrValidate),
		)
		return
	}

	if err := h.uc.DeleteUser(c.Request.Context(), req.UserID); err != nil {
		c.JSON(
			http.StatusConflict,
			enums.NewAppError(http.StatusConflict, err.Error(), enums.ErrConflict),
		)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Deleted user successfully",
	})
}
