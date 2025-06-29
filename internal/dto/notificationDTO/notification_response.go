package notificationdto

type GetAllNotiResponseWrapper struct {
	Code    int                `json:"code"`
	Message string             `json:"message"`
	Data    GetAllNotiResponse `json:"data"`
}

type GetAllNotiResponse struct {
	Notifications []NotificationDTO `json:"notifications"`
	TotalPage     int               `json:"totalPage"`
}
