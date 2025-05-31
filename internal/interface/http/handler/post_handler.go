package handler

import (
	"final_project/internal/application/postapp"
	"final_project/internal/domain/post"
	"final_project/internal/domain/user"
	postdto "final_project/internal/dto/postDTO"
	userdto "final_project/internal/dto/userDTO"
	"final_project/internal/pkg/enums"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	uc *postapp.UseCase
}

func NewPostHandler(uc *postapp.UseCase) *PostHandler {
	return &PostHandler{uc: uc}
}

// @Summary Create a new post
// @Description API tạo mới một post và trả về thông tin post + user + JWT
// @Tags posts
// @Accept json
// @Produce json
// @Param request body postdto.CreatePostRequest true "Post creation payload"
// @Success 201 {object} postdto.GetPostByIDResponseWrapper
// @Failure 400 {object} enums.AppError
// @Failure 409 {object} enums.AppError
// @Router /posts [post]
func (h *PostHandler) CreatePost(c *gin.Context) {
	var (
		req        postdto.CreatePostRequest
		domainPost post.Post
		domainUser user.User
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrValidate),
		)
		return
	}

	domainPost = postdto.CreateDTOToDomain(req)

	domainUser.FullName = domainPost.FullName
	domainUser.Email = domainPost.Email
	domainUser.PhoneNumber = domainPost.PhoneNumber

	JWT, err := h.uc.CreatePost(c.Request.Context(), &domainPost, &domainUser)

	if err != nil {
		c.JSON(
			http.StatusConflict,
			enums.NewAppError(http.StatusConflict, err.Error(), enums.ErrNotFound),
		)

		return
	}

	postDTORes := postdto.DomainToDTO(domainPost)
	userDTORes := userdto.DomainCommonUserToDTO(domainUser)

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Created post successfully",
		"data": postdto.GetPostByIDResponse{
			Post: postDTORes,
			User: userDTORes,
			JWT:  JWT,
		},
	})
}
