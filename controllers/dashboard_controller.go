package controllers

import (
	"musematch/globals"
	"musematch/queries"
	"musematch/utils"
	"musematch/views/pages"

	"github.com/gofiber/fiber/v2"
)

func DashboardArtController(c *fiber.Ctx) error {
	sess, _ := globals.Store.Get(c)
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

func DashboardArtCreateViewController(c *fiber.Ctx) error {
	sess, _ := globals.Store.Get(c)
	id := sess.Get("id").(string)

	user, err := queries.GetUserById(id)
	if err != nil {
		return err
	}

	page := pages.DashboardArtNewPage("대시보드 - 작품 등록하기", user)
	return utils.Render(c, page)
}

func DashboardArtCreateController(c *fiber.Ctx) error {
	return nil
}
