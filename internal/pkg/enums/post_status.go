package enums

type PostStatus int

const (
	PostStatusPending  PostStatus = iota // 0 Người dùng đã gửi yêu cầu duyệt bài và đang chờ xử lý
	PostStatusRejected                   // 1 Admin đã từ chối
	PostStatusApproved                   // 2 Admin duyệt
)

func (s PostStatus) String() string {
	return [...]string{"PENDING", "REJECTED", "APPROVED"}[s]
}
