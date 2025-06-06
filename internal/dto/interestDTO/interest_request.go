package interestdto

import "final_project/internal/pkg/enums"

type GetInterest struct {
	Page   int                `query:"page"`
	Limit  int                `query:"limit"`
	Type   enums.InterestType `query:"type" validate:"omitempty,oneof=1 2" example:"1"`
	Sort   string             `query:"sort" validate:"omitempty,oneof=createdAt" example:"createdAt"`
	Order  string             `query:"order" validate:"omitempty,oneof=ASC DESC" example:"ASC"`
	Search string             `query:"search"`
}

func (r *GetInterest) SetDefault() {
	if r.Page == 0 {
		r.Page = 1
	}
	if r.Limit == 0 {
		r.Limit = 8
	}
	if r.Order == "" {
		r.Order = "ASC"
	}
	if r.Type == 0 {
		r.Type = 1
	}
}

type CreateInterest struct {
	PostID uint `json:"postID" validate:"required"`
}

type DeleteInterest struct {
	InterestID uint `uri:"interestID" binding:"required"`
}
