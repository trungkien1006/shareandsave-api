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
// @Tags interest
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
