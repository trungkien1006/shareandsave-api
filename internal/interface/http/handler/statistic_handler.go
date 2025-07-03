package handler

import (
	"final_project/internal/application/app/statisticapp"
	statisticdto "final_project/internal/dto/statisticDTO"
	"final_project/internal/pkg/enums"
	"net/http"

	"github.com/gin-gonic/gin"
)

type StatisticHandler struct {
	uc *statisticapp.UseCase
}

func NewStatisticHandler(uc *statisticapp.UseCase) *StatisticHandler {
	return &StatisticHandler{uc: uc}
}

// @Summary Get year transaction statistic
// @Description API lấy ra thống kê giao dịch trong năm
// @Security BearerAuth
// @Tags statistics
// @Accept json
// @Produce json
// @Param year path string true "year"
// @Success 200 {object} statisticdto.StatisticTransactionResponseWrapper
// @Failure 400 {object} enums.AppError
// @Failure 404 {object} enums.AppError
// @Router /statistics/transaction/{year} [get]
func (h *StatisticHandler) StatisticTransactionInYear(c *gin.Context) {
	var req statisticdto.GetStatisticTransactionYearRequest

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrValidate),
		)
		return
	}

	totals, err := h.uc.StatisticTransactionInYear(c.Request.Context(), req.Year)
	if err != nil {
		c.JSON(http.StatusNotFound, enums.NewAppError(http.StatusNotFound, err.Error(), enums.ErrNotFound))
		return
	}

	c.JSON(http.StatusOK, statisticdto.StatisticTransactionResponseWrapper{
		Code:    http.StatusOK,
		Message: "Fetched year transaction statistic successfully",
		Data: statisticdto.StatisticTransactionResponse{
			Totals: totals,
		},
	})
}

// @Summary Get transaction statistic
// @Description API lấy ra thống kê giao dịch
// @Security BearerAuth
// @Tags statistics
// @Accept json
// @Produce json
// @Success 200 {object} statisticdto.GetStatisticResponseWrapper
// @Failure 400 {object} enums.AppError
// @Failure 404 {object} enums.AppError
// @Router /statistics/transaction [get]
func (h *StatisticHandler) TotalTransaction(c *gin.Context) {
	total, totalLastMonth, err := h.uc.TotalTransaction(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusNotFound, enums.NewAppError(http.StatusNotFound, err.Error(), enums.ErrNotFound))
		return
	}

	c.JSON(http.StatusOK, statisticdto.GetStatisticResponseWrapper{
		Code:    http.StatusOK,
		Message: "Fetched transaction statistic successfully",
		Data: statisticdto.GetStatisticResponse{
			Total:          uint(total),
			TotalLastMonth: uint(totalLastMonth),
		},
	})
}

// @Summary Get user statistic
// @Description API lấy ra thống kê thành viên
// @Security BearerAuth
// @Tags statistics
// @Accept json
// @Produce json
// @Success 200 {object} statisticdto.GetStatisticResponseWrapper
// @Failure 400 {object} enums.AppError
// @Failure 404 {object} enums.AppError
// @Router /statistics/user [get]
func (h *StatisticHandler) TotalUser(c *gin.Context) {
	total, totalLastMonth, err := h.uc.TotalUser(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusNotFound, enums.NewAppError(http.StatusNotFound, err.Error(), enums.ErrNotFound))
		return
	}

	c.JSON(http.StatusOK, statisticdto.GetStatisticResponseWrapper{
		Code:    http.StatusOK,
		Message: "Fetched user statistic successfully",
		Data: statisticdto.GetStatisticResponse{
			Total:          uint(total),
			TotalLastMonth: uint(totalLastMonth),
		},
	})
}

// @Summary Get post statistic
// @Description API lấy ra thống kê bài viết
// @Security BearerAuth
// @Tags statistics
// @Accept json
// @Produce json
// @Success 200 {object} statisticdto.GetStatisticResponseWrapper
// @Failure 400 {object} enums.AppError
// @Failure 404 {object} enums.AppError
// @Router /statistics/post [get]
func (h *StatisticHandler) TotalPost(c *gin.Context) {
	total, totalLastMonth, err := h.uc.TotalPost(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusNotFound, enums.NewAppError(http.StatusNotFound, err.Error(), enums.ErrNotFound))
		return
	}

	c.JSON(http.StatusOK, statisticdto.GetStatisticResponseWrapper{
		Code:    http.StatusOK,
		Message: "Fetched post statistic successfully",
		Data: statisticdto.GetStatisticResponse{
			Total:          uint(total),
			TotalLastMonth: uint(totalLastMonth),
		},
	})
}

// @Summary Get item warehouse statistic
// @Description API lấy ra thống kê hàng tồn
// @Security BearerAuth
// @Tags statistics
// @Accept json
// @Produce json
// @Success 200 {object} statisticdto.GetStatisticResponseWrapper
// @Failure 400 {object} enums.AppError
// @Failure 404 {object} enums.AppError
// @Router /statistics/item-warehouse [get]
func (h *StatisticHandler) TotalItemWarehouse(c *gin.Context) {
	total, totalLastMonth, err := h.uc.TotalItemWarehouse(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusNotFound, enums.NewAppError(http.StatusNotFound, err.Error(), enums.ErrNotFound))
		return
	}

	c.JSON(http.StatusOK, statisticdto.GetStatisticResponseWrapper{
		Code:    http.StatusOK,
		Message: "Fetched item warehouse statistic successfully",
		Data: statisticdto.GetStatisticResponse{
			Total:          uint(total),
			TotalLastMonth: uint(totalLastMonth),
		},
	})
}
