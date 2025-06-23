package handler

import (
	"context"
	"final_project/internal/application/worker/appointmentapp"
	"final_project/internal/infrastructure/worker"
	"log"
	"time"
)

type AppointmentHandler struct {
	uc      *appointmentapp.UseCase
	cronjob *worker.AppointmentCronJob
}

func NewAppointmentHandler(uc *appointmentapp.UseCase, cronjob *worker.AppointmentCronJob) *AppointmentHandler {
	return &AppointmentHandler{
		uc:      uc,
		cronjob: cronjob,
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
		w.cronjob.ScheduleAppointment(ctx)

		for {
			select {
			case <-ctx.Done():
				log.Println("Stop create appointment.")
				return
			case <-ticker.C:
				log.Println("Checking appointment...")
				w.cronjob.ScheduleAppointment(ctx)
			}
		}
	}()
}
