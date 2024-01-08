package queries

import (
	"musematch/models"
)

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
	artListMap := models.UserArtListMap{
		Item: map[string][]string{},
	}
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

// func GetUserFirstArtList(userId string) (models.UserArtListWithInfo, error) {
// }

func GetUserArtMap(id string) (models.UserArtMap, error) {
	artMap := models.UserArtMap{}
	arts, err := GetArtInfosByUserId(id)
	if err != nil {
		return nil, err
	}

	for _, art := range arts {
		artMap[art.Id] = art
	}
	return artMap, nil
}

func UpdateUser(
	id string,
	name string,
	description string,
	instagramId string,
	facebookId string,
	twitterId string,
	note string,
) error {
	_, err := db.Exec(`
		UPDATE user
		SET name=?, description=?, instagram_id=?, facebook_id=?, twitter_id=?, note=?
		WHERE id=?;`,
		name, description,
		instagramId, facebookId, twitterId,
		note, id,
	)
	return err
}

func DeleteUserLink(id string) error {
	_, err := db.Exec("DELETE FROM user_link WHERE user_id=?", id)
	return err
}

func UpdateUserLink(id string, links []models.UserLink) error {
	err := DeleteUserLink(id)
	if err != nil {
		return err
	}

	if len(links) == 0 {
		return nil
	}

	_, err = db.NamedExec(`
		INSERT INTO user_link (id, user_id, content)
		VALUES (:id, :user_id, :content)
	`, links)

	return err
}

func DeleteUserHistory(userId string) error {
	_, err := db.Exec("DELETE FROM user_history WHERE user_id=?", userId)
	return err
}

func UpdateUserHistory(id string, histories []models.UserHistory) error {
	err := DeleteUserHistory(id)
	if err != nil {
		return err
	}

	if len(histories) == 0 {
		return nil
	}

	_, err = db.NamedExec(`
		INSERT INTO user_history (id, user_id, title, content)
		VALUES (:id, :user_id, :title, :content)
	`, histories)

	return err
}

func GetUserProfile(userId string) (*models.UserProfile, error) {
	user, err := GetUserById(userId)
	if err != nil {
		return nil, err
	}

	link, err := GetUserLink(userId)
	if err != nil {
		return nil, err
	}

	history, err := GetUserHistory(userId)
	if err != nil {
		return nil, err
	}

	artList, err := GetUserArtLists(userId)
	if err != nil {
		return nil, err
	}

	arts, err := GetUserArtMap(userId)
	if err != nil {
		return nil, err
	}

	profile := models.UserProfile{
		User:    user,
		Link:    link,
		History: history,
		ArtList: artList,
		Arts:    arts,
	}

	return &profile, nil
}

func DeleteUserArtList(userId string) error {
	_, err := db.Exec("DELETE FROM user_art_list WHERE user_id=$1", userId)
	return err
}

func UpdateUserArtList(userId string, artList []models.UserArtList) error {
	err := DeleteUserArtList(userId)
	if err != nil {
		return err
	}

	_, err = db.NamedExec(`
		INSERT INTO user_art_list (id, user_id, title)
		VALUES (:id, :user_id, :title)`,
		artList,
	)
	return err
}

func DeleteUserArtListItem() error {
	_, err := db.Exec("DELETE FROM user_art_list_item")
	return err
}

func UpdateUserArtListItem(userId string, listItems []models.UserArtListItem) error {
	err := DeleteUserArtListItem()
	if err != nil {
		return err
	}

	_, err = db.NamedExec(`
		INSERT INTO user_art_list_item (list_id, art_id, idx)
		VALUES (:list_id, :art_id, :idx)`,
		listItems,
	)
	return err
}

func UpdateUserIsOwner(userId string) error {
	_, err := db.Exec(`UPDATE user SET is_owner = 1 WHERE id=?`, userId)
	return err
}
