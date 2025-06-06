package importinvoicedto

import (
	"final_project/internal/pkg/enums"
)

type GetImportInvoiceRequest struct {
	Page        int    `form:"page"`
	Limit       int    `form:"limit"`
	Sort        string `form:"sort" binding:"omitempty,oneof=authorName title createdAT"`
	Order       string `form:"order" binding:"omitempty,oneof=ASC DESC" example:"ASC"` // Default: ASC
	SearchBy    string `form:"searchBy" binding:"omitempty,oneof=title authorName receiverName"`
	SearchValue string `form:"searchValue"`
}

func (r *GetImportInvoiceRequest) SetDefault() {
	if r.Page == 0 {
		r.Page = 1
	}
	if r.Limit == 0 {
		r.Limit = 8
	}
	if r.Sort == "" {
		r.Sort = "id"
	}
	if r.Order == "" {
		r.Order = "ASC"
	}
}

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
