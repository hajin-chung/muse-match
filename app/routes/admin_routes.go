package routes

import (
	"musematch/app/controllers"
	"musematch/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func AdminRoutes(a *fiber.App) {
	a.Get("/admin", middleware.AdminProtected, controllers.AdminController)
	a.Get("/admin/as/:userId", middleware.AdminProtected, controllers.AsController)
	a.Post("/exhibit", middleware.AdminProtected, controllers.CreateExhibitController)
	a.Post("/exhibit/:exhibitId", middleware.AdminProtected, controllers.UpdateExhibitController)
}
