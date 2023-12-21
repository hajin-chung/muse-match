package middleware

import (
	"fmt"
	"musematch/globals"
	"musematch/utils"

	"github.com/gofiber/fiber/v2"
)

func Logger(c *fiber.Ctx) error {
	reqMsg := fmt.Sprintf("REQ %15s %s %s ", c.IP(), c.Method(), c.OriginalURL())
	for key, value := range c.GetReqHeaders() {
		reqMsg += fmt.Sprintf("%s:%s;", key, value)
	}
	reqMsg += fmt.Sprintf("%s", string(c.BodyRaw()))
	utils.Log(reqMsg)

	res := c.Next()

	resMsg := fmt.Sprintf("RES %15s %s %s ", c.IP(), c.Method(), c.OriginalURL())
	for key, value := range c.GetRespHeaders() {
		resMsg += fmt.Sprintf("%s:%s;", key, value)
	}
	utils.Log(resMsg)
	return res
}

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

func ContentTypeHtml(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
	return c.Next()
}

func ContentTypeAuto(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentDPR, "image/svg+xml") 
	return c.Next()
}
