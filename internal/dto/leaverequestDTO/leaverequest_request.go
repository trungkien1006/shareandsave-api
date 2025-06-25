package leaverequestdto

type GetAllLeaveRequestRequest struct {
	Page        int    `form:"page"`
	Limit       int    `form:"limit"`
	Sort        string `form:"sort" binding:"omitempty,oneof=startDate endDate totalDays createdAt"`
	Order       string `form:"order" binding:"omitempty,oneof=ASC DESC" example:"ASC"` // Default: ASC
	SearchBy    string `form:"searchBy" binding:"omitempty,oneof=leaveType reason"`
	SearchValue string `form:"searchValue"`
}

func (r *GetAllLeaveRequestRequest) SetDefault() {
	if r.Page == 0 {
		r.Page = 1
	}
	if r.Limit == 0 {
		r.Limit = 8
	}
	if r.Order == "" {
		r.Order = "ASC"
	}
}
