package admin

type Admin struct {
	ID       uint
	Email    string
	Password string
	FullName string
	Status   int8
	RoleID   uint
}
