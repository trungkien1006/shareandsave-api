package handler

import (
	"final_project/internal/application/transactionapp"
)

type TransactionHandler struct {
	uc *transactionapp.UseCase
}

func NewTransactionHandler(uc *transactionapp.UseCase) *TransactionHandler {
	return &TransactionHandler{uc: uc}
}

// func (h *TransactionHandler) Create(c *gin.Context) {
// 	var (
// 		req               transactiondto.CreateTransactionRequest
// 		domainTransaction transaction.Transaction
// 	)

// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(
// 			http.StatusBadRequest,
// 			enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrValidate),
// 		)
// 		return
// 	}

// 	userID, err := helpers.GetUintFromContext(c, "userID")
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrBadRequest))
// 		return
// 	}

// 	domainTransaction = transactiondto.CreateDTOToDomain(req)
// }
