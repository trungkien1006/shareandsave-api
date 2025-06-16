package commentdto

type GetCommentResponseWrapper struct {
	Code    int                `json:"code"`
	Message string             `json:"message"`
	Data    GetCommentResponse `json:"data"`
}

type GetCommentResponse struct {
	Comments []CommentDTO `json:"messages"`
}
