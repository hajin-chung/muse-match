package controllers

import (
	"fmt"
	"musematch/globals"
	"musematch/utils"

	"github.com/gofiber/fiber/v2"
)

func ErrorController(c *fiber.Ctx, err error) error {
	sess, _ := globals.Store.Get(c)

	currentUserId := sess.Get("id")

	userId, ok := currentUserId.(string)
	if !ok {
		userId = "UNDEFINED"
	}

	errorMessage := fmt.Sprintf("Path: %s\nError: %s\nUserId: %s", c.OriginalURL(), err.Error(), userId)
	go utils.SlackSendMessage(errorMessage)

	return c.Status(500).Render("pages/500", fiber.Map{
		"Theme":   c.Locals("theme"),
		"UserId":  currentUserId,
		"Message": err.Error(),
	}, "layout")
}

func NotFoundController(c *fiber.Ctx) error {
	sess, _ := globals.Store.Get(c)

	currentUserId := sess.Get("id")

	return c.Status(404).Render("pages/404", fiber.Map{
		"Theme":  c.Locals("theme"),
		"UserId": currentUserId,
	}, "layout")
}
