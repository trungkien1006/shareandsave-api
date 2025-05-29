package handler

import (
	"final_project/internal/application/sendrequestapp"
	sendrequest "final_project/internal/domain/send_request"
	"final_project/internal/domain/user"
	sendrequestdto "final_project/internal/dto/sendrequestDTO"
	"final_project/internal/pkg/enums"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SendRequestHandler struct {
	uc *sendrequestapp.UseCase
}

func NewSendRequestHandler(uc *sendrequestapp.UseCase) *SendRequestHandler {
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
// @Router /requests/send-old-item [post]
func (h *SendRequestHandler) CreateSendOldItemRequest(c *gin.Context) {
	var (
		req       sendrequestdto.CreateRequestSendOldItem
		user      user.User
		domainReq sendrequest.SendRequest
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

	domainReq.UserID = req.UserID
	domainReq.Description = req.Description
	domainReq.AppointmentTime = req.AppointmentTime
	domainReq.AppointmentLocation = req.AppointmentLocation

	if err := h.uc.CreateRequest(c.Request.Context(), &domainReq, &user); err != nil {
		c.JSON(
			http.StatusNotFound,
			enums.NewAppError(http.StatusNotFound, err.Error(), enums.ErrNotFound),
		)
		return
	}

	var requestDTO sendrequestdto.RequestSendOldItem

	requestDTO = sendrequestdto.ToRequestDTO(domainReq)

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Create send old item request successfully",
		"data": sendrequestdto.CreateSendOldItemRequestResponse{
			Request: requestDTO,
		},
	})
}
