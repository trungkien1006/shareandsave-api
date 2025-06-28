package handler

import "final_project/internal/application/app/notificationapp"

type NotificationHandler struct {
	uc *notificationapp.UseCase
}

func NewNotificationHandler(uc *notificationapp.UseCase) *NotificationHandler {
	return &NotificationHandler{uc: uc}
}
