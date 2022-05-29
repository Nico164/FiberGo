package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/Nico164/FiberGo/app/controllers"
	"github.com/Nico164/FiberGo/app/middlewares"
)

func SetupUserRoutes(router fiber.Router) {
	user := router.Group("/user")

	user.Get("/", middlewares.Protected, controllers.GetUser)
	user.Post("/login", controllers.LoginUser)
	user.Post("/register", controllers.RegisterUser)
	user.Post("/logout", middlewares.Protected, controllers.LogoutUser)
	user.Post("/Avatar", middlewares.Protected, controllers.ChangeAvatar)
	user.Delete("/Avatar", middlewares.Protected, controllers.RemoveAvatar)
}
