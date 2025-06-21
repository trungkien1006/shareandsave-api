package importinvoicedto

type CreateImportInvoiceResponse struct {
	ImportInvoice ImportInvoiceDTO `json:"importInvoice"`
}

type CreateImportInvoiceResponseWrapper struct {
	Code    int                         `json:"code"`
	Message string                      `json:"message"`
	Data    CreateImportInvoiceResponse `json:"data"`
}

type GetImportInvoiceResponse struct {
	ImInvoices []ImportInvoiceListDTO `json:"importInvoices"`
	TotalPage  int                    `json:"totalPage"`
}

type GetImportInvoiceResponseWrapper struct {
	Code    int                      `json:"code"`
	Message string                   `json:"message"`
	Data    GetImportInvoiceResponse `json:"data"`
}
