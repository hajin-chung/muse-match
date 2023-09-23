package controllers

import (
	"musematch/app/globals"
	"musematch/app/queries"

	"github.com/gofiber/fiber/v2"
)

func ArtController(c *fiber.Ctx) error {
	sess, _ := globals.Store.Get(c)

	currentUserId := sess.Get("id")
	userId := c.Params("userId")
	artId := c.Params("artId")
	user, err := queries.GetUserById(userId)
	if err != nil {
		return err
	}
	art, err := queries.GetArtById(artId)
	if err != nil {
		return err
	}

	return c.Render("pages/art", fiber.Map{
		"Title": art.Name,
		"Header": fiber.Map{
			"ArtistId": currentUserId,
		},
		"IsMyArt": currentUserId == userId,
		"Artist":  user,
		"Art":     art,
	}, "layout")
}
