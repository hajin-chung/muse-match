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

func ArtController(c *fiber.Ctx) error {
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

	artId := c.Params("art_id")
	art, err := queries.GetArtById(artId)
	if err != nil {
		return err
	}
	artist, err := queries.GetUserById(art.UserId)
	if err != nil {
		return err
	}
	tags, err := queries.GetArtTagsById(artId)
	if err != nil {
		return err
	}
	imageIds, err := queries.GetArtImagesById(artId)
	if err != nil {
		return err
	}

	exhibitInfo, place, err := queries.GetArtExhibitInfoById(artId)
	if err != nil {
		return err
	}

	page := pages.ArtPage("sample title", user, art, artist, tags, imageIds, exhibitInfo, place)
	return utils.Render(c, page)
}

func PlaceController(c *fiber.Ctx) error {
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
	
	placeId := c.Params("place_id")
	place, err := queries.GetPlaceById(placeId)
	if err != nil {
		return err
	}

	images, err := queries.GetPlaceImagesById(placeId)
	if err != nil {
		return err
	}

	links, err := queries.GetPlaceLinksById(placeId)
	if err != nil {
		return err
	}

	locations, err := queries.GetPlaceLocationsById(placeId)
	if err != nil {
		return err
	}

	arts ,err := queries.GetPlaceArtsById(placeId)
	if err != nil {
		return err
	}

	page := pages.PlacePage("title", user, place, images, links, locations, arts)
	return utils.Render(c, page)
}

func ArtsController(c *fiber.Ctx) error {
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

	arts, err := queries.GetArtInfos()
	if err != nil {
		return err
	}

	page := pages.ArtsPage("title", user, arts)
	return utils.Render(c, page)
}

func ArtistsController(c *fiber.Ctx) error {
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

	artists, err := queries.GetUserInfos()
	if err != nil {
		return err
	}

	page := pages.ArtistsPage("title", user, artists)
	return utils.Render(c, page)
}
