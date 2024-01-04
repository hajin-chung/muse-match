package controllers

import (
	"encoding/json"
	"log"
	"musematch/globals"
	"musematch/queries"
	"musematch/utils"
	"musematch/views/pages"

	"github.com/gofiber/fiber/v2"
)

func DashboardArtController(c *fiber.Ctx) error {
	sess, _ := globals.Store.Get(c)
	id := sess.Get("id").(string)

	user, err := queries.GetUserById(id)
	if err != nil {
		return err
	}

	arts, err := queries.GetArtInfosByUserId(id)
	if err != nil {
		return err
	}

	page := pages.DashboardArtPage("대시보드 - 작품", user, arts)
	return utils.Render(c, page)
}

func DashboardArtCreateViewController(c *fiber.Ctx) error {
	sess, _ := globals.Store.Get(c)
	id := sess.Get("id").(string)

	user, err := queries.GetUserById(id)
	if err != nil {
		return err
	}

	page := pages.DashboardArtNewPage("대시보드 - 작품 등록하기", user)
	return utils.Render(c, page)
}

type ArtMutatePayload struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Price       int      `json:"price"`
	Info        string   `json:"info"`
	ImageLength int      `json:"imageLength"`
	Tags        []string `json:"tags"`
}

func DashboardArtCreateController(c *fiber.Ctx) error {
	sess, _ := globals.Store.Get(c)
	id := sess.Get("id").(string)

	body := c.Body()
	payload := ArtMutatePayload{}

	err := json.Unmarshal(body, &payload)
	if err != nil {
		return err
	}

	artId := utils.CreateId()

	err = queries.ArtCreate(
		id,
		artId,
		payload.Name,
		payload.Description,
		payload.Price,
		payload.Info,
	)
	if err != nil {
		return err
	}

	err = queries.ArtTagsCreate(artId, payload.Tags)
	if err != nil {
		return err
	}

	imageIds := []string{}
	for i := 0; i < payload.ImageLength; i++ {
		imageId := utils.CreateId()
		imageIds = append(imageIds, imageId)
	}

	err = queries.ArtImagesCreate(artId, imageIds)
	if err != nil {
		log.Println(err)
	}

	uploadUrls := []string{}
	for _, imageId := range imageIds {
		url, err := utils.PresignedPutUrl(imageId)
		if err != nil {
			log.Println(err)
			continue
		}
		uploadUrls = append(uploadUrls, url)
	}

	return c.JSON(fiber.Map{"uploadUrls": uploadUrls})
}

func DashboardArtUpdateViewController(c *fiber.Ctx) error {
	sess, _ := globals.Store.Get(c)
	id := sess.Get("id").(string)
	artId := c.Params("id")

	user, err := queries.GetUserById(id)
	if err != nil {
		return err
	}
	art, err := queries.ArtGetById(artId)
	if err != nil {
		return err
	}
	tags, err := queries.ArtTagsGetById(artId)
	if err != nil {
		return err
	}
	imageIds, err := queries.ArtImageIdsGetById(artId)
	if err != nil {
		return err
	}

	page := pages.DashboardArtUpdatePage("대시보드 - 작품 수정하기", user, art, tags, imageIds)
	return utils.Render(c, page)
}

func DashboardArtUpdateController(c *fiber.Ctx) error {
	artId := c.Params("id")

	body := c.Body()
	payload := ArtMutatePayload{}

	err := json.Unmarshal(body, &payload)
	if err != nil {
		return err
	}

	err = queries.ArtUpdate(
		artId,
		payload.Name,
		payload.Description,
		payload.Price,
		payload.Info,
	)
	if err != nil {
		return err
	}

	err = queries.ArtTagsUpdate(artId, payload.Tags)
	if err != nil {
		return err
	}

	err = queries.ArtImagesDelete(artId)
	if err != nil {
		return err
	}

	imageIds := []string{}
	for i := 0; i < payload.ImageLength; i++ {
		imageId := utils.CreateId()
		imageIds = append(imageIds, imageId)
	}

	err = queries.ArtImagesCreate(artId, imageIds)
	if err != nil {
		log.Println(err)
	}

	uploadUrls := []string{}
	for _, imageId := range imageIds {
		url, err := utils.PresignedPutUrl(imageId)
		if err != nil {
			log.Println(err)
			continue
		}
		uploadUrls = append(uploadUrls, url)
	}

	return c.JSON(fiber.Map{"uploadUrls": uploadUrls})
}

func DashboardArtDeleteController(c *fiber.Ctx) error {
	artId := c.Params("id")
	log.Println(artId)
	err := queries.ArtDelete(artId)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"success": true})
}
