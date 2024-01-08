package queries

import (
	"musematch/models"
	"musematch/utils"
	"time"
)

func GetPlaceById(placeId string) (*models.Place, error) {
	place := models.Place{}
	err := db.Get(&place, "SELECT * FROM place WHERE id=?", placeId)
	if err != nil {
		return nil, err
	}
	return &place, nil
}

func GetPlaceImagesById(placeId string) ([]models.PlaceImage, error) {
	images := []models.PlaceImage{}
	err := db.Select(&images, "SELECT * FROM place_image WHERE place_id=?", placeId)
	if err != nil {
		return nil, err
	}
	return images, nil
}

func GetPlaceLocationsById(placeId string) ([]models.PlaceLocation, error) {
	locations := []models.PlaceLocation{}
	err := db.Select(&locations, "SELECT * FROM place_location WHERE place_id=?", placeId)
	if err != nil {
		return nil, err
	}
	return locations, nil
}

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
	if len(locations) == 0 {
		return nil
	}

	_, err := db.NamedExec(`
		INSERT INTO place_location (id, place_id, title, description)
		VALUES (:id, :place_id, :title, :description)
	`, locations)
	return err
}

func UpdatePlace(placeId string, title string, address string, instagramId string, facebookId string, twitterId string) error {
	_, err := db.Exec(`
		UPDATE place
		SET title=?, address=?, instagram_id=?, facebook_id=?, twitter_id=?
		WHERE id=?
	`, title, address, instagramId, facebookId, twitterId, placeId)
	return err
}

func DeletePlaceLinks(placeId string) error {
	_, err := db.Exec(`
		DELETE FROM place_link WHERE place_id=?
	`, placeId)
	return err
}

func UpdatePlaceLinks(placeId string, linkContents []string) error {
	err := DeletePlaceLinks(placeId)
	if err != nil {
		return err
	}
	err = CreatePlaceLinks(placeId, linkContents)
	return err
}

func DeletePlaceImages(placeId string) error {
	_, err := db.Exec("DELETE FROM place_image WHERE place_id=?", placeId)
	return err
}

func UpdatePlaceImages(placeId string, imageIds []string) error {
	err := DeletePlaceImages(placeId)
	if err != nil {
		return err
	}
	err = CreatePlaceImages(placeId, imageIds)
	return err
}

func DeletePlaceLocations(placeId string) error {
	_, err := db.Exec("DELETE FROM place_location WHERE place_id=?", placeId)
	return err
}

func UpdatePlaceLocations(placeId string, locations []models.PlaceLocation) error {
	err := DeletePlaceLocations(placeId)
	if err != nil {
		return err
	}
	err = CreatePlaceLocations(locations)
	return err
}

func GetPlaceArtsById(placeId string) ([]models.ArtInfo, error) {
	currentDate := utils.DateToString(time.Now())

	arts := []models.ArtInfo{}
	err := db.Select(&arts, `
		SELECT 
			art.id, art.name, art.description, art.user_id, art.price, 
			user.name as artist, 
			art_image.id as thumbnail
		FROM exhibit
			LEFT JOIN place_location ON place_location.id = exhibit.location_id
			LEFT JOIN place ON place.id = place_location.place_id
			LEFT JOIN art ON art.id = exhibit.art_id
			LEFT JOIN art_image ON art.id = art_image.art_id
			LEFT JOIN user ON art.user_id = user.id
		WHERE place.id = $1
			AND art_image.idx = 0
			AND start_date < $2 
			AND end_date > $2
	`, placeId, currentDate)
	if err != nil {
		return nil, err
	}
	return arts, nil
}
