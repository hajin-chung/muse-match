package controllers

import (
	"musematch/models"
	"musematch/utils"
	"musematch/views/pages"

	"github.com/gofiber/fiber/v2"
)

func TestController(c *fiber.Ctx) error {
	page := pages.DashboardArtNewPage("대시보드 - 작품 등록하기", &models.User{})
	return utils.Render(c, page)
}
