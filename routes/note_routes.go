package routes

import (
	"week4/controllers"

	"github.com/gofiber/fiber/v3"
)

func RoutesNote(api fiber.Router) {
	api.Post("/note", controllers.CreateNote)
}