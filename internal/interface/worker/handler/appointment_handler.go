package handler

import (
	"context"
	"final_project/internal/application/app/appointmentapp"
	"final_project/internal/infrastructure/worker"
	"log"

	"github.com/robfig/cron/v3"
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
	// if err := w.cronjob.ScheduleAppointment(ctx); err != nil {
	// 	log.Println("Schedule appointment error: " + err.Error())
	// }
	// Chạy goroutine scan pending định kỳ
	c := cron.New()

	// Schedule vào 11h đêm mỗi ngày
	_, err := c.AddFunc("0 23 * * *", func() {
		log.Println("Running appointment scheduler at 11PM")
		w.cronjob.ScheduleAppointment(ctx)
	})

	if err != nil {
		log.Fatalf("Failed to schedule cron job: %v", err)
	}

	c.Start()

	// Dừng cron khi context kết thúc
	go func() {
		<-ctx.Done()
		log.Println("Stop cron scheduler")
		c.Stop()
	}()
}
