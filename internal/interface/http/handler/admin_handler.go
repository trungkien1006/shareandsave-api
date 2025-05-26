package handler

import (
	"net/http"
	"strconv"

	"final_project/internal/application/adminapp"
	"final_project/internal/domain/admin"
	"final_project/internal/domain/filter"
	adminDTO "final_project/internal/dto/adminDTO"

	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	usecase *adminapp.UseCase
}

func NewAdminHandler(uc *adminapp.UseCase) *AdminHandler {
	return &AdminHandler{usecase: uc}
}

// GetAllAdmins godoc
// @Summary Lấy danh sách admin
// @Description Lấy danh sách admin với phân trang, lọc, sắp xếp
// @Tags admin
// @Accept  json
// @Produce  json
// @Param   page     query    int     false  "Trang"
// @Param   limit    query    int     false  "Số lượng/trang"
// @Param   sort     query    string  false  "Trường sắp xếp"
// @Param   order    query    string  false  "Thứ tự sắp xếp (ASC/DESC)"
// @Param   filter   query    string  false  "Bộ lọc"
// @Success 200 {object} adminDTO.GetAdminResponseWrapper
// @Failure 400 {object} map[string]interface{}
// @Router /admins [get]
func (h *AdminHandler) GetAllAdmins(c *gin.Context) {
	var req filter.FilterRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	var admins []admin.Admin
	totalPage, err := h.usecase.GetAllAdmin(c.Request.Context(), &admins, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	var adminDTOs []adminDTO.AdminDTO
	for _, a := range admins {
		adminDTOs = append(adminDTOs, adminDTO.ToAdminDTO(a))
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
// @Summary Lấy thông tin admin theo ID
// @Description Lấy chi tiết admin theo ID
// @Tags admin
// @Accept  json
// @Produce  json
// @Param   id   path     int  true  "Admin ID"
// @Success 200 {object} adminDTO.GetAdminByIDResponseWrapper
// @Failure 400 {object} enums.AppError
// @Failure 404 {object} enums.AppError
// @Router /admins/{id} [get]
func (h *AdminHandler) GetAdminByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid admin ID"})
		return
	}
	var a admin.Admin
	if err := h.usecase.GetAdminByID(c.Request.Context(), &a, id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, adminDTO.GetAdminByIDResponseWrapper{
		Code:    200,
		Message: "Success",
		Data: adminDTO.GetAdminByIDResponse{
			Admin: adminDTO.ToAdminDTO(a),
		},
	})
}

// CreateAdmin godoc
// @Summary Tạo mới admin
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
	a := admin.NewAdmin(req.Email, req.Password, req.Fullname, int8(req.Status), req.RoleID)
	if err := h.usecase.CreateAdmin(c.Request.Context(), a); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, adminDTO.CreateAdminResponseWrapper{
		Code:    201,
		Message: "Created",
		Data: adminDTO.CreateAdminResponse{
			Admin: adminDTO.ToAdminDTO(*a),
		},
	})
}

// UpdateAdmin godoc
// @Summary Cập nhật admin
// @Description Cập nhật thông tin admin
// @Tags admin
// @Accept  json
// @Produce  json
// @Param   id    path  int  true  "Admin ID"
// @Param   body  body  adminDTO.UpdateAdminRequest  true  "Thông tin cập nhật"
// @Success 200 {object} adminDTO.UpdateAdminResponseWrapper
// @Failure 400 {object} enums.AppError
// @Router /admins/{id} [put]
func (h *AdminHandler) UpdateAdmin(c *gin.Context) {
	var req adminDTO.UpdateAdminRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	a := &admin.Admin{
		ID:       req.ID,
		Fullname: req.Fullname,
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
// @Summary Xóa admin
// @Description Xóa admin theo ID
// @Tags admin
// @Accept  json
// @Produce  json
// @Param   id   path  int  true  "Admin ID"
// @Success 200 {object} adminDTO.UpdateAdminResponseWrapper
// @Failure 400 {object} enums.AppError
// @Router /admins/{id} [delete]
func (h *AdminHandler) DeleteAdmin(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid admin ID"})
		return
	}
	if err := h.usecase.DeleteAdmin(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, adminDTO.UpdateAdminResponseWrapper{
		Code:    200,
		Message: "Deleted",
		Data:    nil,
	})
}
