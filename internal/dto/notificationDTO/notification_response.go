package notificationdto

type GetAllNotiResponseWrapper struct {
	Code    int                `json:"code"`
	Message string             `json:"message"`
	Data    GetAllNotiResponse `json:"data"`
}

type GetAllNotiResponse struct {
	Notifications []NotificationDTO `json:"notifications"`
	TotalPage     int               `json:"totalPage"`
	UnreadCount   int64             `json:"unreadCount"`
}

type ReadAllNotiResponseWrapper struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ReadNotiResponseWrapper struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type CreateNotiResponseWrapper struct {
	Code    int                `json:"code"`
	Message string             `json:"message"`
	Data    CreateNotiResponse `json:"data"`
}

type CreateNotiResponse struct {
	Notification NotificationDTO `json:"notification"`
}
