package controllers

import (
	"log"
	"musematch/app/globals"
	"musematch/app/queries"

	"github.com/gofiber/fiber/v2"
)

func ProfileController(c *fiber.Ctx) error {
	// TODO: handle error
	sess, _ := globals.Store.Get(c)

	currentUserId := sess.Get("id")
	userId := c.Params("userId")
	user, err := queries.GetUserById(userId)
	if err != nil {
		return err
	}
	arts, err := queries.GetArtsByUserId(userId)
	if err != nil {
		return err
	}
	log.Printf("%+v\n", arts)

	return c.Render("pages/profile", fiber.Map{
		"Title": user.Name,
		"Header": fiber.Map{
			"ArtistId": currentUserId,
		},
		"IsMyProfile": userId == currentUserId,
		"User":        user,
		"Arts":        arts,
	}, "layout")
}
