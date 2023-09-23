package queries

import "musematch/app/models"

func GetArtsByUserId(userId string) ([]models.ArtWithThumbnail, error) {
	arts := []models.ArtWithThumbnail{}
	err := db.Select(
		&arts,
		"SELECT art.Id, name, description, price, status, image.id as thumbnail FROM art RIGHT JOIN image ON art.id = image.artId WHERE art.userId = $1",
		userId,
	)
	if err != nil {
		return arts, err
	}

	return arts, nil
}
