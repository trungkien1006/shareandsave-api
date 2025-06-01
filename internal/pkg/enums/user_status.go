package enums

type UserStatus int

const (
	UserStatusAll      UserStatus = iota // 1
	UserStatusInactive                   // 2
	UserStatusActive                     // 3
	UserStatusLocked                     // 4
)

func (s UserStatus) String() string {
	return [...]string{"ALL", "INACTIVE", "ACTIVE", "LOCKED"}[s]
}
