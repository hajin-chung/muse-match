package routes

import (
	"musematch/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(a *fiber.App) {
	a.Get("/", controllers.ExhibitListController)
	a.Get("/exhibit/:exhibitId", controllers.ExhibitController)
	a.Get("/auth/login", controllers.LoginController)
	a.Get("/auth/logout", controllers.LogoutController)
	a.Get("/api/auth/callback/kakao", controllers.KakaoCallbackController)
	a.Get("/art/:userId", controllers.ProfileController)
	a.Get("/art/:userId/:artId", controllers.ArtController)
	a.Get("/image", controllers.GetImageController)
	a.Get("/qr", controllers.GetQrCode)
}
