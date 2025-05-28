package handler

import (
	"net/http"

	"final_project/internal/application/adminapp"
	"final_project/internal/domain/admin"
	"final_project/internal/domain/filter"
	adminDTO "final_project/internal/dto/adminDTO"
	admindto "final_project/internal/dto/adminDTO"

	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	usecase *adminapp.UseCase
}

func NewAdminHandler(uc *adminapp.UseCase) *AdminHandler {
	return &AdminHandler{usecase: uc}
}

// GetAllAdmins godoc
// @Summary Get admin
// @Description Lấy danh sách admin với phân trang, lọc, sắp xếp
// @Tags admin
// @Accept  json
// @Produce  json
// @Param   page     query    int     false  "Trang"
// @Param   limit    query    int     false  "Số lượng/trang"
// @Param   sort     query    string  false  "Trường sắp xếp"
// @Param   order    query    string  false  "Thứ tự sắp xếp (ASC/DESC)"
// @Param   searchBy   query    string  false  "Trường lọc (vd: email, full_name)"
// @Param   searchValue   query    string  false  "Giá trị lọc (vd:abc@gmail.com, John Doe)"
// @Success 200 {object} adminDTO.GetAdminResponseWrapper
// @Failure 400 {object} enums.AppError
// @Router /admins [get]
func (h *AdminHandler) GetAllAdmins(c *gin.Context) {
	var req admindto.GetAllAdminRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	req.SetDefault()

	var domainReq filter.FilterRequest
	domainReq.Page = req.Page
	domainReq.Limit = req.Limit
	domainReq.Sort = req.Sort
	domainReq.Order = req.Order
	domainReq.SearchBy = req.SearchBy
	domainReq.SearchValue = req.SearchValue

	admins, totalPage, err := h.usecase.GetAllAdmin(c.Request.Context(), domainReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	var adminDTOs []adminDTO.AdminDTO
	for _, a := range admins {
		// Lấy roleName từ usecase trả về (nếu cần, có thể trả về []struct{admin, roleName})
		adminDTOs = append(adminDTOs, adminDTO.ToAdminDTO(a.Admin, a.RoleName))
	}

	c.JSON(http.StatusOK, adminDTO.GetAdminResponseWrapper{
		Code:    200,
		Message: "Success",
		Data: adminDTO.GetAdminResponse{
			Admins:    adminDTOs,
			TotalPage: totalPage,
		},
	})
}

// GetAdminByID godoc
// @Summary Get admin by ID
// @Description Lấy chi tiết admin theo ID
// @Tags admin
// @Accept  json
// @Produce  json
// @Param   adminID   path     int  true  "Admin ID"
// @Success 200 {object} adminDTO.GetAdminByIDResponseWrapper
// @Failure 400 {object} enums.AppError
// @Failure 404 {object} enums.AppError
// @Router /admins/{adminID} [get]
func (h *AdminHandler) GetAdminByID(c *gin.Context) {
	var req adminDTO.GetAdminByIDRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	admin, roleName, err := h.usecase.GetAdminByID(c.Request.Context(), uint(req.AdminID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, adminDTO.GetAdminByIDResponseWrapper{
		Code:    200,
		Message: "Success",
		Data: adminDTO.GetAdminByIDResponse{
			Admin: adminDTO.ToAdminDTO(admin, roleName),
		},
	})
}

// CreateAdmin godoc
// @Summary Create admin
// @Description Tạo mới admin
// @Tags admin
// @Accept  json
// @Produce  json
// @Param   body  body  adminDTO.CreateAdminRequest  true  "Thông tin admin"
// @Success 201 {object} adminDTO.CreateAdminResponseWrapper
// @Failure 400 {object} enums.AppError
// @Router /admins [post]
func (h *AdminHandler) CreateAdmin(c *gin.Context) {
	var req adminDTO.CreateAdminRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	// DTO → Domain
	domainAdmin := adminDTO.ToDomainAdmin(req)

	// Gửi domain entity sang usecase
	createdDomainAdmin, roleName, err := h.usecase.CreateAdmin(c.Request.Context(), domainAdmin)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	// Domain → DTO (response)
	resp := adminDTO.ToAdminDTO(createdDomainAdmin, roleName)
	c.JSON(http.StatusCreated, adminDTO.CreateAdminResponseWrapper{
		Code:    201,
		Message: "Created",
		Data:    adminDTO.CreateAdminResponse{Admin: resp},
	})
}

// UpdateAdmin godoc
// @Summary Update admin
// @Description Cập nhật thông tin admin
// @Tags admin
// @Accept  json
// @Produce  json
// @Param   body  body  adminDTO.UpdateAdminRequest  true  "Thông tin cập nhật"
// @Success 200 {object} adminDTO.UpdateAdminResponseWrapper
// @Failure 400 {object} enums.AppError
// @Router /admins [put]
func (h *AdminHandler) UpdateAdmin(c *gin.Context) {
	var req adminDTO.UpdateAdminRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	a := &admin.Admin{
		ID:       req.ID,
		FullName: req.FullName,
		Password: req.Password,
		Status:   int8(req.Status),
		RoleID:   req.RoleID,
	}
	if err := h.usecase.UpdateAdmin(c.Request.Context(), a); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, adminDTO.UpdateAdminResponseWrapper{
		Code:    200,
		Message: "Updated",
		Data:    nil,
	})
}

// DeleteAdmin godoc
// @Summary Delete admin
// @Description Xóa admin theo ID
// @Tags admin
// @Accept  json
// @Produce  json
// @Param   adminID   path  int  true  "Admin ID"
// @Success 200 {object} adminDTO.UpdateAdminResponseWrapper
// @Failure 400 {object} enums.AppError
// @Router /admins/{adminID} [delete]
func (h *AdminHandler) DeleteAdmin(c *gin.Context) {
	var req adminDTO.DeleteAdminRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := h.usecase.DeleteAdmin(c.Request.Context(), req.AdminID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, adminDTO.UpdateAdminResponseWrapper{
		Code:    200,
		Message: "Deleted",
		Data:    nil,
	})
}
