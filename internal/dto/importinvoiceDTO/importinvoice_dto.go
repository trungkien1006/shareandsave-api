package importinvoicedto

import "time"

type ImportInvoiceListDTO struct {
	ID           uint      `json:"id"`
	SenderName   string    `json:"sender_name"`
	ReceiverName string    `json:"receiver_name"`
	ItemCount    int       `json:"item_count"`
	CreatedAt    time.Time `json:"created_at"`
	Classify     string    `json:"classify"`
}
