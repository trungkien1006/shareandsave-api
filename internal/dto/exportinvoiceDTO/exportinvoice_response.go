package exportinvoicedto

type GetExportInvoiceResponse struct {
	ExInvoices []ExportInvoiceListDTO `json:"exportInvoices"`
	TotalPage  int                    `json:"totalPage"`
}

type GetExportInvoiceResponseWrapper struct {
	Code    int                      `json:"code"`
	Message string                   `json:"message"`
	Data    GetExportInvoiceResponse `json:"data"`
}

type CreateExportInvoiceResponse struct {
	ExportInvoice ExportInvoiceDTO `json:"exportInvoice"`
}

type CreateExportInvoiceResponseWrapper struct {
	Code    int                         `json:"code"`
	Message string                      `json:"message"`
	Data    CreateExportInvoiceResponse `json:"data"`
}
