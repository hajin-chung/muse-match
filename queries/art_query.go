package queries

import (
	"musematch/models"
)

func GetArtInfosByUserId(userId string) ([]models.ArtInfo, error) {
	arts := []models.ArtInfo{}
	err := db.Select(
		&arts,
		`SELECT 
			art.id, art.name, art.description, art.user_id, art.price, art_image.id as thumbnail, user.name as artist 
		FROM art
			LEFT JOIN art_image ON art.id = art_image.art_id
			LEFT JOIN user ON art.user_id = user.id
		WHERE art.user_id = $1 AND art_image.idx = 0`,
		userId,
	)

	if err != nil {
		return nil, err
	}
	return arts, nil
}

func ArtCreate(userId string, artId string, name string, description string, price int, info string) error {
	_, err := db.Exec(`
		INSERT INTO art (id, name, description, user_id, price, info)
		VALUES ($1, $2, $3, $4, $5, $6)`,
		artId, name, description, userId, price, info,
	)
	return err
}

func ArtUpdate(artId string, name string, description string, price int, info string) error {
	_, err := db.Exec(`
		UPDATE art
		SET name=$2, description=$3, price=$4, info=$5
		WHERE id=$1`,
		artId, name, description, price, info,
	)
	return err
}

func ArtImagesCreate(artId string, imageIds []string) error {
	if len(imageIds) == 0 {
		return nil
	}

	artImages := []models.ArtImage{}
	for idx, imageId := range imageIds {
		artImages = append(artImages, models.ArtImage{
			Id:    imageId,
			ArtId: artId,
			Idx:   idx,
		})
	}

	_, err := db.NamedExec(`
		INSERT INTO art_image (art_id, id, idx)
		VALUES(:art_id, :id, :idx)`,
		artImages,
	)
	return err
}

func ArtImagesDelete(artId string) error {
	_, err := db.Exec(`DELETE FROM art_image WHERE art_id=$1`, artId)
	return err
}

func ArtImagesUpdate(artId string, imageIds []string) error {
	if len(imageIds) == 0 {
		return nil
	}

	err := ArtImagesDelete(artId)
	if err != nil {
		return err
	}

	err = ArtImagesCreate(artId, imageIds)
	return err
}

func ArtTagsCreate(artId string, tags []string) error {
	if len(tags) == 0 {
		return nil
	}

	artTags := []models.ArtTag{}
	for _, tag := range tags {
		artTags = append(artTags, models.ArtTag{
			ArtId: artId,
			Tag:   tag,
		})
	}

	_, err := db.NamedExec(`
		INSERT INTO art_tag (art_id, tag)
		VALUES (:art_id, :tag)
	`, artTags)

	return err
}

func ArtTagsDelete(artId string) error {
	_, err := db.Exec("DELETE FROM art_Tag WHERE art_id=$1", artId)
	return err
}

func ArtTagsUpdate(artId string, tags []string) error {
	if len(tags) == 0 {
		return nil
	}

	err := ArtTagsDelete(artId)
	if err != nil {
		return err
	}

	err = ArtTagsUpdate(artId, tags)
	return err
}
