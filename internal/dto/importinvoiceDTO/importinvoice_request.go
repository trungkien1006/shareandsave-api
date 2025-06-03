package importinvoicedto

import (
	"final_project/internal/pkg/enums"
	"time"
)

type CreateImportInvoiceRequest struct {
	SenderID          uint                             `json:"senderID" validate:"required"`
	ReceiverID        uint                             `json:"receiverID" validate:"required"`
	Classify          enums.ItemClassify               `json:"classify" validate:"required,oneof=1 2" example:"1"`
	SendDate          time.Time                        `json:"sendDate" validate:"required"`
	Description       string                           `json:"description"`
	ItemImportInvoice []CreateItemImportInvoiceRequest `json:"itemImportInvoice" validate:"required"`
}

type CreateItemImportInvoiceRequest struct {
	ItemID      uint   `json:"itemID" validate:"required"`
	Quantity    int8   `json:"quantity" validate:"required"`
	Description string `json:"description"`
}
