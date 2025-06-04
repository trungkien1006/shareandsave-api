package postdto

// Wrapper cho response lấy 1 Post (thêm code/message)
type CreatePostResponseWrapper struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type GetAdminPostResponse struct {
	Posts     []AdminPostDTO `json:"posts"`
	TotalPage int            `json:"totalPage"`
}

// Wrapper cho response lấy 1 Post (thêm code/message)
type GetPostResponseWrapper struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    GetPostResponse `json:"data"`
}

type GetPostResponse struct {
	Posts     []PostWithCountDTO `json:"posts"`
	TotalPage int                `json:"totalPage"`
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

type GetDetailPostResponse struct {
	Post DetailPostDTO `json:"post"`
}

// Wrapper cho response lấy 1 Post (thêm code/message)
type GetDetailPostResponseWrapper struct {
	Code    int                   `json:"code"`
	Message string                `json:"message"`
	Data    GetDetailPostResponse `json:"data"`
}
