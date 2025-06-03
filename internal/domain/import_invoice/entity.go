package importinvoice

import "time"

type ImportInvoice struct {
	ID          string
	InvoiceNum  int
	SenderID    uint
	ReceiverID  uint
	Classify    string
	SendDate    time.Time
	Description string
	IsLock      bool
	IsAnonymous bool
}
