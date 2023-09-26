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
