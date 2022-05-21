package main

import (
	"fmt"

	"github.com/Nico164/FiberGo/app/routes"
	"github.com/Nico164/FiberGo/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDB()
	app := fiber.New()

	routes.SetupBookRoutes(app)
	routes.SetupStudentRoutes(app)
	routes.SetupTeacherRoutes(app)
	routes.SetupUserRoutes(app)

	app.Listen(":3000")
	fmt.Println("hello world")
}
