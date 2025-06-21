package handler

import (
	"final_project/internal/application/app/exportinvoiceapp"
	exportinvoice "final_project/internal/domain/export_invoice"
	"final_project/internal/domain/filter"
	exportinvoicedto "final_project/internal/dto/exportinvoiceDTO"
	"final_project/internal/pkg/enums"
	"final_project/internal/pkg/helpers"
	"final_project/internal/shared/validator"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ExportInvoiceHandler struct {
	uc *exportinvoiceapp.UseCase
}

// @Summary Get export invoice
// @Description API bao gồm cả lọc, phân trang và sắp xếp
// @Security BearerAuth
// @Tags export invoice
// @Accept json
// @Produce json
// @Param page query int false "Current page" minimum(1) example(1)
// @Param limit query int false "Number record of page" minimum(1) example(10)
// @Param sort query string false "Sort column" example(authorName, title, createdAt)
// @Param order query string false "Sort type" enum(ASC,DESC) example(ASC, DESC)
// @Param   searchBy   query    string  false  "Trường lọc (vd: senderName, receiverName)"
// @Param   searchValue   query    string  false  "Giá trị lọc (vd:abc@gmail.com, John Doe)"
// @Success 200 {object} exportinvoicedto.GetExportInvoiceResponseWrapper
// @Failure 400 {object} enums.AppError
// @Router /export-invoice [get]
func NewExportInvoiceHandler(uc *exportinvoiceapp.UseCase) *ExportInvoiceHandler {
	return &ExportInvoiceHandler{uc: uc}
}

func (h *ExportInvoiceHandler) GetAll(c *gin.Context) {
	var (
		req                 exportinvoicedto.GetExportInvoiceRequest
		domainExportInvoice []exportinvoice.GetExportInvoice
		domainReq           filter.FilterRequest
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

	totalPage, err := h.uc.GetAllExportInvoice(c.Request.Context(), &domainExportInvoice, domainReq)
	if err != nil {
		c.JSON(
			http.StatusNotFound,
			enums.NewAppError(http.StatusNotFound, err.Error(), "ERR_POST_NOT_FOUND"),
		)
		return
	}

	exInvoiceDTOs := make([]exportinvoicedto.ExportInvoiceListDTO, 0)

	for _, value := range domainExportInvoice {
		exInvoiceDTOs = append(exInvoiceDTOs, exportinvoicedto.GetDomainToDTO(value))
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Fetched export invoices successfully",
		"data": exportinvoicedto.GetExportInvoiceResponse{
			ExInvoices: exInvoiceDTOs,
			TotalPage:  int(totalPage),
		},
	})
}

// @Summary Create export invoice
// @Description API tạo phiếu xuất kho kèm lưu kho
// @Security BearerAuth
// @Tags export invoice
// @Accept json
// @Produce json
// @Param request body exportinvoicedto.CreateExportInvoiceRequest true "export invoice creation payload"
// @Success 200 {object} exportinvoicedto.CreateExportInvoiceResponseWrapper
// @Failure 400 {object} enums.AppError
// @Failure 404 {object} enums.AppError
// @Router /export-invoice [post]
func (h *ExportInvoiceHandler) Create(c *gin.Context) {
	var (
		req                 exportinvoicedto.CreateExportInvoiceRequest
		domainExportInvoice exportinvoice.ExportInvoice
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

	domainExportInvoice = exportinvoicedto.ExportInvoiceDTOToDomain(req)

	domainExportInvoice.ReceiverID = userID

	if err := h.uc.Create(c.Request.Context(), &domainExportInvoice); err != nil {
		c.JSON(http.StatusConflict, enums.NewAppError(http.StatusConflict, err.Error(), enums.ErrConflict))
		return
	}

	exportinvoiceDTORes := exportinvoicedto.ExportInvoiceDomainToDTO(domainExportInvoice)

	c.JSON(http.StatusOK, exportinvoicedto.CreateExportInvoiceResponseWrapper{
		Code:    http.StatusOK,
		Message: "Created export invoice successfully",
		Data: exportinvoicedto.CreateExportInvoiceResponse{
			ExportInvoice: exportinvoiceDTORes,
		},
	})
}
