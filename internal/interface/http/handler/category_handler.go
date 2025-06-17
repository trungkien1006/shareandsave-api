package handler

import (
	"final_project/internal/application/app/categoryapp"
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

// @Summary Get categories
// @Description API lấy ra tất cả danh mục đồ đạc
// @Tags categories
// @Accept json
// @Produce json
// @Success 200 {object} categorydto.GetCategoryResponseWrapper
// @Failure 400 {object} enums.AppError
// @Failure 404 {object} enums.AppError
// @Router /categories [get]
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
		Message: "Fetched categories successfully",
		Data: categorydto.GetCategoryResponse{
			Categories: categoriesDTO,
		},
	})
}
