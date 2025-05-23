package enums

type UserStatus int

const (
	UserStatusInactive UserStatus = iota // 0
	UserStatusActive                     // 1
	UserStatusLocked                     // 2
)

func (s UserStatus) String() string {
	return [...]string{"INACTIVE", "ACTIVE", "LOCKED"}[s]
}
