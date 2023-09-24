package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"musematch/app/globals"
	"musematch/app/models"
	"musematch/app/queries"
	"musematch/app/utils"

	"github.com/gofiber/fiber/v2"
)

func ProfileController(c *fiber.Ctx) error {
	// TODO: handle error
	sess, _ := globals.Store.Get(c)

	currentUserId := sess.Get("id")
	userId := c.Params("userId")
	user, err := queries.GetUserById(userId)
	if err != nil {
		return err
	}
	arts, err := queries.GetArtsByUserId(userId)
	if err != nil {
		return err
	}
	log.Printf("%+v\n", arts)

	return c.Render("pages/profile/index", fiber.Map{
		"Title": user.Name,
		"Theme": c.Locals("theme"),
		"Header": fiber.Map{
			"ArtistId": currentUserId,
		},
		"IsMyProfile": userId == currentUserId,
		"User":        user,
		"Arts":        arts,
	}, "layout")
}

func ProfileUpdateViewController(c *fiber.Ctx) error {
	sess, _ := globals.Store.Get(c)

	currentUserId := sess.Get("id").(string)

	user, err := queries.GetUserById(currentUserId)
	if err != nil {
		return err
	}

	return c.Render("pages/profile/update", fiber.Map{
		"Title":  "프로필 수정하기",
		"Theme":  c.Locals("theme"),
		"Header": fiber.Map{
			// "ArtistId": currentUserId,
		},
		"User": user,
	}, "layout")
}

func ProfileUpdateController(c *fiber.Ctx) error {
	sess, _ := globals.Store.Get(c)

	currentUserId := sess.Get("id").(string)

	userInfo := models.UpdateUserInfo{}
	_ = json.Unmarshal(c.Body(), &userInfo)
	err := queries.UpdateUser(currentUserId, &userInfo)
	if err != nil {
		return err
	}

	pictureUrl, _ := utils.PresignedPutUrl(currentUserId)
	headerUrl, _ := utils.PresignedPutUrl(fmt.Sprintf("header-%s", currentUserId))
	return c.JSON(fiber.Map{
		"pictureUrl": pictureUrl,
		"headerUrl":  headerUrl,
	})
}
