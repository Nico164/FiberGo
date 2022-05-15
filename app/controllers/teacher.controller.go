package controllers

import (
	"github.com/Nico164/FiberGo/app/models"
	"github.com/Nico164/FiberGo/database"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetTeacher(c *fiber.Ctx) error {
	db := database.DB
	var teachers []models.Teacher

	db.Find(&teachers)

	if len(teachers) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  404,
			"message": "No Teachers found",
		})
	}
	return c.Status(200).JSON(teachers)
}

func CreateTeacher(c *fiber.Ctx) error {
	db := database.DB
	teacher := new(models.Teacher)

	err := c.BodyParser(teacher)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"Status":  "error",
			"message": "Review your input",
			"data":    err,
		})
	}
	teacher.ID = uuid.New()
	err = db.Create(&teacher).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "could not create",
		})
	}
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Created Teacher",
		"data":    teacher,
	})
}

func UpdateTeacher(c *fiber.Ctx) error {
	type updateTeacherType struct {
		Name   string `json:"name"`
		Major  string `jsone:"major"`
		Phone  string `json:"phone"`
		Email  string `json:"email"`
		Gender string `jsone:"Gender"`
	}

	db := database.DB
	var teacher models.Teacher

	id := c.Params("id")

	db.Find(&teacher, "id = ?", id)

	if teacher.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "No teachers",
		})
	}
	var updateTeacherData updateTeacherType
	err := c.BodyParser(&updateTeacherData)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "review input",
		})
	}

	teacher.Name = updateTeacherData.Name
	teacher.Major = updateTeacherData.Major
	teacher.Phone = updateTeacherData.Phone
	teacher.Email = updateTeacherData.Email
	teacher.Gender = updateTeacherData.Gender
	db.Save((&teacher))

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "changed teacher",
	})
}
