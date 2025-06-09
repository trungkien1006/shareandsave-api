package transactiondto

type CreateTransactionResponseWrapper struct {
	Code    int                       `json:"code"`
	Message string                    `json:"message"`
	Data    CreateTransactionResponse `json:"data"`
}

type CreateTransactionResponse struct {
	Transaction TransactionDTO `json:"transaction"`
}

type UpdateTransactionResponseWrapper struct {
	Code    int                       `json:"code"`
	Message string                    `json:"message"`
	Data    UpdateTransactionResponse `json:"data"`
}

type UpdateTransactionResponse struct {
	Transaction TransactionDTO `json:"transaction"`
}
