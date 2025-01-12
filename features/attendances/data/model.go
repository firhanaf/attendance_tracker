package data

import (
	"attendance_app/features/attendances"
	"gorm.io/gorm"
)

type Attendances struct {
	gorm.Model
	UserID     string
	Date       string
	Checkin    string
	Checkout   string
	IsCheckout bool
}

func AttendanceCoreToModel(input attendances.AttendanceCore) Attendances {
	return Attendances{
		UserID:     input.UserID,
		Date:       input.Date,
		Checkin:    input.Checkin,
		Checkout:   input.Checkout,
		IsCheckout: input.IsCheckout,
	}
}
