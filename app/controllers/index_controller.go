package controllers

import (
	"musematch/app/globals"

	"github.com/gofiber/fiber/v2"
)

func IndexController(c *fiber.Ctx) error {
	// TODO: handle error
	sess, _ := globals.Store.Get(c)

	return c.Render("pages/index", fiber.Map{
		"Title": "Muse Match",
		"Header": fiber.Map{
			"ArtistId": sess.Get("id"),
		},
	}, "layout")
}
