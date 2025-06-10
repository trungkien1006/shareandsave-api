package importinvoicedto

type CreateImportInvoiceResponse struct {
	ImportInvoice ImportInvoiceDTO `json:"importInvoice"`
}

type CreateImportInvoiceResponseWrapper struct {
	Code    int                         `json:"code"`
	Message string                      `json:"message"`
	Data    CreateImportInvoiceResponse `json:"data"`
}

type GetmportInvoiceResponse struct {
	ImInvoices []ImportInvoiceListDTO `json:"importInvoices"`
	TotalPage  int                    `json:"totalPage"`
}

type GetmportInvoiceResponseWrapper struct {
	Code    int                     `json:"code"`
	Message string                  `json:"message"`
	Data    GetmportInvoiceResponse `json:"data"`
}
