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
func UpdateDTOToDomain(dto UpdateTransactionRequest, userID uint, transactionID uint) transaction.Transaction {
	var domainItems []transaction.TransactionItem

	for _, value := range dto.Items {
		domainItems = append(domainItems, transaction.TransactionItem{
			PostItemID: value.PostItemID,
			Quantity:   value.Quantity,
		})
	}

	return transaction.Transaction{
		ID:         transactionID,
		SenderID:   userID,
		Status:     int(dto.Status),
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
		InterestID: domain.InterestID,
		SenderID:   domain.SenderID,
		ReceiverID: domain.ReceiverID,
		Status:     domain.Status,
		Items:      dtoItems,
	}
}

// Domain to DTO
func DomainToDetailDTO(domain transaction.DetailTransaction) DetailTransactionDTO {
	dtoDetailItems := make([]DetailTransactionItemDTO, 0)

	for _, value := range domain.Items {
		dtoDetailItems = append(dtoDetailItems, DetailTransactionItemDTO{
			ItemID:     value.ItemID,
			ItemName:   value.ItemName,
			ItemImage:  value.ItemImage,
			PostItemID: value.PostItemID,
			Quantity:   value.Quantity,
		})
	}

	return DetailTransactionDTO{
		ID:           domain.ID,
		InterestID:   domain.InterestID,
		SenderID:     domain.SenderID,
		ReceiverID:   domain.ReceiverID,
		SenderName:   domain.SenderName,
		ReceiverName: domain.ReceiverName,
		Status:       domain.Status,
		Items:        dtoDetailItems,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
	}
}
