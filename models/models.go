package models

// TODO: add createdAt for all types

type User struct {
	Id          string `db:"id"`
	Name        string `db:"name"`
	Email       string `db:"email"`
	Sub         string `db:"sub"`
	Picture     string `db:"picture"`
	Description string `db:"description"`
	Note        string `db:"note"`
	IsOwner     int    `db:"is_owner"`

	InstagramId string `db:"instagram_id"`
	FacebookId  string `db:"facebook_id"`
	TwitterId   string `db:"twitter_id"`
}

type UserInfo struct {
	Id string `db:"id"`
	Name string `db:"name"`
	ArtCount int `db:"art_count"`
}

type UserLink struct {
	Id      string `db:"id"`
	UserId  string `db:"user_id"`
	Content string `db:"content"`
}

type UserHistory struct {
	Id      string `db:"id"`
	UserId  string `db:"user_id"`
	Title   string `db:"title"`
	Content string `db:"content"`
}

type UserArtList struct {
	Id     string `db:"id"`
	UserId string `db:"user_id"`
	Title  string `db:"title"`
}

type UserArtListItem struct {
	ListId string `db:"list_id"`
	ArtId  string `db:"art_id"`
	Idx    int    `db:"idx"`
}

type UserArtListMap struct {
	List []UserArtList
	Item map[string][]string // ListId -> ArtId
}

type UserArtListWithInfo struct {
	Title string
	Arts  []ArtInfo
}

type UserArtMap map[string]ArtInfo

type UserProfile struct {
	User    *User
	Link    []UserLink
	History []UserHistory
	ArtList *UserArtListMap
	Arts    UserArtMap
}

type UserLikeUser struct {
	UserId     string `db:"user_id"`
	LikeUserId string `db:"like_user_id"`
}

type UserLikeArt struct {
	UserId string `db:"user_id"`
	ArtId  string `db:"art_id"`
}

type UserLikePlace struct {
	UserId  string `db:"user_id"`
	PlaceId string `db:"place_id"`
}

type Art struct {
	Id          string `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
	UserId      string `db:"user_id"`
	Price       int    `db:"price"`
	Info        string `db:"info"`
}

type ArtInfo struct { // used in cards
	Id          string `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
	UserId      string `db:"user_id"`
	Price       int    `db:"price"`
	Artist      string `db:"artist"`
	Thumbnail   string `db:"thumbnail"`
}

type ArtTag struct {
	ArtId string `db:"art_id"`
	Tag   string `db:"tag"`
}

type ArtImage struct {
	Id    string `db:"id"`
	ArtId string `db:"art_id"`
	Idx   int    `db:"idx"`
}

type Place struct {
	Id          string `db:"id"`
	UserId      string `db:"user_id"`
	Title       string `db:"title"`
	Address     string `db:"address"`
	InstagramId string `db:"instagram_id"`
	FacebookId  string `db:"facebook_id"`
	TwitterId   string `db:"twitter_id"`
}

type PlaceInfo struct {
	Id        string `db:"id"`
	Title     string `db:"title"`
	Address   string `db:"address"`
	Thumbnail string `db:"thumbnail"`
}

type PlaceLink struct {
	Id      string `db:"id"`
	PlaceId string `db:"place_id"`
	Content string `db:"content"`
}

type PlaceImage struct {
	Id      string `db:"id"`
	PlaceId string `db:"place_id"`
	Idx     int    `db:"idx"`
}

type PlaceLocation struct {
	Id          string `db:"id"`
	PlaceId     string `db:"place_id"`
	Title       string `db:"title"`
	Description string `db:"description"`
}

type Exhibit struct {
	Id         string `db:"id"`
	LocationId string `db:"location_id"`
	StartDate  string `db:"start_date"`
	EndDate    string `db:"end_date"`
	ArtId      string `db:"art_id"`
	State      string `db:"state"`
}
