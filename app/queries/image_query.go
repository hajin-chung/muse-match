package queries

func CreateImage(imageId string, artId string, idx int) error {
	_, err := db.Exec("INSERT INTO image (id, artId, idx) VALUES ($1, $2, $3)", imageId, artId, idx)
	return err
}
