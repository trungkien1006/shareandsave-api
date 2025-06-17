package handler

import (
	"context"
	"final_project/internal/application/app/interestapp"
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
// @Param type query int false "Interested: 1, FollowedBy: 2" example(1, 2)
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
		"message": "Fetched interests successfully",
		"data": interestdto.GetInterestResponse{
			Interests: interestDTORes,
			TotalPage: totalPage,
		},
	})
}

// @Summary Get interest by ID
// @Description API lấy thông tin interest theo ID
// @Security BearerAuth
// @Tags interests
// @Accept json
// @Produce json
// @Param interestID path int true "ID interest"
// @Success 200 {object} interestdto.GetByIDInterestResponseWrapper
// @Failure 400 {object} enums.AppError
// @Failure 404 {object} enums.AppError
// @Router /interests/{interestID} [get]
func (h *InterestHandler) GetByID(c *gin.Context) {
	var (
		req            interestdto.GetByID
		domainInterest interest.PostInterest
	)

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrValidate))
		return
	}

	err := h.uc.GetInterestByID(context.Background(), &domainInterest, req.InterestID)
	if err != nil {
		c.JSON(http.StatusNotFound, enums.NewAppError(http.StatusNotFound, err.Error(), enums.ErrNotFound))
		return
	}

	var interestDTORes interestdto.PostInterest

	interestDTORes = interestdto.GetDomainToDTO(domainInterest)

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Fetched interest successfully",
		"data": interestdto.GetByIDInterestResponse{
			Interest: interestDTORes,
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

	interestID, err := h.uc.CreateInterest(c.Request.Context(), domainInterest)
	if err != nil {
		c.JSON(http.StatusConflict, enums.NewAppError(http.StatusConflict, err.Error(), enums.ErrConflict))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Created interest successfully",
		"data": interestdto.DeleteInterestResponse{
			InterestID: interestID,
		},
	})
}

// @Summary Delete Interest
// @Description API xóa interest theo ID
// @Security BearerAuth
// @Tags interests
// @Accept json
// @Produce json
// @Param postID path int true "ID post"
// @Success 200 {object} interestdto.DeleteInterestResponseWrapper "Deleted interest successfully"
// @Failure 400 {object} enums.AppError
// @Failure 500 {object} enums.AppError
// @Router /interests/{postID} [delete]
func (h *InterestHandler) Delete(c *gin.Context) {
	var (
		req interestdto.DeleteInterest
	)

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrValidate))
		return
	}

	userID, err := helpers.GetUintFromContext(c, "userID")
	if err != nil {
		c.JSON(http.StatusBadRequest, enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrBadRequest))
		return
	}

	interestID, err := h.uc.DeleteInterest(c.Request.Context(), req.PostID, userID)
	if err != nil {
		c.JSON(
			http.StatusConflict,
			enums.NewAppError(http.StatusConflict, err.Error(), enums.ErrConflict),
		)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Deleted interest successfully",
		"data": interestdto.DeleteInterestResponse{
			InterestID: interestID,
		},
	})
}
