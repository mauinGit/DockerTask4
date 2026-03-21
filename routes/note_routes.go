package routes

import (
	"week4/controllers"

	"github.com/gofiber/fiber/v3"
)

func routesNote(api fiber.Router) {
	api.Post("/note", controllers.CreateNote)
}