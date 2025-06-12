package importinvoicedto

import (
	warehousedto "final_project/internal/dto/warehouseDTO"
	"time"
)

type ImportInvoiceListDTO struct {
	ID           uint      `json:"id"`
	SenderName   string    `json:"senderName"`
	ReceiverName string    `json:"receiverName"`
	ItemCount    int       `json:"itemCount"`
	CreatedAt    time.Time `json:"createdAt"`
	Classify     string    `json:"classify"`
}

type ImportInvoiceDTO struct {
	ID                uint                              `json:"id"`
	InvoiceNum        int                               `json:"invoiceNum"`
	SenderID          uint                              `json:"senderID"`
	SenderName        string                            `json:"senderName"`
	ReceiverID        uint                              `json:"receiverID"`
	ReceiverName      string                            `json:"receiverName"`
	Classify          int                               `json:"classify"`
	Description       string                            `json:"description"`
	IsLock            bool                              `json:"isLock"`
	ItemImportInvoice []ItemImportInvoiceDTO            `json:"itemImportInvoices"`
	Warehouses        []warehousedto.DetailWarehouseDTO `json:"warehouses"`
	ItemCount         int                               `json:"itemCount"`
	CreatedAt         time.Time                         `json:"createdAt"`
}

type ItemImportInvoiceDTO struct {
	ID          uint   `json:"id"`
	InvoiceID   int    `json:"invoiceID"`
	ItemID      uint   `json:"itemID"`
	ItemName    string `json:"itemName"`
	Quantity    int8   `json:"quantity"`
	Description string `json:"description"`
}
