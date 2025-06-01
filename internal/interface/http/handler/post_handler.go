package handler

import (
	"final_project/internal/application/postapp"
	"final_project/internal/domain/post"
	"final_project/internal/domain/user"
	postdto "final_project/internal/dto/postDTO"
	userdto "final_project/internal/dto/userDTO"
	"final_project/internal/pkg/enums"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	uc *postapp.UseCase
}

func NewPostHandler(uc *postapp.UseCase) *PostHandler {
	return &PostHandler{uc: uc}
}

// @Summary Get posts
// @Description API bao gồm cả lọc, phân trang và sắp xếp
// @Tags posts
// @Accept json
// @Produce json
// @Param page query int false "Current page" minimum(1) example(1)
// @Param limit query int false "Number record of page" minimum(1) example(10)
// @Param sort query string false "Sort column (vd: fullName, email)" example(name)
// @Param order query string false "Sort type: ASC hoặc DESC" enum(ASC,DESC) example(ASC)
// @Param status query string false "Pending:0, Rejected:1, Approved:2" example(0, 1, 2)
// @Param type query string false "GiveAwayOldItem:0, FoundItem:1, SeekLoseItem:2, Other:3" example(0, 1, 2, 3)
// @Param   searchBy   query    string  false  "Trường lọc (vd: email, fullName)"
// @Param   searchValue   query    string  false  "Giá trị lọc (vd:abc@gmail.com, John Doe)"
// @Success 200 {object} postdto.GetAdminPostResponseWrapper
// @Failure 400 {object} enums.AppError
// @Router /posts [get]
func (h *PostHandler) GetAllAdminPost(c *gin.Context) {
	var (
		req       postdto.GetAdminPostRequest
		posts     []post.Post
		domainReq post.PostFilterRequest
	)

	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrValidate),
		)
		return
	}

	req.SetDefault()

	domainReq.Page = req.Page
	domainReq.Limit = req.Limit
	domainReq.Sort = req.Sort
	domainReq.Order = req.Order
	domainReq.Status = int(req.Status)
	domainReq.Type = int(req.Type)
	domainReq.SearchBy = req.SearchBy
	domainReq.SearchValue = req.SearchValue

	totalPage, err := h.uc.GetAllAdminPost(c.Request.Context(), &posts, domainReq)

	if err != nil {
		c.JSON(
			http.StatusNotFound,
			enums.NewAppError(http.StatusNotFound, err.Error(), "ERR_POST_NOT_FOUND"),
		)
		return
	}

	postsDTORes := make([]postdto.AdminPostDTO, 0)

	for _, post := range posts {
		postsDTORes = append(postsDTORes, postdto.DomainAdminPostToDTO(post))
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Fetched posts successfully",
		"data": postdto.GetAdminPostResponse{
			Posts:     postsDTORes,
			TotalPage: totalPage,
		},
	})
}

// @Summary Create a new post
// @Description API tạo mới một post và trả về thông tin post + user + JWT
// @Tags posts
// @Accept json
// @Produce json
// @Param request body postdto.CreatePostRequest true "Post creation payload"
// @Success 201 {object} postdto.CreatePostResponseWrapper
// @Failure 400 {object} enums.AppError
// @Failure 409 {object} enums.AppError
// @Router /posts [post]
func (h *PostHandler) CreatePost(c *gin.Context) {
	var (
		req        postdto.CreatePostRequest
		domainPost post.CreatePost
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

	postDTORes := postdto.CreatePostDomainToDTO(domainPost)
	userDTORes := userdto.DomainCommonUserToDTO(domainUser)

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Created post successfully",
		"data": postdto.CreatePostResponse{
			Post: postDTORes,
			User: userDTORes,
			JWT:  JWT,
		},
	})
}

// @Summary Update posts
// @Description API cập nhật bài viết kết hợp với patch
// @Tags posts
// @Accept json
// @Produce json
// @Param postID path int true "ID post"
// @Param request body postdto.UpdatePostRequest true "Update post info"
// @Success 200 {object} postdto.UpdatePostResponseWrapper "Updated post successfully"
// @Failure 400 {object} enums.AppError
// @Router /posts/{postID} [patch]
func (h *PostHandler) UpdatePost(c *gin.Context) {
	var req postdto.UpdatePostRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrValidate),
		)
		return
	}

	var post post.Post

	postID, err := strconv.Atoi(c.Param("postID"))
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrValidate),
		)
		return
	}

	post = postdto.UpdateDTOToDomain(req)

	post.ID = uint(postID)

	if err := h.uc.UpdatePost(c.Request.Context(), &post); err != nil {
		c.JSON(
			http.StatusConflict,
			enums.NewAppError(http.StatusConflict, err.Error(), enums.ErrConflict),
		)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Updated post successfully",
	})
}
