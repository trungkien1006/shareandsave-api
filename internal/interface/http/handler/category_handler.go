package handler

import (
	"final_project/internal/application/app/categoryapp"
	"final_project/internal/domain/category"
	categorydto "final_project/internal/dto/categoryDTO"
	"final_project/internal/pkg/enums"
	"net/http"
	"strconv"

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

// @Summary Create category
// @Description API tạo mới một danh mục đồ đạc
// @Security BearerAuth
// @Tags categories
// @Accept json
// @Produce json
// @Param category body categorydto.CreateCategoryRequest true "Category data"
// @Success 201 {object} categorydto.CreateCategoryResponseWrapper
// @Failure 400 {object} enums.AppError
// @Failure 409 {object} enums.AppError
// @Router /categories [post]
func (h *CategoryHandler) Create(c *gin.Context) {
	var (
		req            categorydto.CreateCategoryRequest
		domainCategory category.Category
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrBadRequest))
		return
	}

	domainCategory = categorydto.CreateCateDTOToDomain(req)

	if err := h.uc.CreateCategory(c.Request.Context(), &domainCategory); err != nil {
		c.JSON(http.StatusConflict, enums.NewAppError(http.StatusConflict, err.Error(), enums.ErrConflict))
		return
	}

	c.JSON(http.StatusCreated, categorydto.CreateCategoryResponseWrapper{
		Code:    http.StatusCreated,
		Message: "Created category successfully",
		Data:    gin.H{},
	})
}

// @Summary Update category
// @Description API cập nhật một danh mục đồ đạc
// @Security BearerAuth
// @Tags categories
// @Accept json
// @Produce json
// @Param categoryID path int true "Category ID"
// @Param category body categorydto.UpdateCategoryRequest true "Category data"
// @Success 200 {object} categorydto.UpdateCategoryResponseWrapper
// @Failure 400 {object} enums.AppError
// @Failure 404 {object} enums.AppError
// @Router /categories/{categoryID} [patch]
func (h *CategoryHandler) Update(c *gin.Context) {
	var (
		req            categorydto.UpdateCategoryRequest
		domainCategory category.Category
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrBadRequest))
		return
	}

	categoryID, err := strconv.Atoi(c.Param("categoryID"))
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrValidate),
		)
		return
	}

	domainCategory.ID = uint(categoryID)
	domainCategory.Name = req.Name

	if err := h.uc.UpdateCategory(c.Request.Context(), &domainCategory); err != nil {
		c.JSON(http.StatusNotFound, enums.NewAppError(http.StatusNotFound, err.Error(), enums.ErrNotFound))
		return
	}

	c.JSON(http.StatusOK, categorydto.UpdateCategoryResponseWrapper{
		Code:    http.StatusOK,
		Message: "Updated category successfully",
		Data:    gin.H{},
	})
}

// @Summary Delete category
// @Description API xóa một danh mục đồ đạc
// @Security BearerAuth
// @Tags categories
// @Accept json
// @Produce json
// @Param categoryID path int true "Category ID"
// @Success 200 {object} categorydto.UpdateCategoryResponseWrapper
// @Failure 400 {object} enums.AppError
// @Failure 404 {object} enums.AppError
// @Router /categories/{categoryID} [delete]
func (h *CategoryHandler) Delete(c *gin.Context) {
	categoryID, err := strconv.Atoi(c.Param("categoryID"))
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrValidate),
		)
		return
	}

	if err := h.uc.DeleteCategory(c.Request.Context(), uint(categoryID)); err != nil {
		c.JSON(http.StatusConflict, enums.NewAppError(http.StatusConflict, err.Error(), enums.ErrConflict))
		return
	}

	c.JSON(http.StatusOK, categorydto.UpdateCategoryResponseWrapper{
		Code:    http.StatusOK,
		Message: "Deleted category successfully",
		Data:    gin.H{},
	})
}
