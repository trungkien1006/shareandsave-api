package handler

import (
	"context"
	"final_project/internal/application/worker/chatapp"
	redisapp "final_project/internal/infrastructure/redis"
	"log"
	"time"
)

type ChatHandler struct {
	uc       *chatapp.UseCase
	consumer *redisapp.StreamConsumer
}

func NewChatHandler(c *redisapp.StreamConsumer, uc *chatapp.UseCase) *ChatHandler {
	return &ChatHandler{
		consumer: c,
		uc:       uc,
	}
}

func (w *ChatHandler) Run() error {
	// Chạy goroutine scan pending định kỳ
	go func() {
		for {
			time.Sleep(30 * time.Second)
			log.Println("Checking pending messages...")
			w.consumer.RecoverPending(func(ctx context.Context, data []map[string]string) error {
				return w.uc.CreateMessage(ctx, data)
			})
		}
	}()

	// Chạy consumer chính
	return w.consumer.Consume(func(ctx context.Context, data []map[string]string) error {
		return w.uc.CreateMessage(ctx, data)
	})
}
