package routes

import (
	"musematch/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(a *fiber.App) {
	a.Get("/", controllers.IndexController)
	a.Get("/auth/login", controllers.LoginController)
	a.Get("/auth/logout", controllers.LogoutController)
	a.Get("/api/auth/callback/kakao", controllers.KakaoCallbackController)
}
