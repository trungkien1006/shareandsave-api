package importinvoicedto

import (
	warehousedto "final_project/internal/dto/warehouseDTO"
)

type CreateImportInvoiceResponse struct {
	Items []warehousedto.ItemWarehouse `json:"items"`
}

type CreateImportInvoiceResponseWrapper struct {
	Code    int                         `json:"code"`
	Message string                      `json:"message"`
	Data    CreateImportInvoiceResponse `json:"data"`
}

type GetmportInvoiceResponse struct {
	ImInvoices []ImportInvoiceListDTO `json:"import_invoices"`
	TotalPage  int                    `json:"total_page"`
}

type GetmportInvoiceResponseWrapper struct {
	Code    int                     `json:"code"`
	Message string                  `json:"message"`
	Data    GetmportInvoiceResponse `json:"data"`
}
