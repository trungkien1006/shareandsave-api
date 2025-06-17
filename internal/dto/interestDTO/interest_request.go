package interestdto

type GetInterest struct {
	Page   int    `form:"page"`
	Limit  int    `form:"limit"`
	Type   int    `form:"type" example:"1"`
	Sort   string `form:"sort" validate:"omitempty,oneof=createdAt" example:"createdAt"`
	Order  string `form:"order" validate:"omitempty,oneof=ASC DESC" example:"ASC"`
	Search string `form:"search"`
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

type GetByID struct {
	InterestID uint `uri:"interestID"`
}

type CreateInterest struct {
	PostID uint `json:"postID" validate:"required"`
}

type DeleteInterest struct {
	PostID uint `uri:"postID" binding:"required"`
}
