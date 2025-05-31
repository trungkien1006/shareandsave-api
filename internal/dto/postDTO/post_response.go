package postdto

import userdto "final_project/internal/dto/userDTO"

type GetPostByIDResponse struct {
	Post AdminPostDTO          `json:"post"`
	User userdto.CommonUserDTO `json:"user"`
	JWT  string                `json:"JWT"`
}

// Wrapper cho response lấy 1 Post (thêm code/message)
type GetPostByIDResponseWrapper struct {
	Code    int                 `json:"code"`
	Message string              `json:"message"`
	Data    GetPostByIDResponse `json:"data"`
}
