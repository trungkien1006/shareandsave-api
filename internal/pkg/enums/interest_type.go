package enums

type InterestType int

const (
	InterestTypeAll        InterestType = iota // 0
	InterestTypeInterested                     // 1 Đã quan tâm
	InterestTypeFollowing                      // 2 Đang quan tâm
)

func (s InterestType) String() string {
	return [...]string{"ALL", "INTERESTED", "FOLLOWING"}[s]
}
