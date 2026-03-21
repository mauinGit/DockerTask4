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
		"message"	: "Note created successfully",
		"data"		:  Note,
	})
}

func GetNotes(c fiber.Ctx) error {
	var Notes []models.Note

	if err := databases.DB.Find(&Notes).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get notes",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message"	: "Success get all notes",
		"data"		:  Notes,
	})
}

func GetNoteByID(c fiber.Ctx) error {
	id := c.Params("id")
	var Note models.Note

	if err := databases.DB.First(&Note, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Note not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message"	: "Success get note",
		"data"		:  Note,
	})
}

func UpdateNote(c fiber.Ctx) error {
	id := c.Params("id")
	var Note models.Note

	if err := databases.DB.First(&Note, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Note not found",
		})
	}

	Title := c.FormValue("title")
	Content := c.FormValue("content")
	Category := c.FormValue("category")

	Note.Title = Title
	Note.Content = Content
	Note.Category = Category

	if err := databases.DB.Save(&Note).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update note",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Note updated successfully",
		"data":    Note,
	})
}

func DeleteNote(c fiber.Ctx) error {
	id := c.Params("id")
	var Note models.Note

	if err := databases.DB.First(&Note, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Note not found",
		})
	}

	if err := databases.DB.Delete(&Note).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete note",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Note deleted successfully",
	})
}