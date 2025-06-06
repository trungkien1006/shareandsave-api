package enums

type InterestStatus int

const (
	InterestStatusAll   InterestStatus = iota // 0
	InterestStatusStart                       // 1 Phiếu quan tâm bắt đầu
	InterestStatusEnd                         // 2 Kết thúc phiếu quan tâm
)

func (s InterestStatus) String() string {
	return [...]string{"ALL", "START", "END"}[s]
}
