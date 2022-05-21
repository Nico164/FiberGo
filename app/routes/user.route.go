package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/Nico164/FiberGo/app/controllers"
)

func SetupUserRoutes(router fiber.Router) {
	user := router.Group("/user")

	user.Get("/", controllers.GetUser)
	user.Post("/login", controllers.LoginUser)
	user.Post("/register", controllers.RegisterUser)
	user.Post("/logout", controllers.LogoutUser)
}
