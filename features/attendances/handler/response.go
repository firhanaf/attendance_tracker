package handler

import "attendance_app/features/attendances"

type AttendanceResponse struct {
	ID       uint   `json:"id"`
	UserID   string `json:"userid"`
	Date     string `json:"date"`
	Checkin  string `json:"checkin"`
	Checkout string `json:"checkout"`
}

func AttendanceCoreToResponse(input attendances.AttendanceCore) AttendanceResponse {
	return AttendanceResponse{
		ID:       input.ID,
		UserID:   input.UserID,
		Date:     input.Date,
		Checkin:  input.Checkin,
		Checkout: input.Checkout,
	}
}
