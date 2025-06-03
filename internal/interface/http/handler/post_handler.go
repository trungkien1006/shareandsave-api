package handler

import (
	"final_project/internal/application/postapp"
	"final_project/internal/domain/post"
	postdto "final_project/internal/dto/postDTO"
	"final_project/internal/pkg/enums"
	"final_project/internal/pkg/helpers"
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
// @Param sort query string false "Sort column" example(authorName, title, createdAt)
// @Param order query string false "Sort type" enum(ASC,DESC) example(ASC, DESC)
// @Param status query string false "Pending:1, Rejected:2, Approved:3" example(1, 2, 3)
// @Param type query string false "GiveAwayOldItem:1, FoundItem:2, SeekLoseItem:3, Other:4" example(1, 2, 3, 4)
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

// @Summary Get detail post
// @Description API lấy bài viết theo id
// @Tags posts
// @Accept json
// @Produce json
// @Param postID path int true "ID post"
// @Success 200 {object} postdto.GetDetailPostResponseWrapper
// @Failure 400 {object} enums.AppError
// @Failure 409 {object} enums.AppError
// @Router /posts/{postID} [get]
func (h *PostHandler) GetPostByID(c *gin.Context) {
	id := c.Param("postID")

	if id == "0" {
		c.JSON(
			http.StatusBadRequest,
			enums.NewAppError(http.StatusBadRequest, "postID phải khác 0", enums.ErrValidate),
		)
		return
	}

	var postDetail post.DetailPost

	postID, _ := strconv.Atoi(id)

	if err := h.uc.GetPostByID(c.Request.Context(), &postDetail, uint(postID)); err != nil {
		c.JSON(
			http.StatusNotFound,
			enums.NewAppError(http.StatusNotFound, err.Error(), "ERR_POST_NOT_FOUND"),
		)
		return
	}

	detailPostDTO := postdto.DetailPostDomainToDTO(postDetail)

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Get detail post successfully",
		"data": postdto.GetDetailPostResponse{
			Post: detailPostDTO,
		},
	})
}

// @Summary Create a new post
// @Description API tạo mới một post và trả về thông tin post + user + JWT
// @Security BearerAuth
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
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrValidate),
		)
		return
	}

	userID, err := helpers.GetUintFromContext(c, "userID")
	if err != nil {
		c.JSON(http.StatusBadRequest, enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrBadRequest))
		return
	}

	domainPost = postdto.CreateDTOToDomain(req)
	domainPost.AuthorID = uint(userID)

	err = h.uc.CreatePost(c.Request.Context(), &domainPost)

	if err != nil {
		c.JSON(
			http.StatusConflict,
			enums.NewAppError(http.StatusConflict, err.Error(), enums.ErrNotFound),
		)

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Created post successfully",
		"data":    gin.H{},
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
