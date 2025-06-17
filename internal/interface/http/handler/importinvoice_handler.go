package handler

import (
	"final_project/internal/application/app/importinvoiceapp"
	"final_project/internal/domain/filter"
	importinvoice "final_project/internal/domain/import_invoice"
	importinvoicedto "final_project/internal/dto/importinvoiceDTO"
	"final_project/internal/pkg/enums"
	"final_project/internal/pkg/helpers"
	"final_project/internal/shared/validator"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ImportInvoiceHandler struct {
	uc *importinvoiceapp.UseCase
}

func NewImportInvoiceHandler(uc *importinvoiceapp.UseCase) *ImportInvoiceHandler {
	return &ImportInvoiceHandler{uc: uc}
}

// @Summary Get import invoice
// @Description API bao gồm cả lọc, phân trang và sắp xếp
// @Security BearerAuth
// @Tags import invoice
// @Accept json
// @Produce json
// @Param page query int false "Current page" minimum(1) example(1)
// @Param limit query int false "Number record of page" minimum(1) example(10)
// @Param sort query string false "Sort column" example(authorName, title, createdAt)
// @Param order query string false "Sort type" enum(ASC,DESC) example(ASC, DESC)
// @Param   searchBy   query    string  false  "Trường lọc (vd: senderName, receiverName)"
// @Param   searchValue   query    string  false  "Giá trị lọc (vd:abc@gmail.com, John Doe)"
// @Success 200 {object} importinvoicedto.GetmportInvoiceResponseWrapper
// @Failure 400 {object} enums.AppError
// @Router /import-invoice [get]
func (h *ImportInvoiceHandler) GetAllImportInvoice(c *gin.Context) {
	var (
		req            importinvoicedto.GetImportInvoiceRequest
		importInvoices []importinvoice.GetImportInvoice
		domainReq      filter.FilterRequest
	)

	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrValidate),
		)
		return
	}

	req.SetDefault()

	domainReq.Page = req.Page
	domainReq.Limit = req.Limit
	domainReq.Sort = req.Sort
	domainReq.Order = req.Order
	domainReq.SearchBy = req.SearchBy
	domainReq.SearchValue = req.SearchValue

	totalPage, err := h.uc.GetAllImportInvoice(c.Request.Context(), &importInvoices, domainReq)

	if err != nil {
		c.JSON(
			http.StatusNotFound,
			enums.NewAppError(http.StatusNotFound, err.Error(), "ERR_POST_NOT_FOUND"),
		)
		return
	}

	imInvoiceDTOs := make([]importinvoicedto.ImportInvoiceListDTO, 0)

	for _, value := range importInvoices {
		imInvoiceDTOs = append(imInvoiceDTOs, importinvoicedto.GetDomainToDTO(value))
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Fetched posts successfully",
		"data": importinvoicedto.GetmportInvoiceResponse{
			ImInvoices: imInvoiceDTOs,
			TotalPage:  totalPage,
		},
	})
}

// @Summary Create import invoice
// @Description API tạo phiếu nhập kho kèm lưu kho
// @Security BearerAuth
// @Tags import invoice
// @Accept json
// @Produce json
// @Param request body importinvoicedto.CreateImportInvoiceRequest true "Import invoice creation payload"
// @Success 200 {object} importinvoicedto.CreateImportInvoiceResponseWrapper
// @Failure 400 {object} enums.AppError
// @Failure 404 {object} enums.AppError
// @Router /import-invoice [post]
func (h *ImportInvoiceHandler) CreateImportInvoice(c *gin.Context) {
	var (
		req                 importinvoicedto.CreateImportInvoiceRequest
		domainImportInvoice importinvoice.ImportInvoice
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrValidate))
		return
	}

	if err := validator.Validate.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrValidate))
		return
	}

	userID, err := helpers.GetUintFromContext(c, "userID")
	if err != nil {
		c.JSON(http.StatusBadRequest, enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrBadRequest))
		return
	}

	domainImportInvoice = importinvoicedto.CreateDTOToDomain(req)

	domainImportInvoice.ReceiverID = userID

	if err := h.uc.CreateImportInvoice(c.Request.Context(), &domainImportInvoice); err != nil {
		c.JSON(http.StatusConflict, enums.NewAppError(http.StatusConflict, err.Error(), enums.ErrConflict))
		return
	}

	importinvoiceDTORes := importinvoicedto.ImportInvoiceDomainToDTO(domainImportInvoice)

	c.JSON(http.StatusOK, importinvoicedto.CreateImportInvoiceResponseWrapper{
		Code:    http.StatusOK,
		Message: "Fetched items successfully",
		Data: importinvoicedto.CreateImportInvoiceResponse{
			ImportInvoice: importinvoiceDTORes,
		},
	})
}
