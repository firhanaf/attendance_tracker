package handler

import (
	"attendance_app/apps/helpers"
	"attendance_app/features/attendances"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"time"
)

type AttendanceHandler struct {
	attendanceService attendances.AttendanceServiceInterface
}

func (handler *AttendanceHandler) CreateAttendance(c *fiber.Ctx) error {
	now := time.Now()
	attendance := AttendanceRequest{
		Date:       now.Format("2006-01-02"),
		Checkin:    now.Format("15:04"),
		Checkout:   "",
		IsCheckout: false,
	}
	data, _ := json.Marshal(attendance)
	var result map[string]interface {
	}
	errJson := json.Unmarshal(data, &result)
	if errJson != nil {
		return errJson
	}
	errBind := c.Bind(result)
	if errBind != nil {
		return c.JSON(helpers.WebResponse(fiber.StatusBadRequest, "Bad Request", nil))
	}
	attendanceCore := AttendanceRequestToCore(attendance)
	err := handler.attendanceService.Add(attendanceCore)
	if err != nil {
		return c.JSON(helpers.WebResponse(fiber.StatusInternalServerError, "Internal Server Error", nil))
	}
	return c.JSON(helpers.WebResponse(fiber.StatusCreated, "Attendance Success", nil))
}

func New(service attendances.AttendanceServiceInterface) *AttendanceHandler {
	return &AttendanceHandler{
		attendanceService: service,
	}
}
