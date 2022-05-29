package routes

import (
	"github.com/Nico164/FiberGo/app/controllers"
	"github.com/Nico164/FiberGo/app/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetupProvinceRoutes(router fiber.Router) {
	province := router.Group("/provinces")

	province.Get("/", controllers.GetAllProvince)
	province.Get("/:id", controllers.GetProvinceById)
	province.Post("/", middlewares.Protected, controllers.CreateProvince)
	province.Put("/:id", middlewares.Protected, controllers.UpdateProvinceById)
	province.Delete("/:id", middlewares.Protected, controllers.DeleteProvinceById)

}
