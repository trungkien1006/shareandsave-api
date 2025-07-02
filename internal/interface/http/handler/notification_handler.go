package handler

import (
	"final_project/internal/application/app/notificationapp"
	"final_project/internal/domain/notification"
	notificationdto "final_project/internal/dto/notificationDTO"
	"final_project/internal/pkg/enums"
	"final_project/internal/pkg/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NotificationHandler struct {
	uc *notificationapp.UseCase
}

func NewNotificationHandler(uc *notificationapp.UseCase) *NotificationHandler {
	return &NotificationHandler{uc: uc}
}

// @Summary Get client notifications
// @Description API bao gồm phân trang
// @Security BearerAuth
// @Tags notifications
// @Accept json
// @Produce json
// @Param page query int false "Current page" minimum(1) example(1)
// @Param limit query int false "Number record per page" minimum(1) example(10)
// @Success 200 {object} notificationdto.GetAllNotiResponseWrapper
// @Failure 400 {object} enums.AppError
// @Failure 404 {object} enums.AppError
// @Router /client/notifications [get]
func (h *NotificationHandler) GetAllClient(c *gin.Context) {
	var (
		req        notificationdto.GetAllNotiRequest
		domainReq  notification.GetAllNotiRequest
		domainNoti []notification.Notification
	)

	if err := c.ShouldBindQuery(&req); err != nil {
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

	req.SetDefault()

	domainReq.Page = req.Page
	domainReq.Limit = req.Limit

	unreadCount, totalPage, err := h.uc.GetAllClientNoti(c.Request.Context(), &domainNoti, domainReq, userID)
	if err != nil {
		c.JSON(http.StatusNotFound, enums.NewAppError(http.StatusNotFound, err.Error(), enums.ErrNotFound))
		return
	}

	notiDTORes := make([]notificationdto.NotificationDTO, 0)

	for _, i := range domainNoti {
		notiDTORes = append(notiDTORes, notificationdto.NotificationDomainToDTO(i))
	}

	c.JSON(http.StatusOK, notificationdto.GetAllNotiResponseWrapper{
		Code:    http.StatusOK,
		Message: "Fetched notifications successfully",
		Data: notificationdto.GetAllNotiResponse{
			Notifications: notiDTORes,
			TotalPage:     totalPage,
			UnreadCount:   unreadCount,
		},
	})
}

// @Summary Get admin notifications
// @Description API bao gồm phân trang
// @Security BearerAuth
// @Tags notifications
// @Accept json
// @Produce json
// @Param page query int false "Current page" minimum(1) example(1)
// @Param limit query int false "Number record per page" minimum(1) example(10)
// @Success 200 {object} notificationdto.GetAllNotiResponseWrapper
// @Failure 400 {object} enums.AppError
// @Failure 404 {object} enums.AppError
// @Router /notifications [get]
func (h *NotificationHandler) GetAllAdmin(c *gin.Context) {
	var (
		req        notificationdto.GetAllNotiRequest
		domainReq  notification.GetAllNotiRequest
		domainNoti []notification.Notification
	)

	if err := c.ShouldBindQuery(&req); err != nil {
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

	req.SetDefault()

	domainReq.Page = req.Page
	domainReq.Limit = req.Limit

	unreadCount, totalPage, err := h.uc.GetAllAdminNoti(c.Request.Context(), &domainNoti, domainReq, userID)
	if err != nil {
		c.JSON(http.StatusNotFound, enums.NewAppError(http.StatusNotFound, err.Error(), enums.ErrNotFound))
		return
	}

	notiDTORes := make([]notificationdto.NotificationDTO, 0)

	for _, i := range domainNoti {
		notiDTORes = append(notiDTORes, notificationdto.NotificationDomainToDTO(i))
	}

	c.JSON(http.StatusOK, notificationdto.GetAllNotiResponseWrapper{
		Code:    http.StatusOK,
		Message: "Fetched notifications successfully",
		Data: notificationdto.GetAllNotiResponse{
			Notifications: notiDTORes,
			TotalPage:     totalPage,
			UnreadCount:   unreadCount,
		},
	})
}
