package controllers

import (
	// "musematch/globals"
	"musematch/views"
	"musematch/utils"

	"github.com/gofiber/fiber/v2"
)

func IndexController(c *fiber.Ctx) error {
	// TODO: handle error
	// _sess, _ := globals.Store.Get(c)

	testComponent := views.Test("test")
	return utils.Render(c, testComponent)
}
