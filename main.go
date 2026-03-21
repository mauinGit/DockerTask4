package main

import (
	"time"
	"week4/config"
	"week4/databases"
	"week4/routes"

	"github.com/gofiber/fiber/v3"
)

func main() {
	config.ENVInit()
	time.Sleep(10 * time.Second)
	databases.DBInit()
	databases.DBMigrate()
	app := fiber.New()
	routes.RoutesNote(app)
	app.Listen(":3000")
}
