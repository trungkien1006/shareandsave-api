package statisticdto

type GetTransactionStatisticResponseWrapper struct {
	Code    int                             `json:"code"`
	Message string                          `json:"message"`
	Data    GetTransactionStatisticResponse `json:"data"`
}

type GetTransactionStatisticResponse struct {
	Total          uint `json:"total"`
	TotalLastMonth uint `json:"totalLastMonth"`
}
