package handler

import (
	"final_project/internal/application/requestapp"
	"final_project/internal/domain/request"
	"final_project/internal/domain/user"
	requestdto "final_project/internal/dto/requestDTO"
	"final_project/internal/pkg/enums"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SendRequestHandler struct {
	uc *requestapp.UseCase
}

func NewSendRequestHandler(uc *requestapp.UseCase) *SendRequestHandler {
	return &SendRequestHandler{uc: uc}
}

// @Summary Create request to send old item
// @Description API gửi yêu cầu gửi đồ cũ
// @Tags requests
// @Accept json
// @Produce json
// @Param request body requestdto.CreateRequestSendOldItem true "Create request send old item"
// @Success 201 {object} requestdto.CreateSendOldItemRequestResponseWrapper "Created request successfully"
// @Failure 400 {object} enums.AppError
// @Router /request-sends [post]
func (h *SendRequestHandler) CreateSendOldItemRequest(c *gin.Context) {
	var (
		req       requestdto.CreateRequestSendOldItem
		user      user.User
		domainReq request.SendRequest
	)
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrValidate),
		)
		return
	}

	if req.UserID == 0 {
		user.FullName = req.FullName
		user.Email = req.Email
		user.PhoneNumber = req.PhoneNumber
	}

	domainReq = requestdto.ToDomainRequest(req)

	if err := h.uc.CreateRequest(c.Request.Context(), &domainReq, &user); err != nil {
		c.JSON(
			http.StatusNotFound,
			enums.NewAppError(http.StatusNotFound, err.Error(), enums.ErrNotFound),
		)
		return
	}

	var requestDTO requestdto.RequestSendOldItem

	requestDTO = requestdto.DomainToDTORequest(domainReq)

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Create send old item request successfully",
		"data": requestdto.CreateSendOldItemRequestResponse{
			Request: requestDTO,
		},
	})
}
