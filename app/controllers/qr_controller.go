package controllers

import (
	"bytes"
	"musematch/app/utils"

	"github.com/gofiber/fiber/v2"
)

func GetQrCode(c *fiber.Ctx) error {
	url := c.Query("url", "https://musematch.app")
	qr, err := utils.GetQrCode(url)
	if err != nil {
		return err
	}

	reader := bytes.NewReader(qr)
	c.Response().Header.Add("Content-Type", "image/png")
	return c.SendStream(reader)
}
