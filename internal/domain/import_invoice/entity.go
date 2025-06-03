package importinvoice

type ImportInvoice struct {
	ID                string
	InvoiceNum        int
	SenderID          uint
	ReceiverID        uint
	Classify          string
	Description       string
	IsLock            bool
	ItemImportInvoice []ItemImportInvoice
}

type ItemImportInvoice struct {
	ID          uint
	InvoiceID   string
	ItemID      uint
	ItemName    string
	Quantity    int8
	Description string
}
