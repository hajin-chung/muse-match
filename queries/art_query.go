package queries

import "musematch/models"

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
