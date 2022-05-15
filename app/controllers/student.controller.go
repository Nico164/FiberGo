package controllers

import (
	"github.com/Nico164/FiberGo/app/models"
	"github.com/Nico164/FiberGo/database"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetStudent(c *fiber.Ctx) error {
	db := database.DB
	var students []models.Student

	db.Find(&students)

	if len(students) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  404,
			"message": "No students found",
		})
	}
	return c.Status(200).JSON(students)
}

func CreateStudent(c *fiber.Ctx) error {
	db := database.DB
	student := new(models.Student)

	err := c.BodyParser(student)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"Status":  "error",
			"message": "Review your input",
			"data":    err,
		})
	}
	student.ID = uuid.New()
	err = db.Create(&student).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "could not create",
		})
	}
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Created Student",
		"data":    student,
	})
}
func DeleteStudent(c *fiber.Ctx) error {
	db := database.DB
	var student models.Student

	id := c.Params("id")

	db.Find(&student, "id = ?", id)

	if student.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "No students",
		})
	}
	err := db.Delete(&student, "id = ?", id).Error

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "failed to delete student",
		})
	}
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "deleted Note",
	})
}
