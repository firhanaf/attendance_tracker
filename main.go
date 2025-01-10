package main

import (
	"attendance_app/apps/config"
	"attendance_app/apps/database"
	"attendance_app/apps/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
)

func main() {
	cfg := config.InitConfig()
	dbMySQL := database.InitMySQL(cfg)
	database.InitialMigration(dbMySQL)

	e := fiber.New()
	e.Use(cors.New())
	e.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	router.InitRouter(dbMySQL, e)

	log.Fatal(e.Listen(":80"))
}
