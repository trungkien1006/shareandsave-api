package persistence

import (
	"gorm.io/gorm"
)

type AppointmentRepoDB struct {
	db *gorm.DB
}

func NewAppointmentRepoDB(db *gorm.DB) *AppointmentRepoDB {
	return &AppointmentRepoDB{db: db}
}

// func (r *AppointmentRepoDB) Create(ctx context.Context, appointments map[uint]appointment.Appointment, appointmentItems map[uint][]appointment.AppointmentItem) error {
// 	var (
// 		dbAppointments []dbmodel.Appointment
// 	)

// 	tx := r.db.Debug().WithContext(ctx).Begin()

// 	for key, value := range appointments {
// 		dbAppointments = append(dbAppointments, dbmodel.AppointmentDomainToDB(value, appointmentItems[key]))
// 	}

// 	if

// 	return nil
// }
