package enums

type UserStatus int

const (
	UserStatusAll      UserStatus = iota // 0
	UserStatusInactive                   // 1
	UserStatusActive                     // 2
	UserStatusLocked                     // 3
)

func (s UserStatus) String() string {
	return [...]string{"ALL", "INACTIVE", "ACTIVE", "LOCKED"}[s]
}
