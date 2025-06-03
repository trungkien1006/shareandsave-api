package importinvoicedto

import warehousedto "final_project/internal/dto/warehouseDTO"

type CreateImportInvoiceResponse struct {
	Items []warehousedto.ItemWarehouse `json:"items"`
}

type CreateImportInvoiceResponseWrapper struct {
	Code    int                         `json:"code"`
	Message string                      `json:"message"`
	Data    CreateImportInvoiceResponse `json:"data"`
}
