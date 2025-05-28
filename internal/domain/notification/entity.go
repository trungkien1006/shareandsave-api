package notification

type Notification struct {
	ID         uint
	SenderID   uint
	ReceiverID uint
	Type       string
	TargetType string
	TargetID   uint
	Message    string
	IsRead     bool
}
