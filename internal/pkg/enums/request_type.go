package enums

type RequestType int

const (
	RequestTypeSendOldItem     RequestType = iota // 0
	RequestTypeSendLoseItem                       // 1
	RequestTypeReceiveOldItem                     // 2
	RequestTypeReceiveLoseItem                    // 3
)

func (s RequestType) String() string {
	return [...]string{"SEND_OLD_ITEM", "SEND_LOSE_ITEM", "RECEIVE_OLD_ITEM", "RECEIVE_LOSE_ITEM"}[s]
}
