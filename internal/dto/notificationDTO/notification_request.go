package notificationdto

type GetAllNotiRequest struct {
	Page  int `form:"page"`
	Limit int `form:"limit"`
}

func (r *GetAllNotiRequest) SetDefault() {
	if r.Page == 0 {
		r.Page = 1
	}
	if r.Limit == 0 {
		r.Limit = 10
	}
}

type ReadNotiRequest struct {
	NotificationID uint `uri:"notificationID"`
}
