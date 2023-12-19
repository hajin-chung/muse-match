package routes

import (
	"musematch/controllers"
	"musematch/middleware"

	"github.com/gofiber/fiber/v2"
)

func PrivateRoutes(a *fiber.App) {
	a.Get("/art/new", middleware.SessionProtected, controllers.NewArtViewController)
	a.Get("/art/:artId/update", middleware.SessionProtected, controllers.UpdateArtViewController)
	a.Post("/art", middleware.SessionProtected, controllers.NewArtController)
	a.Post("/art/:artId", middleware.SessionProtected, controllers.UpdateArtController)
	a.Put("/image", middleware.SessionProtected, controllers.PutImageController)
	a.Get("/me/profile", middleware.SessionProtected, controllers.ProfileUpdateViewController)
	a.Post("/api/me/profile", middleware.SessionProtected, controllers.ProfileUpdateController)
	a.Post("/admin/auth", middleware.SessionProtected, controllers.AdminAuthController)
	a.Get("/admin/auth", middleware.SessionProtected, controllers.AdminAuthViewController)
}
