package enums

type RequestStatus int

const (
	RequestStatusPending     RequestStatus = iota // 0 Người dùng đã gửi yêu cầu và đang chờ xử lý
	RequestStatusWaitingUser                      // 1 Admin đã xác nhận và đang chờ người dùng xác nhận lần cuối
	RequestStatusReject                           // 2 Admin đã từ chối yêu cầu
	RequestStatusApprove                          // 3 Người dùng đồng ý
	RequestStatusFail                             // 4 Người dùng đã hủy
)

func (s RequestStatus) String() string {
	return [...]string{"PENDING", "WAITING_USER", "REJECT", "APPROVE", "FAIL"}[s]
}
