package postdto

import (
	"final_project/internal/pkg/enums"
)

type CreatePostRequest struct {
	Email       string         `json:"email" example:"john@gmail.com"`
	FullName    string         `json:"fullName" example:"John Doe"`
	PhoneNumber string         `json:"phoneNumber" example:"0123456789"` // true: anonymous, false: not anonymous
	AuthorID    uint           `json:"authorID"`
	Type        enums.PostType `json:"type" binding:"oneof=0 1 2 3"`
	Title       string         `json:"title" binding:"required" example:"Tôi muốn tìm đồ thất lạc"`
	Info        string         `json:"info"`
	Images      []string       `json:"images" example:"strbase64, strbase64"`
}

type GetAdminPostRequest struct {
	Page        int              `form:"page"`
	Limit       int              `form:"limit"`
	Sort        string           `form:"sort"`
	Order       string           `form:"order" binding:"omitempty,oneof=ASC DESC" example:"ASC"` // Default: ASC
	Status      enums.PostStatus `form:"status" binding:"oneof=0 1 2"`
	Type        enums.PostType   `form:"type" binding:"oneof=0 1 2 3"`
	SearchBy    string           `form:"searchBy" binding:"oneof=fullName email"`
	SearchValue string           `form:"searchValue"`
}

func (r *GetAdminPostRequest) SetDefault() {
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

type UpdatePostRequest struct {
	Title  string           `json:"title" binding:"omitempty" example:"Tôi muốn tìm đồ thất lạc"`
	Info   string           `json:"info" binding:"omitempty"`
	Status enums.PostStatus `json:"status" binding:"oneof=0 1 2"`
	Images []string         `json:"images" binding:"omitempty" example:"strbase64, strbase64"`
}
