package controllers

import (
	"os"
	"strconv"

	"github.com/Nico164/FiberGo/app/helpers"
	"github.com/Nico164/FiberGo/app/models"
	"github.com/Nico164/FiberGo/database"
	"github.com/gofiber/fiber/v2"
)

func GetUser(c *fiber.Ctx) error {
	user := helpers.ExtractUser(c)
	return c.Status(200).JSON(fiber.Map{
		"message": "success",
		"data":    user,
	})

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

func ChangeAvatar(c *fiber.Ctx) error {
	file, err := c.FormFile("avatar")

	if err != nil {
		return c.JSON(fiber.Map{
			"message": "failed",
		})
	}

	userLocal := helpers.ExtractUser(c)
	id := userLocal.ID
	fileName := strconv.Itoa(int(id)) + ".png"
	pathName := "public/assets/uploads/avatar/" + fileName
	c.SaveFile(file, "./"+pathName)
	db := database.DB
	var user models.User

	db.Find(&user, "id =?", id)
	user.Avatar = pathName
	user.Name = userLocal.Name
	user.Email = userLocal.Email

	db.Set("gorm:association_autoupdate", false).Save(&user)

	return c.JSON(fiber.Map{
		"message": "success",
		"user":    user,
	})

}

func RemoveAvatar(c *fiber.Ctx) error {
	userLocal := helpers.ExtractUser(c)
	id := userLocal.ID
	fileName := strconv.Itoa(int(id)) + ".png"
	pathName := "public/assets/uploads/avatar/" + fileName
	err := os.Remove("./" + pathName)
	if err != nil {
		return c.JSON(fiber.Map{
			"message": "failed",
			"error":   err,
		})
	}
	db := database.DB
	var user models.User

	db.Find(&user, "id =?", id)
	user.Avatar = ""
	user.Name = userLocal.Name
	user.Email = userLocal.Email

	db.Set("gorm:association_autoupdate", false).Save(&user)

	return c.JSON(fiber.Map{
		"message": "success",
		"data":    fileName,
	})

}
