package queries

func CreateImage(imageId string, artId string) error {
	_, err := db.Exec("INSERT INTO image (id, url, artId) VALUES ($1, '', $2)", imageId, artId)
	return err
}
