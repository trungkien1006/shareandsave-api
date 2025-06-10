package handler

import (
	"final_project/internal/application/transactionapp"
	"final_project/internal/domain/transaction"
	transactiondto "final_project/internal/dto/transactionDTO"
	"final_project/internal/pkg/enums"
	"final_project/internal/pkg/helpers"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	uc *transactionapp.UseCase
}

func NewTransactionHandler(uc *transactionapp.UseCase) *TransactionHandler {
	return &TransactionHandler{uc: uc}
}

// @Summary Create transaction
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

	transactionDTORes := transactiondto.DomainToDTO(domainTransaction)

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Created transaction successfully",
		"data": transactiondto.CreateTransactionResponse{
			Transaction: transactionDTORes,
		},
	})
}

// @Summary Update transaction
// @Description API cập nhật một giao dịch và trả về thông tin giao dịch
// @Security BearerAuth
// @Tags transactions
// @Accept json
// @Produce json
// @Param transactionID path int true "ID transaction"
// @Param request body transactiondto.UpdateTransactionRequest true "Transaction update payload"
// @Success 200 {object} transactiondto.UpdateTransactionResponseWrapper
// @Failure 400 {object} enums.AppError
// @Failure 409 {object} enums.AppError
// @Router /transactions/{transactionID} [patch]
func (h *TransactionHandler) Update(c *gin.Context) {
	var (
		req               transactiondto.UpdateTransactionRequest
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

	transactionID, err := strconv.Atoi(c.Param("transactionID"))
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrValidate),
		)
		return
	}

	domainTransaction = transactiondto.UpdateDTOToDomain(req, userID, uint(transactionID))

	// Thực hiện cập nhật giao dịch
	if err := h.uc.UpdateTransaction(c.Request.Context(), &domainTransaction); err != nil {
		c.JSON(
			http.StatusConflict,
			enums.NewAppError(http.StatusConflict, err.Error(), enums.ErrNotFound),
		)
		return
	}

	transactionDTORes := transactiondto.DomainToDTO(domainTransaction)

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Updated transaction successfully",
		"data": transactiondto.UpdateTransactionResponse{
			Transaction: transactionDTORes,
		},
	})
}
