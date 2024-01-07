package controllers

import (
	"encoding/json"
	"log"
	"musematch/globals"
	"musematch/models"
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
	art, err := queries.GetArtById(artId)
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
	err := queries.ArtDelete(artId)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"success": true})
}

func DashboardProfileViewController(c *fiber.Ctx) error {
	sess, _ := globals.Store.Get(c)
	id := sess.Get("id").(string)

	profile, err := queries.GetUserProfile(id)
	if err != nil {
		return err
	}

	page := pages.DashboardProfilePage("sample title", profile)
	return utils.Render(c, page)
}

type HistoryPayload struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type ListPayload struct {
	Title  string   `json:"title"`
	ArtIds []string `json:"artIds"`
}

type ProfileUpdatePayload struct {
	Name        string           `json:"name"`
	Description string           `json:"description"`
	InstagramId string           `json:"instagramId"`
	FacebookId  string           `json:"facebookId"`
	TwitterId   string           `json:"twitterId"`
	Links       []string         `json:"links"`
	Note        string           `json:"note"`
	History     []HistoryPayload `json:"history"`
	List        []ListPayload    `json:"list"`
}

func DashboardProfileController(c *fiber.Ctx) error {
	sess, _ := globals.Store.Get(c)
	id := sess.Get("id").(string)

	body := c.Body()
	payload := ProfileUpdatePayload{}

	err := json.Unmarshal(body, &payload)
	if err != nil {
		return err
	}

	err = queries.UpdateUser(
		id, payload.Name, payload.Description,
		payload.InstagramId, payload.FacebookId, payload.TwitterId,
		payload.Note,
	)
	if err != nil {
		return err
	}

	links := []models.UserLink{}
	for _, content := range payload.Links {
		links = append(links, models.UserLink{
			Id:      utils.CreateId(),
			UserId:  id,
			Content: content,
		})
	}
	err = queries.UpdateUserLink(id, links)
	if err != nil {
		return err
	}

	histories := []models.UserHistory{}
	for _, history := range payload.History {
		histories = append(histories, models.UserHistory{
			Id:      utils.CreateId(),
			UserId:  id,
			Title:   history.Title,
			Content: history.Content,
		})
	}

	err = queries.UpdateUserHistory(id, histories)
	if err != nil {
		return err
	}

	artLists := []models.UserArtList{}
	artListItems := []models.UserArtListItem{}
	for _, artList := range payload.List {
		artListId := utils.CreateId()
		artLists = append(artLists, models.UserArtList{
			Id:     artListId,
			UserId: id,
			Title:  artList.Title,
		})
		for idx, artId := range artList.ArtIds {
			artListItems = append(artListItems, models.UserArtListItem{
				ListId: artListId,
				ArtId:  artId,
				Idx:    idx,
			})
		}
	}

	if len(artLists) > 0 {
		err = queries.UpdateUserArtList(id, artLists)
		if err != nil {
			return err
		}
	}

	if len(artListItems) > 0 {
		err = queries.UpdateUserArtListItem(id, artListItems)
		if err != nil {
			return err
		}
	}

	bannerUrl, err := utils.PresignedPutUrl("banner-" + id)
	pictureUrl, err := utils.PresignedPutUrl(id)
	return c.JSON(fiber.Map{"bannerUrl": bannerUrl, "pictureUrl": pictureUrl})
}

func DashboardPlacesController(c *fiber.Ctx) error {
	sess, _ := globals.Store.Get(c)
	id := sess.Get("id").(string)

	user, err := queries.GetUserById(id)
	if err != nil {
		return err
	}

	placeInfos, err := queries.GetPlaceInfosByUserId(id)
	if err != nil {
		return err
	}

	page := pages.DashboardPlacesPage("대시보드 - 작품", user, placeInfos)
	return utils.Render(c, page)
}

func DashboardPlaceCreateViewController(c *fiber.Ctx) error {
	sess, _ := globals.Store.Get(c)
	id := sess.Get("id").(string)

	user, err := queries.GetUserById(id)
	if err != nil {
		return err
	}

	page := pages.DashboardPlaceNewPage("대시보드 - 작품", user)
	return utils.Render(c, page)
}

type LocationPayload struct {
	Title       string `josn:"title"`
	Description string `josn:"description"`
}

type PlaceMutatePayload struct {
	Title       string            `json:"title"`
	Address     string            `json:"address"`
	InstagramId string            `json:"instagramId"`
	FacebookId  string            `json:"facebookId"`
	TwitterId   string            `json:"twitterId"`
	Links       []string          `json:"links"`
	ImageCount  int               `json:"imageCount"`
	Locations   []LocationPayload `json:"locations"`
}

func DashboardPlaceCreateController(c *fiber.Ctx) error {
	sess, _ := globals.Store.Get(c)
	userId := sess.Get("id").(string)

	body := c.Body()
	payload := PlaceMutatePayload{}
	err := json.Unmarshal(body, &payload)
	if err != nil {
		return err
	}

	placeId := utils.CreateId()
	err = queries.CreatePlace(
		placeId, userId, payload.Title, payload.Address,
		payload.InstagramId, payload.FacebookId, payload.TwitterId)
	if err != nil {
		return err
	}

	err = queries.CreatePlaceLinks(placeId, payload.Links)
	if err != nil {
		return err
	}

	imageIds := []string{}
	for i := 0; i < payload.ImageCount; i++ {
		imageId := utils.CreateId()
		imageIds = append(imageIds, imageId)
	}

	err = queries.CreatePlaceImages(placeId, imageIds)
	if err != nil {
		return err
	}

	locations := []models.PlaceLocation{}
	for _, location := range payload.Locations {
		locations = append(locations, models.PlaceLocation{
			Id:          utils.CreateId(),
			PlaceId:     placeId,
			Title:       location.Title,
			Description: location.Description,
		})
	}
	err = queries.CreatePlaceLocations(locations)
	if err != nil {
		return err
	}

	imageUrls := []string{}
	for _, imageId := range imageIds {
		url, err := utils.PresignedPutUrl(imageId)
		if err != nil {
			return err
		}
		imageUrls = append(imageUrls, url)
	}

	locationImageUrls := []string{}
	for _, location := range locations {
		url, err := utils.PresignedPutUrl(location.Id)
		if err != nil {
			return err
		}
		locationImageUrls = append(locationImageUrls, url)
	}

	return c.JSON(fiber.Map{
		"imageUrls":         imageUrls,
		"locationImageUrls": locationImageUrls,
	})
}

func DashboardPlaceUpdateViewController(c *fiber.Ctx) error {
	sess, _ := globals.Store.Get(c)
	id := sess.Get("id").(string)

	user, err := queries.GetUserById(id)
	if err != nil {
		return err
	}

	page := pages.DashboardPlaceNewPage("대시보드 - 작품", user)
	return utils.Render(c, page)
}
