package queries

import (
	"musematch/app/models"
	"musematch/app/utils"
)

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

func CreateExhibit(newExhibit models.NewExhibitInfo) (string, error) {
	id := utils.CreateId()
	_, err := db.Exec(
		"INSERT INTO exhibit (id, title, location, startDate, endDate) VALUES ($1, $2, $3, $4, $5)",
		id, newExhibit.Title, newExhibit.Location, newExhibit.StartDate, newExhibit.EndDate,
	)
	return id, err
}

func UpdateExhibit(newExhibit models.Exhibit) error {
	_, err := db.Exec(
		"UPDATE exhibit SET title = $1, location = $2, startDate = $3, endDate = $4 WHERE id = $5",
		newExhibit.Title, newExhibit.Location, newExhibit.StartDate, newExhibit.EndDate, newExhibit.Id,
	)

	return err
}
