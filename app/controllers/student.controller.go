package controllers

import "github.com/gofiber/fiber/v2"

func Student(c *fiber.Ctx) error {
	return c.SendString("Hello, " + c.Params("name"))
}

func StudentName(c *fiber.Ctx) error {
	return c.SendString("Hello, " + c.Params("name"))
}
