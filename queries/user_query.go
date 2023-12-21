package queries

import "musematch/models"

func GetUserBySub(sub string) (*models.User, error) {
	user := models.User{}
	err := db.Get(&user, "SELECT * FROM user WHERE sub = $1", sub)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUserById(id string) (*models.User, error) {
	user := models.User{}
	err := db.Get(&user, "SELECT * FROM user WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func CreateUser(newUser *models.User) error {
	_, err := db.NamedExec(`
		INSERT INTO user 
			(id, name, email, sub, picture, description, note, is_owner, instagram_id, facebook_id, twitter_id)
		VALUES 
			(:id, :name, :email, :sub, :picture, :description, :note, :is_owner, :instagram_id, :facebook_id, :twitter_id)`,
		newUser,
	)
	return err
}
