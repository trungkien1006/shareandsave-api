package handler

import (
	"final_project/internal/application/userapp"
	"final_project/internal/domain/filter"
	"final_project/internal/domain/user"
	userdto "final_project/internal/dto/userDTO"
	"final_project/internal/pkg/enums"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	uc *userapp.UseCase
}

func NewUserHandler(uc *userapp.UseCase) *UserHandler {
	return &UserHandler{uc: uc}
}

// @Summary Get admins
// @Description API bao gồm cả lọc, phân trang và sắp xếp
// @Tags admins
// @Accept json
// @Produce json
// @Param page query int false "Current page" minimum(1) example(1)
// @Param limit query int false "Number record of page" minimum(1) example(10)
// @Param sort query string false "Sort column (createdAt goodPoint)"
// @Param order query string false "Sort type: ASC hoặc DESC" enum(ASC,DESC) example(ASC)
// @Param   searchBy   query    string  false  "Trường lọc (fullName email phoneNumber status roleName)"
// @Param   searchValue   query    string  false  "Giá trị lọc (vd:abc@gmail.com, John Doe)"
// @Success 200 {object} userdto.GetAdminResponseWrapper
// @Failure 400 {object} enums.AppError
// @Router /admins [get]
func (h *UserHandler) GetAllAdmin(c *gin.Context) {
	var req userdto.GetAdminRequest

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

	totalPage, err := h.uc.GetAllAdmin(c.Request.Context(), &users, domainReq)

	if err != nil {
		c.JSON(
			http.StatusNotFound,
			enums.NewAppError(http.StatusNotFound, err.Error(), "ERR_USER_NOT_FOUND"),
		)
		return
	}

	adminsDTORes := make([]userdto.AdminDTO, 0)

	for _, user := range users {
		adminsDTORes = append(adminsDTORes, userdto.DomainAdminToDTO(user))
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Fetched admins successfully",
		"data": userdto.GetAdminResponse{
			Admins:    adminsDTORes,
			TotalPage: totalPage,
		},
	})
}

// @Summary Get clients
// @Description API bao gồm cả lọc, phân trang và sắp xếp
// @Tags clients
// @Accept json
// @Produce json
// @Param page query int false "Current page" minimum(1) example(1)
// @Param limit query int false "Number record of page" minimum(1) example(10)
// @Param sort query string false "Sort column (createdAt goodPoint)"
// @Param order query string false "Sort type: ASC hoặc DESC" enum(ASC,DESC) example(ASC)
// @Param   searchBy   query    string  false  "Trường lọc (fullName email phoneNumber status)"
// @Param   searchValue   query    string  false  "Giá trị lọc (vd:abc@gmail.com, John Doe)"
// @Success 200 {object} userdto.GetClientResponseWrapper
// @Failure 400 {object} enums.AppError
// @Router /clients [get]
func (h *UserHandler) GetAllClient(c *gin.Context) {
	var req userdto.GetClientRequest

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

	totalPage, err := h.uc.GetAllClient(c.Request.Context(), &users, domainReq)

	if err != nil {
		c.JSON(
			http.StatusNotFound,
			enums.NewAppError(http.StatusNotFound, err.Error(), "ERR_USER_NOT_FOUND"),
		)
		return
	}

	clientsDTORes := make([]userdto.ClientDTO, 0)

	for _, user := range users {
		clientsDTORes = append(clientsDTORes, userdto.DomainClientToDTO(user))
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Fetched user successfully",
		"data": userdto.GetClientResponse{
			Clients:   clientsDTORes,
			TotalPage: totalPage,
		},
	})
}

// @Summary Get client by ID
// @Description API get client by id
// @Tags clients
// @Accept json
// @Produce json
// @Param clientID path int true "ID client"
// @Success 200 {object} userdto.GetClientByIDResponseWrapper
// @Failure 400 {object} enums.AppError
// @Router /clients/{clientID} [get]
func (h *UserHandler) GetClientByID(c *gin.Context) {
	var req userdto.GetClientByIDRequest

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrValidate),
		)
		return
	}

	var user user.User

	if err := h.uc.GetClientByID(c.Request.Context(), &user, req.ClientID); err != nil {
		c.JSON(
			http.StatusNotFound,
			enums.NewAppError(http.StatusNotFound, err.Error(), enums.ErrNotFound),
		)
		return
	}

	var clientDTORes userdto.ClientDTO

	clientDTORes = userdto.DomainClientToDTO(user)

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Fetched user successfully",
		"data": userdto.GetClientByIDResponse{
			Client: clientDTORes,
		},
	})
}

// @Summary Get admin by ID
// @Description API get admin by id
// @Tags admins
// @Accept json
// @Produce json
// @Param adminID path int true "ID admin"
// @Success 200 {object} userdto.GetAdminByIDResponseWrapper
// @Failure 400 {object} enums.AppError
// @Router /admins/{adminID} [get]
func (h *UserHandler) GetAdminByID(c *gin.Context) {
	var req userdto.GetAdminByIDRequest

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrValidate),
		)
		return
	}

	var user user.User

	if err := h.uc.GetAdminByID(c.Request.Context(), &user, req.AdminID); err != nil {
		c.JSON(
			http.StatusNotFound,
			enums.NewAppError(http.StatusNotFound, err.Error(), enums.ErrNotFound),
		)
		return
	}

	var adminDTORes userdto.AdminDTO

	adminDTORes = userdto.DomainAdminToDTO(user)

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Fetched user successfully",
		"data": userdto.GetAdminByIDResponse{
			Admin: adminDTORes,
		},
	})
}

// @Summary Create admin
// @Description API thêm người dùng
// @Tags admins
// @Accept json
// @Produce json
// @Param request body userdto.CreateAdminRequest true "Create admin info"
// @Success 201 {object} userdto.CreateAdminResponseWrapper "Created admin successfully"
// @Failure 400 {object} enums.AppError
// @Router /admins [post]
func (h *UserHandler) CreateAdmin(c *gin.Context) {
	var req userdto.CreateAdminRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrValidate),
		)
		return
	}

	var user user.User

	user.RoleID = req.RoleID
	user.Email = req.Email
	user.Password = req.Password
	user.Avatar = "" // nếu có field Avatar
	user.FullName = req.FullName
	user.PhoneNumber = req.PhoneNumber
	user.Address = req.Address
	user.Status = int8(req.Status)
	user.GoodPoint = req.GoodPoint

	if err := h.uc.CreateAdmin(c.Request.Context(), &user); err != nil {
		c.JSON(
			http.StatusConflict,
			enums.NewAppError(http.StatusConflict, err.Error(), enums.ErrConflict),
		)
		return
	}

	var adminDTORes userdto.AdminDTO

	adminDTORes = userdto.DomainAdminToDTO(user)

	c.JSON(http.StatusCreated, gin.H{
		"code":    http.StatusCreated,
		"message": "Created admin successfully",
		"data": userdto.CreateAdminResponse{
			Admin: adminDTORes,
		},
	})
}

// @Summary Create client
// @Description API thêm người dùng
// @Tags clients
// @Accept json
// @Produce json
// @Param request body userdto.CreateClientRequest true "Create client info"
// @Success 201 {object} userdto.CreateClientResponseWrapper "Created client successfully"
// @Failure 400 {object} enums.AppError
// @Router /clients [post]
func (h *UserHandler) CreateClient(c *gin.Context) {
	var req userdto.CreateClientRequest

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

	if err := h.uc.CreateClient(c.Request.Context(), &user); err != nil {
		c.JSON(
			http.StatusConflict,
			enums.NewAppError(http.StatusConflict, err.Error(), enums.ErrConflict),
		)
		return
	}

	var clientDTORes userdto.ClientDTO

	clientDTORes = userdto.DomainClientToDTO(user)

	c.JSON(http.StatusCreated, gin.H{
		"code":    http.StatusCreated,
		"message": "Created user successfully",
		"data": userdto.CreateClientResponse{
			Client: clientDTORes,
		},
	})
}

// @Summary Update client
// @Description API cập nhật người dùng
// @Tags clients
// @Accept json
// @Produce json
// @Param clientID path int true "ID client"
// @Param request body userdto.UpdateClientRequest true "Update client info"
// @Success 200 {object} userdto.UpdateClientResponseWrapper "Updated client successfully"
// @Failure 400 {object} enums.AppError
// @Router /clients/{clientID} [patch]
func (h *UserHandler) UpdateClient(c *gin.Context) {
	var req userdto.UpdateClientRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrValidate),
		)
		return
	}

	var user user.User

	userID, err := strconv.Atoi(c.Param("clientID"))
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrValidate),
		)
		return
	}

	user.ID = uint(userID)
	user.Avatar = req.Avatar
	user.FullName = req.FullName
	user.PhoneNumber = req.PhoneNumber
	user.Address = req.Address
	user.Status = int8(req.Status)
	user.GoodPoint = req.GoodPoint
	user.Major = req.Major

	if err := h.uc.UpdateClient(c.Request.Context(), &user); err != nil {
		c.JSON(
			http.StatusConflict,
			enums.NewAppError(http.StatusConflict, err.Error(), enums.ErrConflict),
		)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Updated client successfully",
	})
}

// @Summary Update admin
// @Description API cập nhật người dùng
// @Tags admins
// @Accept json
// @Produce json
// @Param adminID path int true "ID admin"
// @Param request body userdto.UpdateAdminRequest true "Update admin info"
// @Success 200 {object} userdto.UpdateAdminResponseWrapper "Updated admin successfully"
// @Failure 400 {object} enums.AppError
// @Router /admins/{adminID} [patch]
func (h *UserHandler) UpdateAdmin(c *gin.Context) {
	var req userdto.UpdateAdminRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrValidate),
		)
		return
	}

	var user user.User

	userID, err := strconv.Atoi(c.Param("adminID"))
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrValidate),
		)
		return
	}

	user.RoleID = req.RoleID
	user.ID = uint(userID)
	user.Avatar = req.Avatar
	user.FullName = req.FullName
	user.PhoneNumber = req.PhoneNumber
	user.Address = req.Address
	user.Status = int8(req.Status)
	user.GoodPoint = req.GoodPoint
	user.Major = req.Major

	if err := h.uc.UpdateAdmin(c.Request.Context(), &user); err != nil {
		c.JSON(
			http.StatusConflict,
			enums.NewAppError(http.StatusConflict, err.Error(), enums.ErrConflict),
		)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Updated admin successfully",
	})
}

// @Summary Delete client
// @Description API delete client by id
// @Tags clients
// @Accept json
// @Produce json
// @Param clientID path int true "ID client"
// @Success 200 {object} userdto.DeleteClientResponseWrapper
// @Failure 400 {object} enums.AppError
// @Router /clients/{clientID} [delete]
func (h *UserHandler) DeleteClient(c *gin.Context) {
	var req userdto.DeleteClientRequest

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrValidate),
		)
		return
	}

	if err := h.uc.DeleteClient(c.Request.Context(), req.CLientID); err != nil {
		c.JSON(
			http.StatusConflict,
			enums.NewAppError(http.StatusConflict, err.Error(), enums.ErrConflict),
		)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Deleted client successfully",
	})
}

// @Summary Delete admin
// @Description API delete admin by id
// @Tags admins
// @Accept json
// @Produce json
// @Param adminID path int true "ID admin"
// @Success 200 {object} userdto.DeleteAdminResponseWrapper
// @Failure 400 {object} enums.AppError
// @Router /admins/{adminID} [delete]
func (h *UserHandler) DeleteAdmin(c *gin.Context) {
	var req userdto.DeleteAdminRequest

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrValidate),
		)
		return
	}

	if err := h.uc.DeleteAdmin(c.Request.Context(), req.AdminID); err != nil {
		c.JSON(
			http.StatusConflict,
			enums.NewAppError(http.StatusConflict, err.Error(), enums.ErrConflict),
		)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Deleted admin successfully",
	})
}
