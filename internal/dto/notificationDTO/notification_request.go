package notificationdto

type GetAllNotiRequest struct {
	Page  int `form:"page"`
	Limit int `form:"limit"`
}
