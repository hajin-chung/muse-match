package controllers

import (
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

	return c.Render("pages/exhibit/list", fiber.Map{
		"Title":    "상시전시",
		"Theme":    c.Locals("theme"),
		"ArtistId": currentUserId,
		"Exhibits": exhibits,
	}, "layout")
}

func ExhibitController(c *fiber.Ctx) error {
	sess, _ := globals.Store.Get(c)

	currentUserId := sess.Get("id")
	exhibitId := c.Params("exhibitId")
	exhibit, err := queries.GetExhibitById(exhibitId)
	if err != nil {
		return err
	}
	arts, err := queries.GetArtsByExhibitId(exhibitId)
	if err != nil {
		return err
	}

	return c.Render("pages/exhibit/index", fiber.Map{
		"Title":    exhibit.Title,
		"Theme":    c.Locals("theme"),
		"ArtistId": currentUserId,
		"Arts":     arts,
	}, "layout")
}
