package controllers

import (
	"encoding/json"
	"fmt"
	"musematch/globals"
	"musematch/queries"

	"github.com/gofiber/fiber/v2"
)

func AdminController(c *fiber.Ctx) error {
	sess, _ := globals.Store.Get(c)

	currentUserId := sess.Get("id")
	exhibits, err := queries.GetExhibits()
	if err != nil {
		return err
	}
	users, err := queries.GetUsers()
	if err != nil {
		return err
	}

	return c.Render("pages/admin/index", fiber.Map{
		"Theme":    c.Locals("theme"),
		"UserId":   currentUserId,
		"Exhibits": exhibits,
		"Users":    users,
	}, "layout")
}

type AdminSecret struct {
	Secret string `json:"secret"`
}

func AdminAuthController(c *fiber.Ctx) error {
	adminSecret := AdminSecret{}
	json.Unmarshal(c.Body(), &adminSecret)

	sess, _ := globals.Store.Get(c)
	sess.Set("secret", adminSecret.Secret)
	sess.Save()

	return c.JSON(fiber.Map{
		"success": true,
	})
}

func AdminAuthViewController(c *fiber.Ctx) error {
	sess, _ := globals.Store.Get(c)

	currentUserId := sess.Get("id")

	return c.Render("pages/admin/auth", fiber.Map{
		"Theme":  c.Locals("theme"),
		"UserId": currentUserId,
	}, "layout")
}

func AsController(c *fiber.Ctx) error {
	sess, _ := globals.Store.Get(c)

	userId := c.Params("userId")
	sess.Set("id", userId)
	sess.Save()

	return c.Redirect(fmt.Sprintf("/art/%s", userId))
}
