package grpcserver

import (
	"context"
	"errors"
	"final_project/internal/application/app/commentapp"
	"final_project/internal/domain/comment"
	"final_project/internal/infrastructure/grpcpb"
	"log"
)

type MessageHandlerServer struct {
	grpcpb.UnimplementedMessageHandlerServer
	uc *commentapp.UseCase
}

func NewMessageHandlerServer(uc *commentapp.UseCase) grpcpb.MessageHandlerServer {
	return &MessageHandlerServer{uc: uc}
}

func (h *MessageHandlerServer) StoreMessage(ctx context.Context, message *grpcpb.InputMessage) (*grpcpb.OutputMessage, error) {
	log.Println("Received message:", message.Content)

	domainMessage := comment.Comment{
		InterestID: uint(message.InterestId),
		SenderID:   uint(message.SenderId),
		ReceiverID: uint(message.ReceiverId),
		Content:    message.Content,
		IsRead:     uint(message.IsRead),
		CreatedAt:  message.CreatedAt.AsTime(),
	}

	if err := h.uc.CreateComment(ctx, &domainMessage); err != nil {
		return &grpcpb.OutputMessage{
			Code:    400,
			Message: "Stored failed",
		}, errors.New("lưu tin nhắn thất bại")
	}

	return &grpcpb.OutputMessage{
		Code:    200,
		Message: "Stored successfully",
	}, nil
}
