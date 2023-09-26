package models

type User struct {
	Id          string
	Name        string
	Email       string
	Sub         string
	Picture     string
	Description string
	History     string
}

type UpdateUserInfo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	History     string `json:"history"`
}

type Art struct {
	Id          string
	Name        string
	Description string
	UserId      string `db:"userId"`
	Price       int32
	Status      string
}

type NewArtInfo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageCount  int    `json:"imageCount"`
}

type Image struct {
	Id    string
	ArtId string
	Idx   int32
}

type ArtWithThumbnail struct {
	Id          string
	Name        string
	Description string
	UserId      string `db:"userId"`
	Price       int32
	Status      string
	Thumbnail   string
}

type Exhibit struct {
	Id        string
	Title     string
	Location  string
	StartDate string `db:"startDate"`
	EndDate   string `db:"endDate"`
}

type ExhibitArt struct {
	ArtId       string `db:"artId"`
	Name        string
	Description string `db:"description"`
	Thumbnail   string
	Price       int
	Status      string
	UserId      string `db:"userId"`
	UserName    string `db:"userName"`
}
