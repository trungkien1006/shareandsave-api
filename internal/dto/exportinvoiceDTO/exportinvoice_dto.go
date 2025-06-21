package exportinvoicedto

import "time"

type ExportInvoiceListDTO struct {
	ID           uint      `json:"id"`
	InvoiceNum   uint      `json:"invoiceNum"`
	SenderName   string    `json:"senderName"`
	ReceiverName string    `json:"receiverName"`
	ItemCount    int       `json:"itemCount"`
	CreatedAt    time.Time `json:"createdAt"`
	Classify     string    `json:"classify"`
}
