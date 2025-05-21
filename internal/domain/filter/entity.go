package filter

type FilterRequest struct {
	Page   int    `query:"page" binding:"required"`
	Limit  int    `query:"limit" binding:"required"`
	Sort   string `query:"sort" binding:"required; oneof:ASC DESC"`
	Order  string `query:"order" binding:"required"`
	Filter string `query:"filter" binding:"required"`
}
