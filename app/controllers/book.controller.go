package controllers

import (
	"github.com/Nico164/FiberGo/app/models"
	"github.com/Nico164/FiberGo/database"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetBook(c *fiber.Ctx) error {
	db := database.DB
	var books []models.Book

	db.Find(&books)

	if len(books) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  404,
			"message": "No books found",
		})
	}
	return c.Status(200).JSON(books)
}

func CreateBook(c *fiber.Ctx) error {
	db := database.DB
	book := new(models.Book)

	err := c.BodyParser(book)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"Status":  "error",
			"message": "Review your input",
			"data":    err,
		})
	}
	book.ID = uuid.New()
	err = db.Create(&book).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "could not create",
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Created Book",
		"data":    book,
	})

}

func UpdateBook(c *fiber.Ctx) error {
	type updateBookType struct {
		Title   string `json:"title"`
		Author  string `jsone:"author"`
		Summary string `json:"summary"`
	}

	db := database.DB
	var book models.Book

	id := c.Params("id")

	db.Find(&book, "id = ?", id)

	if book.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "No book",
		})
	}
	var updateBookData updateBookType
	err := c.BodyParser(&updateBookData)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "review input",
		})
	}

	book.Title = updateBookData.Title
	book.Author = updateBookData.Author
	book.Summary = updateBookData.Summary
	db.Save((&book))

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "changed book",
	})
}

func DeleteBook(c *fiber.Ctx) error {
	db := database.DB
	var book models.Book

	id := c.Params("id")

	db.Find(&book, "id = ?", id)

	if book.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "No book",
		})
	}
	err := db.Delete(&book, "id = ?", id).Error

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "failed to delete book",
		})
	}
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "deleted Note",
	})
}
