package handler

import (
	"context"
	"final_project/internal/application/interestapp"
	"final_project/internal/domain/interest"
	interestdto "final_project/internal/dto/interestDTO"
	"final_project/internal/pkg/enums"
	"final_project/internal/pkg/helpers"
	"final_project/internal/shared/validator"
	"net/http"

	"github.com/gin-gonic/gin"
)

type InterestHandler struct {
	uc *interestapp.UseCase
}

func NewInterestHandler(uc *interestapp.UseCase) *InterestHandler {
	return &InterestHandler{uc: uc}
}

// @Summary Get interest
// @Description API bao gồm cả lọc, phân trang và sắp xếp
// @Security BearerAuth
// @Tags interests
// @Accept json
// @Produce json
// @Param page query int false "Current page" minimum(1) example(1)
// @Param limit query int false "Number record of page" minimum(1) example(10)
// @Param sort query string false "Sort column (createdAt)"
// @Param order query string false "Sort type" enum(ASC,DESC) example(ASC, DESC)
// @Param type query int false "Interested: 1, Following: 2" example(1, 2)
// @Param search query string false "Search value"
// @Success 200 {object} interestdto.GetInterestResponseWrapper
// @Failure 400 {object} enums.AppError
// @Router /interests [get]
func (h *InterestHandler) GetAll(c *gin.Context) {
	var (
		req            interestdto.GetInterest
		domainReq      interest.GetInterest
		domainInterest []interest.PostInterest
	)

	userID, err := helpers.GetUintFromContext(c, "userID")
	if err != nil {
		c.JSON(http.StatusBadRequest, enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrBadRequest))
		return
	}

	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrValidate))
		return
	}

	if err := validator.Validate.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrValidate))
		return
	}

	req.SetDefault()

	domainReq = interestdto.GetDTOToDomain(req)

	totalPage, err := h.uc.GetAllInterest(context.Background(), &domainInterest, userID, domainReq)

	if err != nil {
		c.JSON(http.StatusNotFound, enums.NewAppError(http.StatusNotFound, err.Error(), enums.ErrNotFound))
		return
	}

	interestDTORes := make([]interestdto.PostInterest, 0)

	for _, value := range domainInterest {
		interestDTORes = append(interestDTORes, interestdto.GetDomainToDTO(value))
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Fetched posts successfully",
		"data": interestdto.GetInterestResponse{
			Interests: interestDTORes,
			TotalPage: totalPage,
		},
	})
}

// @Summary Create interest
// @Description API quan tâm đến bài viết + JWT
// @Security BearerAuth
// @Tags interests
// @Accept json
// @Produce json
// @Param request body interestdto.CreateInterest true "Interest creation payload"
// @Success 201 {object} interestdto.CreateInterestResponseWrapper
// @Failure 400 {object} enums.AppError
// @Failure 409 {object} enums.AppError
// @Router /interests [post]
func (h *InterestHandler) Create(c *gin.Context) {
	var (
		req            interestdto.CreateInterest
		domainInterest interest.Interest
	)

	userID, err := helpers.GetUintFromContext(c, "userID")
	if err != nil {
		c.JSON(http.StatusBadRequest, enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrBadRequest))
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrValidate))
		return
	}

	if err := validator.Validate.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrValidate))
		return
	}

	domainInterest = interestdto.CreateDTOToDomain(req, userID)

	if err := h.uc.CreateInterest(c.Request.Context(), domainInterest); err != nil {
		c.JSON(http.StatusConflict, enums.NewAppError(http.StatusConflict, err.Error(), enums.ErrConflict))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Created post successfully",
		"data":    gin.H{},
	})
}
