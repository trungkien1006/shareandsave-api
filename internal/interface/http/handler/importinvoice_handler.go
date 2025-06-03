package handler

import (
	"final_project/internal/application/importinvoiceapp"
)

type ImportInvoiceHandler struct {
	uc *importinvoiceapp.UseCase
}

func NewImportInvoiceHandler(uc *importinvoiceapp.UseCase) *ImportInvoiceHandler {
	return &ImportInvoiceHandler{uc: uc}
}

// func (h *ImportInvoiceHandler) CreateImportInvoice(c *gin.Context) {
// 	var (
// 		req                 importinvoicedto.CreateImportInvoiceRequest
// 		domainImportInvoice importinvoice.ImportInvoice
// 	)

// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrValidate))
// 		return
// 	}

// 	if err := validator.Validate.Struct(req); err != nil {
// 		c.JSON(http.StatusBadRequest, enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrValidate))
// 		return
// 	}

// 	domainImportInvoice = importinvoicedto.CreateDTOToDomain(req)
// }
