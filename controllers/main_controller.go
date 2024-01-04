package controllers

import (
	"errors"
	"log"
	"musematch/globals"
	"musematch/models"
	"musematch/queries"
	"musematch/utils"
	"musematch/views/pages"

	"github.com/gofiber/fiber/v2"
)

func MainController(c *fiber.Ctx) error {
	// TODO: handle error
	sess, _ := globals.Store.Get(c)

	var user *models.User
	tmpUserId := sess.Get("id")
	userId, ok := tmpUserId.(string)
	if ok {
		_user, err := queries.GetUserById(userId)
		user = _user
		if err != nil {
			return err
		}
	}

	page := pages.Main("this is title", user)
	return utils.Render(c, page)
}

func ErrorController(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	log.Println(err, code)

	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	// Send custom error page
	return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")

}
