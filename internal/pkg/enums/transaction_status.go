package enums

type TransactionStatus int

const (
	TransactionStatusAll       TransactionStatus = iota // 0 all
	TransactionStatusPending                            // 1 đợi xác nhận từ chủ bài viết
	TransactionStatusSuccess                            // 2 thành công
	TransactionStatusCancelled                          // 3 đã hủy bởi 1 trong 2 bên
	TransactionStatusRollBack                           // 4 hoàn tác giao dịch
)

func (s TransactionStatus) String() string {
	return [...]string{"ALL", "PENDING", "SUCCESS", "CANCELLED", "ROLL_BACK"}[s]
}
