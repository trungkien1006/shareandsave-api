package import_invoice

import "time"

type ImportInvoice struct {
	ID          string
	InvoiceNum  int
	AdminID     uint
	SenderID    uint
	ItemType    string
	SendDate    time.Time
	Description string
	IsLock      bool
}
