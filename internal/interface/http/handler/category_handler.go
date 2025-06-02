package handler

import (
	"final_project/internal/application/categoryapp"
	"final_project/internal/domain/category"
	categorydto "final_project/internal/dto/categoryDTO"
	"final_project/internal/pkg/enums"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	uc *categoryapp.UseCase
}

func NewCategoryHandler(uc *categoryapp.UseCase) *CategoryHandler {
	return &CategoryHandler{uc: uc}
}

// Login godoc
// @Summary Đăng nhập
// @Description Đăng nhập người dùng với email và mật khẩu mạnh
// @Tags Auth
// @Accept json
// @Produce json
// @Param login body authdto.LoginRequest true "Dữ liệu đăng nhập"
// @Success 200 {object} authdto.LoginResponseWrapper
// @Failure 400 {object} enums.AppError
// @Failure 401 {object} enums.AppError
// @Router /login [post]
func (h *CategoryHandler) GetAll(c *gin.Context) {
	var (
		categories    []category.Category
		categoriesDTO []categorydto.CategoryDTO
	)

	if err := h.uc.GetAllCategories(c.Request.Context(), &categories); err != nil {
		c.JSON(http.StatusNotFound, enums.NewAppError(http.StatusNotFound, err.Error(), enums.ErrNotFound))
		return
	}

	for _, value := range categories {
		categoriesDTO = append(categoriesDTO, categorydto.CateDomainToDTO(value))
	}

	c.JSON(http.StatusOK, categorydto.GetCategoryResponseWrapper{
		Code:    http.StatusOK,
		Message: "Fetched items successfully",
		Data: categorydto.GetCategoryResponse{
			Categories: categoriesDTO,
		},
	})
}
