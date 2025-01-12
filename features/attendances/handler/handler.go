package handler

import (
	"attendance_app/apps/helpers"
	"attendance_app/features/attendances"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/xuri/excelize/v2"
	"strconv"
	"time"
)

type AttendanceHandler struct {
	attendanceService attendances.AttendanceServiceInterface
}

func (handler *AttendanceHandler) CreateAttendance(c *fiber.Ctx) error {
	now := time.Now()
	attendance := AttendanceRequest{
		UserID:   helpers.UserIDGenerator(),
		Date:     now.Format("2006-01-02"),
		Checkin:  now.Format("15:04:05"),
		Checkout: "",
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
		return c.JSON(helpers.WebResponse(fiber.StatusBadRequest, "bad request", nil))
	}
	attendanceCore := AttendanceRequestToCore(attendance)
	err := handler.attendanceService.Add(attendanceCore)
	if err != nil {
		return c.JSON(helpers.WebResponse(fiber.StatusInternalServerError, "internal server error", nil))
	}
	return c.JSON(helpers.WebResponse(fiber.StatusCreated, "checkin success ", nil))
}

func (handler *AttendanceHandler) UpdateAttendance(c *fiber.Ctx) error {
	userID := c.Params("user_id")
	now := time.Now()
	attendance := AttendanceRequest{
		Checkout:   now.Format("15:04:05"),
		IsCheckout: true,
	}
	data, _ := json.Marshal(attendance)
	var result map[string]interface{}
	errJSON := json.Unmarshal(data, &result)
	if errJSON != nil {
		return errJSON
	}
	errBind := c.Bind(result)
	if errBind != nil {
		return c.JSON(helpers.WebResponse(fiber.StatusBadRequest, "bad request", nil))
	}
	attendanceCore := AttendanceRequestToCore(attendance)
	err := handler.attendanceService.Edit(userID, attendanceCore)
	if err != nil {
		return c.JSON(helpers.WebResponse(fiber.StatusBadRequest, "bad request", nil))
	}
	return c.JSON(helpers.WebResponse(fiber.StatusOK, "checkout success", nil))
}

func (handler *AttendanceHandler) GetAll(c *fiber.Ctx) error {
	result, err := handler.attendanceService.ReadAll()
	if err != nil {
		return c.JSON(helpers.WebResponse(fiber.StatusNotFound, "no data found", nil))
	}
	var resultCore []AttendanceResponse
	for _, v := range result {
		resultCore = append(resultCore, AttendanceResponse{
			ID:       v.ID,
			UserID:   v.UserID,
			Date:     v.Date,
			Checkin:  v.Checkin,
			Checkout: v.Checkout,
		})
	}
	return c.JSON(helpers.WebResponse(fiber.StatusOK, "success get data", resultCore))
}

func (handler *AttendanceHandler) Export(c *fiber.Ctx) error {
	file := excelize.NewFile()
	sheetName := "Attendance"
	file.SetSheetName("Sheet1", sheetName)

	// Write headers to the Excel file
	headers := []string{"ID", "User ID", "Date", "Check In", "Check Out"}
	for i, header := range headers {
		cell := string(rune('A'+i)) + "1" // e.g., A1, B1, C1...
		file.SetCellValue(sheetName, cell, header)
	}

	// Fetch attendance data
	result, err := handler.attendanceService.Download()
	if err != nil || len(result) == 0 {
		return c.JSON(helpers.WebResponse(fiber.StatusNotFound, "no data found", nil))
	}

	// Populate Excel file with attendance data
	for rowIndex, attendance := range result {
		file.SetCellValue(sheetName, "A"+strconv.Itoa(rowIndex+2), attendance.ID)
		file.SetCellValue(sheetName, "B"+strconv.Itoa(rowIndex+2), attendance.UserID)
		file.SetCellValue(sheetName, "C"+strconv.Itoa(rowIndex+2), attendance.Date)
		file.SetCellValue(sheetName, "D"+strconv.Itoa(rowIndex+2), attendance.Checkin)
		file.SetCellValue(sheetName, "E"+strconv.Itoa(rowIndex+2), attendance.Checkout)
	}

	// Serve the file directly
	savePath := "./attendance.xlsx"
	if err := file.SaveAs(savePath); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to save the file on the server",
		})
	}
	return c.Download(savePath)
}

func New(service attendances.AttendanceServiceInterface) *AttendanceHandler {
	return &AttendanceHandler{
		attendanceService: service,
	}
}
