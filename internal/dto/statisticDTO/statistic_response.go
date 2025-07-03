package statisticdto

type GetStatisticResponseWrapper struct {
	Code    int                  `json:"code"`
	Message string               `json:"message"`
	Data    GetStatisticResponse `json:"data"`
}

type GetStatisticResponse struct {
	Total          uint `json:"total"`
	TotalLastMonth uint `json:"totalLastMonth"`
}
