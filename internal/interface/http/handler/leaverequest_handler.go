package handler

import (
	"final_project/internal/application/app/leaverequestapp"
)

type LeaveRequestsHandler struct {
	uc *leaverequestapp.UseCase
}

func NewLeaveRequestsHandler(uc *leaverequestapp.UseCase) *LeaveRequestsHandler {
	return &LeaveRequestsHandler{uc: uc}
}

// func (h *LeaveRequestsHandler) GetAll(c *gin.Context) {
// 	var (
// 		req            leaverequestdto.GetAllLeaveRequestRequest
// 		domainLeaveReq leaverequests.LeaveRequest
// 		domainFilter   filter.FilterRequest
// 	)

// 	if err := c.ShouldBindQuery(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, enums.NewAppError(http.StatusBadRequest, err.Error(), enums.ErrValidate))
// 		return
// 	}

// 	req.SetDefault()

// 	domainFilter.Page = req.Page
// 	domainFilter.Limit = req.Limit
// 	domainFilter.Sort = req.Sort
// 	domainFilter.Order = req.Order
// 	domainFilter.SearchBy = req.SearchBy
// 	domainFilter.SearchValue = req.SearchValue
// }
