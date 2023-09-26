package controllers

import (
	"log"
	"musematch/app/globals"
	"musematch/app/queries"

	"github.com/gofiber/fiber/v2"
)

func ExhibitListController(c *fiber.Ctx) error {
	sess, _ := globals.Store.Get(c)

	currentUserId := sess.Get("id")
	exhibits, err := queries.GetExhibits()
	if err != nil {
		return err
	}
	log.Println(exhibits)

	return c.Render("pages/exhibit/list", fiber.Map{
		"Title":    "상시전시",
		"Theme":    c.Locals("theme"),
		"ArtistId": currentUserId,
		"Exhibits": exhibits,
	}, "layout")
}
