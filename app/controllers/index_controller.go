package controllers

import (
	"log"
	"musematch/app/globals"

	"github.com/gofiber/fiber/v2"
)

func IndexController(c *fiber.Ctx) error {
	// TODO: handle error
	sess, _ := globals.Store.Get(c)
	log.Printf("%s\n", sess.Get("id"))

	return c.Render("pages/index", fiber.Map{
		"Title": "Muse Match",
		"Header": fiber.Map{
			"ArtistId": sess.Get("id"),
		},
	}, "layout")
}
