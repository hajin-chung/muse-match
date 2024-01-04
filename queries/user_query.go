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

func GetUserLink(id string) ([]models.UserLink, error) {
	link := []models.UserLink{}
	err := db.Select(&link, "SELECT * FROM user_link WHERE user_id = $1", id)
	return link, err
}

func GetUserHistory(id string) ([]models.UserHistory, error) {
	history := []models.UserHistory{}
	err := db.Select(&history, "SELECT * FROM user_history WHERE user_id=$1", id)
	return history, err
}

func GetUserArtLists(id string) (*models.UserArtListMap, error) {
	artListMap := models.UserArtListMap{}
	artLists := []models.UserArtList{}

	err := db.Select(&artLists, "SELECT * FROM user_art_list WHERE user_id=$1", id)
	if err != nil {
		return nil, err
	}

	for _, artList := range artLists {
		artListItems := []models.UserArtListItem{}
		err = db.Select(
			&artListItems, 
			"SELECT * FROM user_art_list_item WHERE list_id=$1",
			artList.Id,
		)
		if err != nil {
			return nil, err
		}
		
		for _, item := range artListItems {
			artListMap.Item[artList.Id] = 
				append(artListMap.Item[artList.Id], item.ArtId)
		}
	}

	artListMap.List = artLists
	return &artListMap, nil
}

func GetUserArtMap(id string) (models.UserArtMap, error) {
	artMap := models.UserArtMap{}
	arts, err := GetArtsByUserId(id)
	if err != nil {
		return nil, err
	}

	for _, art := range arts {
		artMap[art.Id] = art
	}
	return artMap, nil
}
