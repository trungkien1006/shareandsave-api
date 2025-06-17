package chatapp

import (
	"context"
	"final_project/internal/domain/comment"
	"strconv"
	"time"
)

type UseCase struct {
	commentRepo comment.Repository
}

func NewUseCase(commentRepo comment.Repository) *UseCase {
	return &UseCase{commentRepo: commentRepo}
}

func (uc *UseCase) CreateMessage(ctx context.Context, messages []map[string]string) error {
	var domainComment []comment.Comment

	timeLayout := "2006-01-02 15:04:05.999999999 -0700 MST"

	for _, value := range messages {
		interestID, _ := strconv.Atoi(value["interestID"])
		senderID, _ := strconv.Atoi(value["senderID"])
		receiverID, _ := strconv.Atoi(value["receiverID"])
		createdAt, _ := time.Parse(timeLayout, value["createdAt"])
		isRead, _ := strconv.Atoi(value["isRead"])

		domainComment = append(domainComment, comment.Comment{
			InterestID: uint(interestID),
			SenderID:   uint(senderID),
			ReceiverID: uint(receiverID),
			Content:    value["content"],
			IsRead:     uint(isRead),
			CreatedAt:  createdAt,
		})
	}

	if err := uc.commentRepo.Create(ctx, &domainComment); err != nil {
		return err
	}

	return nil
}
