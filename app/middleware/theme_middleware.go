package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func ThemeFromCookie(c *fiber.Ctx) error {
	theme := c.Cookies("theme")
	if theme == "" {
		themeCookie := new(fiber.Cookie)
		themeCookie.Name = "theme"
		themeCookie.Value = "light"
		theme = "light"
		c.Cookie(themeCookie)
	}

	c.Locals("theme", theme)
	return c.Next()
}
