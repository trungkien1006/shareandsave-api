package importinvoicedto

import (
	"final_project/internal/pkg/enums"
)

type CreateImportInvoiceRequest struct {
	SenderID          uint                             `json:"senderID" validate:"required"`
	Classify          enums.ItemClassify               `json:"classify" validate:"required,oneof=1 2" example:"1"`
	Description       string                           `json:"description"`
	ItemImportInvoice []CreateItemImportInvoiceRequest `json:"itemImportInvoice" validate:"required"`
}

type CreateItemImportInvoiceRequest struct {
	ItemID      uint   `json:"itemID" validate:"required"`
	Quantity    int8   `json:"quantity" validate:"required,min=1"`
	Description string `json:"description"`
}
