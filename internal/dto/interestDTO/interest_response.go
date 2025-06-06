package interestdto

type GetInterestResponseWrapper struct {
	Code    int                 `json:"code"`
	Message string              `json:"message"`
	Data    GetInterestResponse `json:"data"`
}

type GetInterestResponse struct {
	Interests []PostInterest `json:"interests"`
	TotalPage int            `json:"totalPage"`
}

type CreateInterestResponseWrapper struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Data    CreateInterestResponse `json:"data"`
}

type CreateInterestResponse struct {
	InterestID uint `json:"interestID"`
}

type DeleteInterestResponseWrapper struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Data    DeleteInterestResponse `json:"data"`
}

type DeleteInterestResponse struct {
	InterestID uint `json:"interestID"`
}
