package routes

import (
	"musematch/app/controllers"
	"musematch/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func PrivateRoutes(a *fiber.App) {
	a.Get("/art/new", middleware.SessionProtected, controllers.NewArtViewController)
	a.Post("/api/art", middleware.SessionProtected, controllers.NewArtController)
}
