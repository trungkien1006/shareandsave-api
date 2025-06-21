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

type ExportInvoiceDTO struct {
	ID                uint                   `json:"id"`
	InvoiceNum        int                    `json:"invoiceNum"`
	SenderID          uint                   `json:"senderID"`
	SenderName        string                 `json:"senderName"`
	ReceiverID        uint                   `json:"receiverID"`
	ReceiverName      string                 `json:"receiverName"`
	Classify          int                    `json:"classify"`
	Description       string                 `json:"description"`
	IsLock            bool                   `json:"isLock"`
	ItemExportInvoice []ItemExportInvoiceDTO `json:"itemExportInvoices"`
	CreatedAt         time.Time              `json:"createdAt"`
}

type ItemExportInvoiceDTO struct {
	ID                uint   `json:"id"`
	InvoiceID         int    `json:"invoiceID"`
	ItemWarehouseID   uint   `json:"itemWarehouseID"`
	ItemWarehouseName string `json:"itemWarehouseName"`
	SKU               string `json:"sku"`
}
