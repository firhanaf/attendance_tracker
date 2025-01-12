package handler

import "attendance_app/features/attendances"

type AttendanceRequest struct {
	UserID     string `json:"user_id"`
	Date       string `json:"date"`
	Checkin    string `json:"checkin"`
	Checkout   string `json:"checkout"`
	IsCheckout bool   `json:"is_checkout"`
}

func AttendanceRequestToCore(input AttendanceRequest) attendances.AttendanceCore {
	return attendances.AttendanceCore{
		UserID:     input.UserID,
		Date:       input.Date,
		Checkin:    input.Checkin,
		Checkout:   input.Checkout,
		IsCheckout: input.IsCheckout,
	}
}
