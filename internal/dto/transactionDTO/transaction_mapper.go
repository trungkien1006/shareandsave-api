package transactiondto

import "final_project/internal/domain/transaction"

// DTO to Domain
func CreateDTOToDomain(dto CreateTransactionRequest) transaction.Transaction {
	return transaction.Transaction{
		ID:         dto.ID,
		InterestID: dto.InterestID,
		SenderID:   dto.SenderID,
		ReceiverID: dto.ReceiverID,
	}
}
