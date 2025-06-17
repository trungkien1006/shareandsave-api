package handler

import (
	"final_project/internal/application/worker/chatapp"
)

type ChatHandler struct {
	uc *chatapp.UseCase
}

func NewChatHandler(uc *chatapp.UseCase) *ChatHandler {
	return &ChatHandler{uc: uc}
}
