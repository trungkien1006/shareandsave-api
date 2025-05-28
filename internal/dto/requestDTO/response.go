package requestdto

type CreateSendOldItemRequestWrapper struct {
	Code    int                      `json:"code"`
	Message string                   `json:"message"`
	Data    CreateSendOldItemRequest `json:"data"`
}

type CreateSendOldItemRequest struct {
	Request RequestSendOldItem `json:"request"`
}
