package post

type Post struct {
	ID          uint
	AuthorID    uint
	ItemID      uint
	Title       string
	Description string
	Status      int8
}
