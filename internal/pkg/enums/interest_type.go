package enums

type InterestType int

const (
	InterestTypeAll        InterestType = iota // 0
	InterestTypeInterested                     // 1 Đã quan tâm
	InterestTypeFollowedBy                     // 2 Được quan tâm
)

func (s InterestType) String() string {
	return [...]string{"ALL", "INTERESTED", "FOLLOWED_BY"}[s]
}
