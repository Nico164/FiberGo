package controllers

import (
	"strconv"

	"github.com/Nico164/FiberGo/app/helpers"
	"github.com/Nico164/FiberGo/app/models"
	"github.com/Nico164/FiberGo/database"
	"github.com/gofiber/fiber/v2"
)

func GetUser(c *fiber.Ctx) error {
	// 	db := database.DB
	// 	var users []models.User

	// 	db.Find(&users)

	// 	if len(users) == 0 {
	// 		return c.Status(404).JSON(fiber.Map{
	// 			"status":  404,
	// 			"message": "No users found",
	// 		})
	// 	}
	cookie := c.Cookies("jwt")
	return c.Status(200).JSON(cookie)
}

func RegisterUser(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	password, _ := helpers.GeneratePassWord(data["password"])

	user := models.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: string(password),
	}

	database.DB.Create(&user)

	return c.JSON(user)
}

func LoginUser(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User

	database.DB.Where("email = ?", data["email"]).First(&user)

	if user.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}
	if err := helpers.ComPassWord([]byte(user.Password), []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})

	}

	token, err := helpers.NewClaim(strconv.Itoa(int(user.ID)))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "could not login",
		})
	}

	cookie := helpers.SaveCookie(c, "jwt", token)
	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func LogoutUser(c *fiber.Ctx) error {
	cookie := helpers.SaveCookie(c, "jwt", "")

	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "logout success",
	})
}
