package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/Nico164/FiberGo/app/controllers"
)

func SetupBookRoutes(router fiber.Router) {
	book := router.Group("/books")

	book.Get("/", controllers.GetBook)
	book.Post("/", controllers.CreateBook)
	book.Put("/:id", controllers.UpdateBook)
	book.Delete("/:id", controllers.DeleteBook)
}
