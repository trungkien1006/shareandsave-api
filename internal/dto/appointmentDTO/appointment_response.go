package appointmentdto

type GetAppointmentResponseWrapper struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Data    GetAppointmentResponse `json:"data"`
}

type GetAppointmentResponse struct {
	Appointments []AppointmentDTO `json:"appointments"`
	TotalPage    int              `json:"totalPage"`
}
