package commentdto

type GetAllCommentRequest struct {
	InterestID int    `form:"interestID" binding:"required"`
	Page       int    `form:"page" binding:"required"`
	Limit      int    `form:"limit" binding:"required"`
	Search     string `form:"search"`
}

func (r *GetAllCommentRequest) SetDefault() {
	if r.Page == 0 {
		r.Page = 1
	}
	if r.Limit == 0 {
		r.Limit = 8
	}
}
