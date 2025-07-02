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

// @Summary Create system notification
// @Description API tạo thông báo toàn hệ thống
// @Security BearerAuth
// @Tags notifications
// @Accept json
// @Produce json
// @Param request body notificationdto.CreateNotificationRequest true "Create noti info"
// @Success 200 {object} notificationdto.CreateNotiResponseWrapper
// @Failure 400 {object} enums.AppError
// @Failure 404 {object} enums.AppError
// @Router /notifications [post]
func (h *NotificationHandler) CreateSystemNoti(c *gin.Context) {
	var (
		req        notificationdto.CreateNotificationRequest
		domainNoti notification.Notification
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrValidate))
		return
	}

	domainNoti.Content = req.Content
	domainNoti.Type = "system"

	if err := h.uc.CreateSystemNoti(c.Request.Context(), &domainNoti); err != nil {
		c.JSON(
			http.StatusConflict,
			enums.NewAppError(http.StatusConflict, err.Error(), enums.ErrConflict),
		)
		return
	}

	c.JSON(http.StatusOK, notificationdto.CreateNotiResponseWrapper{
		Code:    http.StatusOK,
		Message: "Created notification successfully",
		Data: notificationdto.CreateNotiResponse{
			Notification: notificationdto.NotificationDTO(domainNoti),
		},
	})
}

// @Summary Update is read notifications
// @Description API cập nhật trạng thái đọc toàn bộ thông báo
// @Security BearerAuth
// @Tags notifications
// @Accept json
// @Produce json
// @Success 200 {object} notificationdto.ReadAllNotiResponseWrapper
// @Failure 400 {object} enums.AppError
// @Failure 404 {object} enums.AppError
// @Router /notifications [patch]
func (h *NotificationHandler) ReadAllNoti(c *gin.Context) {
	userID, err := helpers.GetUintFromContext(c, "userID")
	if err != nil {
		c.JSON(http.StatusBadRequest, enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrBadRequest))
		return
	}

	if err := h.uc.ReadAllNoti(c.Request.Context(), uint(userID)); err != nil {
		c.JSON(
			http.StatusConflict,
			enums.NewAppError(http.StatusConflict, err.Error(), enums.ErrConflict),
		)
		return
	}

	c.JSON(http.StatusOK, notificationdto.ReadAllNotiResponseWrapper{
		Code:    http.StatusOK,
		Message: "Updated is read notifications successfully",
		Data:    gin.H{},
	})
}

// @Summary Update is read notification
// @Description API cập nhật trạng thái đọc thông báo
// @Security BearerAuth
// @Tags notifications
// @Accept json
// @Produce json
// @Success 200 {object} notificationdto.ReadAllNotiResponseWrapper
// @Failure 400 {object} enums.AppError
// @Failure 404 {object} enums.AppError
// @Router /notifications/{notificationID} [patch]
func (h *NotificationHandler) ReadNoti(c *gin.Context) {
	var req notificationdto.ReadNotiRequest

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrValidate))
		return
	}

	if err := h.uc.ReadNoti(c.Request.Context(), uint(req.NotificationID)); err != nil {
		c.JSON(
			http.StatusConflict,
			enums.NewAppError(http.StatusConflict, err.Error(), enums.ErrConflict),
		)
		return
	}

	c.JSON(http.StatusOK, notificationdto.ReadAllNotiResponseWrapper{
		Code:    http.StatusOK,
		Message: "Updated is read notification successfully",
		Data:    gin.H{},
	})
}
