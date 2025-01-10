package handler

import "attendance_app/features/attendances"

type AttendanceRequest struct {
	Date       string `json:"date"`
	Checkin    string `json:"checkin"`
	Checkout   string `json:"checkout"`
	IsCheckout bool   `json:"is_checkout"`
}

func AttendanceRequestToCore(input AttendanceRequest) attendances.AttendanceCore {
	return attendances.AttendanceCore{
		Date:       input.Date,
		Checkin:    input.Checkin,
		Checkout:   input.Checkout,
		IsCheckout: false,
	}
}
