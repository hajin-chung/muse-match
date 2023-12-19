package controllers

import (
	"encoding/json"
	"log"
	"musematch/globals"
	"musematch/models"
	"musematch/queries"
	"musematch/utils"

	"github.com/gofiber/fiber/v2"
)

func ArtController(c *fiber.Ctx) error {
	sess, _ := globals.Store.Get(c)

	currentUserId := sess.Get("id")
	userId := c.Params("userId")
	artId := c.Params("artId")
	user, err := queries.GetUserById(userId)
	if err != nil {
		return err
	}
	art, err := queries.GetArtById(artId)
	if err != nil {
		return err
	}
	images, err := queries.GetImagesByArtId(artId)
	if err != nil {
		return err
	}

	return c.Render("pages/art/index", fiber.Map{
		"Title":   art.Name,
		"Theme":   c.Locals("theme"),
		"UserId":  currentUserId,
		"IsMyArt": currentUserId == userId,
		"Artist":  user,
		"Art":     art,
		"Images":  images,
	}, "layout")
}

func NewArtViewController(c *fiber.Ctx) error {
	sess, _ := globals.Store.Get(c)

	currentUserId := sess.Get("id")

	return c.Render("pages/art/new", fiber.Map{
		"Theme":  c.Locals("theme"),
		"UserId": currentUserId,
	}, "layout")
}

func NewArtController(c *fiber.Ctx) error {
	sess, _ := globals.Store.Get(c)

	userId := sess.Get("id").(string)
	newArt := models.NewArtInfo{}
	_ = json.Unmarshal(c.Body(), &newArt)
	artId, err := queries.CreateArt(newArt, userId)
	if err != nil {
		return err
	}

	uploadUrls := []string{}
	for i := 0; i < newArt.ImageCount; i++ {
		imageId := utils.CreateId()
		queries.CreateImage(imageId, artId, i)
		url, err := utils.PresignedPutUrl(imageId)
		if err != nil {
			log.Println(err)
			continue
		}
		uploadUrls = append(uploadUrls, url)
	}

	return c.JSON(fiber.Map{
		"success":    true,
		"userId":     userId,
		"artId":      artId,
		"uploadUrls": uploadUrls,
	})
}

func UpdateArtViewController(c *fiber.Ctx) error {
	sess, _ := globals.Store.Get(c)

	currentUserId := sess.Get("id")
	artId := c.Params("artId")
	art, err := queries.GetArtById(artId)
	if err != nil {
		return err
	}
	images, err := queries.GetImagesByArtId(artId)
	if err != nil {
		return err
	}

	return c.Render("pages/art/update", fiber.Map{
		"Theme":  c.Locals("theme"),
		"UserId": currentUserId,
		"Art":    art,
		"Images": images,
	}, "layout")
}

func UpdateArtController(c *fiber.Ctx) error {
	sess, _ := globals.Store.Get(c)

	userId := sess.Get("id").(string)
	artId := c.Params("artId")
	newArt := models.NewArtInfo{}
	err := json.Unmarshal(c.Body(), &newArt)
	if err != nil {
		return err
	}

	err = queries.UpdateArt(artId, newArt, userId)
	if err != nil {
		return err
	}

	err = queries.DeleteImageByArtId(artId)
	if err != nil {
		return err
	}
	uploadUrls := []string{}
	for i := 0; i < newArt.ImageCount; i++ {
		imageId := utils.CreateId()
		queries.CreateImage(imageId, artId, i)
		url, err := utils.PresignedPutUrl(imageId)
		if err != nil {
			log.Println(err)
			continue
		}
		uploadUrls = append(uploadUrls, url)
	}

	return c.JSON(fiber.Map{
		"success":    true,
		"userId":     userId,
		"artId":      artId,
		"uploadUrls": uploadUrls,
	})
}
