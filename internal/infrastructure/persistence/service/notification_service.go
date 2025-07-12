package persistence

import (
	"context"
	"errors"
	"final_project/internal/domain/notification"
	"final_project/internal/domain/redis"
	"final_project/internal/domain/user"
	"fmt"
	"strconv"
)

type NotificationService struct {
	repo      notification.Repository
	redisRepo redis.Repository
	userRepo  user.Repository
	fcm       notification.Notifier
}

func NewNotificationService(repo notification.Repository, redisRepo redis.Repository, userRepo user.Repository, fcm notification.Notifier) *NotificationService {
	return &NotificationService{
		repo:      repo,
		redisRepo: redisRepo,
		userRepo:  userRepo,
		fcm:       fcm,
	}
}

func (s *NotificationService) CreateAndPushSocket(ctx context.Context, noti *notification.Notification) error {
	var (
		senderName   string = ""
		receiverName string = ""
		err          error
	)

	if err := s.repo.Create(ctx, noti); err != nil {
		return err
	}

	if noti.SenderID != nil {
		senderName, err = s.userRepo.GetUserNameByID(ctx, *noti.SenderID)
		if err != nil {
			return err
		}
	}

	if noti.ReceiverID != nil {
		receiverName, err = s.userRepo.GetUserNameByID(ctx, *noti.ReceiverID)
		if err != nil {
			return err
		}
	}

	notiMap := map[string]interface{}{
		"ID":           noti.ID,
		"senderID":     noti.SenderID,
		"senderName":   senderName,
		"receiverID":   noti.ReceiverID,
		"receiverName": receiverName,
		"type":         noti.Type,
		"targetType":   noti.TargetType,
		"targetID":     noti.TargetID,
		"content":      noti.Content,
		"isRead":       noti.IsRead,
		"createdAt":    noti.CreatedAt,
	}

	fmt.Println("----Thời gian tạo thông báo: " + noti.CreatedAt.String())

	if err := s.redisRepo.InsertToStream(ctx, "notistream", notiMap); err != nil {
		return err
	}

	if noti.ReceiverID != nil {
		token, err := s.redisRepo.GetFromRedis(ctx, "user:"+strconv.Itoa(int(*noti.ReceiverID))+":fcmToken")
		if err != nil {
			return err
		}

		isRead := ""

		if noti.IsRead {
			isRead = "true"
		} else {
			isRead = "false"
		}

		senderID := ""
		receiverID := ""

		if noti.SenderID != nil {
			senderID = strconv.Itoa(int(*noti.SenderID))
		}

		if noti.ReceiverID != nil {
			receiverID = strconv.Itoa(int(*noti.ReceiverID))
		}

		notiMapStr := map[string]string{
			"ID":           strconv.Itoa(int(noti.ID)),
			"senderID":     senderID,
			"senderName":   senderName,
			"receiverID":   receiverID,
			"receiverName": receiverName,
			"type":         noti.Type,
			"targetType":   noti.TargetType,
			"targetID":     strconv.Itoa(int(noti.TargetID)),
			"content":      noti.Content,
			"isRead":       isRead,
			"createdAt":    noti.CreatedAt.String(),
		}

		if err := s.fcm.SendToToken(token, "Bạn có thông báo mới!", noti.Content, notiMapStr); err != nil {
			return errors.New("Có lỗi khi gửi thông báo đẩy: " + err.Error())
		}
	}

	return nil
}
