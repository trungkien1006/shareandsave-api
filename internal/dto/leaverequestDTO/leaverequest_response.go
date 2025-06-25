package leaverequestdto

type GetLeaveRequestResponseWrapper struct {
	Code    int                     `json:"code"`
	Message string                  `json:"message"`
	Data    GetLeaveRequestResponse `json:"data"`
}

type GetLeaveRequestResponse struct {
	LeaveRequests []LeaveRequestDTO `json:"leaveRequests"`
	TotalPage     int               `json:"totalPage"`
}
