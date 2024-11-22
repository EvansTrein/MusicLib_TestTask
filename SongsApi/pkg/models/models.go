package models

import "gorm.io/gorm"

type Song struct {
	gorm.Model
	MusicGroup  string `gorm:"not null"`
	SongName    string `gorm:"not null"`
	ReleaseDate string
	Text        string
	Link        string
}

type SongData struct {
	MusicGroup  string `json:"group"`
	SongName    string `json:"song"`
	ReleaseDate string `json:"releaseDate"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}

type RequestData struct {
	MusicGroup string `json:"group" binding:"required"`
	Song       string `json:"song" binding:"required"`
}
