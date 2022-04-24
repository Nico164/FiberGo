package main

import (
	"fmt"

	"github.com/Nico164/FiberGo/app/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	routes.SetupBookRoutes(app)
	routes.SetupStudentRoutes(app)

	app.Listen(":3000")
	fmt.Println("hello world")
}
