package queries

import "musematch/app/models"

func CreateImage(imageId string, artId string, idx int) error {
	_, err := db.Exec("INSERT INTO image (id, artId, idx) VALUES ($1, $2, $3)", imageId, artId, idx)
	return err
}

func GetImagesByArtId(artId string) ([]models.Image, error) {
	images := []models.Image{}

	err := db.Select(&images, "SELECT * FROM image WHERE artId = $1", artId)
	return images, err
}
