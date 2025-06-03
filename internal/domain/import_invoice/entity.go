package importinvoice

import "time"

type ImportInvoice struct {
	ID                string
	InvoiceNum        int
	SenderID          uint
	ReceiverID        uint
	Classify          string
	SendDate          time.Time
	Description       string
	IsLock            bool
	ItemImportInvoice []ItemImportInvoice
}

type ItemImportInvoice struct {
	ID          uint
	InvoiceID   string
	ItemID      uint
	Quantity    int8
	Description string
}
