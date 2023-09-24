package controllers

import (
	"musematch/app/utils"

	"github.com/gofiber/fiber/v2"
)

func GetImageController(c *fiber.Ctx) error {
	imageId := c.Query("id")

	imageUrl, err := utils.PresignedGetUrl(imageId)
	if err != nil {
		return err
	}

	return c.Redirect(imageUrl)
}