package handler

import (
	"final_project/internal/application/roleapp"
	rolepermission "final_project/internal/domain/role_permission"
	roledto "final_project/internal/dto/roleDTO"
	"final_project/internal/pkg/enums"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RoleHandler struct {
	uc *roleapp.UseCase
}

func NewRoleHandler(uc *roleapp.UseCase) *RoleHandler {
	return &RoleHandler{uc: uc}
}

// @Summary Get roles
// @Description API lấy ra tất cả chức vụ
// @Tags roles
// @Accept json
// @Produce json
// @Success 200 {object} roledto.GetRoleResponseWrapper
// @Failure 400 {object} enums.AppError
// @Failure 404 {object} enums.AppError
// @Router /roles [get]
func (h *RoleHandler) GetAll(c *gin.Context) {
	var (
		roles    []rolepermission.Role
		roleDTOs []roledto.RoleDTO
	)

	if err := h.uc.GetAllRoles(c.Request.Context(), &roles); err != nil {
		c.JSON(http.StatusNotFound, enums.NewAppError(http.StatusNotFound, err.Error(), enums.ErrNotFound))
		return
	}

	for _, value := range roles {
		roleDTOs = append(roleDTOs, roledto.RoleDomainToDTO(value))
	}

	c.JSON(http.StatusOK, roledto.GetRoleResponseWrapper{
		Code:    http.StatusOK,
		Message: "Fetched roles successfully",
		Data: roledto.GetRoleResponse{
			Roles: roleDTOs,
		},
	})
}
