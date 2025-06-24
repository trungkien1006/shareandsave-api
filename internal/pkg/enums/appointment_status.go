package enums

type AppointmentStatus int

const (
	AppointmentStatusAll      AppointmentStatus = iota // 0
	AppointmentStatusSended                            // 1 Phiếu hẹn gửi thành công
	AppointmentStatusCanceled                          // 2 Phiếu hẹn bị hủy
)

func (s AppointmentStatus) String() string {
	return [...]string{"ALL", "SENDED", "CANCELED"}[s]
}
