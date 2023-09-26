package queries

import "musematch/app/models"

func GetExhibits() ([]models.Exhibit, error) {
	exhibits := []models.Exhibit{}

	err := db.Select(&exhibits, `SELECT * FROM exhibit`)
	if err != nil {
		return nil, err
	}

	return exhibits, nil
}

func GetExhibitById(exhibitId string) (models.Exhibit, error) {
	exhibit := models.Exhibit{}

	err := db.Get(&exhibit, "SELECT * FROM exhibit WHERE id = $1", exhibitId)
	return exhibit, err
}
