package models

import "gorm.io/gorm"

type Song struct {
	gorm.Model
	Group       string `gorm:"not null"`
	SongName    string `gorm:"not null"`
	ReleaseDate string
	Text        string
	Link        string
}

type SongData struct {
	Group       string `json:"group"`
	SongName    string `json:"song"`
	ReleaseDate string `json:"releaseDate"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}

type RequestData struct {
	Group string `json:"group" binding:"required"`
	Song  string `json:"song" binding:"required"`
}
