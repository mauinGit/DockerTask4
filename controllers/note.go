package controllers

import (
	"week4/databases"
	"week4/models"

	"github.com/gofiber/fiber/v3"
)

func CreateNote(c fiber.Ctx) error {
	var Note models.Note

	Title := c.FormValue("title")
	Content := c.FormValue("content")
	Category := c.FormValue("category")

	Note.Title = Title
	Note.Content = Content
	Note.Category = Category

	if err := databases.DB.Create(&Note).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create note",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Note created successfully",
		"data":    Note,
	})
}
