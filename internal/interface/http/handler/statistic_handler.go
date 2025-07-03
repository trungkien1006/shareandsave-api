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

// @Summary Get transaction statistic
// @Description API lấy ra thống kê giao dịch
// @Security BearerAuth
// @Tags statistics
// @Accept json
// @Produce json
// @Success 200 {object} statisticdto.GetTransactionStatisticResponseWrapper
// @Failure 400 {object} enums.AppError
// @Failure 404 {object} enums.AppError
// @Router /statistics/transaction [get]
func (h *StatisticHandler) TotalTransaction(c *gin.Context) {
	total, totalLastMonth, err := h.uc.TotalTransaction(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusNotFound, enums.NewAppError(http.StatusNotFound, err.Error(), enums.ErrNotFound))
		return
	}

	c.JSON(http.StatusOK, statisticdto.GetTransactionStatisticResponseWrapper{
		Code:    http.StatusOK,
		Message: "Fetched transaction statistic successfully",
		Data: statisticdto.GetTransactionStatisticResponse{
			Total:          uint(total),
			TotalLastMonth: uint(totalLastMonth),
		},
	})
}
