package routes

import (
	"musematch/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(a *fiber.App) {
	a.Get("/", controllers.IndexController)
}
