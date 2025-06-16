package handler

import (
	"final_project/internal/application/commentapp"
)

type CommentHandler struct {
	uc *commentapp.UseCase
}

func NewCommentHandler(uc *commentapp.UseCase) *CommentHandler {
	return &CommentHandler{uc: uc}
}

// func (h *CommentHandler) GetAll(c *gin.Context) {
// 	var (
// 		req           commentdto.GetAllCommentRequest
// 		filter        comment.GetComment
// 		domainComment comment.Comment
// 	)

// }
