package queries

import "musematch/app/models"

func GetUserBySub(sub string) (models.User, error) {
	user := models.User{}
	err := db.Get(&user, "SELECT * FROM user WHERE sub = $1", sub)
	if err != nil {
		return user, err
	}
	return user, nil
}

func CreateUser(newUser *models.User) error {
	_, err := db.NamedExec(
		"INSERT INTO user (id, name, email, sub, picture, description, history) VALUES (:id, :name, :email, :picture, :description, :history)",
		newUser,
	)
	if err != nil {
		return err
	}

	return nil
}
