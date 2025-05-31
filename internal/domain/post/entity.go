package post

type Post struct {
	ID          uint
	AuthorID    uint
	AuthorName  string
	FullName    string
	Email       string
	PhoneNumber string
	Type        int
	Slug        string
	Title       string
	Content     string
	Info        string
	Status      int8
	Images      []string
}
