package postdto

type AdminPostDTO struct {
	ID         uint   `json:"id"`
	AuthorName string `json:"authorName"`
	Type       int    `json:"type"`
	Slug       string `json:"slug"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Status     int8   `json:"status"`
}
