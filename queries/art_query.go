package queries

import (
	"log"
	"musematch/models"
	"musematch/utils"
)

func GetArtsByUserId(userId string) ([]models.ArtWithThumbnail, error) {
	arts := []models.ArtWithThumbnail{}
	err := db.Select(
		&arts,
		`SELECT art.id, art.name, art.description, art.userId, art.price, art.status, image.id as thumbnail from art left join image ON art.id = image.artId where art.userId=$1 AND image.idx = 0`,
		userId,
	)
	if err != nil {
		log.Println(err)
		return nil, err
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
		"INSERT INTO art (id, name, description, userId, price, status) VALUES ($1, $2, $3, $4, NULL, '')",
		id, newArtInfo.Title, newArtInfo.Description, userId)
	if err != nil {
		return "", err
	}

	return id, nil
}

func UpdateArt(artId string, newArtInfo models.NewArtInfo, userId string) error {
	_, err := db.Exec(
		"UPDATE art SET name = $1, description = $2, price = $3 WHERE id = $4 AND userId = $5",
		newArtInfo.Title, newArtInfo.Description, newArtInfo.Price, artId, userId)

	return err
}

func GetArtsByExhibitId(exhibitId string) ([]models.ExhibitArt, error) {
	arts := []models.ExhibitArt{}
	err := db.Select(
		&arts,
		`SELECT 
			art.id as artId, 
			art.name, 
			art.description, 
			image.id as thumbnail ,
			art.price, 
			art.status, 
			art.userId, 
			user.name as userName
		FROM exhibitArts 
		LEFT JOIN art ON exhibitArts.artId = art.id
		LEFT JOIN image ON art.id = image.artId 
		LEFT JOIN user ON art.userId = user.id
		WHERE exhibitArts.exhibitId = $1 AND image.idx = 0`,
		exhibitId,
	)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return arts, nil
}
