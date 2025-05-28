package export_invoice

import "time"

type ExportInvoice struct {
	ID          string
	InvoiceNum  int
	AdminID     uint
	ReceiverID  uint
	ItemType    string
	ReceiveDate time.Time
	Description string
	IsLock      bool
}
