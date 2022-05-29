package middlewares

import (
	"github.com/Nico164/FiberGo/app/helpers"
	"github.com/Nico164/FiberGo/app/models"
	"github.com/Nico164/FiberGo/database"
	"github.com/gofiber/fiber/v2"
)

func Protected(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := helpers.ReadJWT(cookie)

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}
	claims, _ := helpers.ExtractToken(token)
	var user models.User

	database.DB.Where("id = ?", claims.Issuer).First(&user)

	c.Locals("user", user)
	return c.Next()
}
