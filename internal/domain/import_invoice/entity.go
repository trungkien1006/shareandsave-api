package importinvoice

import "time"

type ImportInvoice struct {
	ID                uint
	InvoiceNum        int
	SenderID          uint
	SenderName        string
	ReceiverID        uint
	ReceiverName      string
	Classify          int
	Description       string
	IsLock            bool
	ItemImportInvoice []ItemImportInvoice
	ItemCount         int
	CreatedAt         time.Time
}

type ItemImportInvoice struct {
	ID          uint
	InvoiceID   string
	ItemID      uint
	ItemName    string
	Quantity    int8
	Description string
}

type GetImportInvoice struct {
	ID           uint
	SenderName   string
	ReceiverName string
	Classify     string
	ItemCount    int
	CreatedAt    time.Time
}
