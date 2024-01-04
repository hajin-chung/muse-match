package controllers

import (
	"musematch/models"
	"musematch/utils"
	"musematch/views/pages"

	"github.com/gofiber/fiber/v2"
)

func TestController(c *fiber.Ctx) error {
	profile := models.UserProfile {
		User: &models.User{},
		Link: []models.UserLink{},
		History: []models.UserHistory{},
		ArtList: &models.UserArtListMap{},
		Arts: models.UserArtMap{},
	}
	page := pages.DashboardProfilePage("대시보드 - 작품 등록하기", &profile)
	return utils.Render(c, page)
}
