package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/Nico164/FiberGo/app/controllers"
	"github.com/Nico164/FiberGo/app/middlewares"
)

func SetupBookRoutes(router fiber.Router) {
	book := router.Group("/books")

	book.Get("/", controllers.GetBook)
	book.Post("/", middlewares.Protected, controllers.CreateBook)
	book.Put("/:id", middlewares.Protected, controllers.UpdateBook)
	book.Delete("/:id", middlewares.Protected, controllers.DeleteBook)
}
