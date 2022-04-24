package routes

import (
	"github.com/Nico164/FiberGo/app/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupStudentRoutes(router fiber.Router) {
	student := router.Group("/student")

	student.Get("/:name", controllers.Student)
	student.Get("/:name?", controllers.StudentName)

}
