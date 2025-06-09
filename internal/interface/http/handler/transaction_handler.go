package handler

import (
	"final_project/internal/application/transactionapp"
	"final_project/internal/domain/transaction"
	transactiondto "final_project/internal/dto/transactionDTO"
	"final_project/internal/pkg/enums"
	"final_project/internal/pkg/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	uc *transactionapp.UseCase
}

func NewTransactionHandler(uc *transactionapp.UseCase) *TransactionHandler {
	return &TransactionHandler{uc: uc}
}

// @Summary Create a new transaction
// @Description API tạo mới một giao dịch và trả về thông tin giao dịch
// @Security BearerAuth
// @Tags transactions
// @Accept json
// @Produce json
// @Param request body transactiondto.CreateTransactionRequest true "Transaction creation payload"
// @Success 201 {object} transactiondto.CreateTransactionResponseWrapper
// @Failure 400 {object} enums.AppError
// @Failure 409 {object} enums.AppError
// @Router /transactions [post]
func (h *TransactionHandler) Create(c *gin.Context) {
	var (
		req               transactiondto.CreateTransactionRequest
		domainTransaction transaction.Transaction
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrValidate),
		)
		return
	}

	userID, err := helpers.GetUintFromContext(c, "userID")
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrBadRequest))
		return
	}

	domainTransaction = transactiondto.CreateDTOToDomain(req, userID)

	if err := h.uc.CreateTransaction(c.Request.Context(), &domainTransaction); err != nil {
		c.JSON(
			http.StatusConflict,
			enums.NewAppError(http.StatusConflict, err.Error(), enums.ErrNotFound),
		)
		return
	}

	var transactionDTORes transactiondto.TransactionDTO

	transactionDTORes = transactiondto.DomainToDTO(domainTransaction)

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Created transaction successfully",
		"data": transactiondto.CreateTransactionResponse{
			Transaction: transactionDTORes,
		},
	})
}
