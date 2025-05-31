package postdto

import (
	"final_project/internal/pkg/enums"
)

type CreatePostRequest struct {
	Email       string         `json:"email" example:"john@gmail.com"`
	FullName    string         `json:"fullName" example:"John Doe"`
	PhoneNumber string         `json:"phoneNumber" example:"0123456789"` // true: anonymous, false: not anonymous
	AuthorID    uint           `json:"authorID"`
	Type        enums.PostType `json:"type" oneof:""`
	Title       string         `json:"title" binding:"required" example:"Tôi muốn tìm đồ thất lạc"`
	Info        string         `json:"info" binding:"required" example:"{"a": "oke", "b": "hihi"}"`
	Images      []string       `json:"images" example:"["strbase64", "strbase64"]"`
}
