package transactiondto

import "final_project/internal/domain/transaction"

// DTO to Domain
func CreateDTOToDomain(dto CreateTransactionRequest, userID uint) transaction.Transaction {
	var domainItems []transaction.TransactionItem

	for _, value := range dto.Items {
		domainItems = append(domainItems, transaction.TransactionItem{
			PostItemID: value.PostItemID,
			Quantity:   value.Quantity,
		})
	}

	return transaction.Transaction{
		// PostID:     dto.PostID,
		InterestID: dto.InterestID,
		// SenderID:   dto.SenderID,
		ReceiverID: userID,
		Items:      domainItems,
	}
}

// DTO to Domain
func DomainToDTO(domain transaction.Transaction) TransactionDTO {
	dtoItems := make([]TransactionItemDTO, 0)

	for _, value := range domain.Items {
		dtoItems = append(dtoItems, TransactionItemDTO{
			PostItemID: value.PostItemID,
			Quantity:   value.Quantity,
		})
	}

	return TransactionDTO{
		ID:         domain.ID,
		PostID:     domain.PostID,
		InterestID: domain.InterestID,
		SenderID:   domain.SenderID,
		ReceiverID: domain.ReceiverID,
		Items:      dtoItems,
	}
}
