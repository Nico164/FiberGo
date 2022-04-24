package controllers

import "github.com/gofiber/fiber/v2"

func Tree(c *fiber.Ctx) error {
	return c.SendString("Hello, " + c.Params("name"))
}
