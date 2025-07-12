package handler

import (
	"final_project/internal/application/app/userapp"
	"final_project/internal/domain/filter"
	"final_project/internal/domain/user"
	usergooddeed "final_project/internal/domain/user_good_deed"
	userdto "final_project/internal/dto/userDTO"
	"final_project/internal/pkg/enums"
	"final_project/internal/pkg/helpers"
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

// @Summary Get users report
// @Description API lấy ra báo cáo thành tích của sinh viên
// @Security BearerAuth
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} userdto.GetUserReportResponseWrapper
// @Failure 400 {object} enums.AppError
// @Router /users/report [get]
func (h *UserHandler) GetUserReport(c *gin.Context) {
	var (
		domainUserReport []user.UserReport
	)

	if err := h.uc.GetUserReport(c.Request.Context(), &domainUserReport); err != nil {
		c.JSON(http.StatusNotFound, enums.NewAppError(http.StatusNotFound, err.Error(), enums.ErrNotFound))
		return
	}

	userReportDTORes := make([]userdto.UserReportDTO, 0)

	for _, value := range domainUserReport {
		userReportDTORes = append(userReportDTORes, userdto.UserReportDomainToDTO(value))
	}

	c.JSON(http.StatusOK, userdto.GetUserReportResponseWrapper{
		Code:    http.StatusOK,
		Message: "Fetched user reports successfully",
		Data: userdto.GetUserReportResponse{
			UserReports: userReportDTORes,
		},
	})
}

// @Summary Get users rank
// @Description API bao gồm phân trang
// @Security BearerAuth
// @Tags users
// @Accept json
// @Produce json
// @Param page query int false "Current page" minimum(1) example(1)
// @Param limit query int false "Number record of page" minimum(1) example(10)
// @Success 200 {object} userdto.GetUserRankResponseWrapper
// @Failure 400 {object} enums.AppError
// @Router /client/users/ranks [get]
func (h *UserHandler) GetUserRanks(c *gin.Context) {
	var (
		req             userdto.GetUserRankRequest
		domainUserRanks []user.UserRank
		domainFilter    filter.FilterRequest
	)

	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrValidate),
		)
		return
	}

	userID, err := helpers.GetUintFromContext(c, "userID")
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrBadRequest))
		return
	}

	req.SetDefault()

	domainFilter.Page = req.Page
	domainFilter.Limit = req.Limit

	totalPage, userInfo, userRank, err := h.uc.GetAllUserRank(c.Request.Context(), &domainUserRanks, domainFilter, userID)
	if err != nil {
		c.JSON(http.StatusNotFound, enums.NewAppError(http.StatusNotFound, err.Error(), enums.ErrNotFound))
		return
	}

	userRankDTOs := make([]userdto.UserRankDTO, 0)
	var userRankDTO userdto.UserRankDTO

	for _, value := range domainUserRanks {
		userRankDTOs = append(userRankDTOs, userdto.UserGoodDeedDomainToDTO(value))
	}

	userRankDTO = userdto.UserGoodDeedDomainToDTO(*userInfo)

	c.JSON(http.StatusOK, userdto.GetUserRankResponseWrapper{
		Code:    http.StatusOK,
		Message: "Fetched users rank successfully",
		Data: userdto.GetUserRankResponse{
			YourInfo:  userRankDTO,
			YourRank:  userRank,
			UserRanks: userRankDTOs,
			TotalPage: totalPage,
		},
	})
}

// @Summary Get user good deeds
// @Description Get all good deeds of a user by user ID
// @Security BearerAuth
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} userdto.GetUserGoodDeedResponseWrapper
// @Failure 400 {object} enums.AppError
// @Failure 404 {object} enums.AppError
// @Router /client/users/my-good-deeds [get]
func (h *UserHandler) GetUserGoodDeed(c *gin.Context) {
	var goodDeeds []usergooddeed.UserGoodDeedDetail

	userID, err := helpers.GetUintFromContext(c, "userID")
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrBadRequest))
		return
	}

	if err := h.uc.GetUserGoodDeed(c.Request.Context(), &goodDeeds, int(userID)); err != nil {
		c.JSON(
			http.StatusNotFound,
			enums.NewAppError(http.StatusNotFound, err.Error(), "ERR_USER_NOT_FOUND"),
		)
		return
	}

	goodDeedsDTORes := make([]userdto.UserGoodDeedDetail, 0)

	for _, goodDeed := range goodDeeds {
		goodDeedsDTORes = append(goodDeedsDTORes, userdto.DomainUserGoodDeedToDTO(goodDeed))
	}

	c.JSON(http.StatusOK, userdto.GetUserGoodDeedResponseWrapper{
		Code:    http.StatusOK,
		Message: "Fetched user good deeds successfully",
		Data: userdto.GetUserGoodDeedResponse{
			GoodDeeds: goodDeedsDTORes,
		},
	})
}

// @Summary Get user good deeds by ID
// @Description Get all good deeds of a user by user ID
// @Security BearerAuth
// @Tags users
// @Accept json
// @Produce json
// @Param userID path int true "ID user"
// @Success 200 {object} userdto.GetUserGoodDeedResponseWrapper
// @Failure 400 {object} enums.AppError
// @Failure 404 {object} enums.AppError
// @Router /users/{userID}/my-good-deeds [get]
func (h *UserHandler) GetUserGoodDeedByID(c *gin.Context) {
	var goodDeeds []usergooddeed.UserGoodDeedDetail

	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrValidate),
		)
		return
	}

	if err := h.uc.GetUserGoodDeed(c.Request.Context(), &goodDeeds, int(userID)); err != nil {
		c.JSON(
			http.StatusNotFound,
			enums.NewAppError(http.StatusNotFound, err.Error(), "ERR_USER_NOT_FOUND"),
		)
		return
	}

	goodDeedsDTORes := make([]userdto.UserGoodDeedDetail, 0)

	for _, goodDeed := range goodDeeds {
		goodDeedsDTORes = append(goodDeedsDTORes, userdto.DomainUserGoodDeedToDTO(goodDeed))
	}

	c.JSON(http.StatusOK, userdto.GetUserGoodDeedResponseWrapper{
		Code:    http.StatusOK,
		Message: "Fetched user good deeds by id successfully",
		Data: userdto.GetUserGoodDeedResponse{
			GoodDeeds: goodDeedsDTORes,
		},
	})
}

// @Summary Create a good deed
// @Description Create a new good deed for a user
// @Security BearerAuth
// @Tags users
// @Accept json
// @Produce json
// @Param request body userdto.CreateGoodDeedRequest true "Create good deed request"
// @Success 200 {object} userdto.CreateUserGoodDeedResponseWrapper
// @Failure 400 {object} enums.AppError
// @Failure 500 {object} enums.AppError
// @Router /users/good-deeds [post]
func (h *UserHandler) CreateGoodDeed(c *gin.Context) {
	var req userdto.CreateGoodDeedRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrValidate),
		)
		return
	}

	goodDeed := usergooddeed.UserGoodDeed{
		UserID:        req.UserID,
		GoodDeedType:  req.GoodDeedType,
		TransactionID: &req.TransactionID,
	}

	if err := h.uc.CreateGoodDeed(c.Request.Context(), &goodDeed); err != nil {
		c.JSON(
			http.StatusConflict,
			enums.NewAppError(http.StatusConflict, err.Error(), enums.ErrConflict),
		)
		return
	}

	c.JSON(http.StatusOK, userdto.CreateUserGoodDeedResponseWrapper{
		Code:    http.StatusOK,
		Message: "Created user good deed successfully",
		Data:    gin.H{},
	})
}

// @Summary Delete a good deed
// @Description API delete user good deed
// @Security BearerAuth
// @Tags users
// @Accept json
// @Produce json
// @Param goodDeedID path int true "ID good deed"
// @Success 200 {object} userdto.DeleteUserGoodDeedResponseWrapper
// @Failure 400 {object} enums.AppError
// @Router /users/good-deeds/{goodDeedID} [delete]
func (h *UserHandler) DeleteGoodDeed(c *gin.Context) {
	var (
		req userdto.DeleteGoodDeedRequest
	)

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrValidate),
		)
		return
	}

	if err := h.uc.DeleteGoodDeed(c.Request.Context(), uint(req.GoodDeedID)); err != nil {
		c.JSON(
			http.StatusConflict,
			enums.NewAppError(http.StatusConflict, err.Error(), enums.ErrConflict),
		)
		return
	}

	c.JSON(http.StatusOK, userdto.DeleteUserGoodDeedResponseWrapper{
		Code:    http.StatusOK,
		Message: "Deleted user good deeds successfully",
		Data:    gin.H{},
	})
}

// @Summary Get admins
// @Description API bao gồm cả lọc, phân trang và sắp xếp
// @Security BearerAuth
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
// @Security BearerAuth
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
// @Security BearerAuth
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
// @Security BearerAuth
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

	clientID, err := strconv.Atoi(c.Param("clientID"))
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrValidate),
		)
		return
	}

	user.ID = uint(clientID)
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

	userDTORes := userdto.DomainUpdateUserToDTO(user)

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Updated client successfully",
		"data": userdto.UpdateClientResponse{
			Client: userDTORes,
		},
	})
}

// @Summary Update admin
// @Description API cập nhật người dùng
// @Security BearerAuth
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
		"data":    gin.H{},
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
		"data":    gin.H{},
	})
}

// @Summary Delete admin
// @Description API delete admin by id
// @Security BearerAuth
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
		"data":    gin.H{},
	})
}
