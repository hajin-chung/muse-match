package models

import "database/sql"

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
	Price       sql.NullInt32
	Status      string
}

type NewArtInfo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	ImageCount  int    `json:"imageCount"`
}

type Image struct {
	Id    string
	ArtId string `db:"artId"`
	Idx   int32
}

type ArtWithThumbnail struct {
	Id          string
	Name        string
	Description string
	UserId      string `db:"userId"`
	Price       sql.NullInt32
	Status      string
	Thumbnail   string
}

type NewExhibitInfo struct {
	Title     string `db:"title" json:"title"`
	Location  string `db:"location" json:"location"`
	StartDate string `db:"startDate" json:"startDate"`
	EndDate   string `db:"endDate" json:"endDate"`
}

type Exhibit struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	Location  string `json:"location"`
	StartDate string `db:"startDate" json:"startDate"`
	EndDate   string `db:"endDate" json:"endDate"`
}

type ExhibitArt struct {
	ArtId       string `db:"artId"`
	Name        string
	Description string `db:"description"`
	Thumbnail   string
	Price       sql.NullInt32
	Status      string
	UserId      string `db:"userId"`
	UserName    string `db:"userName"`
}
