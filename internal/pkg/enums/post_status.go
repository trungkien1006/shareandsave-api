package enums

type PostStatus int

const (
	PostStatusAll      PostStatus = iota // 0 Người dùng đã gửi yêu cầu duyệt bài và đang chờ xử lý
	PostStatusPending                    // 1 Người dùng đã gửi yêu cầu duyệt bài và đang chờ xử lý
	PostStatusRejected                   // 2 Admin đã từ chối
	PostStatusApproved                   // 3 Admin duyệt
)

func (s PostStatus) String() string {
	return [...]string{"ALL", "PENDING", "REJECTED", "APPROVED"}[s]
}
