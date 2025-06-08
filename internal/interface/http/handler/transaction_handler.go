package handler

import (
	"final_project/internal/application/transactionapp"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	uc *transactionapp.UseCase
}

func NewTransactionHandler(uc *transactionapp.UseCase) *TransactionHandler {
	return &TransactionHandler{uc: uc}
}

func (h *TransactionHandler) Create(c *gin.Context) {

}
