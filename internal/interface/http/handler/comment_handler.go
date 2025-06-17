package handler

import (
	"final_project/internal/application/app/commentapp"
	"final_project/internal/domain/comment"
	commentdto "final_project/internal/dto/commentDTO"
	"final_project/internal/pkg/enums"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CommentHandler struct {
	uc *commentapp.UseCase
}

func NewCommentHandler(uc *commentapp.UseCase) *CommentHandler {
	return &CommentHandler{uc: uc}
}

// @Summary Get messages
// @Description API bao gồm cả tìm kiếm và phân trang
// @Security BearerAuth
// @Tags messages
// @Accept json
// @Produce json
// @Param interestID query int true "InterestID"
// @Param page query int false "Current page" minimum(1) example(1)
// @Param limit query int false "Number record per page" minimum(1) example(10)
// @Param   search   query    string  false "Search message content"
// @Success 200 {object} commentdto.GetCommentResponseWrapper
// @Failure 400 {object} enums.AppError
// @Failure 404 {object} enums.AppError
// @Router /messages [get]
func (h *CommentHandler) GetAll(c *gin.Context) {
	var (
		req           commentdto.GetAllCommentRequest
		filter        comment.GetComment
		domainComment []comment.Comment
	)

	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrValidate))
		return
	}

	req.SetDefault()

	filter.InterestID = req.InterestID
	filter.Page = req.Page
	filter.Limit = req.Limit
	filter.Search = req.Search

	if err := h.uc.GetAllComment(c.Request.Context(), &domainComment, filter); err != nil {
		c.JSON(
			http.StatusNotFound,
			enums.NewAppError(http.StatusNotFound, err.Error(), enums.ErrNotFound),
		)
		return
	}

	commentDTORes := make([]commentdto.CommentDTO, 0)

	for _, value := range domainComment {
		commentDTORes = append(commentDTORes, commentdto.CommentDomainToDTO(value))
	}

	c.JSON(http.StatusOK, commentdto.GetCommentResponseWrapper{
		Code:    http.StatusOK,
		Message: "Fetched messages successfully",
		Data: commentdto.GetCommentResponse{
			Comments: commentDTORes,
		},
	})
}

// @Summary Update is read message
// @Description API cập nhật trạng thái đọc tin nhắn
// @Security BearerAuth
// @Tags messages
// @Accept json
// @Produce json
// @Param interestID path int true "ID interest"
// @Success 200 {object} commentdto.UpdateReadMessageResponseWrapper
// @Failure 400 {object} enums.AppError
// @Failure 404 {object} enums.AppError
// @Router /messages/{interestID} [patch]
func (h *CommentHandler) UpdateReadMessage(c *gin.Context) {
	var req commentdto.GetAllCommentRequest

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrValidate))
		return
	}

	if err := h.uc.UpdateReadMessage(c.Request.Context(), uint(req.InterestID)); err != nil {
		c.JSON(
			http.StatusConflict,
			enums.NewAppError(http.StatusConflict, err.Error(), enums.ErrConflict),
		)
		return
	}

	c.JSON(http.StatusOK, commentdto.UpdateReadMessageResponseWrapper{
		Code:    http.StatusOK,
		Message: "Update is read messages successfully",
		Data:    gin.H{},
	})
}
