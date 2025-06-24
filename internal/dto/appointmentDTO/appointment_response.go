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

type GetAppointmentByIDResponseWrapper struct {
	Code    int                        `json:"code"`
	Message string                     `json:"message"`
	Data    GetAppointmentByIDResponse `json:"data"`
}

type GetAppointmentByIDResponse struct {
	Appointment AppointmentDTO `json:"appointment"`
}

type UpdateAppointmentResponseWrapper struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
