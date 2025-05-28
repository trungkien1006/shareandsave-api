package filter

type FilterRequest struct {
	Page        int    `query:"page" binding:"required"`
	Limit       int    `query:"limit" binding:"required"`
	Sort        string `query:"sort" `
	Order       string `query:"order" binding:"required; oneof:ASC DESC"`
	SearchBy    string `form:"searchBy"`
	SearchValue string `form:"searchValue"`
}

func (f *FilterRequest) NewFilterRequest(page int, limit int, sort string, order string, filter string, searchBy, searchValue string) *FilterRequest {
	return &FilterRequest{
		Page:        page,
		Limit:       limit,
		Sort:        sort,
		Order:       order,
		SearchBy:    searchBy,
		SearchValue: searchValue,
	}
}
