package controllers

import "github.com/gofiber/fiber/v2"

func About(c *fiber.Ctx) error {
	return c.SendString("about")
}
