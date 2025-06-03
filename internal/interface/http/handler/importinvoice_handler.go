package handler

import (
	"final_project/internal/application/importinvoiceapp"
	importinvoice "final_project/internal/domain/import_invoice"
	"final_project/internal/domain/warehouse"
	importinvoicedto "final_project/internal/dto/importinvoiceDTO"
	warehousedto "final_project/internal/dto/warehouseDTO"
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
		domainWarehouse     []warehouse.Warehouse
		DTOItemWarehouse    []warehousedto.ItemWarehouse
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

	if err := h.uc.CreateImportInvoice(c.Request.Context(), domainImportInvoice, &domainWarehouse); err != nil {
		c.JSON(http.StatusConflict, enums.NewAppError(http.StatusConflict, err.Error(), enums.ErrConflict))
		return
	}

	for _, value := range domainWarehouse {
		for _, v := range value.ItemWareHouse {
			DTOItemWarehouse = append(DTOItemWarehouse, warehousedto.ItemWarehouseDomainToDTO(v))
		}
	}

	c.JSON(http.StatusOK, importinvoicedto.CreateImportInvoiceResponseWrapper{
		Code:    http.StatusOK,
		Message: "Fetched items successfully",
		Data: importinvoicedto.CreateImportInvoiceResponse{
			Items: DTOItemWarehouse,
		},
	})
}
