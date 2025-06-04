package importinvoicedto

import "time"

type ImportInvoiceListDTO struct {
	ID           uint      `json:"id"`
	SenderName   string    `json:"senderName"`
	ReceiverName string    `json:"receiverName"`
	ItemCount    int       `json:"itemCount"`
	CreatedAt    time.Time `json:"createdAt"`
	Classify     string    `json:"classify"`
}
