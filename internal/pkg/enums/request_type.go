package enums

type RequestType int

const (
	RequestTypeSendOldItem    RequestType = iota // 0
	RequestTypeReceiveOldItem                    // 1
	RequestTypeSendLoseItem                      // 2
)

func (s RequestType) String() string {
	return [...]string{"SEND_OLD_ITEM", "RECEIVE_OLD_ITEM", "SEND_LOSE_ITEM"}[s]
}
