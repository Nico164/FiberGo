package routes

import (
	"github.com/Nico164/FiberGo/app/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupTeacherRoutes(router fiber.Router) {
	teacher := router.Group("/teacher")

	teacher.Get("/", controllers.GetTeacher)
	teacher.Post("/", controllers.CreateTeacher)
	teacher.Put("/:id", controllers.UpdateTeacher)

}
