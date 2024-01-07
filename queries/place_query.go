package queries

import (
	"musematch/models"
	"musematch/utils"
)

func GetPlaceInfosByUserId(userId string) ([]models.PlaceInfo, error) {
	places := []models.PlaceInfo{}
	err := db.Select(
		&places,
		`SELECT 
			place.id, place.title, place.address, place_image.id as thumbnail
		FROM place LEFT JOIN place_image ON place.id = place_image.place_id 
		WHERE place.user_id=?`,
		userId,
	)
	if err != nil {
		return nil, err
	}

	return places, nil
}

func GetPlaceLinksById(placeId string) ([]models.PlaceLink, error) {
	placeLinks := []models.PlaceLink{}
	err := db.Select(&placeLinks, "SELECT * FROM place_link WHERE place_id=?", placeId)
	if err != nil {
		return nil, err
	}

	return placeLinks, nil
}

func CreatePlace(placeId string, userId string, title string, address string, instagramId string, facebookId string, twitterId string) error {
	_, err := db.Exec(`
		INSERT INTO place (id, user_id, title, address, instagram_id, facebook_id, twitter_id)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`, placeId, userId, title, address, instagramId, facebookId, twitterId)
	return err
}

func CreatePlaceLinks(placeId string, linkContents []string) error {
	if len(linkContents) == 0 {
		return nil
	}

	links := []models.PlaceLink{}
	for _, content := range linkContents {
		links = append(links, models.PlaceLink{
			Id:      utils.CreateId(),
			PlaceId: placeId,
			Content: content,
		})
	}

	_, err := db.NamedExec(`
		INSERT INTO place_link (id, place_id, content)
		VALUES (:id, :place_id, :content)
	`, links)

	return err
}

func CreatePlaceImages(placeId string, imageIds []string) error {
	if len(imageIds) == 0 {
		return nil
	}

	placeImages := []models.PlaceImage{}
	for idx, imageId := range imageIds {
		placeImages = append(placeImages, models.PlaceImage{
			Id:      imageId,
			PlaceId: placeId,
			Idx:     idx,
		})
	}

	_, err := db.NamedExec(`
		INSERT INTO place_image (id, place_id, idx)
		VALUES (:id, :place_id, :idx)
	`, placeImages)
	return err
}

func CreatePlaceLocations(locations []models.PlaceLocation) error {
	_, err := db.NamedExec(`
		INSERT INTO place_location (id, place_id, title, description)
		VALUES (:id, :place_id, :title, :description)
	`, locations)
	return err
}
