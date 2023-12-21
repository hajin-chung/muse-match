package controllers

import (
	"log"
	"musematch/globals"
	"musematch/queries"
	"musematch/utils"
	"musematch/views/pages"

	"github.com/gofiber/fiber/v2"
)

func DashboardArtController(c *fiber.Ctx) error {
	sess, err := globals.Store.Get(c)
	if err != nil {
		log.Println(err)
		return err
	}
	id := sess.Get("id").(string)

	user, err := queries.GetUserById(id)
	if err != nil {
		return err
	}

	arts, err := queries.GetArtInfosByUserId(id)
	if err != nil {
		return err
	}

	page := pages.DashboardArtPage("대시보드 - 작품", user, arts)
	return utils.Render(c, page)
}
