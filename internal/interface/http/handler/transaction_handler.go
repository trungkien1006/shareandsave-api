package handler

import (
	"final_project/internal/application/app/transactionapp"
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

// @Summary Get transaction
// @Description API bao gồm cả lọc, phân trang và sắp xếp
// @Security BearerAuth
// @Tags transactions
// @Accept json
// @Produce json
// @Param page query int false "Current page" minimum(1) example(1)
// @Param limit query int false "Number record per page" minimum(1) example(10)
// @Param sort query string false "Sort column (vd: name)" example(name)
// @Param order query string false "Sort type: ASC hoặc DESC" enum(ASC,DESC) example(ASC)
// @Param status query string false "Pending:1, Success:2, Cancelled:3" example(1, 2, 3)
// @Param postID query int false "Id bài viết"
// @Param   searchBy   query    string  false  "Trường lọc (senderID, senderName, receiverID, receiverName, interestID)"
// @Param   searchValue   query    string  false  "Giá trị lọc:"
// @Success 200 {object} transactiondto.FilterTransactionResponseWrapper
// @Failure 400 {object} enums.AppError
// @Failure 404 {object} enums.AppError
// @Router /transactions [get]
func (h *TransactionHandler) GetAll(c *gin.Context) {
	var (
		req                transactiondto.GetTransactionRequest
		filter             transaction.FilterTransaction
		domainTransactions []transaction.DetailTransaction
	)

	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrValidate),
		)
		return
	}

	req.SetDefault()

	filter.Page = req.Page
	filter.Limit = req.Limit
	filter.Sort = req.Sort
	filter.Order = req.Order
	filter.PostID = req.PostID
	filter.Status = req.Status
	filter.SearchBy = req.SearchBy
	filter.SearchValue = req.SearchValue

	totalPage, err := h.uc.GetAllTransaction(c.Request.Context(), &domainTransactions, filter)
	if err != nil {
		c.JSON(
			http.StatusNotFound,
			enums.NewAppError(http.StatusNotFound, err.Error(), enums.ErrNotFound),
		)
		return
	}

	transactionDTORes := make([]transactiondto.DetailTransactionDTO, 0)

	for _, value := range domainTransactions {
		transactionDTORes = append(transactionDTORes, transactiondto.DomainToDetailDTO(value))
	}

	c.JSON(http.StatusOK, transactiondto.FilterTransactionResponseWrapper{
		Code:    http.StatusOK,
		Message: "Fetched transactions successfully",
		Data: transactiondto.FilterTransactionResponse{
			Transactions: transactionDTORes,
			TotalPage:    totalPage,
		},
	})
}

// @Summary Get detail pending transaction
// @Description API lấy chi tiết giao dịch đang chờ
// @Security BearerAuth
// @Tags transactions
// @Accept json
// @Produce json
// @Param interestID path int true "ID interest"
// @Success 201 {object} transactiondto.GetPendingTransactionResponseWrapper
// @Failure 400 {object} enums.AppError
// @Failure 409 {object} enums.AppError
// @Router /transactions/{interestID} [get]
func (h *TransactionHandler) GetDetailPendingTransaction(c *gin.Context) {
	var (
		req               transactiondto.GetPendingTransaction
		domainTransaction transaction.DetailTransaction
	)

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrValidate),
		)
		return
	}

	if err := h.uc.GetDetailPendingTransaction(c.Request.Context(), &domainTransaction, req.InterestID); err != nil {
		c.JSON(
			http.StatusConflict,
			enums.NewAppError(http.StatusConflict, err.Error(), enums.ErrNotFound),
		)
		return
	}

	transactionDTORes := transactiondto.DomainToDetailDTO(domainTransaction)

	c.JSON(http.StatusOK, transactiondto.GetPendingTransactionResponseWrapper{
		Code:    http.StatusOK,
		Message: "Fetched detail pending transaction successfully",
		Data: transactiondto.GetPendingTransactionResponse{
			Transaction: transactionDTORes,
		},
	})
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
