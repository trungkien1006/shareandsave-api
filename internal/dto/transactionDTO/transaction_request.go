package transactiondto

import "final_project/internal/pkg/enums"

type GetTransactionRequest struct {
	Page        int                     `form:"page"`
	Limit       int                     `form:"limit"`
	Sort        string                  `form:"sort" binding:"omitempty,oneof=createdAt" example:"createdAt"`
	Order       string                  `form:"order" binding:"omitempty,oneof=ASC DESC" example:"ASC"`
	PostID      uint                    `form:"postID"`
	Status      enums.TransactionStatus `form:"status"`
	SearchBy    string                  `form:"searchBy" binding:"omitempty,oneof=senderID receiverID senderName receiverName interestID"`
	SearchValue string                  `form:"searchValue"`
}

func (r *GetTransactionRequest) SetDefault() {
	if r.Page == 0 {
		r.Page = 1
	}
	if r.Limit == 0 {
		r.Limit = 8
	}
	if r.Order == "" {
		r.Order = "ASC"
	}
}

type CreateTransactionRequest struct {
	// PostID     uint                    `json:"postID" binding:"required"`
	InterestID uint   `json:"interestID" binding:"required"`
	Method     string `json:"method" binding:"required"`
	// SenderID   uint                    `json:"senderID" binding:"required"`
	Items []CreateTransactionItem `json:"items" binding:"required"`
}

type CreateTransactionItem struct {
	PostItemID uint `json:"postItemID" binding:"required"`
	Quantity   int  `json:"quantity" binding:"required"`
}

type UpdateTransactionRequest struct {
	Method string                  `json:"method"`
	Status enums.TransactionStatus `json:"status"`
	Items  []UpdateTransactionItem `json:"items"`
}

type UpdateTransactionItem struct {
	PostItemID uint `json:"postItemID" binding:"required"`
	Quantity   int  `json:"quantity" binding:"required"`
}
