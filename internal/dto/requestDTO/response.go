package requestdto

type CreateSendOldItemRequestResponseWrapper struct {
	Code    int                              `json:"code"`
	Message string                           `json:"message"`
	Data    CreateSendOldItemRequestResponse `json:"data"`
}

type CreateSendOldItemRequestResponse struct {
	Request RequestSendOldItem `json:"request"`
}
