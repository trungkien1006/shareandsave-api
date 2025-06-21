package exportinvoice

import (
	"time"
)

type GetExportInvoice struct {
	ID           uint
	InvoiceNum   uint
	SenderName   string
	ReceiverName string
	Classify     int
	ItemCount    int
	CreatedAt    time.Time
}

type ExportInvoice struct {
	ID                 uint
	InvoiceNum         int
	SenderID           uint
	SenderName         string
	ReceiverID         uint
	ReceiverName       string
	Classify           int
	Description        string
	IsLock             bool
	CreatedAt          time.Time
	ItemExportInvoices []ItemExportInvoice
}

type ItemExportInvoice struct {
	ID              uint
	InvoiceID       int
	ItemWarehouseID uint
	SKU             string
}
