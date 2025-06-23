package handler

import (
	"context"
	"final_project/internal/application/worker/appointmentapp"
	"log"
	"time"
)

type AppointmentHandler struct {
	uc *appointmentapp.UseCase
}

func NewAppointmentHandler(uc *appointmentapp.UseCase) *AppointmentHandler {
	return &AppointmentHandler{
		uc: uc,
	}
}

func (w *AppointmentHandler) Run(ctx context.Context) {
	// Chạy goroutine scan pending định kỳ
	go func() {
		ticker := time.NewTicker(30 * time.Minute)
		defer ticker.Stop()

		// Chạy lần đầu luôn (không cần đợi 30 phút)
		log.Println("Checking appointment...")
		// w.consumer.RecoverPending(func(ctx context.Context, data []map[string]string) error {
		// 	return w.uc.CreateMessage(ctx, data)
		// })

		for {
			select {
			case <-ctx.Done():
				log.Println("Stop create appointment.")
				return
			case <-ticker.C:
				log.Println("Checking appointment...")
				// err := w.consumer.RecoverPending(func(ctx context.Context, data []map[string]string) error {
				// 	return w.uc.CreateMessage(ctx, data)
				// })
				// if err != nil {
				// 	log.Printf("RecoverPending error: %v\n", err)
				// }
			}
		}
	}()
}
