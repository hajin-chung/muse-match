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

func ArtMapToList(artMap models.UserArtMap) []models.ArtInfo {
	list := []models.ArtInfo{}
	for _, artInfo := range artMap {
		list = append(list, artInfo)
	}
	return list
}

func ArtistController(c *fiber.Ctx) error {
	sess, _ := globals.Store.Get(c)

	var user *models.User
	userId, ok := sess.Get("id").(string)
	if ok {
		_user, err := queries.GetUserById(userId)
		user = _user
		if err != nil {
			return err
		}
	}

	artistId := c.Params("user_id")
	profile, err := queries.GetUserProfile(artistId)
	if err != nil {
		return err
	}
	artGrid := ArtMapToList(profile.Arts)

	page := pages.ArtistPage("sample title", user, profile, artGrid)
	return utils.Render(c, page)
}
