package handler

import (
	"context"
	"final_project/internal/application/worker/chatapp"
	"final_project/internal/infrastructure/worker"
	"log"
	"time"
)

type ChatHandler struct {
	uc       *chatapp.UseCase
	consumer *worker.StreamConsumer
}

func NewChatHandler(c *worker.StreamConsumer, uc *chatapp.UseCase) *ChatHandler {
	return &ChatHandler{
		consumer: c,
		uc:       uc,
	}
}

func (w *ChatHandler) Run(ctx context.Context) error {
	// Chạy goroutine scan pending định kỳ
	go func() {
		ticker := time.NewTicker(30 * time.Minute)
		defer ticker.Stop()

		// Chạy lần đầu luôn (không cần đợi 30 phút)
		log.Println("Checking pending messages...")
		w.consumer.RecoverPending(func(ctx context.Context, data []map[string]string) error {
			return w.uc.CreateMessage(ctx, data)
		})

		for {
			select {
			case <-ctx.Done():
				log.Println("Stop recovering pending messages.")
				return
			case <-ticker.C:
				log.Println("Checking pending messages...")
				err := w.consumer.RecoverPending(func(ctx context.Context, data []map[string]string) error {
					return w.uc.CreateMessage(ctx, data)
				})
				if err != nil {
					log.Printf("RecoverPending error: %v\n", err)
				}
			}
		}
	}()

	// Chạy consumer chính
	return w.consumer.Consume(func(ctx context.Context, data []map[string]string) error {
		return w.uc.CreateMessage(ctx, data)
	})
	// return nil
}
