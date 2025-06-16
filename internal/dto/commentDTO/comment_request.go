package commentdto

type GetAllCommentRequest struct {
	Page   uint   `json:"page" binding:"required"`
	Limit  uint   `json:"limit" binding:"required"`
	Search string `json:"search"`
}

func (r *GetAllCommentRequest) SetDefault() {
	if r.Page == 0 {
		r.Page = 1
	}
	if r.Limit == 0 {
		r.Limit = 8
	}
}
