package controllers

import (
	"musematch/utils"

	"github.com/gofiber/fiber/v2"
)

func ImageGetController(c *fiber.Ctx) error {
	imageId := c.Query("id")

	imageUrl, err := utils.PresignedGetUrl(imageId)
	if err != nil {
		return err
	}

	return c.Redirect(imageUrl)
}
