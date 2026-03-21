package routes

import (
	"week4/controllers"

	"github.com/gofiber/fiber/v3"
)

func RoutesNote(api fiber.Router) {
	api.Post("/note", controllers.CreateNote)
	api.Get("/note", controllers.GetNotes)
	api.Get("/note/:id", controllers.GetNoteByID)
	api.Put("/note/:id", controllers.UpdateNote)
	api.Delete("/note/:id", controllers.DeleteNote)
}