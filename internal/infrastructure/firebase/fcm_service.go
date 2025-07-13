package firebase

import (
	"context"
	"final_project/internal/domain/notification"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
)

type fcmService struct {
	client *messaging.Client
}

func NewFCMService(app *firebase.App) (notification.Notifier, error) {
	client, err := app.Messaging(context.Background())

	if err != nil {
		return nil, err
	}

	return &fcmService{client: client}, nil
}

func (f *fcmService) SendToToken(token, title, body string, noti map[string]string) error {
	log.Println("-------token: " + token)
	log.Println("-------body: " + body)

	for key, value := range noti {
		log.Println("---key: " + key)
		log.Println("---value: " + value)
	}

	msg := &messaging.Message{
		Token: token,
		Notification: &messaging.Notification{
			Title: title,
			Body:  body,
		},
		Data: noti,
	}

	_, err := f.client.Send(context.Background(), msg)
	if err != nil {
		log.Println("Gửi thất bại:", err)
		return err
	}
	return nil
}
