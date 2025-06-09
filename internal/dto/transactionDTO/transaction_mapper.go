package transactiondto

import "final_project/internal/domain/transaction"

// DTO to Domain
func CreateDTOToDomain(dto CreateTransactionRequest) transaction.Transaction {
	var domainItems []transaction.TransactionItem

	for _, value := range dto.Items {
		domainItems = append(domainItems, transaction.TransactionItem{
			ItemID:   value.ItemID,
			Quantity: value.Quantity,
		})
	}

	return transaction.Transaction{
		ID:         dto.ID,
		PostID:     dto.PostID,
		InterestID: dto.InterestID,
		SenderID:   dto.SenderID,
		ReceiverID: dto.ReceiverID,
		Items:      domainItems,
	}
}
