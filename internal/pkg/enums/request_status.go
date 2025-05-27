package enums

type RequestStatus int

const (
	RequestStatusSendOldItem    RequestStatus = iota // 0
	RequestStatusReceiveOldItem                      // 1
	RequestStatusSendLoseItem                        // 2
)

func (s RequestStatus) String() string {
	return [...]string{"SEND_OLD_ITEM", "RECEIVE_OLD_ITEM", "SEND_LOSE_ITEM"}[s]
}
