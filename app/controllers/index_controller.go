package controllers

import "github.com/gofiber/fiber/v2"

func IndexController(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title":  "214554",
		"Header": fiber.Map{
			// "ArtistId": "123",
		},
	}, "layout")
}
