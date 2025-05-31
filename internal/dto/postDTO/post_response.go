package postdto

import userdto "final_project/internal/dto/userDTO"

type CreatePostResponse struct {
	Post PostDTO               `json:"post"`
	User userdto.CommonUserDTO `json:"user"`
	JWT  string                `json:"JWT"`
}

// Wrapper cho response lấy 1 Post (thêm code/message)
type CreatePostResponseWrapper struct {
	Code    int                `json:"code"`
	Message string             `json:"message"`
	Data    CreatePostResponse `json:"data"`
}

type GetAdminPostResponse struct {
	Posts     []AdminPostDTO `json:"posts"`
	TotalPage int            `json:"totalPage"`
}

// Wrapper cho response lấy 1 Post (thêm code/message)
type GetAdminPostResponseWrapper struct {
	Code    int                  `json:"code"`
	Message string               `json:"message"`
	Data    GetAdminPostResponse `json:"data"`
}

// Wrapper cho response lấy 1 Post (thêm code/message)
type UpdatePostResponseWrapper struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
