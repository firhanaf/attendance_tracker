package router

import (
	"attendance_app/features/attendances/data"
	"attendance_app/features/attendances/handler"
	"attendance_app/features/attendances/service"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, c *fiber.App) {
	attendanceData := data.New(db)
	attendanceService := service.New(attendanceData)
	attendanceHandlerAPI := handler.New(attendanceService)
	c.Post("/attend", attendanceHandlerAPI.CreateAttendance)
	c.Post("/:user_id/checkout", attendanceHandlerAPI.UpdateAttendance)
	c.Get("/attendances", attendanceHandlerAPI.GetAll)
	c.Get("/export", attendanceHandlerAPI.Export)
}
