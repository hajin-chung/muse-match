package controllers

import (
	"encoding/json"
	"musematch/app/globals"
	"musematch/app/models"
	"musematch/app/queries"
	"musematch/app/utils"

	"github.com/gofiber/fiber/v2"
)

func ExhibitListController(c *fiber.Ctx) error {
	sess, _ := globals.Store.Get(c)

	currentUserId := sess.Get("id")
	exhibits, err := queries.GetExhibits()
	if err != nil {
		return err
	}

	return c.Render("pages/exhibit/list", fiber.Map{
		"Title":    "상시전시",
		"Theme":    c.Locals("theme"),
		"UserId":   currentUserId,
		"Exhibits": exhibits,
	}, "layout")
}

func ExhibitController(c *fiber.Ctx) error {
	sess, _ := globals.Store.Get(c)

	currentUserId := sess.Get("id")
	exhibitId := c.Params("exhibitId")
	exhibit, err := queries.GetExhibitById(exhibitId)
	if err != nil {
		return err
	}
	arts, err := queries.GetArtsByExhibitId(exhibitId)
	if err != nil {
		return err
	}

	return c.Render("pages/exhibit/index", fiber.Map{
		"Title":   exhibit.Title,
		"Theme":   c.Locals("theme"),
		"UserId":  currentUserId,
		"Exhibit": exhibit,
		"Arts":    arts,
	}, "layout")
}

func CreateExhibitController(c *fiber.Ctx) error {
	newExhibit := models.NewExhibitInfo{}
	json.Unmarshal(c.Body(), &newExhibit)
	exhibitId, err := queries.CreateExhibit(newExhibit)
	if err != nil {
		return err
	}

	uploadUrl, err := utils.PresignedPutUrl(exhibitId)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"url": uploadUrl,
	})
}

func UpdateExhibitController(c *fiber.Ctx) error {
	newExhibit := models.Exhibit{}
	json.Unmarshal(c.Body(), &newExhibit)
	err := queries.UpdateExhibit(newExhibit)
	if err != nil {
		return err
	}

	uploadUrl, err := utils.PresignedPutUrl(newExhibit.Id)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"url": uploadUrl,
	})
}
