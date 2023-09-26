package routes

import (
	"musematch/app/controllers"
	"musematch/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func PrivateRoutes(a *fiber.App) {
	a.Get("/art/new", middleware.SessionProtected, controllers.NewArtViewController)
	a.Post("/api/art", middleware.SessionProtected, controllers.NewArtController)
	a.Put("/image", middleware.SessionProtected, controllers.PutImageController)
	a.Get("/me/profile", middleware.SessionProtected, controllers.ProfileUpdateViewController)
	a.Post("/api/me/profile", middleware.SessionProtected, controllers.ProfileUpdateController)
	a.Post("/admin/auth", middleware.SessionProtected, controllers.AdminAuthController)
	a.Get("/admin/auth", middleware.SessionProtected, controllers.AdminAuthViewController)
}
