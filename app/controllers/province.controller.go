package controllers

import (
	"github.com/Nico164/FiberGo/app/models"
	"github.com/Nico164/FiberGo/database"
	"github.com/gofiber/fiber/v2"
)

func GetAllProvince(c *fiber.Ctx) error {
	db := database.DB
	var provinces []models.Province

	db.Find(&provinces)

	if len(provinces) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  404,
			"message": "No Provinces found",
		})
	}
	return c.Status(200).JSON(provinces)
}

func GetProvinceById(c *fiber.Ctx) error {
	db := database.DB
	var provinces models.Province
	id := c.Params("id")

	db.Where("id = ?", id).First(&provinces)

	return c.Status(200).JSON(provinces)
}

func CreateProvince(c *fiber.Ctx) error {
	db := database.DB
	province := new(models.Province)

	err := c.BodyParser(province)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"Status":  "error",
			"message": "Review your input",
			"data":    err,
		})
	}
	province.ID = 0
	err = db.Create(&province).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "could not create",
		})
	}
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Created Province",
		"data":    province,
	})
}

func UpdateProvinceById(c *fiber.Ctx) error {
	type updateProvinceType struct {
		Name string `json:"name"`
	}

	db := database.DB
	var province models.Province

	id := c.Params("id")

	db.Find(&province, "id = ?", id)

	if province.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "No provinces",
		})
	}
	var updateProvinceData updateProvinceType
	err := c.BodyParser(&updateProvinceData)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "review input",
		})
	}
	province.Name = updateProvinceData.Name
	db.Save(&province)
	return c.JSON(fiber.Map{
		"status": "success",
		"data":   province,
	})
}

func DeleteProvinceById(c *fiber.Ctx) error {
	db := database.DB
	var province models.Province

	id := c.Params("id")

	db.Find(&province, "id = ?", id)

	if province.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "No province",
		})
	}
	err := db.Delete(&province, "id = ?", id).Error

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "failed to delete province",
		})
	}
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "deleted Province",
	})
}
