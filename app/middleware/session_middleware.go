package middleware

import (
	"musematch/app/globals"

	"github.com/gofiber/fiber/v2"
)

func SessionProtected(c *fiber.Ctx) error {
	sess, _ := globals.Store.Get(c)

	userId := sess.Get("id")
	if userId == nil {
		return c.Status(500).Redirect("/auth/login")
	}

	return c.Next()
}

func AdminProtected(c *fiber.Ctx) error {
	sess, _ := globals.Store.Get(c)

	userId := sess.Get("id")
	if userId == nil {
		return c.Status(500).Redirect("/auth/login")
	}
	secret := sess.Get("secret")
	if secret != globals.Env.ADMIN {
		return c.Status(500).Redirect("/admin/auth")
	}

	return c.Next()
}
