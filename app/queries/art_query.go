package queries

import (
	"log"
	"musematch/app/models"
	"musematch/app/utils"
)

func GetArtsByUserId(userId string) ([]models.ArtWithThumbnail, error) {
	arts := []models.ArtWithThumbnail{}
	err := db.Select(
		&arts,
		`SELECT 
			art.Id, art.name, art.description, art.price, art.status, image.id as thumbnail 
			FROM image LEFT JOIN art ON art.id = image.artId WHERE art.userId = $1 GROUP BY art.id HAVING COUNT(image.id) = 1`,
		userId,
	)
	if err != nil {
		return arts, err
	}

	return arts, nil
}

func GetArtById(artId string) (models.Art, error) {
	art := models.Art{}
	err := db.Get(&art, "SELECT * FROM art WHERE id = $1", artId)
	if err != nil {
		log.Println(err)
		return art, err
	}

	return art, nil
}

func CreateArt(newArtInfo models.NewArtInfo, userId string) (string, error) {
	id := utils.CreateId()
	_, err := db.Exec(
		"INSERT INTO art (id, name, description, userId, price, status) VALUES ($1, $2, $3, $4, 0, '')",
		id, newArtInfo.Title, newArtInfo.Description, userId)
	if err != nil {
		return "", err
	}

	return id, nil
}
