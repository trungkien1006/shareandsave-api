package user

type User struct {
	ID          uint
	RoleID      uint
	RoleName    string
	Email       string
	Password    string
	Avatar      string
	Active      bool
	FullName    string
	PhoneNumber string
	Address     string
	Status      int8
	GoodPoint   int
	Major       string
}
