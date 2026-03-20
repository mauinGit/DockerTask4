package main

import (
	"week4/config"
	"week4/databases"

	"github.com/gofiber/fiber/v3"
)

func main() {
	config.ENVInit()
	databases.DBInit()
	app := fiber.New()
	app.Listen(":3000")
}