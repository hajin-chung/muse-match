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

type Art struct {
	Id          string
	Name        string
	Description string
	UserId      string
	Price       int32
	Status      string
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
	UserId      string
	Price       int32
	Status      string
	Thumbnail   string
}
